package yammm_test

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/jzon"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/validation"
	"github.com/wyrth-io/whit/internal/yammm"
	"github.com/wyrth-io/whit/parser"
)

func Test_Validator_errors_on_non_completed_context(t *testing.T) {
	tt := testutils.NewTester(t)

	// Need a schema and an instance
	ctx := yammm.NewContext()
	graph := map[string]any{}
	v := yammm.NewValidator(ctx, t.Name(), graph)
	tt.CheckNotNil(v)

	ic := validation.NewIssueCollector()
	result := v.Validate(ic)
	tt.CheckFalse(result)
	tt.CheckEqual(1, ic.Count())
	issues := []validation.Issue{}
	ic.EachIssue(func(issue validation.Issue) { issues = append(issues, issue) })
	tt.CheckMatches("Cannot validate instance: no model defined", issues[0].Message())

	ic = validation.NewIssueCollector()
	issues = []validation.Issue{}
	model := yammm.NewModel("testmodel")
	model.Source = t.Name()
	model.Line = 1
	model.Column = 2
	_ = ctx.SetMainModel(model)
	result = v.Validate(ic)
	tt.CheckFalse(result)
	ic.EachIssue(func(issue validation.Issue) { issues = append(issues, issue) })
	// The source (name of this test) plus line and column are included in the error message if set.
	tt.CheckMatches(fmt.Sprintf(`\[%s:1:2\] Cannot validate instance: model not completed`, t.Name()), issues[0].Message())

	ic = validation.NewIssueCollector()
	ok := ctx.Complete(ic)
	tt.CheckTrue(ok)
}

func Test_Validator_is_ok_with_empty_model_but_warns(t *testing.T) {
	tt := testutils.NewTester(t)

	// Need a schema and an instance
	ctx := yammm.NewContext()
	_ = ctx.SetMainModel(yammm.NewModel("testmodel"))
	ic := validation.NewIssueCollector()
	ok := ctx.Complete(ic)
	tt.CheckTrue(ok)

	graph := map[string]any{}
	v := yammm.NewValidator(ctx, t.Name(), graph)

	ic = validation.NewIssueCollector()
	result := v.Validate(ic)
	tt.CheckFalse(result)
	tt.CheckEqual(1, ic.Count())
	tt.CheckTrue(ic.HasWarnings())
	ic.EachIssueAtLevel(validation.Warning, func(issue validation.Issue) {
		tt.CheckMatches("The graph top level is empty", issue.Message())
	})
}

