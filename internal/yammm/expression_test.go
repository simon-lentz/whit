package yammm_test

import (
	"reflect"
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/tc"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/yammm"
)

// TestOperators tests that the yammm.Expressions of SExpr type calls the evaluator
// as expected for given operators. If an operator was not correctly implemented and
// called some other evaluator method the test would panic. For example "+" should call Add().
// A pluggable Evaluator implementation (below) is used and a function is poked into it
// for the given field name - for example "FAdd".
func TestOperators(t *testing.T) {
	runBinaryValueOp("+", "FAdd", t)
	runBinaryValueOp("-", "FSub", t)
	runBinaryValueOp("*", "FMul", t)
	runBinaryValueOp("/", "FDiv", t)
	runBinaryBoolOp("==", "FEqual", true, t)
	runBinaryBoolOp("!=", "FEqual", false, t)
	runBinaryBoolOp("=~", "FMatch", true, t)
	runBinaryBoolOp("!~", "FMatch", false, t)
	runBinaryBoolOp("in", "FIn", true, t)
	runUnaryBoolOp("!", "FIsTrue", false, t)
	runUnaryBoolOp("||", "FIsTrue", true, t) // unary since left and right are evaluated separately.
	runUnaryBoolOp("&&", "FIsTrue", true, t) // unary since left and right are evaluated separately.

	runCompareOp("<", "FCompare", true, t)
	runCompareOp("<=", "FCompare", true, t)
	runCompareOp(">", "FCompare", false, t)
	runCompareOp(">=", "FCompare", false, t)
}
func TestLiteral(t *testing.T) {
	tt := testutils.NewTester(t)
	lit := yammm.NewLiteral("hello")
	actual := lit.Eval(nil) // no evaluator needed (unused).
	tt.CheckEqual("hello", actual)
}

func TestReduce(t *testing.T) {
	tt := testutils.NewTester(t)
	// DSL: [1,2,3] reduce { $0 + $1 }
	e := expr("reduce",
		lit([]int{1, 2, 3}),
		lit([]any{}),            // args
		lit([]string{"0", "1"}), // parameter names
		expr("+",
			variable("0"),
			variable("1"),
		))
	ev := yammm.NewEvaluator()
	actual := e.Eval(ev)
	tt.CheckEqual(6, actual)

	// DSL: [1,2,3]->reduce |$memo, $val| { $memo + $val }
	e = expr("reduce",
		lit([]int{1, 2, 3}),
		lit([]any{}),                 // args
		lit([]string{"memo", "val"}), // parameter names
		expr("+",
			variable("memo"),
			variable("val"),
		))
	ev = yammm.NewEvaluator()
	actual = e.Eval(ev)
	tt.CheckEqual(6, actual)
}
func TestThen(t *testing.T) {
	tt := testutils.NewTester(t)
	e := expr("then",
		lit(1),
		lit([]any{}),
		lit([]string{}),
		expr("+",
			variable("0"),
			lit(10),
		))
	ev := yammm.NewEvaluator()
	actual := e.Eval(ev)
	tt.CheckEqual(11, actual)

	e = expr("then",
		lit(nil),
		lit([]any{}),
		lit([]string{}),
		expr("+",
			variable("0"),
			lit(10),
		))
	ev = yammm.NewEvaluator()
	actual = e.Eval(ev)
	tt.CheckNil(actual)
}
func TestLest(t *testing.T) {
	tt := testutils.NewTester(t)
	e := expr("lest",
		lit(1),
		lit([]any{}),
		lit([]string{}),
		lit(10),
	)
	ev := yammm.NewEvaluator()
	actual := e.Eval(ev)
	tt.CheckEqual(1, actual)

	e = expr("lest",
		lit(nil),
		lit([]any{}),
		lit([]string{}),
		lit(10),
	)
	ev = yammm.NewEvaluator()
	actual = e.Eval(ev)
	tt.CheckEqual(10, actual)
}

