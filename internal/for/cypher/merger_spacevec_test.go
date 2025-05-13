package cypher_test

import (
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/for/cypher"
)

func Test_MergerSpacevector(t *testing.T) {
	tt := testutils.NewTester(t)
	model := `schema "testing"
	type Car { regNbr String
		embedding Spacevector[3]
	}
	`
	data := `{
		"Cars": [ {"id": "$$local", "regNbr": "ABC123", "embedding": [1.1, 2.1, 3.1]}]
	}`
	ctx, graph := ValidateModelMessages(t, model, data)
	statements := cypher.NewMergeGenerator().Process(ctx, graph)

	tt.CheckTruef(len(statements) == 1, "one statement should have been generated, got: %d", len(statements))
	props := statements[0].Parameters["props"].(map[string]any)
	tt.CheckEqual([]float32{1.1, 2.1, 3.1}, props["embedding"])
}
