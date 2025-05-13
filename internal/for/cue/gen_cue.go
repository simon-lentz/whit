// Package cue contains logic to transform a yammm model to Cue.
package cue

import (
	"fmt"
	"sort"
	"strings"

	"github.com/wyrth-io/whit/internal/pio"
	"github.com/wyrth-io/whit/internal/tc"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/yammm"
)

// Marshal marshals the model in the given context in Cue format written to the given writer.
func Marshal(ctx yammm.Context, out *pio.Writer) {
	if !ctx.IsCompleted() {
		return
	}
	// import cue packages (ony strings, time and list for now)
	out.Printf("import (\n")
	indented := pio.IndentedWriter(out)
	indented.Indentedf(`"strings"` + "\n")
	indented.Indentedf(`"time"` + "\n")
	indented.Indentedf(`"list"` + "\n")
	out.Printf(")\n")

	marshalModel(ctx, out)

	// All REF_TO_ and EDGE_
	typesRequiringRef := utils.NewSet[string]()
	for _, t := range ctx.Model().Types {
		for _, a := range t.Associations {
			marshalAssociationCueRefs(a, out)
			typesRequiringRef.Add(a.To)
		}
	}
	// Sort the slices since a set is unordered and if output directly will cause diffs between generated schemas
	// when there are none.
	types := typesRequiringRef.Slices()
	sort.Slice(types, func(i, j int) bool {
		return types[i] < types[j]
	})
	for _, name := range types {
		t := ctx.LookupType(name)
		marshalTypeCueRef(t, out)
	}
	// Schema must have constraint on document's contents
	out.Printf("graph: #Graph\n")
}

func marshalModel(ctx yammm.Context, out *pio.Writer) {
	model := ctx.Model()
	// All types
	for _, t := range model.Types {
		marshalType(ctx, t, out)
	}
	marshalModelGraph(ctx, out)
	// All data types
	for _, dt := range model.DataTypes {
		marshalDataType(dt, out)
	}
}

// marshalModelGraph marshals the document definition to the out writer.
func marshalModelGraph(ctx yammm.Context, out *pio.Writer) {
	m := ctx.Model()
	out.Indentedf("#Graph: {\n")
	indented := pio.IndentedWriter(out)
	for _, t := range utils.Filter(m.Types, func(t *yammm.Type) bool {
		return !(t.IsAbstract || t.IsPart)
	}) {
		indented.Indentedf("%s?: [...#%s]\n", t.PluralName, t.Name)
	}
	out.Indentedf("}\n")
}

// marshalDataType produces a cue representation of the data type in the given writer.
func marshalDataType(dt *yammm.DataType, out *pio.Writer) {
	out.Printf("#%s: %s\n", dt.Name, ConstraintAsCue(dt.Constraint))
}

// ConstraintAsCue returns a cue language string being the cue representation of a yammm type constraint.
func ConstraintAsCue(constraint []string) string {
	if len(constraint) < 1 {
		return tc.LstringS // nothing set, default to string
	}
	c := constraint
	switch c[0] {
	case tc.IntegerS:
		min := ""
		max := ""
		switch len(c) {
		case 1:
		case 2:
			min = c[1]
		case 3:
			min = c[1]
			max = c[2]
		}
		if max == "" {
			if min == "" {
				return tc.LintS
			}
			return "int & >= " + min
		}
		if min == "" {
			return "int & <= " + max
		}
		return "int & >= " + min + " && " + "<= " + max

	case tc.FloatS:
		min := ""
		max := ""
		switch len(c) {
		case 1:
		case 2:
			min = c[1]
		case 3:
			min = c[1]
			max = c[2]
		}
		if max == "" {
			if min == "" {
				return tc.LfloatS
			}
			return "float & >= " + min
		}
		if min == "" {
			return "float & <= " + max
		}
		return "float & >= " + min + " && " + "<= " + max
	case tc.StringS:
		min := ""
		max := ""
		switch len(c) {
		case 1:
		case 2:
			min = c[1]
		case 3:
			min = c[1]
			max = c[2]
		}
		if max == "" {
			if min == "" {
				return tc.LstringS
			}
			return fmt.Sprintf("string & strings.MinRunes(%s)", min)
		}
		if min == "" {
			return fmt.Sprintf("string & strings.MaxRunes(%s)", max)
		}
		return fmt.Sprintf("string & strings.MinRunes(%s) & strings.MaxRunes(%s)", min, max)
	case tc.BooleanS:
		return "bool"
	case tc.DateS:
		return "string & time.Format(time.RFC3339Date)"
	case tc.TimestampS:
		if len(c) == 1 {
			return "string & time.Format(time.RFC3339)"
		}
		return fmt.Sprintf("string & time.Format(%s)", c[1])
	case tc.EnumS:
		quoted := utils.Map(c[1:], func(s string) string { return fmt.Sprintf("%q", s) })
		return fmt.Sprintf("string & (%s)", strings.Join(quoted, " | "))
	case "Pattern":
		quoted := utils.Map(c[1:], func(s string) string { return fmt.Sprintf("=~ %q", s) })
		return fmt.Sprintf("string & (%s)", strings.Join(quoted, " | "))
	case tc.AliasS:
		// TODO: Not sure about this
		return fmt.Sprintf("#%s", c[1])
	case tc.Spacevector:
		return fmt.Sprintf("list.Repeat([ float32 ], %s)", c[1])
	}
	return tc.LstringS
}

