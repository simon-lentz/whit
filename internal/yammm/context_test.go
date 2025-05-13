package yammm_test

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/validation"
	"github.com/wyrth-io/whit/internal/yammm"
)

func Test_Context_New_a_context_can_be_created(t *testing.T) {
	tt := testutils.NewTester(t)
	defer testutils.ShouldNotPanic(t)
	ctx := yammm.NewContext()
	tt.CheckNotNil(ctx)
}

func Test_Context_SetModel(t *testing.T) {
	tt := testutils.NewTester(t)
	defer testutils.ShouldNotPanic(t)
	ctx := yammm.NewContext()
	// Ok to define an empty model
	err := ctx.SetMainModel(&yammm.Model{Name: "myfirstmodel"})
	tt.CheckNotError(err)

	// Error on redefine
	err = ctx.SetMainModel(&yammm.Model{Name: "myfirstmodel"})
	tt.CheckError(err)
}

func Test_Context_SetModelWithContent(t *testing.T) {
	tt := testutils.NewTester(t)
	defer testutils.ShouldNotPanic(t)
	ctx := yammm.NewContext()
	err := ctx.SetMainModel(&yammm.Model{
		Name: "myfirstmodel",
		Types: []*yammm.Type{
			{Name: "Person", PluralName: "People",
				Properties: []*yammm.Property{
					{Name: "name", DataType: []string{"String"}},
				}},
		},
	})
	tt.CheckNotError(err)
	ic := validation.NewIssueCollector()
	// Test a Complete to be able to lookup the type to verify the model is used.
	ok := ctx.Complete(ic)
	tt.CheckTrue(ok)
	tt.CheckEqual(0, ic.Count())

	person := ctx.LookupType("Person")
	tt.CheckNotNil(person)
}

func Test_a_model_can_be_roundtripped_via_JSON(t *testing.T) {
	tt := testutils.NewTester(t)
	model := yammm.Model{
		Name: "myfirstmodel",
		Types: []*yammm.Type{
			{Name: "Person", PluralName: "People",
				Properties: []*yammm.Property{
					{Name: "name", DataType: []string{"String"}},
				}},
		},
	}
	bytes, err := json.Marshal(&model)
	fmt.Printf("%s\n", bytes)
	tt.CheckNotError(err)
	// roundtripped := new(yammm.Model)
	var roundtripped yammm.Model
	_ = json.Unmarshal(bytes, &roundtripped)
	tt.CheckEqual("myfirstmodel", roundtripped.Name)
}

var modelJSONBlob = //
`{
	"name": "example",
	"types": [
		{
			"name": "Student",
			"plural_name": "Students",
			"inherits": [
				"Person"
			]
		},
		{
			"name": "Person",
			"plural_name": "People",
			"properties": [
				{
					"name": "name",
					"datatype": [
						"String"
					],
					"primary": true
				}
			],
			"inherits": [
				"Location"
			],
			"compositions": [
				{
					"name": "HAS",
					"to": "Head",
					"optional": true,
					"many": false
				},
				{
					"name": "HAS",
					"to": "Limb",
					"optional": true,
					"many": true
				}
			],
			"associations": [
				{
					"name": "MOTHER",
					"to": "Person",
					"optional": true,
					"many": false
				},
				{
					"name": "SIBLINGS",
					"to": "Person",
					"optional": true,
					"many": true,
					"properties": [
						{
							"name": "since",
							"datatype": [
								"String"
							]
						}
					]
				}
			]
		},
		{
			"name": "Head",
			"plural_name": "Heads",
			"properties": [
				{
					"name": "hasHair",
					"datatype": [
						"Boolean"
					]
				},
				{
					"name": "id",
					"datatype": [
						"String"
					],
					"primary": true
				}
			],
			"is_part": true
		},
		{
			"name": "Limb",
			"plural_name": "Limbs",
			"properties": [
				{
					"name": "type",
					"datatype": [
						"String"
					]
				},
				{
					"name": "id",
					"datatype": [
						"String"
					],
					"primary": true
				}
			],
			"is_part": true
		},
		{
			"name": "Location",
			"plural_name": "Locations",
			"properties": [
				{
					"name": "long",
					"datatype": [
						"String"
					]
				},
				{
					"name": "lat",
					"datatype": [
						"String"
					]
				}
			],
			"is_abstract": true
		}
	],
	"data_types": [
		{
			"name": "xdate",
			"constraint": [
				"Date"
			]
		}
	]
}`

func Test_Context_SetModelFromJSON(t *testing.T) {
	tt := testutils.NewTester(t)
	ctx := yammm.NewContext()
	reader := strings.NewReader(modelJSONBlob)
	err := ctx.SetModelFromJSON(reader)
	tt.CheckNotError(err)
	ic := validation.NewIssueCollector()
	ok := ctx.Complete(ic)
	if !ok {
		present := validation.NewColorPresentor()
		_ = present.Present(ic, validation.Info, os.Stdout)
	}
	tt.CheckTrue(ok)
	tt.CheckEqual(0, ic.Count())
	person := ctx.LookupType("Person")
	tt.CheckNotNil(person)
}

