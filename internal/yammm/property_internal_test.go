package yammm

import (
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/validation"
)

func Test_Property_validate(t *testing.T) {
	tt := testutils.NewTester(t)
	ctx := NewContext()
	ic := validation.NewIssueCollector()
	dt := &Property{Name: "x", DataType: []string{"Integer"}}

	dt.validate(ctx, ic)
	tt.CheckEqual(0, ic.Count())

	// must start with lower case name
	ic = validation.NewIssueCollector()
	dt = &Property{Name: "X", DataType: []string{"Integer"}}
	dt.validate(ctx, ic)
	tt.CheckEqual(1, ic.Count())
	ms := validation.MessageSet(ic)
	tt.CheckStringSlicesEqual([]string{"invalid property name 'X' - property name must start with lower case 'a'-'z'"}, ms.Slices())

	// name cannot be empty
	ic = validation.NewIssueCollector()
	dt = &Property{Name: "", DataType: []string{"Integer"}}
	dt.validate(ctx, ic)
	tt.CheckEqual(1, ic.Count())
	ms = validation.MessageSet(ic)
	tt.CheckStringSlicesEqual([]string{"a property with empty name having constraint 'Integer' - must have a name"}, ms.Slices())

	// cannot have error in constraint
	ic = validation.NewIssueCollector()
	dt = &Property{Name: "x", DataType: []string{":::"}}
	dt.validate(ctx, ic)
	tt.CheckEqual(1, ic.Count())
	ms = validation.MessageSet(ic)
	tt.CheckStringSlicesEqual([]string{"cannot compute type checker for property 'x' from its constraint '[:::]'"}, ms.Slices())

	ic = validation.NewIssueCollector()
	dt = &Property{Name: "xyz:234", DataType: []string{"String"}}
	dt.validate(ctx, ic)
	tt.CheckEqual(1, ic.Count())
	ms = validation.MessageSet(ic)
	tt.CheckStringSlicesEqual([]string{"invalid property name 'xyz:234' - contains illegal characters"}, ms.Slices())
}
