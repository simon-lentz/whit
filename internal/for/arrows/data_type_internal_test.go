package arrows

import (
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/validation"
)

// TESTS TODO:
//   * check validity of resulting Cue expression (labels joined with &)

// DONE:
//   * cannot be named after a built in type
//   * Datatypes cannot inherit
//   * Datatypes are unique
//   * Datatypes cannot have properties (Could allow "#description")
//   * Datatypes cannot have relations
//   * must have lower case name

func Test_Graph_data_type_cannot_inherit(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewTerminatingIssueCollector()
	emptyLabels := []string{}
	n1 := newNode("n1", "something", []string{"#DataType", "string"})
	n2 := newNode("n2", "SuperSomething", emptyLabels)
	n3 := newNode("n3", "superDataType", []string{"#DataType", "string"})
	r1 := newRelation("r1", "#INHERITS", n1, n2)
	r2 := newRelation("r2", "#INHERITS", n2, n3)
	g := newGraph(
		[]*Node{n1, n2, n3},
		[]*Relationship{r1, r2},
	)
	g.prepareMeta(ic)
	tt.CheckEqual(2, ic.Count()) // assert two reported issues
	m := utils.NewSet[string]()
	ic.EachIssue(func(issue validation.Issue) {
		m.Add(issue.Message())
	})
	tt.CheckEqual(2, m.Size()) // assert not identical messages
	tt.CheckTrue(m.Contains("a #DataType cannot inherit - 'something' inherits from 'SuperSomething'"))
	tt.CheckTrue(m.Contains("a type cannot inherit from a #DataType - 'SuperSomething' inherits from 'superDataType'"))
}

func Test_Graph_data_types_must_be_unique(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewTerminatingIssueCollector()
	n1 := newNode("n1", "something", []string{"#DataType", "string"})
	n2 := newNode("n2", "something", []string{"#DataType", "string"})
	g := newGraph(
		[]*Node{n1, n2},
		[]*Relationship{},
	)
	validation.Do(func() {
		g.prepareMeta(ic)
	})
	m := utils.NewSet[string]()
	ic.EachIssue(func(issue validation.Issue) {
		m.Add(issue.Message())
	})

	tt.CheckEqual(2, m.Size()) // assert not identical messages
	tt.CheckTrue(m.Contains("node with id='n2' has the same caption as node='n1' - must have unique name"))
	tt.CheckTrue(m.Contains("earlier errors prevents further processing and production of output"))
}

func Test_Graph_data_types_cannot_have_properties(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewTerminatingIssueCollector()
	n1 := newNode("n1", "something", []string{"#DataType", "string"})
	addProperty(n1, "contraband", "snus")
	g := newGraph(
		[]*Node{n1},
		[]*Relationship{},
	)
	validation.Do(func() {
		g.prepareMeta(ic)
	})
	m := utils.NewSet[string]()
	ic.EachIssue(func(issue validation.Issue) {
		m.Add(issue.Message())
	})

	tt.CheckEqual(2, m.Size()) // assert not identical messages
	tt.CheckTrue(m.Contains("node: 'something' is a DataType and cannot have properties"))
	tt.CheckTrue(m.Contains("earlier errors prevents further processing and production of output"))
}

func Test_Graph_data_types_cannot_have_relations(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewTerminatingIssueCollector()
	n1 := newNode("n1", "something", []string{"#DataType", "string"})
	n2 := newNode("n2", "Thing", []string{})
	addProperty(n2, "id+", "string")
	r1 := newRelation("r1", "LIKES", n1, n2)
	g := newGraph(
		[]*Node{n1, n2},
		[]*Relationship{r1},
	)
	validation.Do(func() {
		g.prepareMeta(ic)
	})
	m := utils.NewSet[string]()
	ic.EachIssue(func(issue validation.Issue) {
		m.Add(issue.Message())
	})

	tt.CheckEqual(1, m.Size()) // assert not identical messages
	expected := "a #DataType cannot be in a relation - DataType 'something' has relation 'LIKES' to 'Thing'"
	tt.CheckTruef(m.Contains(expected), "Expected: '%s', was not in the set: %v", expected, messages(m))
}

func Test_Graph_data_types_cannot_have_name_of_built_in_type(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewTerminatingIssueCollector()
	n1 := newNode("n1", "string", []string{"#DataType", "string"})
	g := newGraph(
		[]*Node{n1},
		[]*Relationship{},
	)
	validation.Do(func() {
		g.prepareMeta(ic)
	})
	m := utils.NewSet[string]()
	ic.EachIssue(func(issue validation.Issue) {
		m.Add(issue.Message())
	})

	tt.CheckEqual(2, m.Size()) // assert not identical messages
	tt.CheckTrue(m.Contains("node with ID 'n1' for DataType 'string' - duplicates build in data type"))
	tt.CheckTrue(m.Contains("earlier errors prevents further processing and production of output"))
}

func Test_Graph_data_types_must_start_with_lower_case(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewTerminatingIssueCollector()
	n1 := newNode("n1", "Date", []string{"#DataType", "string"})
	g := newGraph(
		[]*Node{n1},
		[]*Relationship{},
	)
	validation.Do(func() {
		g.prepareMeta(ic)
	})
	m := utils.NewSet[string]()
	ic.EachIssue(func(issue validation.Issue) {
		m.Add(issue.Message())
	})

	tt.CheckEqual(2, m.Size()) // assert not identical messages
	tt.CheckTrue(m.Contains("node with ID 'n1' for DataType 'Date' - must start with lower case letter"))
	tt.CheckTrue(m.Contains("earlier errors prevents further processing and production of output"))
}
