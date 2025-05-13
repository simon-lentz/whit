package cypher

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/google/uuid"
	"github.com/tada/catch"
	"github.com/wyrth-io/whit/internal/jzon"
	"github.com/wyrth-io/whit/internal/tc"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/yammm"
	"golang.org/x/exp/slices"
)

// MergeGenerator generates merge Statements creation/merge data into a Cypher/Neo4j DB.
type MergeGenerator interface {
	// Process a graph of data that has been validated against the model described by the context.
	// If the graph was not validated prior to this call the behaviour is undefined and will most likely panic.
	Process(ctx yammm.Context, graph any) []Statement
}

type merger struct {
	yammm.BaseGraphListener
	replacementMap map[string]uuid.UUID
	links          []link
	statements     []Statement
}

type link struct {
	propName string // property name in fromType
	fromName string
	fromID   string
	toName   string
	toID     string
	propMap  map[string]any
}

// NewMergeGenerator returns a MergeGenerator for Cypher/Neo4j.
func NewMergeGenerator() MergeGenerator {
	return &merger{}
}

func (m *merger) Process(ctx yammm.Context, graph any) []Statement {
	// Get the ID replacements for the graph
	idMapper := yammm.NewIDMapper()
	var err error
	if m.replacementMap, err = idMapper.Map(ctx, graph); err != nil {
		panic(catch.Error("could not create UUID replacement map: %s", err))
	}
	// Walk the graph, create a statement for each instance to create (done in OnProperties).
	// All relationships are collected in the form of links and their creation is deferred until
	// all statements for all instances have been generated.
	walker := yammm.NewGraphWalker(ctx, m)
	walker.Walk(graph)

	// Generate statements for all (deferred) relationships.
	m.makeLinkStatements()
	return m.statements
}

func (m *merger) OnProperties(_ yammm.Context, t *yammm.Type, propMap map[string]any) {
	var uid uuid.UUID
	var err error
	if uid, err = t.InstanceID(propMap, m.replacementMap); err != nil {
		panic(catch.Error("Cannot create id for type %s: %s", t.Name, err))
	}
	// For float properties convert integer values in the map to float64.
	for k, v := range propMap {
		if t.PropByName(k).BaseType().Kind() == tc.FloatKind {
			_, ok := utils.GetFloat64(v)
			if !ok {
				// value is neither float32 nor float64
				// convert to float if value is a kind of integer
				i, ok := utils.GetInt64(v)
				if ok {
					propMap[k] = float64(i)
				} else {
					panic(catch.Error("expected float for '%s.%s': got type %t\n", t.Name, k, v))
				}
			}
		}
	}
	m.makeInstanceStatement(t, uid, propMap)
}

func (m *merger) OnEdge(
	ctx yammm.Context,
	a *yammm.Association,
	assocPropMap map[string]any,
	fromType *yammm.Type,
	fromPks map[string]any,
	toType *yammm.Type,
	toPks map[string]any,
) {
	var fromID, toID uuid.UUID
	var err error
	if fromID, err = fromType.InstanceID(fromPks, m.replacementMap); err != nil {
		panic(catch.Error("Cannot create id for from-type %s: %s", fromType.Name, err))
	}
	if toID, err = toType.InstanceID(toPks, m.replacementMap); err != nil {
		panic(catch.Error("Cannot create id for to-type %s: %s", toType.Name, err))
	}
	m.addLink(link{
		propName: a.PropertyName(ctx),
		fromName: fromType.Name,
		fromID:   fromID.String(),
		toName:   toType.Name,
		toID:     toID.String(),
		propMap:  assocPropMap,
	})
}

