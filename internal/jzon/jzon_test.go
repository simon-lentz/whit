package jzon_test

import (
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/jzon"
	"github.com/wyrth-io/whit/internal/utils"
)

func Test_Context_Line(t *testing.T) {
	tt := testutils.NewTester(t)
	ctx := jzon.NewContext(t.Name(), "abc\ndef\nghi\n")
	for offset := 0; offset < 3; offset++ {
		tt.CheckEqual(1, ctx.Line(offset))
	}
	for offset := 4; offset < 7; offset++ {
		tt.CheckEqual(2, ctx.Line(offset))
	}
	for offset := 8; offset < 11; offset++ {
		tt.CheckEqual(3, ctx.Line(offset))
	}
}

func Test_Context_Col(t *testing.T) {
	tt := testutils.NewTester(t)
	ctx := jzon.NewContext(t.Name(), "abc\ndef\nghi\n")
	for offset := 0; offset < 3; offset++ {
		tt.CheckEqual(offset+1, ctx.Col(offset))
	}
	for offset := 4; offset < 7; offset++ {
		tt.CheckEqual(offset-3, ctx.Col(offset))
	}
	for offset := 8; offset < 11; offset++ {
		tt.CheckEqual(offset-7, ctx.Col(offset))
	}
}

func Test_UnmarshalNode(t *testing.T) {
	tt := testutils.NewTester(t)
	source := `{"a": [ "b", "c" ], "x": [42, 43]}`
	ctx := jzon.NewContext(t.Name(), source)
	node, err := ctx.UnmarshalNode()
	tt.CheckNil(err)
	tt.CheckNotNil(node)
	tt.CheckTrue(node.IsObject())
	for k, v := range node.ObjectValue() {
		tt.CheckTrue(k.IsString())
		tt.CheckTrue(v.IsArray())
	}
}
func Test_LineIsSetOnError(t *testing.T) {
	tt := testutils.NewTester(t)
	source := `{"a":
		[ "b", "c" }]`
	ctx := jzon.NewContext(t.Name(), source)
	node, err := ctx.UnmarshalNode()
	tt.CheckNotNil(err)
	tt.CheckEqual("[Test_LineIsSetOnError:2:14] Invalid JSON: invalid character '}' after array element", err.Error())
	tt.CheckNil(node)
}
func Test_TrailingCommaErrorArray(t *testing.T) {
	tt := testutils.NewTester(t)
	source := `{"a":
		[ "b", "c",]`
	ctx := jzon.NewContext(t.Name(), source)
	node, err := ctx.UnmarshalNode()
	tt.CheckNotNil(err)
	tt.CheckEqual("[Test_TrailingCommaErrorArray:2:14] Invalid JSON: invalid character ']' looking for beginning of value", err.Error())
	tt.CheckNil(node)
}
func Test_TrailingCommaErrorObj(t *testing.T) {
	tt := testutils.NewTester(t)
	source := `{"a":
		[ "b", "c"],}`
	ctx := jzon.NewContext(t.Name(), source)
	node, err := ctx.UnmarshalNode()
	tt.CheckNotNil(err)
	tt.CheckEqual(
		"[Test_TrailingCommaErrorObj:2:15] Invalid JSON: invalid character '}' looking for beginning of object key string",
		err.Error())
	tt.CheckNil(node)
}

func Test_ValueNodes(t *testing.T) {
	tt := testutils.NewTester(t)
	source := `"a"`
	ctx := jzon.NewContext(t.Name(), source)
	node, err := ctx.UnmarshalNode()
	tt.CheckNil(err)
	tt.CheckTrue(node.IsString())
	tt.CheckEqual("a", node.StringValue())

	source = `42`
	ctx = jzon.NewContext(t.Name(), source)
	node, err = ctx.UnmarshalNode()
	tt.CheckNil(err)
	tt.CheckTrue(node.IsInt())
	tt.CheckEqual(42, node.IntValue())

	source = `4.2`
	ctx = jzon.NewContext(t.Name(), source)
	node, err = ctx.UnmarshalNode()
	tt.CheckNil(err)
	tt.CheckTrue(node.IsFloat())
	tt.CheckEqual(4.2, node.FloatValue())

	source = `true`
	ctx = jzon.NewContext(t.Name(), source)
	node, err = ctx.UnmarshalNode()
	tt.CheckNil(err)
	tt.CheckTrue(node.IsBool())
	tt.CheckTrue(node.BoolValue())
}
func Test_HandleNull(t *testing.T) {
	tt := testutils.NewTester(t)
	t.Run("null not accepted", func(t *testing.T) {
		source := `null`
		ctx := jzon.NewContext(t.Name(), source)
		node, err := ctx.UnmarshalNode()
		tt.CheckNotNil(err)
		tt.CheckEqual("[Test_HandleNull/null_not_accepted:1:1] JSON Null value is not accepted", err.Error())
		tt.CheckNil(node)
	})

	t.Run("null accepted", func(t *testing.T) {
		source := `null`
		ctx := jzon.NewContext(t.Name(), source, jzon.Option("accept-null"))
		node, err := ctx.UnmarshalNode()
		tt.CheckNil(err)
		tt.CheckTrue(node.IsNull())
	})
}

func Test_ObjNode_Property(t *testing.T) {
	tt := testutils.NewTester(t)
	source := `{"a": 10, "b": 20}`
	ctx := jzon.NewContext(t.Name(), source)
	node, err := ctx.UnmarshalNode()
	tt.CheckNotError(err)
	tt.CheckTrue(node.IsObject())
	tt.CheckEqual(10, node.Property("a").IntValue())
	tt.CheckEqual(20, node.Property("b").IntValue())
}
func Test_ObjNode_PropertyNames(t *testing.T) {
	tt := testutils.NewTester(t)
	source := `{"a": 10, "b": 20}`
	ctx := jzon.NewContext(t.Name(), source)
	node, err := ctx.UnmarshalNode()
	tt.CheckNotError(err)
	tt.CheckTrue(node.IsObject())
	props := node.PropertyNames()
	tt.CheckTrue(utils.Any(props, func(s string) bool { return s == "a" }))
	tt.CheckTrue(utils.Any(props, func(s string) bool { return s == "b" }))
	tt.CheckEqual(2, len(props))
}
func Test_ArraysOfObj(t *testing.T) {
	tt := testutils.NewTester(t)
	source := `
	{
		"Cars": [
			{ "regNbr": "ABC123" }	
		],
		"Contrabands": { "xxx": 42 }
	}`
	ctx := jzon.NewContext(t.Name(), source)
	node, err := ctx.UnmarshalNode()
	tt.CheckNotError(err)
	tt.CheckTrue(node.IsObject())
}

func Test_LineAndCol(t *testing.T) {
	tt := testutils.NewTester(t)
	source := `{ "Cars": [ { "regNbr": 42 } ] }`
	ctx := jzon.NewContext(t.Name(), source)
	node, err := ctx.UnmarshalNode()
	tt.CheckNotError(err)
	tt.CheckTrue(node.IsObject())
	tt.CheckEqual(1, node.Line())
	tt.CheckEqual(1, node.Col())

	cars := node.PropertyName("Cars")
	tt.CheckEqual(1, cars.Line())
	tt.CheckEqual(3, cars.Col())
}
