package arrows

import (
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/validation"
)

// TESTS TODO:
//   * inheritance cannot have properties
//   * multiplicity is irrelevant on inheritance, will have default values

// DONE:
//   * inheritance cannot be duplicate
//   * circularity

func Test_Graph_inheritance_must_be_unique(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewTerminatingIssueCollector()

	emptyLabels := []string{}
	n1 := newNode("n0", "Car", emptyLabels)
	n2 := newNode("n1", "Registered", emptyLabels)
	n3 := newNode("n2", "Vehicle", emptyLabels)

	r1 := newRelation("r1", "#INHERITS", n1, n2)
	r2 := newRelation("r2", "#INHERITS", n1, n2)

	g := newGraph(
		[]*Node{n1, n2, n3},
		[]*Relationship{r1, r2},
	)
	g.prepareMeta(ic)
	tt.CheckEqual(1, ic.Count())

	m := utils.NewSet[string]()
	ic.EachIssue(func(issue validation.Issue) {
		m.Add(issue.Message())
	})
	tt.CheckEqual(1, m.Size())
	tt.CheckTrue(m.Contains("duplicate relationship 'INHERITS' between 'Car' and 'Registered'"))
}
func Test_Graph_inheritance_cannot_be_circular(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewTerminatingIssueCollector()

	emptyLabels := []string{}
	n1 := newNode("n1", "Car", emptyLabels)
	n2 := newNode("n2", "Registered", emptyLabels)
	n3 := newNode("n3", "Vehicle", emptyLabels)

	r1 := newRelation("r1", "#INHERITS", n1, n2)
	r2 := newRelation("r2", "#INHERITS", n1, n1)

	g := newGraph(
		[]*Node{n1, n2, n3},
		[]*Relationship{r1, r2},
	)
	g.prepareMeta(ic)
	tt.CheckEqual(1, ic.Count())

	m := utils.NewSet[string]()
	ic.EachIssue(func(issue validation.Issue) {
		m.Add(issue.Message())
	})
	tt.CheckEqual(1, m.Size())
	tt.CheckTrue(m.Contains("circular inheritance detected for node 'Car'"))
}