func marshalProperty(p *yammm.Property, out *pio.Writer) {
	var oS string
	if p.Optional {
		oS = "?"
	}
	out.Indent()
	out.Printf("%s%s: %s\n", p.Name, oS, ConstraintAsCue(p.DataType))
}
func marshalAssociation(ctx yammm.Context, a *yammm.Association, out *pio.Writer) {
	// Level 1 - instance
	// SIBLING_People:
	//    - type: "sister"
	//      Where:
	//          name: "Jane Doe"
	//    - type: "brother"
	//      Where:
	//          name: "Dumb Doe"
	// // Level 2 - meta
	// #Person: {
	//     name+: string
	//     SIBLING_People?: [...#EDGE_SIBLING_Person]
	// }
	// #EDGE_SIBLING_Person: {
	//     type: "sister"|"brother"|"..."
	//     Where: #REF_TO_PERSON
	// }

	// SIBLING_People, MOTHER_Person
	toType := ctx.LookupType(a.To)
	out.Indentedf("%s_%s%s: ", a.Name, utils.IfTrue(a.Many, toType.PluralName, a.To), utils.IfTrue(a.Optional, "?", ""))
	if a.Many || len(a.Properties) > 0 {
		if a.Optional {
			out.Printf("[...#EDGE_%s_%s]\n", a.Name, a.To)
		} else {
			out.Printf("[#EDGE_%[1]s_%[2]s, ...#EDGE_%[1]s_%[2]s]\n", a.Name, a.To)
		}
	} else {
		// max 1 and no properties
		out.Printf("{\n")
		pio.IndentedWriter(out).Indentedf("Where: #REF_TO_%s\n", strings.ToUpper(utils.CamelToSnake(a.To)))
		out.Indentedf("}\n")
	}
}
func marshalAssociationCueRefs(a *yammm.Association, out *pio.Writer) {
	out.Indentedf("#EDGE_%s_%s: {\n", a.Name, a.To)
	indented := pio.IndentedWriter(out)
	for _, p := range a.Properties {
		marshalProperty(p, indented)
	}
	indented.Indentedf("Where: #REF_TO_%s\n", strings.ToUpper(utils.CamelToSnake(a.To)))
	out.Indent()
	out.Printf("}\n")
}

// marshalComposition creates the upper case property being one or more composed objects.
// A composition will get a property named after the composition and
// type in Name or plural depending on the multiplicity:
//
//	HAS_Engine?: #Engine
//	HAS_Engine: #Engine
//	HAS_Engines?: [...#Engine]
//	HAS_Engines: [#Engine, ...#Engine]
func marshalComposition(ctx yammm.Context, c *yammm.Composition, out *pio.Writer) {
	out.Indent()
	propName := c.PropertyName(ctx)
	switch {
	case c.Optional && !c.Many: // "01"
		out.Printf("%s?: #%s // Composed 01 Part\n", propName, c.To)
	case !c.Optional && !c.Many: // "11"
		out.Printf("%s: #%s // Composed 11 Part\n", propName, c.To)
	case c.Optional && c.Many: // "0M"
		out.Printf("%s?: [...#%s] // Composed 0M Part\n", propName, c.To)
	case !c.Optional && c.Many: // "1M":
		out.Printf("%s: [#%s, ...#%s] // Composed 1M Part\n", propName, c.To, c.To)
	}
}

func marshalType(ctx yammm.Context, t *yammm.Type, out *pio.Writer) {
	out.Indentedf("#%s: {\n", t.Name)
	indented := pio.IndentedWriter(out)
	for _, st := range t.Inherits {
		indented.Indentedf("#%s\n", st)
	}
	for _, p := range t.Properties {
		marshalProperty(p, indented)
	}
	for _, c := range t.Compositions {
		marshalComposition(ctx, c, indented)
	}
	for _, a := range t.Associations {
		marshalAssociation(ctx, a, indented)
	}

	out.Indentedf("}\n")
}

func marshalTypeCueRef(t *yammm.Type, out *pio.Writer) {
	snakedName := strings.ToUpper(utils.CamelToSnake(t.Name))

	out.Indentedf("#REF_TO_%s: {\n", snakedName)
	indented := pio.IndentedWriter(out)
	pk := make([]*yammm.Property, len(t.AllPrimaryKeys()))
	copy(pk, t.AllPrimaryKeys())
	sort.Slice(pk, func(i, j int) bool {
		return pk[i].Name < pk[j].Name
	})
	for _, p := range pk {
		indented.Indentedf("%s: %s\n", p.Name, ConstraintAsCue(p.DataType))
	}
	out.Indentedf("}\n")
}
