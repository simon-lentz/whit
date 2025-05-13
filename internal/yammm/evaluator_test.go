package yammm_test

import (
	"math"
	"regexp"
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/yammm"
)

func TestEvaluatorAdd(t *testing.T) {
	tt := testutils.NewTester(t)
	e := yammm.NewEvaluator()
	tt.CheckEqual(3, e.Add(1, 2))
	tt.CheckEqual(3.2, e.Add(1.1, 2.1))
	tt.CheckEqual(3.1, e.Add(1, 2.1))
	tt.CheckEqual(3.1, e.Add(1.1, 2))
}

func TestEvaluatorSub(t *testing.T) {
	tt := testutils.NewTester(t)
	e := yammm.NewEvaluator()
	tt.CheckEqual(8, e.Sub(10, 2))
	tt.CheckEqual(8.1, e.Sub(10.2, 2.1))
	tt.CheckEqual(7.9, e.Sub(10, 2.1))
	tt.CheckEqual(8.1, e.Sub(10.1, 2))
}

func TestEvaluatorMul(t *testing.T) {
	round := func(x any) float64 { return math.Round(x.(float64)*100.0) / 100.0 }
	tt := testutils.NewTester(t)
	e := yammm.NewEvaluator()
	tt.CheckEqual(20, e.Mul(10, 2))
	tt.CheckEqual(21.42, round(e.Mul(10.2, 2.1)))
	tt.CheckEqual(21, round(e.Mul(10, 2.1)))
	tt.CheckEqual(20.2, round(e.Mul(10.1, 2)))
}
func TestEvaluatorDiv(t *testing.T) {
	round := func(x any) float64 { return math.Round(x.(float64)*100.0) / 100.0 }
	tt := testutils.NewTester(t)
	e := yammm.NewEvaluator()
	tt.CheckEqual(5, e.Div(10, 2))
	tt.CheckEqual(4.86, round(e.Div(10.2, 2.1)))
	tt.CheckEqual(4.76, round(e.Div(10, 2.1)))
	tt.CheckEqual(5.05, round(e.Div(10.1, 2)))
}
func TestEvaluatorMod(t *testing.T) {
	tt := testutils.NewTester(t)
	e := yammm.NewEvaluator()
	tt.CheckEqual(1, e.Mod(10, 3))
}

func TestEvaluatorEqual(t *testing.T) {
	tt := testutils.NewTester(t)
	e := yammm.NewEvaluator()
	tt.CheckTrue(e.Equal(1, 1))
	tt.CheckFalse(e.Equal(1, 2))

	tt.CheckTrue(e.Equal(1.1, 1.1))
	tt.CheckFalse(e.Equal(1.1, 2.1))

	tt.CheckTrue(e.Equal("abc", "abc"))
	tt.CheckFalse(e.Equal("abc", "xyz"))

	tt.CheckTrue(e.Equal(true, true))
	tt.CheckTrue(e.Equal(false, false))
	tt.CheckFalse(e.Equal(true, false))
	tt.CheckFalse(e.Equal(false, true))
}

func TestEvaluatorNot(t *testing.T) {
	tt := testutils.NewTester(t)
	e := yammm.NewEvaluator()
	tt.CheckTrue(e.Not(false))
	tt.CheckFalse(e.Not(true))
}

func TestEvaluatorIsTrue(t *testing.T) {
	tt := testutils.NewTester(t)
	e := yammm.NewEvaluator()
	tt.CheckTrue(e.IsTrue(true))
	tt.CheckFalse(e.IsTrue(false))
}

func TestEvaluatorIn(t *testing.T) {
	tt := testutils.NewTester(t)
	e := yammm.NewEvaluator()
	arr := []any{1, 2, 3, 4}
	tt.CheckTrue(e.In(1, arr))
	tt.CheckFalse(e.In(5, arr))
	arr = []any{"a", "b", "c"}
	tt.CheckTrue(e.In("b", arr))
	tt.CheckFalse(e.In("x", arr))
}