func TestMap(t *testing.T) {
	tt := testutils.NewTester(t)
	// DSL: [1,2,3] map { $0 * 2 }
	e := expr("map",
		lit([]int{1, 2, 3}),
		lit([]any{}),
		lit([]string{}), // default parameter names
		expr("*",
			variable("0"),
			lit(2),
		))
	ev := yammm.NewEvaluator()
	actual := e.Eval(ev).([]any)
	tt.CheckEqual(2, actual[0])
	tt.CheckEqual(4, actual[1])
	tt.CheckEqual(6, actual[2])

	// DSL: [1,2,3] |$memo, $val| reduce { $memo + $val }
	e = expr("map",
		lit([]int{1, 2, 3}),
		lit([]any{}),
		lit([]string{"val"}), // parameter names
		expr("*",
			variable("val"),
			lit(2),
		))
	ev = yammm.NewEvaluator()
	actual = e.Eval(ev).([]any)
	tt.CheckEqual(2, actual[0])
	tt.CheckEqual(4, actual[1])
	tt.CheckEqual(6, actual[2])
}
func TestFilter(t *testing.T) {
	tt := testutils.NewTester(t)
	// DSL: [1,2,3] filter { $0 != 2 }
	e := expr("filter",
		lit([]int{1, 2, 3}),
		lit([]any{}),
		lit([]string{}), // default parameter names
		expr("!=",
			variable("0"),
			lit(2),
		))
	ev := yammm.NewEvaluator()
	actual := e.Eval(ev).([]any)
	tt.CheckEqual(2, len(actual))
	tt.CheckEqual(1, actual[0])
	tt.CheckEqual(3, actual[1])

	// DSL: [1,2,3] |$memo, $val| reduce { $memo + $val }
	e = expr("filter",
		lit([]int{1, 2, 3}),
		lit([]any{}),
		lit([]string{"val"}), // parameter names
		expr("!=",
			variable("val"),
			lit(2),
		))
	ev = yammm.NewEvaluator()
	actual = e.Eval(ev).([]any)
	tt.CheckEqual(2, len(actual))
	tt.CheckEqual(1, actual[0])
	tt.CheckEqual(3, actual[1])
}
func TestCompact(t *testing.T) {
	tt := testutils.NewTester(t)
	// DSL: [1,_,3,_] compact
	e := expr("compact",
		lit([]any{1, nil, 3, nil}),
	)
	ev := yammm.NewEvaluator()
	actual := e.Eval(ev).([]any)
	tt.CheckEqual(2, len(actual))
	tt.CheckEqual(1, actual[0])
	tt.CheckEqual(3, actual[1])
}

func TestCount(t *testing.T) {
	tt := testutils.NewTester(t)
	// DSL: [1,2,3] count { $0 != 2 }
	e := expr("count",
		lit([]any{1, 2, 3}),
		lit([]any{}),
		lit([]string{}),
		expr("!=",
			variable("0"),
			lit(2),
		),
	)
	ev := yammm.NewEvaluator()
	actual := e.Eval(ev)
	tt.CheckEqual(2, actual)
}

func TestAll(t *testing.T) {
	tt := testutils.NewTester(t)
	// DSL: [1,2,3] count { $0  != 0 }
	e := expr("all",
		lit([]any{1, 2, 3}),
		lit([]any{}),
		lit([]string{}),
		expr("!=",
			variable("0"),
			lit(0),
		),
	)
	ev := yammm.NewEvaluator()
	actual := e.Eval(ev).(bool)
	tt.CheckTrue(actual)

	e = expr("all",
		lit([]any{1, 2, 3}),
		lit([]any{}),
		lit([]string{}),
		expr("==",
			variable("0"),
			lit(0),
		),
	)
	ev = yammm.NewEvaluator()
	actual = e.Eval(ev).(bool)
	tt.CheckFalse(actual)
}

func TestAny(t *testing.T) {
	tt := testutils.NewTester(t)
	// DSL: [1,2,3] count { $0  != 0 }
	e := expr("any",
		lit([]any{1, 2, 3}),
		lit([]any{}),
		lit([]string{}),
		expr("!=",
			variable("0"),
			lit(2),
		),
	)
	ev := yammm.NewEvaluator()
	actual := e.Eval(ev).(bool)
	tt.CheckTrue(actual)

	e = expr("any",
		lit([]any{1, 2, 3}),
		lit([]any{}),
		lit([]string{}),
		expr("==",
			variable("0"),
			lit(0),
		),
	)
	ev = yammm.NewEvaluator()
	actual = e.Eval(ev).(bool)
	tt.CheckFalse(actual)
}