func Test_Validator_UnknownType(t *testing.T) {
	model := `
	schema "testing"
	type Car { regNbr String primary}
	`
	instance := `
	{
		"Cars": [
			{ "regNbr": "ABC123" }	
		],
		"Contrabands": { "xxx": 42 }
	}`
	validateMessages(t, model, instance,
		"[Test_Validator_UnknownType:6:3] The name 'Contrabands' is not the plural name of a type in the model",
	)
}
func Test_Validator_MissingPrimary(t *testing.T) {
	model := `schema "testing" type Car { regNbr String primary}`
	instance := `{ "Cars": [ {  } ] }`
	validateMessages(t, model, instance,
		"[Test_Validator_MissingPrimary:1:13] Property value of 'Car.regNbr' is required and is missing",
	)
}
func Test_Validator_IdSetWhenThereAreOtherPks(t *testing.T) {
	model := `schema "testing"
	type Person { name String primary }	`
	instance := `{ "People": [{ "id": "$$illegal", "name": "Henrik" } ] }`
	validateMessages(t, model, instance,
		"[Test_Validator_IdSetWhenThereAreOtherPks:1:14] Property value of 'Person.id' should not be set: not allowed when there are other primary keys", //nolint:lll
	)
	instance = `{ "People": [{ "id": "$$illegal" } ] }`
	validateMessages(t, model, instance,
		"[Test_Validator_IdSetWhenThereAreOtherPks:1:14] Property value of 'Person.id' should not be set: not allowed when there are other primary keys", //nolint:lll
	)
}
func Test_Validator_WrongDataType(t *testing.T) {
	model := `schema "testing" type Car { regNbr String primary}`
	instance := `{ "Cars": [ { "regNbr": 42 } ] }`
	validateMessages(t, model, instance,
		"[Test_Validator_WrongDataType:1:25] Property value of 'Car.regNbr' must have string base type: got int",
	)
}
func Test_Validator_DataTypeConstraints(t *testing.T) {
	model := `schema "testing" type Subject { sid Integer[0,1] primary}`
	instance := `{ "Subjects": [ { "sid": 0 } ] }`
	validateMessages(t, model, instance) // expect no messages

	instance = `{ "Subjects": [ { "sid": -1 } ] }`
	validateMessages(t, model, instance,
		"[Test_Validator_DataTypeConstraints:1:26] Integer value -1 is < min 0",
	)

	instance = `{ "Subjects": [ { "sid": 2 } ] }`
	validateMessages(t, model, instance,
		"[Test_Validator_DataTypeConstraints:1:26] Integer value 2 is > max 1",
	)

	model = `schema "testing" type Subject { sid Float[0,1] primary}`
	instance = `{ "Subjects": [ { "sid": 0.0 } ] }`
	validateMessages(t, model, instance) // expect no messages

	model = `schema "testing" type Subject { sid Float[0,1] primary}`
	instance = `{ "Subjects": [ { "sid": 1.1 } ] }`
	validateMessages(t, model, instance,
		"[Test_Validator_DataTypeConstraints:1:26] Float value 1.1 is > max 1",
	)
	instance = `{ "Subjects": [ { "sid": -1.0 } ] }`
	validateMessages(t, model, instance,
		"[Test_Validator_DataTypeConstraints:1:26] Float value -1 is < min 0",
	)
	model = `schema "testing" type Subject { sid Enum["a","b"] primary}`
	instance = `{ "Subjects": [ { "sid": "a" } ] }`
	validateMessages(t, model, instance) // expect no messages
	instance = `{ "Subjects": [ { "sid": "c" } ] }`
	validateMessages(t, model, instance,
		`[Test_Validator_DataTypeConstraints:1:26] String value 'c' does not match Enum["a", "b"]`,
	)

	model = `schema "testing" type Subject { sid Pattern["a.","b."] primary}`
	instance = `{ "Subjects": [ { "sid": "ax" } ] }`
	validateMessages(t, model, instance) // expect no messages
	instance = `{ "Subjects": [ { "sid": "ca" } ] }`
	validateMessages(t, model, instance,
		`[Test_Validator_DataTypeConstraints:1:26] String value 'ca' does not match Pattern["a.", "b."]`,
	)

	model = `schema "testing" type Subject { sid Boolean primary}`
	instance = `{ "Subjects": [ { "sid": true } ] }`
	validateMessages(t, model, instance) // expect no messages
	instance = `{ "Subjects": [ { "sid": "yes" } ] }`
	validateMessages(t, model, instance,
		`[Test_Validator_DataTypeConstraints:1:26] Property value of 'Subject.sid' must have bool base type: got string`,
	)

	model = `schema "testing" type Subject { sid Date primary}`
	instance = `{ "Subjects": [ { "sid": "2023-06-21" } ] }`
	validateMessages(t, model, instance) // expect no messages
	instance = `{ "Subjects": [ { "sid": "rosebud" } ] }`
	validateMessages(t, model, instance,
		`[Test_Validator_DataTypeConstraints:1:26] value 'rosebud' does not match Date format '2006-01-02'`,
	)
}
func Test_ValidatorTimestampProperty(t *testing.T) {
	model := `schema "testing" type Subject { it Timestamp primary}`
	instance := `{ "Subjects": [ { "it": "2023-06-21T01:02:50Z" } ] }`
	validateMessages(t, model, instance) // expect no messages
	instance = `{ "Subjects": [ { "it": "rosebud" } ] }`
	validateMessages(t, model, instance,
		`[Test_ValidatorTimestampProperty:1:25] value does not match Timestamp format `+
			`: parsing time "rosebud" as "2006-01-02T15:04:05Z07:00": cannot parse "rosebud" as "2006"`,
	)
	model = `schema "testing" type Subject { it Timestamp["15:04:05"] primary}`
	instance = `{ "Subjects": [ { "it": "01:02:50" } ] }`
	validateMessages(t, model, instance) // expect no messages
	instance = `{ "Subjects": [ { "it": "30:61:61" } ] }`
	validateMessages(t, model, instance,
		`[Test_ValidatorTimestampProperty:1:25] value does not match Timestamp format : parsing time "30:61:61": hour out of range`,
	)
}

