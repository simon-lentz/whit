// Package arrows contains utilities for operating on neo4j Arrows graphs.
package arrows

import (
	"github.com/gertd/go-pluralize"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/validation"
)

// Graph is the top object unarshalled from Arrows json data.
type Graph struct {
	Style         map[string]interface{} `json:"style"` // mix of string and numeric values
	Nodes         []*Node                `json:"nodes"`
	Relationships []*Relationship        `json:"relationships"`
	nodeByID      map[string]*Node
	nodeByCaption map[string]*Node
}

const (
	// MIXIN the "#Mixin" label.
	MIXIN = "#Mixin"

	// ABSTRACT the "#Abstract" label.
	ABSTRACT = "#Abstract"
)

var thePluralizer *pluralize.Client

func (g *Graph) prepareMeta(ic validation.IssueCollector) {
	thePluralizer = pluralize.NewClient()

	// create index to make it possible to lookup Node by Caption and by ID
	g.indexCaptionAndID(ic)

	// convert the separate relationships to direct pointers between Nodes
	g.convertRelationshipsToEdges(ic)

	// transitively follow all INHERITS relationships for all nodes and collect pointers
	// to all super types per node to have a quick lookup of the set of supertypes for an type.
	g.computeSuperTypes(ic)

	g.computePrimaryKeys(ic)
}

// computePrimaryKeys computes the set of primary keys for eveery node.
func (g *Graph) computePrimaryKeys(ic validation.IssueCollector) {
	// Make all nodes have their own primary properties
	for _, n := range g.Nodes {
		n.allProperties = make(map[string]*Property)
		// add all primary keys in the node itself
		for _, p := range n.nProperties {
			// Cannot be duplicate within the node since it is a map and UnMarshaling the json will make it
			// unique within the node.
			n.allProperties[p.Name] = p
			if p.primary {
				n.primaryKeys = append(n.primaryKeys, p)
			}
		}
	}
	// Then check all superTypes and mixins. (This works because all supertypes have been collected per node)
	for _, n := range g.Nodes {
		n.superTypes.Each(func(typeName string) {
			st := g.nodeByCaption[typeName]
			for _, p := range st.nProperties {
				if _, ok := n.allProperties[p.Name]; ok {
					ic.Collectf(validation.Error, "Node '%s' property '%s' defined more than once, (inherited from '%s')", n.Caption, p.Name, st.Caption) //nolint:lll
				} else {
					n.allProperties[p.Name] = p
				}
			}
			n.primaryKeys = append(n.primaryKeys, st.primaryKeys...)
		})
		n.possibleMixins.Each(func(typeName string) {
			mixin := g.nodeByCaption[typeName]
			for _, p := range mixin.nProperties {
				if _, ok := n.allProperties[p.Name]; ok {
					ic.Collectf(validation.Error, "Node '%s' property '%s' defined more than once, (mixed in from '%s')", n.Caption, p.Name, mixin.Caption) //nolint:lll
				} else {
					n.allProperties[p.Name] = p
				}
			}
			// TODO: Primary keys in mixins? Needed if there are relations to a mixin; they act as an interface
			// as the target for the mixin must have this primary key property.
			// Something like: n.primaryKeys = append(n.primaryKeys, mixin.primaryKeys...)
		})
		// // TODO: Log these in debug mode
		// keyNames := []string{}
		// for _, pk := range n.primaryKeys {
		// 	keyNames = append(keyNames, pk.Name)
		// }
		// fmt.Printf("// %s primary keys; %s\n", n.Caption, strings.Join(keyNames, ", "))
	}
	for _, n := range g.Nodes {
		if !n.HasIncomingAssociations() {
			continue
		}
		if len(n.primaryKeys) < 1 {
			ic.Collectf(validation.Error, "node '%s' has no primary key(s) - required for incoming relationships", n.Caption)
		}
	}
}