func TestCompare(t *testing.T) {
	tt := testutils.NewTester(t)
	// DSL: 1->compare(2)
	e := expr("compare",
		lit(1),
		lit([]any{lit(2)}),
		lit([]string{}),
		nil,
	)
	ev := yammm.NewEvaluator()
	actual := e.Eval(ev).(int)
	tt.CheckEqual(-1, actual)
}
func TestLen(t *testing.T) {
	tt := testutils.NewTester(t)
	// DSL: [1,2,3] len
	e := expr("len",
		lit([]any{1, 2, 3}),
	)
	ev := yammm.NewEvaluator()
	actual := e.Eval(ev).(int)
	tt.CheckEqual(3, actual)
}
func TestUnique(t *testing.T) {
	tt := testutils.NewTester(t)
	// DSL: [1, 2, 3, 2, 1] unique
	e := expr("unique",
		lit([]any{1, 2, 3, 2, 1}),
	)
	ev := yammm.NewEvaluator()
	actual := e.Eval(ev)
	tt.CheckEqual([]any{1, 2, 3}, actual)

	// DSL: 1 unique
	e = expr("unique",
		lit(1),
	)
	ev = yammm.NewEvaluator()
	actual = e.Eval(ev)
	tt.CheckEqual([]any{1}, actual)

	// DSL: [1,2,"a", [1,2], "a", [1,2]] unique
	e = expr("unique",
		lit([]any{1, 2, "a", []any{1, 2}, "a", []any{1, 2}}),
	)
	ev = yammm.NewEvaluator()
	actual = e.Eval(ev)
	tt.CheckEqual([]any{1, 2, "a", []any{1, 2}}, actual)
}

func TestSlice_SingleValue(t *testing.T) {
	tt := testutils.NewTester(t)
	// DSL: [1, 2, 3][2]
	e := expr("@",
		lit([]any{1, 2, 3}),
		lit(2),
	)
	ev := yammm.NewEvaluator()
	actual := e.Eval(ev)
	tt.CheckEqual(3, actual)
}
func TestSlice_Range(t *testing.T) {
	tt := testutils.NewTester(t)
	// DSL: [1, 2, 3, 4, 5][2,4]
	e := expr("@",
		lit([]any{1, 2, 3, 4, 5}),
		lit(2),
		lit(4),
	)
	ev := yammm.NewEvaluator()
	actual := e.Eval(ev)
	tt.CheckEqual([]any{3, 4}, actual)
}
func TestSlice_Type(t *testing.T) {
	tt := testutils.NewTester(t)
	// DSL: Integer[2,4]
	e := expr("@",
		dtlit("Integer"),
		lit(2),
		lit(4),
	)
	ev := yammm.NewEvaluator()
	actual := e.Eval(ev)
	tc := actual.(tc.TypeChecker)
	ok, _ := tc.Check(2)
	tt.CheckTrue(ok)
	ok, _ = tc.Check(4)
	tt.CheckTrue(ok)
	ok, _ = tc.Check(1)
	tt.CheckFalse(ok)
	ok, _ = tc.Check(5)
	tt.CheckFalse(ok)
}

func TestMod(t *testing.T) {
	tt := testutils.NewTester(t)
	e := expr("%",
		lit(10),
		lit(3),
	)
	ev := yammm.NewEvaluator()
	actual := e.Eval(ev)
	tt.CheckEqual(1, actual)
}

func TestXor(t *testing.T) {
	tt := testutils.NewTester(t)
	e := expr("^",
		lit(false),
		lit(true),
	)
	ev := yammm.NewEvaluator()
	actual := e.Eval(ev)
	tt.CheckEqual(true, actual)
}

func expr(op yammm.Op, args ...yammm.Expression) (e yammm.Expression) {
	se := yammm.SExpr{op}
	return append(se, args...)
}
func lit(v any) yammm.Expression {
	return yammm.NewLiteral(v)
}
func dtlit(v string) yammm.Expression {
	return yammm.DatatypeLiteral(v)
}
func op(s string) yammm.Op {
	return yammm.Op(s)
}
func variable(name string) yammm.Expression {
	return expr("$", lit(name))
}

