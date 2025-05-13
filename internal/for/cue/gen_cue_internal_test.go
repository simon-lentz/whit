package cue

import (
	"bytes"
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/pio"
	"github.com/wyrth-io/whit/internal/validation"
	"github.com/wyrth-io/whit/internal/yammm"
)

func Test_Property_marshalCue(t *testing.T) {
	tt := testutils.NewTester(t)

	// required
	buf := &bytes.Buffer{}
	out := pio.IndentedWriter(pio.WriterOn(buf))
	dt := &yammm.Property{Name: "example", DataType: []string{"Integer", "0", ""}}
	marshalProperty(dt, out)
	tt.CheckEqual("    example: int & >= 0\n", buf.String())

	// optional
	buf = &bytes.Buffer{}
	out = pio.IndentedWriter(pio.WriterOn(buf))
	dt = &yammm.Property{Name: "example", DataType: []string{"Integer", "0", ""}, Optional: true}
	marshalProperty(dt, out)
	tt.CheckEqual("    example?: int & >= 0\n", buf.String())
}

func Test_Composition_marshalCue(t *testing.T) {
	tt := testutils.NewTester(t)
	buf := new(bytes.Buffer)
	out := pio.WriterOn(buf)
	ctx := yammm.NewContext()
	_ = ctx.SetMainModel(&yammm.Model{Name: "testmodel"})
	_, _ = ctx.AddType("Goat", []*yammm.Property{{Name: "name", DataType: []string{"String"}}})
	ic := validation.NewIssueCollector()
	ctx.Complete(ic)
	tt.CheckEqual(0, ic.Count())

	c := yammm.Composition{Relationship: yammm.Relationship{To: "Goat",
		Optional: false, Many: false, Name: "HAS_A"}}
	marshalComposition(ctx, &c, out)
	tt.CheckEqual("HAS_A_Goat: #Goat // Composed 11 Part\n", buf.String())

	// indents
	buf = new(bytes.Buffer)
	out = pio.IndentedWriter(pio.WriterOn(buf))
	marshalComposition(ctx, &c, out)
	tt.CheckEqual("    HAS_A_Goat: #Goat // Composed 11 Part\n", buf.String())

	// optional/many cases (not already covered)
	c = yammm.Composition{Relationship: yammm.Relationship{To: "Goat",
		Optional: true, Many: false, Name: "HAS_A"}}
	buf = new(bytes.Buffer)
	out = pio.IndentedWriter(pio.WriterOn(buf))
	marshalComposition(ctx, &c, out)
	tt.CheckEqual("    HAS_A_Goat?: #Goat // Composed 01 Part\n", buf.String())

	c = yammm.Composition{Relationship: yammm.Relationship{To: "Goat",
		Optional: true, Many: true, Name: "HAS_SEVERAL"}}
	buf = new(bytes.Buffer)
	out = pio.IndentedWriter(pio.WriterOn(buf))
	marshalComposition(ctx, &c, out)
	tt.CheckEqual("    HAS_SEVERAL_Goats?: [...#Goat] // Composed 0M Part\n", buf.String())

	c = yammm.Composition{Relationship: yammm.Relationship{To: "Goat",
		Optional: false, Many: true, Name: "HAS_SEVERAL"}}
	buf = new(bytes.Buffer)
	out = pio.IndentedWriter(pio.WriterOn(buf))
	marshalComposition(ctx, &c, out)
	tt.CheckEqual("    HAS_SEVERAL_Goats: [#Goat, ...#Goat] // Composed 1M Part\n", buf.String())
}

func Test_Association_marshalCue(t *testing.T) {
	tt := testutils.NewTester(t)
	ctx := yammm.NewContext()
	_ = ctx.SetMainModel(&yammm.Model{Name: "testmodel"})
	_, _ = ctx.AddType("Goat", []*yammm.Property{{Name: "name", DataType: []string{"String"}}})
	ic := validation.NewIssueCollector()
	ctx.Complete(ic)
	tt.CheckEqual(0, ic.Count())
	mkA := func() yammm.Association {
		return yammm.Association{
			Relationship: yammm.Relationship{
				Name:     "OWNED",
				To:       "Goat",
				Optional: false, Many: false},
		}
	}
	buf := new(bytes.Buffer)
	out := pio.WriterOn(buf)
	c := mkA()
	marshalAssociation(ctx, &c, out)
	tt.CheckEqual("OWNED_Goat: {\n    Where: #REF_TO_GOAT\n}\n", buf.String())

	// indents
	buf = new(bytes.Buffer)
	out = pio.IndentedWriter(pio.WriterOn(buf))
	marshalAssociation(ctx, &c, out)
	tt.CheckEqual("    OWNED_Goat: {\n        Where: #REF_TO_GOAT\n    }\n", buf.String())

	// optional/many cases (not already covered)
	c = mkA()
	c.Optional = true
	c.Many = false
	buf = new(bytes.Buffer)
	out = pio.IndentedWriter(pio.WriterOn(buf))
	marshalAssociation(ctx, &c, out)
	tt.CheckEqual("    OWNED_Goat?: {\n        Where: #REF_TO_GOAT\n    }\n", buf.String())

	c.Optional = true
	c.Many = true
	buf = new(bytes.Buffer)
	out = pio.IndentedWriter(pio.WriterOn(buf))
	marshalAssociation(ctx, &c, out)
	tt.CheckEqual("    OWNED_Goats?: [...#EDGE_OWNED_Goat]\n", buf.String())

	c.Optional = false
	c.Many = true
	buf = new(bytes.Buffer)
	out = pio.IndentedWriter(pio.WriterOn(buf))
	marshalAssociation(ctx, &c, out)
	tt.CheckEqual("    OWNED_Goats: [#EDGE_OWNED_Goat, ...#EDGE_OWNED_Goat]\n", buf.String())
}
