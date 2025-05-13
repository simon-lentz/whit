package yammm

import (
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/validation"
)

func Test_Association_validate(t *testing.T) {
	tt := testutils.NewTester(t)
	mkAssociation := func() Association {
		return Association{
			Relationship: Relationship{
				To:   "Goat",
				Name: "OWNED_GOATS",
			},
			Properties: []*Property{{Name: "since", DataType: []string{"String"}}},
		}
	}
	// valid case
	a := mkAssociation()
	ctx := NewContext()
	ic := validation.NewIssueCollector()
	ok := a.validate(ctx, ic)
	tt.CheckTrue(ok)

	// invalid cases
	a = mkAssociation()
	a.Name = "owned_cows"
	ic = validation.NewIssueCollector()
	ok = a.validate(ctx, ic)
	tt.CheckFalse(ok)

	a = mkAssociation()
	a.To = "goats"
	ic = validation.NewIssueCollector()
	ok = a.validate(ctx, ic)
	tt.CheckFalse(ok)

	a = mkAssociation()
	a.Properties[0].Name = "SINCE"
	ic = validation.NewIssueCollector()
	ok = a.validate(ctx, ic)
	tt.CheckFalse(ok)
}

func Test_Composition_validate(t *testing.T) {
	tt := testutils.NewTester(t)
	ctx := NewContext()

	// valid composition
	c := Composition{
		Relationship: Relationship{To: "Goat", Name: "GOAT"},
	}
	ic := validation.NewIssueCollector()
	ok := c.validate(ctx, ic)
	tt.CheckTrue(ok)

	// invalid composition
	c = Composition{
		Relationship: Relationship{To: "goat"},
	}
	ic = validation.NewIssueCollector()
	ok = c.validate(ctx, ic)
	tt.CheckFalse(ok)
}
