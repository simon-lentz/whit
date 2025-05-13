package cypher_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/for/cypher"
	"github.com/wyrth-io/whit/internal/tc"
	"github.com/wyrth-io/whit/internal/utils"
)

// Tests:
// Generates statements for composition.
// Generates statements for association.

func Test_MergerDoesUUIDReplacement(t *testing.T) {
	tt := testutils.NewTester(t)
	model := `schema "testing"
	type Car { regNbr String }
	`
	data := `{
		"Cars": [ {"id": "$$local", "regNbr": "ABC123"}]
	}`
	ctx, graph := ValidateModelMessages(t, model, data)
	statements := cypher.NewMergeGenerator().Process(ctx, graph)

	tt.CheckTruef(len(statements) == 1, "one statement should have been generated, got: %d", len(statements))
	tt.CheckTrue(utils.Any(statements, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `(n:Car {uid: $props.uid})`)
	}))
	params := statements[0].Parameters
	props, ok := params["props"].(map[string]any)
	tt.CheckTruef(ok, "params should have a 'props' key")
	uid, ok := props["uid"]
	uids := uid.(string)
	tt.CheckTruef(ok, "props should have a 'uid' key")
	ok, _ = tc.DefaultUUIDChecker.Check(uids)
	theUUID, err := tc.DefaultUUIDChecker.GetUUID(uids)
	tt.CheckNotError(err)
	tt.CheckTruef(ok, "should be typechecked as an UUID")
	tt.CheckTruef(!strings.HasPrefix(uids, "$$"), "should not start with local $$")
	var nullUUID uuid.UUID
	tt.CheckNotEqual(nullUUID, theUUID)
}
func Test_MergerIncludesAllSuperTypeLables(t *testing.T) {
	tt := testutils.NewTester(t)
	model := `schema "testing"
	abstract type Vehicle { v Integer}
	abstract type Motorized extends Vehicle { m Integer}
	type Car extends Vehicle, Motorized { regNbr String }
	`
	data := `{
		"Cars": [ {"id": "$$local", "regNbr": "ABC123"}]
	}`
	ctx, graph := ValidateModelMessages(t, model, data)
	statements := cypher.NewMergeGenerator().Process(ctx, graph)

	tt.CheckTruef(len(statements) == 1, "one statement should have been generated, got: %d", len(statements))
	tt.CheckTruef(strings.Contains(
		statements[0].Source, `(n:Car:Motorized:Vehicle {uid: $props.uid})`),
		"should have found label in alpha order")
}

func Test_MergerIncludesSuperTypeProperties(t *testing.T) {
	tt := testutils.NewTester(t)
	model := `schema "testing"
	abstract type Vehicle { v Integer}
	abstract type Motorized extends Vehicle { v Integer}
	type Car extends Vehicle, Motorized { regNbr String }
	`
	data := `{
		"Cars": [ {"id": "$$local", "regNbr": "ABC123"}]
	}`
	ctx, graph := ValidateModelMessages(t, model, data)
	statements := cypher.NewMergeGenerator().Process(ctx, graph)

	tt.CheckTruef(len(statements) == 1, "one statement should have been generated, got: %d", len(statements))
	tt.CheckTrue(utils.Any(statements, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `(n:Car:Motorized:Vehicle {uid: $props.uid})`)
	}))
	params := statements[0].Parameters
	props, ok := params["props"].(map[string]any)
	tt.CheckTruef(ok, "params should have a 'props' key")
	tt.CheckTruef(len(props) == 2, "should have id and v as properties: got %v", utils.Keys(props))
}

func Test_MergerCastsIntsToFloatWhereNeeded(t *testing.T) {
	tt := testutils.NewTester(t)
	model := `schema "testing"
	type Thing { ratio Float }
	`
	data := `{
		"Things": [ {"id": "$$local", "ratio": 1}]
	}`
	ctx, graph := ValidateModelMessages(t, model, data)
	statements := cypher.NewMergeGenerator().Process(ctx, graph)

	tt.CheckTruef(len(statements) == 1, "one statement should have been generated, got: %d", len(statements))
	params := statements[0].Parameters
	props, ok := params["props"].(map[string]any)
	tt.CheckTruef(ok, "params should have a 'props' key")
	tt.CheckTruef(len(props) == 2, "should have id and ratio as properties: got %v", utils.Keys(props))
	val := props["ratio"]
	tt.CheckTruef(reflect.ValueOf(val).Kind() == reflect.Float64, "value must be a float64, got %t", val)
}
