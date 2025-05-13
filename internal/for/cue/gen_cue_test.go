package cue_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/hlindberg/testutils"
	yammmcue "github.com/wyrth-io/whit/internal/for/cue"
	"github.com/wyrth-io/whit/internal/pio"
	"github.com/wyrth-io/whit/internal/validation"
	"github.com/wyrth-io/whit/internal/yammm"
)

func Test_Context_MarshalCue_sample(t *testing.T) {
	tt := testutils.NewTester(t)
	ctx := yammm.NewContext()
	reader := strings.NewReader(modelJSONBlob)
	err := ctx.SetModelFromJSON(reader)
	tt.CheckNotError(err)
	ic := validation.NewIssueCollector()
	ok := ctx.Complete(ic)
	tt.CheckTrue(ok)
	tt.CheckEqual(0, ic.Count())

	actual := new(bytes.Buffer)
	out := pio.WriterOn(actual)
	yammmcue.Marshal(ctx, out)

	spacedModelBlob := strings.ReplaceAll(modelCueBlob, "\t", "    ")

	actualSlices := strings.Split(actual.String(), "\n")
	expectedSlices := strings.Split(spacedModelBlob, "\n")
	tt.CheckStringSlicesEqual(expectedSlices, actualSlices)
	tt.CheckTextEqual(spacedModelBlob, actual.String())
}

func Test_Context_MarshalCue_basic_functionality(t *testing.T) {
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
		DataTypes: []*yammm.DataType{
			{Name: "example", Constraint: []string{"Integer", "0", ""}},
		},
	})
	tt.CheckNotError(err)
	ic := validation.NewIssueCollector()
	ok := ctx.Complete(ic)
	tt.CheckTrue(ok)
	tt.CheckEqual(0, ic.Count())

	actual := new(bytes.Buffer)
	out := pio.WriterOn(actual)
	yammmcue.Marshal(ctx, out)

	expected := new(bytes.Buffer)
	eout := pio.WriterOn(expected)
	eout.Printf("import (\n")
	eout.Printf(`    "strings"` + "\n")
	eout.Printf(`    "time"` + "\n")
	eout.Printf(`    "list"` + "\n")
	eout.Printf(")\n")
	eout.Printf("#Person: {\n")
	eout.Printf("    name: string\n")
	eout.Printf("}\n")
	eout.Printf("#Graph: {\n")
	eout.Printf("    People?: [...#Person]\n")
	eout.Printf("}\n")
	eout.Printf("#example: int & >= 0\n")
	eout.Printf("graph: #Graph\n")

	actualSlices := strings.Split(actual.String(), "\n")
	expectedSlices := strings.Split(expected.String(), "\n")
	tt.CheckStringSlicesEqual(expectedSlices, actualSlices)
	tt.CheckTextEqual(expected.String(), actual.String())
}

// note this result has tabs that will need to be expanded to 4 spaces.
var modelCueBlob = //
`import (
	"strings"
	"time"
	"list"
)
#Student: {
	#Person
}
#Person: {
	#Location
	name: string
	HAS_Head?: #Head // Composed 01 Part
	HAS_Limbs?: [...#Limb] // Composed 0M Part
	MOTHER_Person?: {
		Where: #REF_TO_PERSON
	}
	SIBLINGS_People?: [...#EDGE_SIBLINGS_Person]
}
#Head: {
	hasHair: bool
	id: string
}
#Limb: {
	type: string
	id: string
}
#Location: {
	long: string
	lat: string
}
#Graph: {
	Students?: [...#Student]
	People?: [...#Person]
}
#xdate: string & time.Format(time.RFC3339Date)
#EDGE_MOTHER_Person: {
	Where: #REF_TO_PERSON
}
#EDGE_SIBLINGS_Person: {
	since: string
	Where: #REF_TO_PERSON
}
#REF_TO_PERSON: {
	name: string
}
graph: #Graph
`

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
