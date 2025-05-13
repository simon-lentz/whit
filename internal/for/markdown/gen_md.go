package markdown

import (
	"fmt"
	"sort"
	"strings"

	"github.com/wyrth-io/whit/internal/pio"
	"github.com/wyrth-io/whit/internal/yammm"
)

// Marshal marshals/transforms the given context model to markdown text and outputs it to the given pio.Writer.
func Marshal(ctx yammm.Context, out *pio.Writer) { //nolint:gocognit
	if !ctx.IsCompleted() {
		return
	}
	model := ctx.Model()
	packageName := model.Name // TODO: this is too simplistic, may want different versions etc.
	out.FormatLn("# Package %s", packageName).
		FormatLn("%s\n\n", model.Documentation).
		FormatLn("```mermaid").
		FormatLn("classDiagram").
		FormatLn("    directionTB")

	// All types
	for _, t := range model.Types {
		marshalType(t, out)
	}

	out.Println("```")
	out.Println("")

	out.Println("## Types")
	types := model.Types
	sort.Slice(types, func(i, j int) bool {
		return types[i].Name < types[j].Name
	})

	for i := range types {
		t := types[i]
		needBlankLine := false
		out.Printf("### %s\n", t.Name)
		if t.IsAbstract {
			out.Println("* Abstract: true")
			needBlankLine = true
		}
		if t.IsPart {
			out.Println("* Part: true")
			needBlankLine = true
		}
		hyperlink := func(entries []string) string {
			var links []string
			for _, s := range entries {
				links = append(links, fmt.Sprintf("[%s](#%s)", s, strings.ToLower(s)))
			}
			return strings.Join(links, ", ")
		}
		if len(t.Inherits) > 0 {
			out.Printf("* Extends: %s\n", hyperlink(t.Inherits))
			needBlankLine = true
		}
		if t.Documentation != "" {
			if needBlankLine {
				out.Println("")
			}
			out.Println(t.Documentation)
			out.Println("")
			needBlankLine = false
		}
		propKind := func(p *yammm.Property) string {
			if p.IsPrimaryKey {
				return "primary"
			}
			if !p.Optional {
				return "required"
			}
			return ""
		}

		if len(t.Properties) > 0 {
			if needBlankLine {
				out.Println("")
			}
			out.Printf("| property | type | kind | description |\n")
			out.Printf("| -------- | ---- | ---- | ----------- |\n")
			for _, p := range t.Properties {
				out.Printf("| %s | %s | %s | %s |\n",
					p.Name,
					strings.ReplaceAll(p.DataTypeString(), "|", "&#124;"),
					propKind(p),
					htmlPara(p.Documentation),
				)
			}
			needBlankLine = true
		}
		if len(t.Associations) > 0 {
			if needBlankLine {
				out.Println("")
			}
			out.Println("#### Associations")
			for _, a := range t.Associations {
				hasProperties := len(a.Properties) > 0
				hasDocumentation := a.Documentation != ""
				mp := convertMultiplicity(a.Relationship)
				// * THE_NAME_OF_THE_ASSOC --> (one:many) OtherType
				out.Printf("\n* **%s** --> %s [%s](#%s)<br>\n", a.Name, mp, a.To, strings.ToLower(a.To))
				// Followed by the documentation as part of the bullet point
				if hasDocumentation {
					out.Printf("%s\n", a.Documentation)
				}
				// Followed by a property table indented 4 spaces to be part of the bullet point
				if hasProperties {
					out.Printf("    | property | type | kind | description |\n")
					out.Printf("    | -------- | ---- | ---- | ----------- |\n")
					for _, p := range a.Properties {
						out.Printf("    | %s | %s | %s | %s |\n",
							p.Name,
							strings.ReplaceAll(p.DataTypeString(), "|", "&#124;"),
							propKind(p),
							htmlPara(p.Documentation),
						)
					}
				}
			}
		}
		if len(t.Compositions) > 0 {
			out.Printf("#### Compositions\n")
			for _, c := range t.Compositions {
				mp := convertMultiplicity(c.Relationship)
				out.Printf("* **%s** &#9670;-> %s [%s](#%s)\n", c.Name, mp, c.To, strings.ToLower(c.To))
			}
		}
		if len(t.Invariants) > 0 {
			out.Printf("#### Invariants\n")
			for i := range t.Invariants {
				out.Printf("* %s\n", t.Invariants[i].Name)
			}
		}
	}
}

func htmlPara(s string) string {
	s = strings.ReplaceAll(s, "|", "&#124;")
	s = strings.ReplaceAll(s, "\n", "<br>")
	return fmt.Sprintf("<p>%s</p>", s)
}

func marshalType(t *yammm.Type, out *pio.Writer) {
	stereotypes := []string{}
	if t.IsAbstract {
		stereotypes = append(stereotypes, "Abstract")
	}
	if t.IsPart {
		stereotypes = append(stereotypes, "Part")
	}
	// Note: mermaid bombs if body is present without properties
	needsBody := len(t.Properties) > 0 || len(stereotypes) > 0
	out.Indentedf("class %s", t.Name)
	if needsBody {
		out.Indentedf("{\n")
	}
	indented := pio.IndentedWriter(out)

	if len(stereotypes) > 0 {
		indented.Indentedf("<<%s>>\n", strings.Join(stereotypes, ", "))
	}
	for i := range t.Properties {
		marshalProperty(t.Properties[i], indented)
	}

	if needsBody {
		out.Indentedf("}\n")
	} else {
		out.Println("")
	}
	for _, st := range t.Inherits {
		// Role --|> SuperRole : INHERITS
		out.Indentedf("%s --|> %s : %s\n", t.Name, st, "Extends")
	}
	for _, c := range t.Compositions {
		// Composition *-- Part : t (singular or plural)
		// *--
		out.Indentedf("%s *-- \"%s\" %s : %s\n", t.Name, convertToMermaidMultiplicity(c.Relationship), c.To, c.Name)
	}
	for _, a := range t.Associations {
		// properties on an association is singnalled by adding ... to the name
		name := a.Name
		if len(a.Properties) > 0 {
			name += " ..."
		}
		out.Indentedf("%s --> \"%s\" %s : %s\n", t.Name, convertToMermaidMultiplicity(a.Relationship), a.To, name)
	}
}
func marshalProperty(p *yammm.Property, out *pio.Writer) {
	kind := ""
	switch {
	case p.IsPrimaryKey:
		kind = " +"
	case p.Optional:
		kind = " ?"
	}
	// Note this does violence on UML as type and name is put in reverse order for readability.
	// Also note the type constraint can be quite long and should probably be truncated in tables.
	out.FormatLn("%s %s%s", p.Name, p.ShortDataTypeString(), kind)
}

func convertMultiplicity(a yammm.Relationship) (result string) {
	opt := a.Optional
	many := a.Many
	switch {
	case opt && many:
		result = "(many)"
	case !opt && many:
		result = "(one:many)"
	case opt && !many:
		result = "(one)"
	case !opt && !many:
		result = "(one:one)"
	}
	return
}
func convertToMermaidMultiplicity(a yammm.Relationship) (result string) {
	opt := a.Optional
	many := a.Many
	switch {
	case opt && many:
		result = "0..*"
	case !opt && many:
		result = "1..*"
	case opt && !many:
		result = "0..1"
	case !opt && !many:
		result = "1..1"
	}
	return
}