func TestEvaluatorMatch(t *testing.T) {
	tt := testutils.NewTester(t)
	e := yammm.NewEvaluator()
	rx := regexp.MustCompile("^b$")
	tt.CheckTrue(e.Match("b", rx))
	tt.CheckFalse(e.Match("x", rx))
}
func TestEvaluatorAbs(t *testing.T) {
	tt := testutils.NewTester(t)
	e := yammm.NewEvaluator()
	tt.CheckEqual(1, e.Abs(-1))
	tt.CheckEqual(1, e.Abs(1))
	tt.CheckEqual(1.1, e.Abs(-1.1))
	tt.CheckEqual(1.1, e.Abs(1.1))
}
func TestEvaluatorFloor(t *testing.T) {
	tt := testutils.NewTester(t)
	e := yammm.NewEvaluator()
	tt.CheckEqual(-1, e.Floor(-1))
	tt.CheckEqual(1, e.Floor(1))
	tt.CheckEqual(-2, e.Floor(-1.1))
	tt.CheckEqual(1, e.Floor(1.1))
}
func TestEvaluatorCeil(t *testing.T) {
	tt := testutils.NewTester(t)
	e := yammm.NewEvaluator()
	tt.CheckEqual(-1, e.Ceil(-1))
	tt.CheckEqual(1, e.Ceil(1))
	tt.CheckEqual(-1, e.Ceil(-1.1))
	tt.CheckEqual(2, e.Ceil(1.1))
}
func TestEvaluatorMin(t *testing.T) {
	tt := testutils.NewTester(t)
	e := yammm.NewEvaluator()
	tt.CheckEqual(-1, e.Min(-1, 1)) // ok
	tt.CheckEqual(-1, e.Min(1, -1)) // TODO: debug !!!
}
func TestEvaluatorMax(t *testing.T) {
	tt := testutils.NewTester(t)
	e := yammm.NewEvaluator()
	tt.CheckEqual(1, e.Max(-1, 1))
	tt.CheckEqual(1, e.Max(1, -1))
}
func TestEvaluatorProperty(t *testing.T) {
	tt := testutils.NewTester(t)
	e := yammm.NewEvaluator(map[string]any{"a": 10, "b": 20})
	tt.CheckEqual(10, e.Property("a"))
	tt.CheckEqual(20, e.Property("b"))
	tt.CheckNil(e.Property("c"))
}

// BaseEvaluator returns nil on all operations returning any, and false for all boolean
// operations. For operations where a slice would be returned it returns an empty []any{}.
type BaseEvaluator struct {
	yammm.Evaluator
}

func (*BaseEvaluator) Add(_, _ any) any { return nil }
func (*BaseEvaluator) Mul(_, _ any) any { return nil }
func (*BaseEvaluator) Sub(_, _ any) any { return nil }
func (*BaseEvaluator) Div(_, _ any) any { return nil }

func (*BaseEvaluator) IsTrue(_ any) bool   { return false }
func (*BaseEvaluator) Not(_ any) bool      { return false }
func (*BaseEvaluator) Equal(_, _ any) bool { return false }
func (*BaseEvaluator) Match(_, _ any) bool { return false }
func (*BaseEvaluator) In(_, _ any) bool    { return false }

func (*BaseEvaluator) Var(_ string) any                                           { return nil }
func (*BaseEvaluator) Reduce(_ any, _ []string, _ yammm.Expression, _ ...any) any { return nil }
func (*BaseEvaluator) Then(_ any, _ []string, _ yammm.Expression) any             { return nil }
func (*BaseEvaluator) Lest(_ any, _ yammm.Expression) any                         { return nil }
func (*BaseEvaluator) SetVar(_ string, _ any) any                                 { return nil }
func (*BaseEvaluator) Map(_ any, _ []string, _ yammm.Expression) any              { return []any{} }
func (*BaseEvaluator) Filter(_ any, _ []string, _ yammm.Expression) any           { return []any{} }
func (*BaseEvaluator) Compact(_ any) any                                          { return []any{} }
func (*BaseEvaluator) Count(_ any, _ []string, _ yammm.Expression) any            { return []any{} }
func (*BaseEvaluator) Any(_ any, _ []string, _ yammm.Expression) bool             { return false }
func (*BaseEvaluator) All(_ any, _ []string, _ yammm.Expression) bool             { return false }
func (*BaseEvaluator) Compare(_, _ any) int                                       { return 0 }
