package arrows

import (
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/validation"
)

func Test_Relationship_normalize_association(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewTerminatingIssueCollector()
	r := newRelationship("r1", "SOMETHING", "n0", "n1")
	validation.Do(func() {
		r.normalize(ic)
	})
	tt.CheckEqual(0, ic.Count()) // no warnings or errors
	tt.CheckTrue(r.isAssociation)
	tt.CheckFalse(r.isInherits)
	tt.CheckFalse(r.isComposition)

	// assert internal flag and method are the same (only tested here)
	tt.CheckEqual(r.isInherits, r.IsInheritance())
	tt.CheckEqual(r.isAssociation, r.IsAssociation())
	tt.CheckEqual(r.isComposition, r.IsComposition())

	tt.CheckFalse(r.reversedEdge)
	tt.CheckEqual("n0", r.FromID)
	tt.CheckEqual("n1", r.ToID)
	tt.CheckEqual(Zero, r.fromMin)
	tt.CheckEqual(Zero, r.toMin)
	tt.CheckEqual(One, r.fromMax)
	tt.CheckEqual(One, r.toMax)
}

func Test_Relationship_normalize_inherits(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewTerminatingIssueCollector()
	r := newRelationship("r1", "#INHERITS", "n0", "n1")
	validation.Do(func() {
		r.normalize(ic)
	})
	tt.CheckEqual(0, ic.Count()) // no warnings or errors
	tt.CheckFalse(r.isAssociation)
	tt.CheckTrue(r.isInherits)
	tt.CheckFalse(r.isComposition)
	tt.CheckFalse(r.reversedEdge)
	tt.CheckEqual("n0", r.FromID)
	tt.CheckEqual("n1", r.ToID)
	tt.CheckEqual(Zero, r.fromMin)
	tt.CheckEqual(Zero, r.toMin)
	tt.CheckEqual(One, r.fromMax)
	tt.CheckEqual(One, r.toMax)
	tt.CheckEqual("(n0)[0..1]--#INHERITS-->[0..1](n1)", r.String())
}

func Test_Relationship_normalize_composes_into(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewTerminatingIssueCollector()
	r := newRelationship("r1", "#COMPOSES_INTO", "n0", "n1")
	addRelProperty(r, "#to", "1M")
	validation.Do(func() {
		r.normalize(ic)
	})
	_, ok := r.Properties["#to"]
	tt.CheckFalse(ok)
	tt.CheckEqual(0, ic.Count()) // no warnings or errors
	tt.CheckFalse(r.isAssociation)
	tt.CheckFalse(r.isInherits)
	tt.CheckTrue(r.isComposition)
	tt.CheckTrue(r.reversedEdge)
	// assert node ids are reversed
	tt.CheckEqual("n1", r.FromID)
	tt.CheckEqual("n0", r.ToID)
	// assert multiplicity is reversed as well
	tt.CheckEqual(One, r.fromMin)
	tt.CheckEqual(Zero, r.toMin)
	tt.CheckEqual(Multiplicity(Many), r.fromMax)
	tt.CheckEqual(One, r.toMax)
	tt.CheckEqual("(n0)[0..1]--#COMPOSES_INTO-->[1..M](n1)", r.String())
}

func Test_Relationship_normalize_composes_of(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewTerminatingIssueCollector()
	r := newRelationship("r1", "#COMPOSED_OF", "n0", "n1")
	addRelProperty(r, "#to", "1M")
	validation.Do(func() {
		r.normalize(ic)
	})
	tt.CheckEqual(0, ic.Count()) // no warnings or errors
	tt.CheckFalse(r.isAssociation)
	tt.CheckFalse(r.isInherits)
	tt.CheckTrue(r.isComposition)
	tt.CheckFalse(r.reversedEdge)
	// assert node ids are not reversed
	tt.CheckEqual("n0", r.FromID)
	tt.CheckEqual("n1", r.ToID)
	// assert multiplicity is also not reversed
	tt.CheckEqual(Zero, r.fromMin)
	tt.CheckEqual(One, r.toMin)
	tt.CheckEqual(One, r.fromMax)
	tt.CheckEqual(Multiplicity(Many), r.toMax)
	tt.CheckEqual("(n0)[0..1]--#COMPOSED_OF-->[1..M](n1)", r.String())
}

// HELPERS.

func newRelationship(id string, relType string, from, to string) *Relationship {
	return &Relationship{
		ID:         id,
		Type:       relType,
		Style:      make(map[string]interface{}),
		Properties: make(map[string]string),
		FromID:     from,
		ToID:       to,
	}
}

func addRelProperty(r *Relationship, name, value string) {
	r.Properties[name] = value
}

// TODO:
//   * Errors - #UNKNOWN (tested in graph - wrong place)
//   * Warning INHERITS, IS_A, COMPOSED_OF, COMPOSES_INTO without #
//   * Errors on bad multiplicity
//   * Assert #from property is dropped
