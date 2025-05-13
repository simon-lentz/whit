package arrows

import (
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/validation"
)

func Test_Graph_properties_from_mixins_are_unique_simple_case(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewTerminatingIssueCollector()
	n1 := newNode("n1", "Target", []string{"Role"})
	addProperty(n1, "name", "string")

	n2 := newNode("n2", "Role", []string{"#Mixin"})
	addProperty(n2, "name", "string")
	g := newGraph(
		[]*Node{n1, n2},
		[]*Relationship{},
	)
	validation.Do(func() {
		g.prepareMeta(ic)
	})
	m := messageSet(ic)
	tt.CheckEqual(1, m.Size()) // assert not identical messages
	tt.CheckTrue(m.Contains("Node 'Target' property 'name' defined more than once, (mixed in from 'Role')"))

	tt.CheckEqual(1, ic.Count())
}
func Test_Graph_properties_from_mixins_are_unique_nested_mixin_case(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewTerminatingIssueCollector()
	n1 := newNode("n1", "Target", []string{"Role"})
	addProperty(n1, "name", "string")

	n2 := newNode("n2", "Role", []string{"#Mixin", "NestedMixin"})
	n3 := newNode("n3", "NestedMixin", []string{"#Mixin"})
	addProperty(n2, "name", "string")

	g := newGraph(
		[]*Node{n1, n2, n3},
		[]*Relationship{},
	)
	validation.Do(func() {
		g.prepareMeta(ic)
	})
	m := messageSet(ic)
	sz := m.Size()
	tt.CheckEqual(1, sz) // assert not identical messages
	tt.CheckTrue(m.Contains("Node 'Target' property 'name' defined more than once, (mixed in from 'Role')"))

	tt.CheckEqual(1, ic.Count())
}

func Test_Graph_properties_from_mixins_are_unique_superType_case(t *testing.T) {
	// Target <- Role(#Mixin) ->INHERITS -> SuperRole <- ExtraRole(#Mixin)
	tt := testutils.NewTester(t)
	ic := validation.NewTerminatingIssueCollector()
	n1 := newNode("n1", "Target", []string{"Role"})
	addProperty(n1, "name", "string")

	n2 := newNode("n2", "Role", []string{"#Mixin"})
	n3 := newNode("n3", "SuperRole", []string{"ExtraRole"})

	n4 := newNode("n4", "ExtraRole", []string{"#Mixin"})
	addProperty(n4, "name", "string")

	r1 := newRelation("r1", "#INHERITS", n2, n3)

	g := newGraph(
		[]*Node{n1, n2, n3, n4},
		[]*Relationship{r1},
	)
	validation.Do(func() {
		g.prepareMeta(ic)
	})
	m := messageSet(ic)
	tt.CheckEqual(1, m.Size()) // assert not identical messages
	tt.CheckTrue(m.Contains("Node 'Target' property 'name' defined more than once, (mixed in from 'ExtraRole')"))
}

func Test_Graph__mixins_must_be_unique(t *testing.T) {
	// Target <- Role(#Mixin) <- ExtraRole(#Mixin)
	// Target <- ExtraRole(#Mixin)
	tt := testutils.NewTester(t)
	ic := validation.NewTerminatingIssueCollector()
	n1 := newNode("n1", "Target", []string{"Role", "ExtraRole"})
	n2 := newNode("n2", "Role", []string{"#Mixin", "ExtraRole"})
	n3 := newNode("n3", "ExtraRole", []string{"#Mixin"})
	g := newGraph(
		[]*Node{n1, n2, n3},
		[]*Relationship{},
	)
	validation.Do(func() {
		g.prepareMeta(ic)
	})
	m := messageSet(ic)
	tt.CheckEqual(1, m.Size()) // assert not identical messages
	tt.CheckTrue(m.Contains("mixin 'ExtraRole' already possible mixin for type 'Target'"))
}