// computeSuperTypes buids a set of all super types and mixins (transitively) and stores them per node.
// This set serves as an index for further processing.
func (g *Graph) computeSuperTypes(ic validation.IssueCollector) { //nolint:gocognit
	// For all nodes n, if they have an outgoing IS_A relation, then the to type is a super type of n.
	// All supertypes of the to node (recursively) are also super types of n
	var depthFirst func(n *Node)
	depthFirst = func(n *Node) {
		// If node does not have computed supertypes, compute them first oterwise just return and stop recursion
		if n.superTypes == nil {
			n.superTypes = utils.NewSet[string]()
			for _, r := range n.outEdges {
				if r.isInherits {
					superNode := g.nodeByID[r.ToID]
					n.superTypes.Add(superNode.Caption)
					depthFirst(superNode)
					n.superTypes = n.superTypes.Union(superNode.superTypes)
				}
			}
			// // TODO: If loglevel is debug
			// superTypeNames := []string{}
			// n.superTypes.Each(func(s string) {
			// 	superTypeNames = append(superTypeNames, s)
			// })
			// fmt.Printf("// %s supertypes: %s\n", n.Caption, strings.Join(superTypeNames, ", "))
		}
	}
	var depthFirstMixin func(n *Node)
	depthFirstMixin = func(n *Node) {
		// If node does not have computed mixins, compute them first otherwise just return and stop recursion
		if n == nil {
			panic("no node !!")
		}
		if n.possibleMixins == nil {
			n.possibleMixins = utils.NewSet[string]()
			if n.isDataType {
				return
			}
			for _, label := range n.Labels {
				mixinNode, ok := g.nodeByCaption[label]
				if !ok {
					switch label {
					case ABSTRACT, MIXIN: // all regular labels are references to mixins
						continue
					default:
						ic.Collectf(validation.Error, "mixin type '%s' not found", label)
						continue
					}
				}
				if !n.possibleMixins.Add(label) { // node's mixins
					ic.Collectf(validation.Error, "mixin '%s' already possible mixin for type '%s'", label, n.Caption)
				}
				depthFirstMixin(mixinNode) // mixin's into node's mixins (i.e. nested)
				mixinNode.possibleMixins.Each((func(s string) {
					if !n.possibleMixins.Add(s) {
						ic.Collectf(validation.Error, "mixin '%s' already possible mixin for type '%s'", s, n.Caption)
					}
				}))
			}
			n.superTypes.Each(func(name string) {
				st, ok := g.nodeByCaption[name]
				if !ok {
					ic.Collectf(validation.Fatal, "super type '%s' not found", name)
				}
				depthFirstMixin(st) // mixin's into node superType
				st.possibleMixins.Each((func(s string) {
					if !n.possibleMixins.Add(s) {
						ic.Collectf(validation.Error, "mixin '%s' already possible mixin for type '%s'", s, n.Caption)
					}
				}))
			})

			// // TODO: If loglevel is debug
			// superTypeNames := []string{}
			// n.superTypes.Each(func(s string) {
			// 	superTypeNames = append(superTypeNames, s)
			// })
			// fmt.Printf("// %s supertypes: %s\n", n.Caption, strings.Join(superTypeNames, ", "))
		}
	}
	for _, n := range g.Nodes {
		depthFirst(n)
		if n.superTypes.Contains(n.Caption) {
			ic.Collectf(validation.Error, "circular inheritance detected for node '%s'", n.Caption)
		}
	}
	// all super types must be done before computing mixins
	for _, n := range g.Nodes {
		depthFirstMixin(n)
	}
}

// CueMarshalInstance interprets the graph as an instance graph and outputs to a cue schema (on stdout in first version).
// The intention is for the instance graph to be compliant with an already created schema.
// TODO: There needs to be validation, but this simply outputs content for the time being.
func (g *Graph) CueMarshalInstance(ic validation.IssueCollector) {
	// TODO:

	// Every Node is some sort of type. The Caption defines the name of the type
	// It may have this name as a Label as well (which is ignored)
	for _, n := range g.Nodes {
		n.normalize(ic, false) // normalize for instance mode instead of meta mode...
	}
	// TODO: Unfinished (or possibly remove if feature never is asked for)
}

// indexCaptionAndID creates the indexes for caption->Node and node-id -> Node.
func (g *Graph) indexCaptionAndID(ic validation.IssueCollector) {
	state := newMetaParseState()

	// Every Node is some sort of type. The Caption defines the name of the type
	// It may have this name as a Label as well (which is ignored)
	for _, n := range g.Nodes {
		n.normalize(ic, true)
		state.addCaption(ic, n)
		state.addID(ic, n)
	}
	if ic.HasErrors() {
		ic.Collectf(validation.Fatal, "earlier errors prevents further processing and production of output")
	}
	g.nodeByID = state.idToNode
	g.nodeByCaption = state.captionToNode
}

