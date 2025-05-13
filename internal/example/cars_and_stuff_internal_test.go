// @formatter:off
//nolint:all
// @formatter:on

package example

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"testing"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"

	"github.com/hlindberg/testutils"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/tada/catch"
	"github.com/wyrth-io/whit/internal/pio"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/yammm"
)

type Car struct {
	RegisteredVehicle `json:"RegisteredVehicle"`
}
type Person struct {
	Entity
	Birthday string `json:"birthday"`
	Name     string `json:"name"`
}
type REF_TO_Retirement struct { //nolint:all
	ID string `json:"id"`
}
type Registered struct {
	RegNbr              string `json:"regNbr"` // primary key
	RegistrationDate    string `json:"registrationDate"`
	DeregisterationDate string `json:"deregistrationDate"`
}
type Vehicle struct {
	Model string `json:"model"`
	Color string `json:"color"`
}
type RegisteredVehicle struct {
	Registered
	Vehicle
}
type Entity struct {
	MotorVehicleOwner *MotorVehicleOwner `json:"MotorVehicleOwner"`
}

type MotorVehicleOwner struct {
	OWNS_VEHICLE_RegisteredVehicle struct { //nolint:all
		FromDate string                    `json:"fromDate"`
		ToDate   string                    `json:"toDate"`
		Where    *REF_TO_RegisteredVehicle `json:"Where"`
	} `json:"OWNS_VEHICLE_RegisteredVehicle"`
}
type REF_TO_RegisteredVehicle struct { //nolint:all
	RegNbr string `json:"regNbr"`
}

// Overall Graph - a document.
type Graph struct {
	Cars   []*Car    `json:"Cars"`
	People []*Person `json:"People"`
}

func Test_Unmarshal(t *testing.T) {
	tt := testutils.NewTester(t)
	blob := `
	{
	"Cars": [
		{
			"regNbr": "ABC123",
			"registrationDate": "May 2017",
			"model": "Skoda"
		}
	],
	"People": [
		{
			"name": "Henrik",
			"birthday": "march 8",
			"MotorVehicleOwner": {
				"OWNS_VEHICLE_RegisteredVehicle": {
					"fromDate": "today",
					"Where": {
						"regNbr": "ABC123"
					}
				}
			}
		}
	]
    }`
	var g Graph
	if err := json.Unmarshal([]byte(blob), &g); err != nil {
		tt.Fatalf("could not unmarshal: %s", err.Error())
	}
	// TODO: Assertions
	fmt.Println(g)
}

type Type struct { // Done
	Singular     string
	Plural       string
	Properties   []Property
	Inherits     []string
	Mixins       []Type
	Compositions []Composition
	Associations []Association
	IsMixin      bool
	IsPart       bool
	IsAbstract   bool
}
type Relationship struct { // DONE
	ToSingular string
	ToPlural   string
	Optional   bool
	Many       bool
}
type Association struct { // DONE
	Relationship
	Name       string
	Properties []Property
	Where      []PrimaryKey
}
type Composition struct { // DONE
	Relationship
}
type PrimaryKey struct { // DONE
	Name string
	Type string
}
type Property struct { // DDONE
	Name     string
	Type     string
	Optional bool
}
type GraphT struct {
	Types     []Type
	DataTypes []yammm.DataType
}

// cueContext needs to be the same between calls to BaseDataType according to the cuelang documentation.
// This and `baseTypes` variable should be in some sort of context and not as package globals.
var cueContext *cue.Context

// baseTypes is a table of base type constraints which are initialized in BaseDataType function.
// This is done since it takes some time to compile the base time (not much, but would be bad to do this thousands of times).
var baseTypes map[string]cue.Value