type pluggableEvaluator struct {
	BaseEvaluator
	FAdd func(left, right any) any
	FMul func(left, right any) any
	FSub func(left, right any) any
	FDiv func(left, right any) any

	FIsTrue func(value any) bool
	FNot    func(value any) bool
	FEqual  func(left, right any) bool
	FMatch  func(left, right any) bool
	FIn     func(left, right any) bool

	FReduce  func(left any, params []string, expr yammm.Expression, start ...any) any
	FCompare func(left, right any) int
}

func (p *pluggableEvaluator) Add(left, right any) any { return p.FAdd(left, right) }
func (p *pluggableEvaluator) Mul(left, right any) any { return p.FMul(left, right) }
func (p *pluggableEvaluator) Sub(left, right any) any { return p.FSub(left, right) }
func (p *pluggableEvaluator) Div(left, right any) any { return p.FDiv(left, right) }

func (p *pluggableEvaluator) IsTrue(value any) bool      { return p.FIsTrue(value) }
func (p *pluggableEvaluator) Not(value any) bool         { return p.FNot(value) }
func (p *pluggableEvaluator) Equal(left, right any) bool { return p.FEqual(left, right) }
func (p *pluggableEvaluator) Match(left, right any) bool { return p.FMatch(left, right) }
func (p *pluggableEvaluator) In(left, right any) bool    { return p.FIn(left, right) }
func (p *pluggableEvaluator) Reduce(left any, params []string, expr yammm.Expression, start ...any) any {
	return p.FReduce(left, params, expr, start...)
}
func (p *pluggableEvaluator) Compare(left, right any) int { return p.FCompare(left, right) }

// runBinaryValueOp creats a pluggable evaluator and pokes a function into
// the evaluator based on the given fname. This is used to test a binary op that
// returns a value.
func runBinaryValueOp(opstr string, fname string, t *testing.T) { //nolint:thelper
	t.Helper()
	tt := testutils.NewTester(t)
	e := expr(op(opstr), lit(1), lit(2))
	ev := pluggableEvaluator{}
	f := func(left, right any) any { return []any{left, right} }
	reflect.ValueOf(&ev).Elem().FieldByName(fname).Set(reflect.ValueOf(f))
	actual := e.Eval(&ev).([]any)
	tt.CheckEqual(1, actual[0])
	tt.CheckEqual(2, actual[1])
}

// runBinaryBoolOp creats a pluggable evaluator and pokes a function into
// the evaluator based on the given fname. This is used to test a binary op that
// returns a boolean value.
func runBinaryBoolOp(opstr string, fname string, expected bool, t *testing.T) { //nolint:thelper
	t.Helper()
	tt := testutils.NewTester(t)
	e := expr(op(opstr), lit(1), lit(2))
	ev := pluggableEvaluator{}
	f := func(left, right any) bool { return true }
	reflect.ValueOf(&ev).Elem().FieldByName(fname).Set(reflect.ValueOf(f))
	actual := e.Eval(&ev).(bool)
	tt.CheckEqual(expected, actual)
}

// runCompareOp creats a pluggable evaluator and pokes a function into
// the evaluator based on the given fname. This is used to test a binary op that
// returns a boolean value.
func runCompareOp(opstr string, fname string, expected bool, t *testing.T) { //nolint:thelper
	t.Helper()
	tt := testutils.NewTester(t)
	e := expr(op(opstr), lit(1), lit(2))
	ev := pluggableEvaluator{}
	f := utils.ValueOrder
	reflect.ValueOf(&ev).Elem().FieldByName(fname).Set(reflect.ValueOf(f))
	actual := e.Eval(&ev).(bool)
	tt.CheckEqual(expected, actual)
}

// runUnaryBoolOp creats a pluggable evaluator and pokes a function into
// the evaluator based on the given fname. This is used to test an unary op that
// returns a boolean value.
func runUnaryBoolOp(opstr string, fname string, expected bool, t *testing.T) { //nolint:thelper
	t.Helper()
	tt := testutils.NewTester(t)
	e := expr(op(opstr), lit(1), lit(2))
	ev := pluggableEvaluator{}
	f := func(left any) bool { return true }
	reflect.ValueOf(&ev).Elem().FieldByName(fname).Set(reflect.ValueOf(f))
	actual := e.Eval(&ev).(bool)
	tt.CheckEqual(expected, actual)
}