func Test_Validator_AbstractAsConcrete(t *testing.T) {
	model := `schema "testing" abstract type Thing { it String primary} `
	instance := `{ "Things": [ { "it": "ABC123" } ] }`
	validateMessages(t, model, instance,
		"[Test_Validator_AbstractAsConcrete:1:3] The type 'Thing' is not a concrete type in the model",
	)
}
func Test_Validator_UnknownProperty(t *testing.T) {
	model := `schema "testing" type Car { regNbr String primary}`
	instance := `{ "Cars": [
		{ "regNbr": "ABC123", "notRegNbr": "str", "anotherBad": 123 }
		]
	}`
	validateMessages(t, model, instance,
		"[Test_Validator_UnknownProperty:2:3] Object of type 'Car' has excess properties: [anotherBad notRegNbr]",
	)
}

func Test_Validator_NotListOfInstances(t *testing.T) {
	model := `schema "testing" type Car { regNbr String primary}`
	instance := `{ "Cars": {  } }`
	validateMessages(t, model, instance,
		"[Test_Validator_NotListOfInstances:1:3] The value of 'Cars' is not a list of instances",
	)
}

func Test_Validator_RequiredCompositionMissing(t *testing.T) {
	model := `schema "testing"
	type Person { name String primary *-> HAS (one:one) Body }
	part type Body { it String }
	`
	instance := `{ "People": [{ "name": "Henrik" } ] }`
	validateMessages(t, model, instance,
		"[Test_Validator_RequiredCompositionMissing:1:14] Composition 'Person.HAS_Body' is required and is missing",
	)
}
func Test_Validator_CompositionWithManyNotAList(t *testing.T) {
	model := `schema "testing"
	type Person { name String primary *-> HAS (one:many) Body }
	part type Body { it String }
	`
	instance := `{ "People": [{ "name": "Henrik", "HAS_Bodies": { "id": "123" } } ] }`
	validateMessages(t, model, instance,
		"[Test_Validator_CompositionWithManyNotAList:1:48] Composition 'Person.HAS_Body' must be a list since relation is to 'many'",
	)
}
func Test_Validator_CompositionWithOneInList(t *testing.T) {
	model := `schema "testing"
	type Person { name String primary *-> HAS (one:one) Body }
	part type Body { it String }
	`
	instance := `{ "People": [{ "name": "Henrik", "HAS_Body": [{ "it": "123" }] } ] }`
	validateMessages(t, model, instance,
		"[Test_Validator_CompositionWithOneInList:1:46] Composition 'Person.HAS_Body' cannot be a list since relation is to 'one'",
	)
}
func Test_Validator_CompositionNestedObjectIsValidated(t *testing.T) {
	model := `schema "testing"
	type Person { name String primary *-> HAS (one:one) Body }
	part type Body { it String primary }
	`
	instance := `{ "People": [{ "name": "Henrik", "HAS_Body": {  } } ] }`
	validateMessages(t, model, instance,
		"[Test_Validator_CompositionNestedObjectIsValidated:1:46] Property value of 'Body.it' is required and is missing",
	)
}
func Test_Validator_RequiredAssociationMissing(t *testing.T) {
	model := `schema "testing"
	type Person { name String primary --> LIKES (one:one) Food }
	type Food { it String primary }
	`
	instance := `{ "People": [{ "name": "Henrik" } ] }`
	validateMessages(t, model, instance,
		"[Test_Validator_RequiredAssociationMissing:1:14] Association 'Person.LIKES_Food' is required and is missing",
	)
}

