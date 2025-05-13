package yammm_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/hlindberg/testutils"
	"github.com/lithammer/shortuuid/v4"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/yammm"
)

func Test_ErrorIfSuperTypeDoesNotExist(t *testing.T) {
	model := `
	schema "testing"
	type Car extends Vehicle { regNbr String primary}
	`
	instance := `
	{
		"Cars": [
			{ "regNbr": "ABC123" }	
		]
	}`
	validateModelMessages(t, model, instance,
		fmt.Sprintf("[%s:3:6] type 'Car' inherits from 'Vehicle' - super type not found", t.Name()),
		"Cannot validate instance: no context defined",
	)
}

func Test_MultipleDuplicateExtends(t *testing.T) {
	model := `
	schema "testing"
	type Car extends Thing, Thing { regNbr String primary}
	abstract type Thing {}
	`
	instance := `{ "Cars": [ { "regNbr": "ABC123" }] }`
	validateModelMessages(t, model, instance) // TODO: This is ok - but should warn as something else must have been intended
}
func Test_ExtendsDoesNotExist(t *testing.T) {
	model := `
	schema "testing"
	type Car extends Thing { regNbr String primary}
	`
	instance := `{ "Cars": [ { "regNbr": "ABC123" }] }`
	validateModelMessages(t, model, instance,
		fmt.Sprintf("[%s:3:6] type 'Car' inherits from 'Thing' - super type not found", t.Name()),
		"Cannot validate instance: no context defined",
	)
}

func Test_MultipleIndirectMixins(t *testing.T) {
	model := `
	schema "testing"
	type Car extends Thing, OtherThing { regNbr String primary}
	abstract type Thing extends OtherThing {}
	abstract type OtherThing {}
	`
	instance := `{ "Cars": [ { "regNbr": "ABC123" }] }`
	validateModelMessages(t, model, instance) // is ok
}
func Test_InheritsSeveralTimes(t *testing.T) {
	model := `
	schema "testing"
	type Car extends Vehicle, Thing { regNbr String primary}
	abstract type Thing {}
	type Vehicle extends Thing {}
	`
	instance := `{ "Cars": [ { "regNbr": "ABC123" }] }`
	validateModelMessages(t, model, instance)
}
func Test_InheritsIncompatibleProperty(t *testing.T) {
	model := `
	schema "testing"
	type Car extends Vehicle, Thing { regNbr String primary}
	abstract type Thing { regNbr Integer}
	type Vehicle extends Thing {}
	`
	instance := `{ "Cars": [ { "regNbr": "ABC123" }] }`
	validateModelMessages(t, model, instance,
		"Cannot validate instance: no context defined",
		"[Test_InheritsIncompatibleProperty:4:23] property 'regNbr' definition clashes with definition at [Test_InheritsIncompatibleProperty:3:35]", //nolint
	)
}

// This test is obsolete since the added always present id primary key makes all types have one when
// Created with context.AddType(). To actually test this, the model must be set from JSON with
// the id missing.
func Test_PartWithoutPrimaryKeys(t *testing.T) {
	t.Skip()
	model := `
	schema "testing"
	type Car { regNbr String primary }
	part type CarPart { partNbr String }
	`
	instance := `{ "Cars": [ { "regNbr": "ABC123" }] }`
	validateModelMessages(t, model, instance,
		fmt.Sprintf("[%s:4:11] part type 'CarPart' does not have any primary keys - required for parts", t.Name()),
		"Cannot validate instance: no context defined",
	)
}
func Test_CompositionToSelfIsOk(t *testing.T) {
	model := `
	schema "testing"
	type Something {}
	part type Car { regNbr String primary *-> HAS (many) Car }
	`
	instance := `{ "Somethings":[ {}] }`
	validateModelMessages(t, model, instance)
}

func Test_CompositionMustBeToConcretePart(t *testing.T) {
	model := `
	schema "testing"
	type Car { regNbr String primary *-> HAS (many) Car }
	`
	instance := `{ "Cars": [ { "regNbr": "ABC123" }] }`
	validateModelMessages(t, model, instance,
		fmt.Sprintf("[%s:3:6] composition 'Car.HAS_Cars' must reference a concrete part type",
			t.Name()),
		"Cannot validate instance: no context defined",
	)
}

func Test_TypeDefinedMultipleTimes(t *testing.T) {
	model := `
	schema "testing"
	type Car { regNbr String primary}
	type Car { regNbr String primary}
	`
	instance := `{ "Cars": [ { "regNbr": "ABC123" }] }`
	validateModelMessages(t, model, instance,
		fmt.Sprintf("[%s:4:6] type 'Car' is defined multiple times", t.Name()),
		"Cannot validate instance: no context defined",
	)
}
func Test_AssociationTypeDoesNotExist(t *testing.T) {
	model := `
	schema "testing"
	type Car { regNbr String primary --> HAS (many) Other }
	`
	instance := `{ "Cars": [ { "regNbr": "ABC123" }] }`
	validateModelMessages(t, model, instance,
		fmt.Sprintf("[%s:3:6] type 'Other' referenced in association 'HAS' does not exist", t.Name()),
		"Cannot validate instance: no context defined",
	)
}