// convertRelationshipsToEdges stores "to node" pointer in "from node" for all relationships.
// This makes it possible to find "out edges" given a Node.
func (g *Graph) convertRelationshipsToEdges(ic validation.IssueCollector) { //nolint:gocognit
	// Set all relationships as edges
	for i, r := range g.Relationships {
		r.normalize(ic)

		var fromNode, toNode *Node
		var ok bool
		if fromNode, ok = g.nodeByID[r.FromID]; !ok {
			ic.Collectf(validation.Fatal, "the Relationship with id='%s' references 'from node' id='%s' - node not found", r.ID, r.FromID)
		}
		if toNode, ok = g.nodeByID[r.ToID]; !ok {
			ic.Collectf(validation.Fatal, "the Relationship with id='%s' references 'from node' id='%s' - node not found", r.ID, r.FromID)
		}

		// validate inheritance
		if r.IsInheritance() {
			if fromNode.isDataType {
				ic.Collectf(validation.Error, "a #DataType cannot inherit - '%s' inherits from '%s'", fromNode.Caption, toNode.Caption)
			}
			if toNode.isDataType {
				ic.Collectf(validation.Error, "a type cannot inherit from a #DataType - '%s' inherits from '%s'", fromNode.Caption, toNode.Caption) //nolint:lll
			}
		}
		// validate composition
		if r.IsComposition() {
			if fromNode.isDataType {
				ic.Collectf(validation.Error, "a #DataType cannot be a composition - '%s' is composed of '%s'", fromNode.Caption, toNode.Caption) //nolint:lll
			}
			if toNode.isDataType {
				ic.Collectf(validation.Error, "a composition cannot have #DataType parts - '%s' is composed of '%s'", fromNode.Caption, toNode.Caption) //nolint:lll
			}
		}
		if r.IsAssociation() {
			if fromNode.isDataType {
				ic.Collectf(validation.Error, "a #DataType cannot be in a relation - DataType '%s' has relation '%s' to '%s'", fromNode.Caption, r.Type, toNode.Caption) //nolint:lll
			}
			if toNode.isDataType {
				ic.Collectf(validation.Error, "a #DataType cannot be in a relation - '%s' has relation '%s' to DataType '%s'", fromNode.Caption, r.Type, toNode.Caption) //nolint:lll
			}
		}

		// Link nodes with pointers
		fromNode.outEdges = append(fromNode.outEdges, g.Relationships[i])
		toNode.inEdges = append(toNode.inEdges, g.Relationships[i])
	}
	// Once all nodes have received all of their edges, each node needs an index of relationship type to target nodes.
	// This then needs to be validated - for example A--[R]->B appearing more than once
	// TODO: CHECK that outEdges contains all relations, also those from supertypes? If not (because they are inherited in
	// the cue output, the logic here needs to find all to be able to verify)
	for _, n := range g.Nodes {
		relIndex := make(map[string]*utils.Set[string], len(n.outEdges))
		for _, r := range n.outEdges {
			if set, ok := relIndex[r.Type]; !ok {
				set = utils.NewSet[string]()
				relIndex[r.Type] = set
				_ = set.Add(r.ToID) // new set so cannot already have valey
			} else {
				// This type of relationship already seen. It is ok unless it is to the same type
				// of node as an earlier relation of the same type.
				toNodeName := g.nodeByID[r.ToID].Caption
				isNewTarget := set.Add(r.ToID)
				if !isNewTarget {
					ic.Collectf(validation.Error, "duplicate relationship '%s' between '%s' and '%s'", r.Type, n.Caption, toNodeName)
				}
			}
		}
		n.relIndex = relIndex
		// Mixins cannot have incoming relations
		if n.isMixin && len(n.inEdges) > 0 {
			ic.Collectf(validation.Error, "node '%s' is a Mixin - incoming relations not allowed", n.Caption)
		}
	}
}

type metaParseState struct {
	captionToNode map[string]*Node
	idToNode      map[string]*Node
}

func newMetaParseState() *metaParseState {
	state := &metaParseState{
		captionToNode: make(map[string]*Node),
		idToNode:      make(map[string]*Node),
	}
	return state
}

// addCaption adds a caption->*Node to the index, collects errors if caption is not unique.
// Only used for meta schema - for instances, indexing by Caption is not meaningful.
func (st *metaParseState) addCaption(ic validation.IssueCollector, n *Node) {
	// Assert: Caption is not empty
	if len(n.Caption) < 1 {
		ic.Collectf(validation.Error, "node with id='%s' must have a caption to define a node type", n.ID)
	} else if n2, ok := st.captionToNode[n.Caption]; ok {
		// Assert: Caption is unique (adds to index if not present)
		ic.Collectf(validation.Error, "node with id='%s' has the same caption as node='%s' - must have unique name", n.ID, n2.ID)
	} else {
		st.captionToNode[n.Caption] = n
	}
}

// addID adds the given node's ID to the index of id->Node, and returns an error if the node did not have an id
// or if id is not unique. (Applies to both meta and instance models).
func (st *metaParseState) addID(ic validation.IssueCollector, n *Node) {
	// Assert: ID is not empty
	if len(n.ID) < 1 {
		ic.Collectf(validation.Error, "node with empty id not allowed - node has caption '%s'", n.Caption)
	} else if n2, ok := st.idToNode[n.ID]; ok {
		ic.Collectf(validation.Error, "node with id='%s' must have unique id - found in nodes with captions: '%s' and '%s'", n.ID, n.Caption, n2.Caption) //nolint:lll
	} else {
		st.captionToNode[n.Caption] = n
		st.idToNode[n.ID] = n
	}
}