func Test_Validator_AssociationPropertyMissing(t *testing.T) {
	model := `schema "testing"
	type Person { name String primary
		--> LIKES (one:one) Food { since String required }
	}
	type Food { it String primary }
	`
	instance := `{
		"People": [
		{ 	"name": "Henrik", 
			"LIKES_Food": {
				"Where": {
					"it": "banana"
				}
			}
		}
	]}`
	validateMessages(t, model, instance,
		"[Test_Validator_AssociationPropertyMissing:4:18] Association 'Person.LIKES_Food' is missing required properties [since]",
	)
}
func Test_Validator_AssociationPropertyWrongType(t *testing.T) {
	model := `schema "testing"
	type Person { name String primary
		--> LIKES (one:one) Food { since String required }
	}
	type Food { it String primary }
	`
	instance := `{
		"People": [
		{ 	"name": "Henrik", 
			"LIKES_Food": {
				"since": 42,
				"Where": {
					"it": "banana"
				}
			}
		}
	]}`
	validateMessages(t, model, instance,
		"[Test_Validator_AssociationPropertyWrongType:5:5] Property value of 'Person.LIKES_Food.since' must have string base type: got int",
	)
}
func Test_Validator_AssociationExcessProperty(t *testing.T) {
	model := `schema "testing"
	type Person { name String primary
		--> LIKES (one:one) Food { since String required }
	}
	type Food { it String primary }
	`
	instance := `{
		"People": [
		{ 	"name": "Henrik", 
			"LIKES_Food": {
				"since": "forever",
				"extra": "should be reported",
				"Where": {
					"it": "banana"
				}
			}
		}
	]}`
	validateMessages(t, model, instance,
		"[Test_Validator_AssociationExcessProperty:4:18] Association 'Person.LIKES_Food' has excess properties [extra]",
	)
}

func Test_Validator_AssociationExcessPrimaryKey(t *testing.T) {
	model := `schema "testing"
	type Person { name String primary
		--> LIKES (one:one) Food { since String required }
	}
	type Food { it String primary }
	`
	instance := `{
		"People": [
		{ 	"name": "Henrik", 
			"LIKES_Food": {
				"since": "forever",
				"Where": {
					"it": "banana",
					"extra": "should be reported"
				}
			}
		}
	]}`
	validateMessages(t, model, instance,
		"[Test_Validator_AssociationExcessPrimaryKey:6:14] Association 'Person.LIKES_Food.Where' has excess primary key [extra]",
	)
}

func Test_Validator_AssociationPrimaryKeyMissing(t *testing.T) {
	model := `schema "testing"
	type Person { name String primary
		--> LIKES (one:one) Food { since String required }
	}
	type Food { it String primary }
	`
	instance := `{
		"People": [
		{ 	"name": "Henrik", 
			"LIKES_Food": {
				"since": "forever",
				"Where": {
				}
			}
		}
	]}`
	validateMessages(t, model, instance,
		"[Test_Validator_AssociationPrimaryKeyMissing:6:14] Association 'Person.LIKES_Food' is missing primary key 'id' or all of [it]",
	)
}
func Test_Validator_AssociationIdOrKeysAccepted(t *testing.T) {
	model := `schema "testing"
	type Person { name String primary
		--> LIKES (one:one) Food { since String required }
	}
	type Food { it String primary  it2 String primary}
	`
	instance := `{
		"People": [
		{ 	"name": "Henrik", 
			"LIKES_Food": {
				"since": "forever",
				"Where": {
					"it": "banana",
					"it2": "vulgaris"
				}
			}
		}
	]}`
	// ok since 'it' and 'it2' are set.
	validateMessages(t, model, instance)

	instance = `{
		"People": [
		{ 	"name": "Henrik", 
			"LIKES_Food": {
				"since": "forever",
				"Where": {
					"id": "$$local"
				}
			}
		}
	]}`
	// ok since 'id' is set.
	validateMessages(t, model, instance)
}
func Test_Validator_AliasDataType(t *testing.T) {
	model := `schema "testing" type Subject { it Color primary} type Color = Enum["red", "green"]`
	instance := `{ "Subjects": [ { "it": "red" } ] }`
	validateMessages(t, model, instance) // expect no messages

	instance = `{ "Subjects": [ { "it": "blue" } ] }`
	validateMessages(t, model, instance,
		`[Test_Validator_AliasDataType:1:25] String value 'blue' does not match Enum["green", "red"]`,
	)
}

