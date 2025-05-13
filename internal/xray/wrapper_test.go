package xray_test

import (
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/xray"
)

func Test_WrapperFromMaps(t *testing.T) {
	tt := testutils.NewTester(t)
	xyMap := map[string]any{"x": 42, "y": 84}
	testMap := map[string]any{"a": "hello", "b": []map[string]any{xyMap}}
	doc := xray.NewWrapper(testMap)

	// There are two properties a and b.
	tt.CheckEqual(2, doc.Len())
	tt.CheckStringSlicesEqual([]string{"a", "b"}, doc.FeatureNames())

	// Get value "a" - i.e. as from an object having property "a".
	fa := doc.Value("a")
	tt.CheckEqual("hello", fa)

	// Get feature "b" - i.e. "b" is an Object or Slice.
	fb := doc.Feature("b")

	// It is known to be a Slice so Len() can be obtained.
	tt.CheckEqual(1, fb.Len())
	// There should be a feature at index 0.
	fb0 := fb.FeatureAtIndex(0)
	// And this feature has values for properties "x" and "y"
	fbx := fb0.Value("x")
	tt.CheckEqual(42, fbx)
	fby := fb0.Value("y")
	tt.CheckEqual(84, fby)
}

func Test_WrapperFromStruct(t *testing.T) {
	type XY struct {
		X int
		Y *int
	}
	type TestGraph struct {
		A string
		B []*XY
		C []float32
	}
	tt := testutils.NewTester(t)
	xyMap := XY{X: 42, Y: utils.Ptr(84)}
	testMap := TestGraph{A: "hello", B: []*XY{&xyMap}, C: []float32{1, 2, 3}}
	doc := xray.NewWrapper(testMap)

	// There are two properties A and B.
	tt.CheckEqual(3, doc.Len())
	tt.CheckStringSlicesEqual([]string{"A", "B", "C"}, doc.FeatureNames())

	// Get value "a" - i.e. as from an object having property "a".
	fa := doc.Value("A")
	tt.CheckEqual("hello", fa)

	// Get feature "b" - i.e. "b" is an Object or Slice.
	fb := doc.Feature("B")

	// It is known to be a Slice so Len() can be obtained.
	tt.CheckEqual(1, fb.Len())
	// There should be a feature at index 0.
	fb0 := fb.FeatureAtIndex(0)
	// And this feature has values for properties "x" and "y"
	fbx := fb0.Value("X")
	tt.CheckEqual(42, fbx)
	fby := fb0.Value("Y")
	tt.CheckEqual(84, fby)

	// Get feature "c" - i.e. "c" is an Object or Slice.
	fc := doc.Value("C")
	tt.CheckEqual([]float32{1, 2, 3}, fc)
}
