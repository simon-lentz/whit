package arrows

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/validation"
)

func Test_Graph_prepares_meta_and_validates_properties_are_unique(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewTerminatingIssueCollector()

	emptyLabels := []string{}
	n1 := newNode("n0", "Car", emptyLabels)
	addProperty(n1, "regNbr+", "string")

	n2 := newNode("n1", "Registered", emptyLabels)
	addProperty(n2, "regNbr+", "string")

	n3 := newNode("n2", "Vehicle", emptyLabels)
	addProperty(n3, "regNbr+", "string")

	r1 := newRelation("r1", "#INHERITS", n1, n2)
	r2 := newRelation("r2", "#INHERITS", n2, n3)

	g := newGraph(
		[]*Node{n1, n2, n3},
		[]*Relationship{r1, r2},
	)
	g.prepareMeta(ic)
	tt.CheckEqual(3, ic.Count())

	m := utils.NewSet[string]()
	ic.EachIssue(func(issue validation.Issue) {
		m.Add(issue.Message())
	})
	tt.CheckEqual(3, m.Size())
	tt.CheckTrue(m.Contains("Node 'Car' property 'regNbr' defined more than once, (inherited from 'Registered')"))
	tt.CheckTrue(m.Contains("Node 'Car' property 'regNbr' defined more than once, (inherited from 'Vehicle')"))
	tt.CheckTrue(m.Contains("Node 'Registered' property 'regNbr' defined more than once, (inherited from 'Vehicle')"))
}

func Test_Graph_prepares_meta_to_have_primaryKeys_set_from_inherited_types(t *testing.T) {
	tt := testutils.NewTester((t))
	ic := validation.NewTerminatingIssueCollector()
	emptyLabels := []string{}
	n1 := newNode("n0", "Car", emptyLabels)
	n2 := newNode("n1", "Registered", emptyLabels)
	n3 := newNode("n2", "Vehicle", emptyLabels)
	n4 := newNode("n4", "SuperRegistered", emptyLabels)
	addProperty(n4, "regNbr+", "string")
	r1 := newRelation("r1", "#INHERITS", n1, n2)
	r2 := newRelation("r2", "#INHERITS", n2, n3)
	r3 := newRelation("r3", "#INHERITS", n2, n4)

	// assert emptyLabels is still empty (for sanity)
	tt.CheckEqual(0, len(emptyLabels))

	g := newGraph(
		[]*Node{n1, n2, n3, n4},
		[]*Relationship{r1, r2, r3},
	)
	g.prepareMeta(ic)
	// Check super types of Car
	tt.CheckTrue(n1.superTypes.Contains("Registered"))
	tt.CheckTrue(n1.superTypes.Contains("SuperRegistered"))
	tt.CheckTrue(n1.superTypes.Contains("Vehicle"))

	// SuperRegistered: Check that regNbr is a primary key
	primaryKeyNames := utils.NewSetFrom(n4.primaryKeys, func(p *Property) string {
		return p.Name
	})
	tt.CheckTrue(primaryKeyNames.Contains("regNbr"))

	// Vehicle: Check that regNbr is a primary key in Vehicle
	primaryKeyNames = utils.NewSetFrom(n2.primaryKeys, func(p *Property) string {
		return p.Name
	})
	tt.CheckTrue(primaryKeyNames.Contains("regNbr"))

	// Car: Check that regNbr is a primary key in Car
	primaryKeyNames = utils.NewSetFrom(n1.primaryKeys, func(p *Property) string {
		return p.Name
	})
	tt.CheckTrue(primaryKeyNames.Contains("regNbr"))
}

func Test_Graph_errors_on_unknown_meta_relationship(t *testing.T) {
	tt := testutils.NewTester((t))
	ic := validation.NewTerminatingIssueCollector()
	emptyLabels := []string{}
	n1 := newNode("n0", "Car", emptyLabels)
	n2 := newNode("n1", "Registered", emptyLabels)
	r1 := newRelation("r1", "#UNKNOWN", n1, n2)

	g := newGraph(
		[]*Node{n1, n2},
		[]*Relationship{r1},
	)
	validation.Do(func() {
		g.prepareMeta(ic)
	})
	tt.CheckEqual(1, ic.Count())
	tt.CheckTrue(ic.HasErrors())

	m := messageSet(ic)
	expected := "relationship id 'r1' - unknown special relationship '#UNKNOWN'"
	tt.CheckTruef(m.Contains(expected), "Expected: '%s', was not in the set: %v", expected, messages(m))
}

func Test_Graph_Duplicate_relationship_not_allowed(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewIssueCollector()

	person := newNode("p", "Person", []string{})
	addProperty(person, "name+", "string")

	hair := newNode("hair", "Hair", []string{})
	addProperty(hair, "color+", "string")

	r1 := newRelationship("r1", "HAS_HAIR", "p", "hair") // 01 Hair
	r2 := newRelationship("r2", "HAS_HAIR", "p", "hair") // 01 Hair

	g := newGraph(
		[]*Node{person, hair},
		[]*Relationship{r1, r2})

	g.prepareMeta(ic)
	tt.CheckEqual(1, ic.Count())
	m := messageSet(ic)
	expected := "duplicate relationship 'HAS_HAIR' between 'Person' and 'Hair'"
	tt.CheckTruef(m.Contains(expected), "Expected: '%s', was not in the set: %v", expected, messages(m))
}

func Test_Graph_Relationship_requires_to_node_primary_key(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewIssueCollector()

	person := newNode("p", "Person", []string{})
	hair := newNode("hair", "Hair", []string{})

	r1 := newRelationship("r1", "HAIR_OF", "hair", "p")

	g := newGraph(
		[]*Node{person, hair},
		[]*Relationship{r1})

	g.prepareMeta(ic)
	tt.CheckEqual(1, ic.Count())
	m := messageSet(ic)
	expected := "node 'Person' has no primary key(s) - required for incoming relationships"
	tt.CheckTruef(m.Contains(expected), "Expected: '%s', was not in the set: %v", expected, messages(m))
}

// HELPER FUNCS.

func newGraph(nodes []*Node, rels []*Relationship) *Graph {
	return &Graph{
		Style:         make(map[string]interface{}),
		Nodes:         nodes,
		Relationships: rels,
		nodeByID:      nil,
		nodeByCaption: nil,
	}
}
func newRelation(id string, relType string, from *Node, to *Node) *Relationship {
	return &Relationship{
		ID:         id,
		Type:       relType,
		Style:      make(map[string]interface{}),
		Properties: make(map[string]string),
		FromID:     from.ID,
		ToID:       to.ID,
	}
}
func addProperty(n *Node, name, value string) {
	n.Properties[name] = value
}
func newNode(id string, caption string, labels []string) *Node {
	return &Node{
		ID:         id,
		Caption:    caption,
		Labels:     labels,
		Properties: map[string]string{},
	}
}

func messageSet(ic validation.IssueCollector) *utils.Set[string] {
	m := utils.NewSet[string]()
	ic.EachIssue(func(issue validation.Issue) {
		m.Add(issue.Message())
	})
	return m
}
func messages(set *utils.Set[string]) string {
	m := []string{}
	set.Each(func(s string) {
		m = append(m, fmt.Sprintf(`"%s"`, s))
	})
	return strings.Join(m, ", ")
}