// BaseDataType returns the name of the base data type as a string. The returned value is one of
// "string", "bool", "int", "float". If the  given cue type expression does not unify with any of
// these base types, an error is returned. The `typeString` argument should be a valid cue value constraint expression.
// An error is returned if it is not.
// In case of a non nil error, the returned type string is "".
//
// Note: For a constraint like >10, the result is "float" since there are numbers that are floats that are larger than 10.
// It is therefore always best to include the base type "int" if an int type is wanted.
func BaseDataType(typeString string) (string, error) {
	// set up cue context if not already initialized. TODO: change API to create a stateful object
	if cueContext == nil {
		cueContext = cuecontext.New()
	}
	// initialize the base type premises. TODO: change API to create a stateful object.
	if baseTypes == nil {
		baseTypes = map[string]cue.Value{
			"string": cueContext.CompileString("{t: string}"),
			"int":    cueContext.CompileString("{t: int}"),
			"float":  cueContext.CompileString("{t: number}"),
			"bool":   cueContext.CompileString("{t: bool}"),
		}
	}
	// Compile the user given data type as a constraint of the variable t.
	cueTypeString := cueContext.CompileString(fmt.Sprintf("t: %s", typeString))
	if err := cueTypeString.Err(); err != nil {
		return "", catch.Error("data type constraint '%s', is not a valid cue expression: %s", typeString, err)
	}
	// Check if the user given constraint is Subsumed
	for t, val := range baseTypes {
		if val.Subsume(cueTypeString) == nil {
			return t, nil
		}
	}
	return "", catch.Error("foobar")
}

// Creates a new samle graph containing examples of all model features.
func NewGraphT() *GraphT {
	return &GraphT{
		Types: []Type{
			{Singular: "Vehicle", Plural: "Vehicles"},
			{Singular: "Person", Plural: "People",
				Properties: []Property{
					{Name: "name", Type: "string"},
				},
				Associations: []Association{
					{
						Relationship: Relationship{ToSingular: "Car", ToPlural: "Cars", Optional: true},
						Name:         "LIKES",
						Where: []PrimaryKey{
							{Name: "regNbr", Type: "string"},
						},
					},
					{
						Relationship: Relationship{ToSingular: "Car", ToPlural: "Cars", Optional: true, Many: true},
						Name:         "OWNS",
						Where: []PrimaryKey{
							{Name: "regNbr", Type: "string"},
						},
					},
				},
			},
			{Singular: "Engine", Plural: "Engines", IsPart: true},
			{Singular: "Car", Plural: "Cars",
				Inherits: []string{"Vehicle"},
				Compositions: []Composition{
					{Relationship: Relationship{ToSingular: "Engine", ToPlural: "Engines"}},
				},
				Properties: []Property{
					{Name: "regNbr", Type: "string"},
					{Name: "model", Type: "string", Optional: true},
				},
			},
		},
	} // GraphT end
}

func ToGo(g *GraphT, pkgName string, w io.Writer) {
	out := pio.WriterOn(w)
	out.Printf("// package %s is a generated package.\n", pkgName)
	out.Printf("package %s\n", pkgName)
	out.Printf("//nolint\n")
	// Types
	for _, t := range g.Types {
		out.Printf("type %s struct {\n", t.Singular)
		for _, it := range t.Inherits {
			out.Printf("    %s\n", it)
		}
		for _, p := range t.Properties {
			out.Printf("    %s %s %s\n", p.Name, p.Type, jsontag(p.Name+utils.IfTrue(p.Optional, ",omitempty", ""))) // UC Go name, LC tag
		}
		// Mixins
		for _, m := range t.Mixins {
			out.Printf("    %s *%s %s\n", m.Singular, m.Singular, jsontag(m.Singular+",omitempty"))
		}
		for _, c := range t.Compositions {
			if c.Many {
				to := c.ToPlural
				out.Printf("    %s []*%s %s\n", to, c.ToSingular, jsontag(to+utils.IfTrue(c.Optional, ",omitempty", "")))
			} else {
				to := c.ToSingular
				out.Printf("    %s *%s %s\n", to, to, jsontag(to+utils.IfTrue(c.Optional, ",omitempty", "")))
			}
		}
		for _, a := range t.Associations {
			if a.Many {
				singular := fmt.Sprintf("%s_%s", a.Name, a.ToSingular)
				plural := fmt.Sprintf("%s_%s", a.Name, a.ToPlural)
				out.Printf("    %s []*EDGE_%s %s\n", plural, singular, jsontag(utils.IfTrue(a.Optional, plural+",omitempty", plural)))
			} else {
				x := fmt.Sprintf("%s_%s", a.Name, a.ToSingular)
				out.Printf("    %s *EDGE_%s %s\n", x, x, jsontag(utils.IfTrue(a.Optional, x+",omitempty", x)))
			}
		}
		// Associations
		out.Printf("}\n")
	}
	// Document
	out.Printf("type Graph struct {\n")
	for _, t := range utils.Filter(g.Types, func(t Type) bool {
		return !(t.IsAbstract || t.IsMixin || t.IsPart)
	}) {
		out.Printf("    %s []*%s %s\n", t.Plural, t.Singular, jsontag(t.Plural+",omitempty"))
	}
	out.Println("}")
	// EDGES
	for _, t := range g.Types {
		for _, a := range t.Associations {
			out.Printf("type EDGE_%s_%s struct {\n", a.Name, a.ToSingular)
			for _, p := range a.Properties {
				out.Printf("    %s %s %s\n", p.Name, p.Type, jsontag(p.Name+utils.IfTrue(p.Optional, ",omitempty", ""))) // UC Go name, LC tag
			}
			out.Printf("    Where struct {\n")
			for _, pk := range a.Where {
				out.Printf("        %s %s %s\n", pk.Name, pk.Type, jsontag(pk.Name)) // UC Go name, LC tag
			}
			out.Printf("    } %s\n", jsontag("Where"))
			out.Printf("}\n")
		}
	}
	// DATATYPES
}