func Test_Context_AddType(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewIssueCollector()
	ctx := yammm.NewContext()
	model := yammm.Model{Name: "testmodel"}
	err := ctx.SetMainModel(&model)
	tt.CheckNotError(err)
	_, err = ctx.AddType("Person", []*yammm.Property{})
	tt.CheckNotError(err)
	ctx.Complete(ic)
	p := ctx.LookupType("Person")
	tt.CheckNotNil(p)
}
func Test_Context_AddDataType(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewIssueCollector()
	ctx := yammm.NewContext()
	model := yammm.Model{Name: "testmodel"}
	err := ctx.SetMainModel(&model)
	tt.CheckNotError(err)
	_, err = ctx.AddDataType("date", []string{"Date"})
	tt.CheckNotError(err)
	ctx.Complete(ic)
	dt := ctx.LookupDataType("date")
	tt.CheckNotNil(dt)
	tt.CheckEqual("date", dt.Name)
	tt.CheckEqual([]string{"Date"}, dt.Constraint)
}

func Test_Context_AddAbstractType(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewIssueCollector()
	ctx := yammm.NewContext()
	model := yammm.Model{Name: "testmodel"}
	err := ctx.SetMainModel(&model)
	tt.CheckNotError(err)
	_, err = ctx.AddAbstractType("Common", []*yammm.Property{})
	tt.CheckNotError(err)
	ctx.Complete(ic)
	p := ctx.LookupType("Common")
	tt.CheckNotNil(p)
	tt.CheckTrue(p.IsAbstract)
}
func Test_Context_AddCompositionPartType(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewIssueCollector()
	ctx := yammm.NewContext()
	model := yammm.Model{Name: "testmodel"}
	err := ctx.SetMainModel(&model)
	tt.CheckNotError(err)
	_, err = ctx.AddCompositionPartType("Common", []*yammm.Property{})
	tt.CheckNotError(err)
	ctx.Complete(ic)
	p := ctx.LookupType("Common")
	tt.CheckNotNil(p)
	tt.CheckTrue(p.IsPart)
}
func Test_Context_AddProperty(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewIssueCollector()
	ctx := yammm.NewContext()
	model := yammm.Model{Name: "testmodel"}
	err := ctx.SetMainModel(&model)
	tt.CheckNotError(err)
	_, err = ctx.AddType("Person", []*yammm.Property{})
	tt.CheckNotError(err)
	_, err = ctx.AddProperty("Person", "name", []string{"String"}, false, true)
	tt.CheckNotError(err)
	ctx.Complete(ic)
	p := ctx.LookupType("Person")
	tt.CheckNotNil(p)
	tt.CheckTrue(utils.Any(p.Properties, func(p *yammm.Property) bool {
		return p.Name == "name"
	}))
}

func Test_Context_AddAssociation(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewIssueCollector()
	ctx := yammm.NewContext()
	model := yammm.Model{Name: "testmodel"}
	err := ctx.SetMainModel(&model)
	tt.CheckNotError(err)
	_, err = ctx.AddType("GoatHerder", []*yammm.Property{})
	tt.CheckNotError(err)
	_, err = ctx.AddType("Goat", []*yammm.Property{{Name: "name", DataType: []string{"String"}, IsPrimaryKey: true}})
	tt.CheckNotError(err)
	_, err = ctx.AddAssociation("GoatHerder", "HERDS", "Goat", true, true, []*yammm.Property{}, "")
	tt.CheckNotError(err)
	ctx.Complete(ic)
	goatHearder := ctx.LookupType("GoatHerder")
	tt.CheckNotNil(goatHearder)
	isHERDSTheOnlyAssociation := utils.All(goatHearder.Associations, func(a *yammm.Association) bool {
		return a.Name == "HERDS"
	})
	tt.CheckTrue(isHERDSTheOnlyAssociation)
}

func Test_Context_AddComposition(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewIssueCollector()
	ctx := yammm.NewContext()
	model := yammm.Model{Name: "testmodel"}
	err := ctx.SetMainModel(&model)
	tt.CheckNotError(err)
	_, err = ctx.AddType("Building", []*yammm.Property{})
	tt.CheckNotError(err)
	_, err = ctx.AddCompositionPartType("Room", []*yammm.Property{{Name: "name", DataType: []string{"String"}, IsPrimaryKey: true}})
	tt.CheckNotError(err)
	_, err = ctx.AddComposition("Building", "HAS_ROOM", "Room", true, true, "")
	tt.CheckNotError(err)
	ctx.Complete(ic)
	building := ctx.LookupType("Building")
	tt.CheckNotNil(building)
	isRoomTheOnlyPart := utils.All(building.Compositions, func(c *yammm.Composition) bool {
		return c.To == "Room"
	})
	tt.CheckTrue(isRoomTheOnlyPart)
}

func Test_Context_AddInherits(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := validation.NewIssueCollector()
	ctx := yammm.NewContext()
	model := yammm.Model{Name: "testmodel"}
	err := ctx.SetMainModel(&model)
	tt.CheckNotError(err)
	_, err = ctx.AddType("GoatHerder", []*yammm.Property{})
	tt.CheckNotError(err)
	_, err = ctx.AddAbstractType("Herder", []*yammm.Property{{Name: "name", DataType: []string{"String"}, IsPrimaryKey: true}})
	tt.CheckNotError(err)
	err = ctx.AddInherits("GoatHerder", "Herder")
	tt.CheckNotError(err)
	ctx.Complete(ic)
	goatHearder := ctx.LookupType("GoatHerder")
	tt.CheckNotNil(goatHearder)
	isHerderInherited := utils.All(goatHearder.Inherits, func(name string) bool {
		return name == "Herder"
	})
	tt.CheckTrue(isHerderInherited)
}

// TODO:
// Add Abstract, Part
// Add relationships (Composition, Association)
// Add data type
// Here lint - have a cookie: .
