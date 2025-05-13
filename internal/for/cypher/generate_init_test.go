// Contains generated code. For those parts: DO NOT EDIT.
package cypher_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/for/cypher"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/validation"
	"github.com/wyrth-io/whit/internal/yammm"
)

func Test_GenerateInit(t *testing.T) {
	tt := testutils.NewTester(t)

	// Create a Context and set it up with the generated Yammm meta model that is completed.
	ctx := yammm.NewContext()
	err := ctx.SetModelFromJSON(strings.NewReader(SerializedModel))
	tt.CheckNotError(err)
	ic := validation.NewIssueCollector()
	ctx.Complete(ic)
	tt.CheckEqual(0, ic.Count())

	// Generate the cypher statements
	generated := cypher.GenerateInit(ctx)

	tt.CheckTrue(utils.Any(generated, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `CREATE CONSTRAINT Person_primary_keys IF NOT EXISTS FOR (n:Person) REQUIRE (n.name) IS NODE KEY`)
	}))
	for i := 0; i < len(generated); i++ {
		if strings.Contains(generated[i].Source, `Person_required_properties`) {
			fmt.Println(generated[i].Source)
		}
	}
	tt.CheckTrue(utils.Any(generated, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `CREATE CONSTRAINT Person_required_property_birthday IF NOT EXISTS FOR (n:Person) REQUIRE (n.birthday) IS NOT NULL`)
	}))
}

func Test_GenerateInitIdHasNodeKeyConstraint(t *testing.T) {
	model := `schema "test"
	type Person {
		name String
	}`
	ctx,_ := makeContext(t.Name(), model, true)
	// Generate the cypher statements
	generated := cypher.GenerateInit(ctx)
	tt := testutils.NewTester(t)
	tt.CheckFalse(utils.Any(generated, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `CREATE CONSTRAINT Person_primary_keys IF NOT EXISTS FOR (n:Person) REQUIRE (n.name) IS NODE KEY`)
	}))
	tt.CheckTrue(utils.Any(generated, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `CREATE CONSTRAINT Person_primary_uid IF NOT EXISTS FOR (n:Person) REQUIRE (n.uid) IS NODE KEY`)
	}))
}

func Test_GenerateInitHasNodeKeyConstraintsForOtherPrimaryKeysInCombination(t *testing.T) {
	model := `schema "test"
	type Person {
		name String primary
		country String primary
	}`
	ctx,_ := makeContext(t.Name(), model, true)
	// Generate the cypher statements
	generated := cypher.GenerateInit(ctx)
	tt := testutils.NewTester(t)
	tt.CheckTruef(len(generated) >= 2, "att least 2 statements should have been generated: got %d", len(generated))
	tt.CheckTrue(utils.Any(generated, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `CREATE CONSTRAINT Person_primary_uid IF NOT EXISTS FOR (n:Person) REQUIRE (n.uid) IS NODE KEY`)
	}))
	tt.CheckTrue(utils.Any(generated, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `CREATE CONSTRAINT Person_primary_keys IF NOT EXISTS FOR (n:Person) REQUIRE (n.name, n.country) IS NODE KEY`)
	}))
}
func Test_GenerateInitHasDatatypeConstraintsPerProperty(t *testing.T) {
	model := `schema "test"
	type Person {
		name String primary
		country String primary
		x Integer
		y Float
		z Boolean
	}`
	ctx,_ := makeContext(t.Name(), model, true)
	// Generate the cypher statements
	generated := cypher.GenerateInit(ctx)
	tt := testutils.NewTester(t)
	tt.CheckTruef(len(generated) == 8, "8 statements should have been generated: got %d", len(generated))
	tt.CheckTrue(utils.Any(generated, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `CREATE CONSTRAINT Person_uid_dt IF NOT EXISTS FOR (n:Person) REQUIRE n.uid IS :: STRING`)
	}))
	tt.CheckTrue(utils.Any(generated, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `CREATE CONSTRAINT Person_name_dt IF NOT EXISTS FOR (n:Person) REQUIRE n.name IS :: STRING`)
	}))
	tt.CheckTrue(utils.Any(generated, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `CREATE CONSTRAINT Person_country_dt IF NOT EXISTS FOR (n:Person) REQUIRE n.country IS :: STRING`)
	}))
	tt.CheckTrue(utils.Any(generated, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `CREATE CONSTRAINT Person_x_dt IF NOT EXISTS FOR (n:Person) REQUIRE n.x IS :: INTEGER`)
	}))
	tt.CheckTrue(utils.Any(generated, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `CREATE CONSTRAINT Person_y_dt IF NOT EXISTS FOR (n:Person) REQUIRE n.y IS :: FLOAT`)
	}))
	tt.CheckTrue(utils.Any(generated, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `CREATE CONSTRAINT Person_z_dt IF NOT EXISTS FOR (n:Person) REQUIRE n.z IS :: BOOLEAN`)
	}))
}

func Test_GenerateInitSupportSpacevector(t *testing.T) {
	model := `schema "test"
	type Person {
		embedding Spacevector[3]
	}`
	ctx,_ := makeContext(t.Name(), model, true)
	// Generate the cypher statements
	generated := cypher.GenerateInit(ctx)
	tt := testutils.NewTester(t)
	tt.CheckTruef(len(generated) == 4, "4 statements should have been generated: got %d", len(generated))
	tt.CheckTrue(utils.Any(generated, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `CREATE CONSTRAINT Person_primary_uid IF NOT EXISTS FOR (n:Person) REQUIRE (n.uid) IS NODE KEY`)
	}))
	tt.CheckTrue(utils.Any(generated, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `CREATE CONSTRAINT Person_uid_dt IF NOT EXISTS FOR (n:Person) REQUIRE n.uid IS :: STRING`)
	}))
	tt.CheckTrue(utils.Any(generated, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `CREATE CONSTRAINT Person_embedding_dt IF NOT EXISTS FOR (n:Person) REQUIRE n.embedding IS :: LIST<FLOAT NOT NULL>`)
	}))
	tt.CheckTrue(utils.Any(generated, func(st cypher.Statement) bool {
		return strings.Contains(st.Source, `CALL db.index.vector.createNodeIndex($idxName, $nodeLabel, $propertyName, $dim, "cosine")`)
	}))
	// Check that $idxName is corret, $nodeLabel is set to Person and propertyName to "embedding"
	vectorStmnt := generated[3]
	params := vectorStmnt.Parameters
	tt.CheckEqual("Person", params["nodeLabel"])
	tt.CheckEqual("Person_embedding_idx", params["idxName"])
	tt.CheckEqual("embedding", params["propertyName"])
	tt.CheckEqual(3, params["dim"])
}