func Test_CompositionTypeDoesNotExist(t *testing.T) {
	model := `
	schema "testing"
	type Car { regNbr String primary *-> HAS (many) Other }
	`
	instance := `{ "Cars": [ { "regNbr": "ABC123" }] }`
	validateModelMessages(t, model, instance,
		fmt.Sprintf("[%s:3:6] type 'Other' referenced in composition 'HAS' does not exist", t.Name()),
		"Cannot validate instance: no context defined",
	)
}

func Test_InstanceID_Deterministic(t *testing.T) {
	tt := testutils.NewTester(t)
	model := `
	schema "testing"
	type Car { regNbr String primary }
	`
	ctx, ic := makeContext(t.Name(), model, false)
	tt.CheckFalse(ic.HasErrors() || ic.HasFatal())
	carType := ctx.LookupType("Car")
	propMap := map[string]any{"regNbr": "ABC123"}
	replacements := map[string]uuid.UUID{}
	uid, err := carType.InstanceID(propMap, replacements)
	tt.CheckNotError(err)
	// Manually construct the deterministic UUID
	payload := "ABC123"
	expected := uuid.NewSHA1(yammm.WhitNamespace, []byte(carType.IDPreamble()+payload))
	tt.CheckEqual(expected, uid)
	// Check it is a version 5 UUID
	tt.CheckEqual(uuid.Version('\005'), uid.Version())
	tt.CheckEqual("RFC4122", uid.Variant().String())
}

func Test_InstanceID_Random(t *testing.T) {
	tt := testutils.NewTester(t)
	model := `
	schema "testing"
	type Car { }
	`
	ctx, ic := makeContext(t.Name(), model, false)
	tt.CheckFalse(ic.HasErrors() || ic.HasFatal())
	carType := ctx.LookupType("Car")
	propMap := map[string]any{}
	replacements := map[string]uuid.UUID{}
	uid, err := carType.InstanceID(propMap, replacements)
	tt.CheckNotError(err)
	// check that the uid is a valid uid
	tt.CheckEqual(uuid.Version('\004'), uid.Version())
	tt.CheckEqual("RFC4122", uid.Variant().String())
}

func Test_InstanceID_Replaced(t *testing.T) {
	tt := testutils.NewTester(t)
	model := `
	schema "testing"
	type Car { }
	`
	ctx, ic := makeContext(t.Name(), model, false)
	tt.CheckFalse(ic.HasErrors() || ic.HasFatal())
	carType := ctx.LookupType("Car")
	expected, err := uuid.NewRandom()
	tt.CheckNotError(err)
	short := shortuuid.DefaultEncoder.Encode(expected)

	checkID := func(s string, t *testing.T) { //nolint:thelper
		t.Helper()
		propMap := map[string]any{"id": s}
		replacements := map[string]uuid.UUID{s: expected}
		uid, err := carType.InstanceID(propMap, replacements)
		tt.CheckNotError(err)
		tt.CheckEqual(expected, uid)
		// check that the uid is a valid uid
		tt.CheckEqual(uuid.Version('\004'), uid.Version())
		tt.CheckEqual("RFC4122", uid.Variant().String())
	}
	checkID("$$local1", t)
	checkID("$$:"+short+":local1", t)
}
func Test_Type_AllSubtypes(t *testing.T) {
	tt := testutils.NewTester(t)
	model := `
	schema "testing"
	abstract type Machine {}
	abstract type Vehicle extends Machine {}
	type Car extends Vehicle {}
	type MC extends Vehicle {}
	`
	ctx, ic := makeContext(t.Name(), model, false)
	tt.CheckFalse(ic.HasErrors() || ic.HasFatal())
	carType := ctx.LookupType("Car")
	tt.CheckEqual(0, len(carType.AllSubTypes()))

	vehicleType := ctx.LookupType("Vehicle")
	st := vehicleType.AllSubTypes()
	tt.CheckEqual(2, len(st))
	tt.CheckNotNil(utils.Find(st, func(s string) bool { return s == "Car" }))
	tt.CheckNotNil(utils.Find(st, func(s string) bool { return s == "MC" }))

	machineType := ctx.LookupType("Machine")
	st = machineType.AllSubTypes()
	tt.CheckEqual(3, len(st))
	tt.CheckNotNil(utils.Find(st, func(s string) bool { return s == "Car" }))
	tt.CheckNotNil(utils.Find(st, func(s string) bool { return s == "MC" }))
	tt.CheckNotNil(utils.Find(st, func(s string) bool { return s == "Vehicle" }))
}

// Type plural is mising, is not initial UC, has illegal chars - impossible to test via grammar.
// Type name is missing, is not initial UC, has illegal chars - imposible to test via grammar.
// Type cannot be both Part and Mixin - impossible to test with grammar.
// Mixin with empty name, non UC name, illegal chars in name - impossible to test with grammar.
// Property with non LC start in name, empty or with illegal chars - impossible to test with grammar.

// Property with custom datatype
// Association/Composition Property with custom datatype
// Lint, chill and have a lint-snack .