// validatedMessages validates the instance against an assummed correct model.
func validateMessages(t *testing.T, model, instance string, expected ...string) {
	t.Helper()
	tt := testutils.NewTester(t)
	// ignore ic from model validation, asserted to not have errors
	ctx, _ := makeContext(t.Name(), model, true)
	inst := makeJzonInstance(t.Name(), instance)
	// v := yammm.NewValidator(ctx, t.Name(), makeInstance(instance))
	v := yammm.NewValidator(ctx, t.Name(), inst)
	ic := validation.NewIssueCollector()
	result := v.Validate(ic)
	actual := utils.NewSet[string]()
	ic.EachIssue(func(issue validation.Issue) { actual.Add(issue.Message()) })
	expectedMessages := utils.NewSet(expected...)
	diff := actual.Diff(expectedMessages)
	if diff.Size() != 0 {
		fmt.Printf("%v\n", diff)
	}
	// tt.CheckEqual(0, diff.Size())
	diff = expectedMessages.Diff(actual)
	if diff.Size() != 0 {
		fmt.Printf("%v\n", diff)
	}
	tt.CheckEqual(0, diff.Size())
	if len(expected) == 0 {
		tt.CheckTrue(result)
	} else {
		tt.CheckFalse(result)
	}
}

// validatedModelMessages validates the model (no assumption it is correct) together with
// instance validation. Thus messages can also come from validation when model is broken.
func validateModelMessages(t *testing.T, model, instance string, expected ...string) {
	t.Helper()
	tt := testutils.NewTester(t)
	// use false to stop panic on bad model from makeContext and use ic from model validation
	ctx, ic := makeContext(t.Name(), model, false)
	v := yammm.NewValidator(ctx, t.Name(), makeInstance(instance))
	// ic := validation.NewIssueCollector()
	result := v.Validate(ic)
	if len(expected) == 0 {
		tt.CheckTrue(result)
	} else {
		tt.CheckFalse(result)
	}
	actual := utils.NewSet[string]()
	ic.EachIssue(func(issue validation.Issue) { actual.Add(issue.Message()) })
	expectedMessages := utils.NewSet(expected...)
	diffA := actual.Diff(expectedMessages)
	diffE := expectedMessages.Diff(actual)
	diff := diffA.Union(diffE)
	if diff.Size() != 0 {
		fmt.Printf("actual<->expected:\n%v\n", diff)
	}
	tt.CheckEqual(0, diffE.Size())
}

// makeInstance makes an interface for json input.
func makeInstance(jsonString string) any {
	var m interface{}

	d := json.NewDecoder(strings.NewReader(jsonString))
	d.UseNumber() // Numbers will be a json.Number which is easily checked if it is int or float
	err := d.Decode(&m)
	if err != nil {
		panic(fmt.Sprintf("Bad json given to makeInstance: %s", err))
	}
	return m
}

// makeJzonInstance makes an interface for json input using jzon parser.
func makeJzonInstance(sourceName string, jsonString string) any {
	ctx := jzon.NewContext(sourceName, jsonString)
	node, err := ctx.UnmarshalNode()
	// d := json.NewDecoder(strings.NewReader(jsonString))
	// d.UseNumber() // Numbers will be a json.Number which is easily checked if it is int or float
	// err := d.Decode(&m)
	if err != nil {
		panic(fmt.Sprintf("Bad json given to makeJzonInstance: %s", err))
	}
	return node
}

// makeContext makes a context from yammm model source.
func makeContext(sourceRef string, source string, assertModel bool) (yammm.Context, validation.IssueCollector) {
	yammmCtx, ic := parser.ParseString(sourceRef, source)
	if yammmCtx == nil {
		if assertModel {
			err := validation.NewColorPresentor().Present(ic, validation.Info, os.Stderr)
			if err != nil {
				panic(err)
			}
			panic("Bad input to makeContext - error in test itself.")
		}
	}
	return yammmCtx, ic
}

type Obj map[string]any
type Objs []map[string]any

// Test [name:line:col] in error messages from validation.
