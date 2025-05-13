package arrows

import (
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/validation"
)

// TODO:
// * A composition that is 11 from one type cannot be in another composition (people sharing the same "Head"). All "from"
//   must then be 01 to show this. If allowed, the ability to be Part of different compositions also means that a Part can
//   float off and not be part of any composition (the 01 constraint is met), which requires a validation that a Part must
//   be part of some composition - it can only float off if explicitly created without a Composition and this cannot happen
//   from instance documents unless there are bugs in the implementation. THEREFORE:
//
//     * Forbid Part to be composed into different compositions and have this discussion later if it should be allowed.
//
// DONE:
//   * Compositions cannot have properties
//   * Duplicate compositions not allowed

func Test_Graph_Compositions_cannot_have_properties(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewIssueCollector()

	person := newNode("p", "Person", []string{}) // Composition
	addProperty(person, "name+", "string")

	hair := newNode("hair", "Hair", []string{})
	addProperty(hair, "color", "string")

	r1 := newRelationship("r1", "#COMPOSED_OF", "p", "hair") // 01 Hair
	addRelProperty(r1, "#to", "01")
	addRelProperty(r1, "nope", "01")

	g := newGraph(
		[]*Node{person, hair},
		[]*Relationship{r1})

	g.prepareMeta(ic)
	tt.CheckEqual(1, ic.Count())
	m := messageSet(ic)
	expected := "relationship properties other than #to/#from not allowed in COMPOSED type relationship - id 'r1' from node id 'p'"
	tt.CheckTruef(m.Contains(expected), "Expected: '%s', was not in the set: %v", expected, messages(m))
}

func Test_Graph_Duplicate_Composition_not_allowed(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewIssueCollector()

	person := newNode("p", "Person", []string{})
	addProperty(person, "name+", "string")

	hair := newNode("hair", "Hair", []string{})
	addProperty(hair, "color", "string")

	r1 := newRelationship("r1", "#COMPOSED_OF", "p", "hair") // 01 Hair
	r2 := newRelationship("r2", "#COMPOSED_OF", "p", "hair") // 01 Hair

	g := newGraph(
		[]*Node{person, hair},
		[]*Relationship{r1, r2})

	g.prepareMeta(ic)
	tt.CheckEqual(1, ic.Count())
	m := messageSet(ic)
	expected := "duplicate relationship 'COMPOSED_OF' between 'Person' and 'Hair'"
	tt.CheckTruef(m.Contains(expected), "Expected: '%s', was not in the set: %v", expected, messages(m))
}