func jsontag(s string) string {
	return fmt.Sprintf("`json:\"%s\"`", s)
}
func Test_ToGo(t *testing.T) {
	tt := testutils.NewTester(t)
	tt.CheckTrue(true)

	g := NewGraphT()
	w := new(bytes.Buffer)

	ToGo(g, "wyrth", w)

	expected := []string{
		`// package wyrth is a generated package.`,
		`package wyrth`,
		`//nolint`,
		`type Vehicle struct {`,
		`}`,
		`type Person struct {`,
		`    name string ` + jsontag("name"),
		`    LIKES_Car *EDGE_LIKES_Car ` + jsontag("LIKES_Car,omitempty"),
		`    OWNS_Cars []*EDGE_OWNS_Car ` + jsontag("OWNS_Cars,omitempty"),
		`}`,
		`type Engine struct {`,
		`}`,
		`type Car struct {`,
		`    Vehicle`,
		`    regNbr string ` + jsontag("regNbr"),
		`    model string ` + jsontag("model,omitempty"),
		`    Engine *Engine ` + jsontag("Engine"),
		`}`,
		`type Graph struct {`,
		`    Vehicles []*Vehicle ` + jsontag("Vehicles,omitempty"),
		`    People []*Person ` + jsontag("People,omitempty"),
		`    Cars []*Car ` + jsontag("Cars,omitempty"),
		`}`,
		`type EDGE_LIKES_Car struct {`,
		`    Where struct {`,
		`        regNbr string ` + jsontag("regNbr"),
		`    } ` + jsontag("Where"),
		`}`,
		`type EDGE_OWNS_Car struct {`,
		`    Where struct {`,
		`        regNbr string ` + jsontag("regNbr"),
		`    } ` + jsontag("Where"),
		`}`,
	}

	expectedString := strings.Join(expected, "\n")
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(expectedString, w.String(), false)
	fmt.Println(dmp.DiffPrettyText(diffs))

	actuaLinesOfText := strings.Split(w.String(), "\n")

	tt.CheckStringSlicesEqual(expected, actuaLinesOfText)
}

// TODO:
// * Create a package with the simplistic Wyrth meta-meta model "yammm" Yet Another Meta-Meta Model
// * Implement output of data type from yammm
// * Add the ability to specify "baseType" in Arrows & yamm. Allow only the base types there. For example `uint8` is not
//   a base type but a constraint on `int`

func Test_BaseDataType(t *testing.T) {
	tt := testutils.NewTester(t)

	x, err := BaseDataType(`>10`)
	tt.CheckEqualAndNoError(x, "float", err)

	x, err = BaseDataType(`"a"|"b"|"c"`)
	tt.CheckEqualAndNoError(x, "string", err)

	x, err = BaseDataType(`=~"[a-z]"`)
	tt.CheckEqualAndNoError(x, "string", err)

	x, err = BaseDataType(`bool`)
	tt.CheckEqualAndNoError(x, "bool", err)

	x, err = BaseDataType(`float`)
	tt.CheckEqualAndNoError(x, "float", err)

	x, err = BaseDataType(`3.1412`)
	tt.CheckEqualAndNoError(x, "float", err)
}