func (m *merger) OnCompositionEdge(ctx yammm.Context, comp *yammm.Composition,
	fromType *yammm.Type, fromPkMap map[string]any,
	toType *yammm.Type, toPkMap map[string]any) {
	var fromID, toID uuid.UUID
	var err error
	if fromID, err = fromType.InstanceID(fromPkMap, m.replacementMap); err != nil {
		panic(catch.Error("Cannot create id for from-type %s: %s", fromType.Name, err))
	}
	if toID, err = toType.InstanceID(toPkMap, m.replacementMap); err != nil {
		panic(catch.Error("Cannot create id for to-type %s: %s", toType.Name, err))
	}
	m.addLink(link{
		propName: comp.PropertyName(ctx),
		fromName: fromType.Name,
		fromID:   fromID.String(),
		toName:   toType.Name,
		toID:     toID.String(),
		propMap:  map[string]any{}, // cannot have properties.
	})
}
func (m *merger) addLink(aLink link) {
	m.links = append(m.links, aLink)
}
func getFloat32Slice(val any) []float32 {
	if val == nil {
		return nil
	}
	if reflect.ValueOf(val).Kind() != reflect.Slice {
		// Should not happen since validation should have been done.
		panic(catch.Error("spacevector did not get a slice value"))
	}
	kv := reflect.ValueOf(val)
	// the elements are interface as array could hold mixed types of values in JSON
	// TODO: Need to check that they are all either float 32 or 64.
	// TODO: Transform value to []float32 as this is needed for checking
	transformed := []float32{}
	for i := 0; i < kv.Len(); i++ {
		ki := kv.Index(i).Interface()
	again:
		switch n := ki.(type) {
		case int:
			transformed = append(transformed, float32(n))
		case int8:
			transformed = append(transformed, float32(n))
		case int16:
			transformed = append(transformed, float32(n))
		case int32:
			transformed = append(transformed, float32(n))
		case int64:
			transformed = append(transformed, float32(n))
		case float32:
			transformed = append(transformed, n)
			continue
		case float64:
			transformed = append(transformed, float32(n))
			continue
		case jzon.Node:
			ki = n.RawValue()
			goto again
		default:
			// Should not happen since this should have been validated.
			panic(catch.Error("spacevector got slice with non float value"))
		}
	}
	return transformed
}
func (m *merger) makeInstanceStatement(t *yammm.Type, uid uuid.UUID, propMap map[string]any) {
	// Change "id" to "uid" for Cypher
	props := utils.FilterMap(propMap, func(k string, _ any) bool { return k != "id" })
	props[UID] = uid.String()

	// Dig out vector values
	for _, p := range t.AllProperties() {
		if p.BaseType().Kind() == tc.SpacevectorKind {
			props[p.Name] = getFloat32Slice(props[p.Name])
		}
	}
	sortit := func(s []string) []string { slices.Sort(s); return s }
	// One label for the type and one for each type it extends (sorted due to stability of output).
	labels := append([]string{t.Name}, sortit(t.AllSuperTypes())...)
	labelsString := strings.Join(labels, ":")

	// Generate a MERGE statement matching on all labels and on the uid and then either setting
	// all properties, or merging properties if the node for the uid exists.
	template := `MERGE (n:%s {uid: $props.uid}) ON CREATE SET n = $props ON MATCH SET n += $props`
	// Save the statement.
	m.statements = append(m.statements, Statement{
		Source:     fmt.Sprintf(template, labelsString),
		Parameters: map[string]any{"props": props},
	})
}

// linkTemplate is filled out to become (for example):
//
//	MATCH(from:Person), (to:Car) WHERE from.uid = $fromUid AND to.uid = $toUid
//	CREATE (from)-[r:OWNES_Cars {since: $props.since}]->to
//
// .
var linkTemplate = `MATCH (from:%s), (to:%s)
WHERE from.uid = $fromUid AND to.uid = $toUid
CREATE (from)-[r:%s%s]->(to)`

// makeLinkStatements adds one Statement to the returned statements per deferred relationship
// link. This function uses the linkTemplate.
func (m *merger) makeLinkStatements() {
	for i := range m.links {
		theLink := m.links[i]

		// make a string out of the association properties. The template has
		// the enclosing {}, and it needs entries on the form `apropname: $prop.apropname, ...`
		mappedProperties := strings.Join(
			utils.Map(utils.Keys(theLink.propMap), func(p string) string {
				return fmt.Sprintf("%s: $props.%s", p, p)
			}),
			", ")
		// Need to skip the { properties } part if there are none, and also
		// skip including an empty "props" in the parameters.
		parameters := map[string]any{
			"fromUid": theLink.fromID,
			"toUid":   theLink.toID,
		}
		linkProperties := ""
		if len(theLink.propMap) > 0 {
			linkProperties = " {" + mappedProperties + "}"
			parameters["props"] = theLink.propMap
		}
		m.statements = append(m.statements, Statement{
			Source: fmt.Sprintf(linkTemplate,
				theLink.fromName,
				theLink.toName,
				theLink.propName, // for example "HAS", "OWNES"
				linkProperties,
			),
			Parameters: parameters,
		})
	}
}
