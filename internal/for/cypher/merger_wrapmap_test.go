package cypher_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/for/cypher"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/validation"
	"github.com/wyrth-io/whit/internal/xray"
	"github.com/wyrth-io/whit/internal/yammm"
)

func Test_CyperMergeFromWrappedGeneratedGo(t *testing.T) {
	tt := testutils.NewTester(t)
	car := Car{}
	car.RegNbr = "ABC123"

	d := xray.NewWrapper(&car)
	names := d.FeatureNames()
	tt.CheckTrue(len(names) > 0)

	d = xray.NewWrapper(&car)
	names = d.FeatureNames()
	tt.CheckTrue(len(names) > 0)
}

func Test_CypherMergeFromMap(t *testing.T) {
	tt := testutils.NewTester(t)

	var g map[string]any
	if err := json.Unmarshal([]byte(blob), &g); err != nil {
		tt.Fatalf("could not unmarshal: %s", err.Error())
	}
	GenerateAndAssert(t, g)
}

func Test_CypherGenerate2FromStructs(t *testing.T) {
	tt := testutils.NewTester(t)

	var g Graph
	if err := json.Unmarshal([]byte(blob), &g); err != nil {
		tt.Fatalf("could not unmarshal: %s", err.Error())
	}
	GenerateAndAssert(t, g)
}

func GenerateAndAssert(t *testing.T, g any) { //nolint:thelper
	tt := testutils.NewTester(t)
	// Create a Context
	ctx := yammm.NewContext()
	// ... and set it up with the generated Yammm meta model that is completed.
	err := ctx.SetModelFromJSON(strings.NewReader(SerializedModel))
	tt.CheckNotError(err)
	ic := validation.NewIssueCollector()
	ctx.Complete(ic)
	tt.CheckEqual(0, ic.Count())

	// Generate the cypher statements
	generator := cypher.NewMergeGenerator()
	generated := generator.Process(ctx, g)

	tt.CheckTrue(utils.Any(generated, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `(n:Car:Registered:RegisteredVehicle:Vehicle {uid: $props.uid})`)
	}))
	tt.CheckTrue(utils.Any(generated, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `(n:Person:Entity:MotorVehicleOwner {uid: $props.uid})`)
	}))
	tt.CheckTrue(utils.Any(generated, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `[r:HAS_Head]`)
	}))
	tt.CheckTrue(utils.Any(generated, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `[r:HAS_Limbs]`)
	}))
	tt.CheckTrue(utils.Any(generated, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `[r:MOTHER_Person]`)
	}))

	tt.CheckTrue(utils.Any(generated, func(st cypher.Statement) bool {
		params := st.Parameters
		props := params["props"]
		switch x := props.(type) {
		case map[string]interface{}:
			if !(utils.HasKey(x, "fromDate")) {
				return false
			}
		default:
			return false
		}
		return strings.Contains(st.Source, `[r:OWNS_VEHICLE_RegisteredVehicle {fromDate: $props.fromDate}]`) &&
			strings.Contains(st.Source, `from.uid = $fromUid`) &&
			strings.Contains(st.Source, `to.uid = $toUid`)
	}))
}
