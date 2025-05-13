package yammm

import (
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/validation"
)

func Test_Type_invalid_names(t *testing.T) {
	tt := testutils.NewTester(t)
	ctx := NewContext()
	x := Type{
		Name:       "Ok",
		PluralName: "Oks",
	}
	ic := validation.NewIssueCollector()
	x.validate(ctx, ic)
	tt.CheckEqual(0, ic.Count())

	ic = validation.NewIssueCollector()
	x = Type{
		Name:       "N:tok",
		PluralName: "Oks",
	}
	x.validate(ctx, ic)
	tt.CheckEqual(1, ic.Count())

	ic = validation.NewIssueCollector()
	x = Type{
		Name:       "Ok",
		PluralName: "N:toks",
	}
	x.validate(ctx, ic)
	tt.CheckEqual(1, ic.Count())
}
