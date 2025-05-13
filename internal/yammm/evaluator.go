package yammm

import (
	"math"
	"reflect"
	"regexp"

	"github.com/tada/catch"
	"github.com/wyrth-io/whit/internal/tc"
	"github.com/wyrth-io/whit/internal/utils"
)

// Evaluator defines an interface supporting evaluation of expressions. In contrast to
// supporting all kinds of operations, it provides the foundation - for example Compare() handles
// all kinds of comparisons which is used to support <, >, <=, >= etc.
type Evaluator interface {
	// NewWithProperties returns the same ind of evaluator as this with given properties
	// set.
	NewWithProperties(props map[string]any) Evaluator

	// SetVar sets the value of a variable.
	SetVar(name string, value any) any

	// Add adds the left and right value.
	Add(left, right any) any

	// Mul multiplies the left with right value.
	Mul(left, right any) any

	// Sub Subtracts the right from the left value.
	Sub(left, right any) any

	// Div divides the left by the right value.
	Div(left, right any) any

	// Mod produces the left modulo the right value.
	Mod(left, right any) any

	// IsTrue tests if the value is boolean true.
	IsTrue(value any) bool

	// Not returns the boolean inverse of the given bool value.
	Not(value any) bool

	// Equal tests if the left and right values are equal
	Equal(left, right any) bool

	// Match tests if the left value matches the right value.
	Match(left, right any) bool

	// MatchRegexp matches string with regexp and returns slice of matching
	// overall and each submatch.
	MatchRegexp(left, right any) []any

	// In tests if the left value is in the array of values in right
	In(left, right any) bool

	// Var returns the value of a variable (or nil if not set)
	Var(name string) any

	// Property returns the value of the property (or nil if it is not set).
	Property(name string) any

	// Reduce reduces the left array by the given Expression. The params
	// describe the names of the two variables to use in the expression.
	// If left empty they will be available as "$0" and "$1" (in DSL syntax).
	// Reduce feeds "memo" (the result of the previous evaluation of the expr),
	// as the first argument, and each respective value iteratively as the second.
	// If a starting value is provided it will be used as the value of "memo" in the
	// first iteration, otherwise the first value of the given array. The result of
	// evaluating the expr for the last element in the array is returned.
	// For example DSL `[1,2,3] reduce {$0 + $1}` would return the result 6.
	Reduce(left any, params []string, expr Expression, start ...any) any

	// Then evaluates the expr after having assigned the left value to the given
	// parameter, or to "0" if no params are given. If the left value is nil,
	// the expression is not evaluated and nil is returned otherwise the result of
	// evaluating the expression.
	Then(left any, params []string, expr Expression) any

	// Lest evaluates the expr if the left value is nil and returns the result.
	// If the left value is not nil it is returned without evaluation of the expression.
	Lest(left any, expr Expression) any

	// With evaluates the expr after having assigned the left value to the given
	// parameter, or to "0" if no params are given.
	With(left any, params []string, expr Expression) any

	// Map evaluates the given expression for each value in the left array set as
	// a variable name given by params, or "0" if no param names are given. A new array
	// of each evaluation of the expression is returned.
	Map(left any, params []string, expr Expression) any

	// Filter evaluates the given expression for each value in the left array set as
	// a variable name given by params, or "0" if no param names are given. A new array
	// of each element in the left array for which the evaluation of the expression returns
	// true is returned.
	Filter(left any, params []string, expr Expression) any

	// Compact filters returns a new array with all non nil values in left.
	// all
	Compact(left any) any

	// Count evaluates the given expression for each value in the left array set as
	// a variable name given by params, or "0" if no param names are given. A count of
	// the number of times the evaluated exression returned true is returned.
	Count(left any, params []string, expr Expression) any

	// Any evaluates the given expression for each value in the left array set as
	// a variable name given by params, or "0" if no param names are given. Evaluation stops
	// and returns true when the evaluation of the expression returns true for the first time.
	Any(left any, params []string, expr Expression) bool

	// All evaluates the given expression for each value in the left array set as
	// a variable name given by params, or "0" if no param names are given. Evaluation stops
	// and returns false when the evaluation of the expression returns false for the first time.
	// Boolean true is returned if all evaluations of the expression returned true.
	All(left any, params []string, expr Expression) bool

	// AllOrNone performs the equivalent of `left count == left len || left count == 0
	AllOrNone(left any, params []string, expr Expression) bool

	// Compare compares left and right and returns 0 if equal, 1 if left is greater and -1 if right
	// is greater.
	Compare(left, right any) int

	// Len computes the length of a string or a slice.
	Len(left any) int

	// Unique filters a slice based on uniqueness. If the value is not a slice a new slice with
	// the single value is returned.
	Unique(left any) []any

	// ChangeSign changes the sign of a numeric value.
	ChangeSign(val any) any

	// Abs returns the absolute value of a numeric value.
	Abs(val any) any

	// Floor returns the math.Floor value for a floating point value. If given an integer value
	// the value is returned.
	Floor(val any) any

	// Floor returns the math.Floor value for a floating point value. If given an integer value
	// the value is returned.
	Ceil(val any) any

	// Round returns the math.RoundToEven value for a floating point value. If given an integer value
	// the value is returned. Panics for non numerical value.
	Round(val any) any

	// Min returns the smaller of two compareable values.
	Min(left any, right ...any) any

	// Max returns the greater of two compareable values.
	Max(left any, right ...any) any
}

type evaluator struct {
	vars   map[string]any
	parent *evaluator
	props  map[string]any
}

// NewEvaluator returns a new evaluator, optionally with a map of property values.
func NewEvaluator(props ...map[string]any) Evaluator {
	var theProps map[string]any
	if len(props) > 0 {
		theProps = props[0]
	}
	return &evaluator{vars: make(map[string]any), props: theProps}
}
func newParentedEvaluator(e *evaluator, props ...map[string]any) Evaluator { //nolint:unparam
	var theProps map[string]any
	if len(props) > 0 {
		theProps = props[0]
	}
	e2 := &evaluator{vars: make(map[string]any), parent: e, props: theProps}
	e2.setVar("self", theProps)
	return e2
}
func (e *evaluator) NewWithProperties(props map[string]any) Evaluator {
	e2 := &evaluator{vars: make(map[string]any), parent: e, props: props}
	e2.setVar("self", props)
	return e2
}
func (e *evaluator) SetVar(name string, val any) any {
	if name == "self" {
		panic(catch.Error("$self is a reserved variable name"))
	}
	return e.setVar(name, val)
}

// setVar is an unchecked package private variable setting.
func (e *evaluator) setVar(name string, val any) any {
	e.vars[name] = val
	return val
}
func (e *evaluator) getInt64(val any) (int64, bool) {
	switch x := val.(type) {
	case int:
		return int64(x), true
	case int8:
		return int64(x), true
	case int16:
		return int64(x), true
	case int32:
		return int64(x), true
	case int64:
		return x, true
	}
	return 0, false
}
func (e *evaluator) getFloat64(val any) (float64, bool) {
	switch x := val.(type) {
	case float32:
		return float64(x), true
	case float64:
		return x, true
	}
	return 0, false
}

// ternary state.
type ternary int

// Ternary constants NEITHER, TRUE, FALSE indicate a one of the ternary states.
const (
	NEITHER = ternary(0)
	TRUE    = ternary(1)
	FALSE   = ternary(-1)
)

func (e *evaluator) getPromoted(left, right any) (li, ri int64, lf, rf float64, floatop ternary) {
	var ok bool
	lf, ok = e.getFloat64(left)
	if ok {
		floatop = TRUE
		if rf, ok = e.getFloat64(right); ok { // case FF
			return 0, 0, lf, rf, floatop
		}
		if ri, ok = e.getInt64(right); ok { // case FI
			return 0, 0, lf, float64(ri), floatop
		}
		return 0, 0, 0, 0, NEITHER
	}
	// left is not a float at this point, is it an int?
	li, ok = e.getInt64(left)
	if ok {
		ri, ok = e.getInt64(right)
		if ok {
			return li, ri, 0, 0, FALSE // case II
		}
		rf, ok = e.getFloat64(right)
		if ok {
			return 0, 0, float64(li), rf, TRUE // case IF
		}
	}
	return 0, 0, 0, 0, NEITHER
}
func (e *evaluator) Add(left, right any) any {
	li, ri, lf, rf, floatop := e.getPromoted(left, right)
	switch floatop {
	case TRUE:
		return lf + rf
	case FALSE:
		return li + ri
	default:
		// TODO: concat of strings and slices
		panic(catch.Error("+ of non numeric values"))
	}
}
func (e *evaluator) Mul(left, right any) any {
	li, ri, lf, rf, floatop := e.getPromoted(left, right)
	switch floatop {
	case TRUE:
		return lf * rf
	case FALSE:
		return li * ri
	default:
		panic(catch.Error("* of non numeric values"))
	}
}
func (e *evaluator) Div(left, right any) any {
	li, ri, lf, rf, floatop := e.getPromoted(left, right)
	switch floatop {
	case TRUE:
		return lf / rf
	case FALSE:
		return li / ri
	default:
		panic(catch.Error("/ of non numeric values"))
	}
}
func (e *evaluator) Mod(left, right any) any {
	li, ri, _, _, floatop := e.getPromoted(left, right)
	switch floatop {
	case TRUE:
		panic(catch.Error("% of floating point value"))
	case FALSE:
		return li % ri
	default:
		panic(catch.Error("/ of non numeric values"))
	}
}
func (e *evaluator) Sub(left, right any) any {
	li, ri, lf, rf, floatop := e.getPromoted(left, right)
	switch floatop {
	case TRUE:
		return lf - rf
	case FALSE:
		return li - ri
	default:
		panic(catch.Error("- of non numeric values"))
	}
}
func (e *evaluator) ChangeSign(x any) any {
	return e.Sub(0, x)
}
func (e *evaluator) Equal(left, right any) bool {
	li, ri, lf, rf, floatop := e.getPromoted(left, right)
	switch floatop {
	case TRUE:
		return lf == rf
	case FALSE:
		return li == ri
	default:
	}
	switch x := left.(type) {
	case int, int8, int16, int32, int64:
		return false // since getPromoted would have solved this
	case float32, float64:
		return false // since getPromoted would have solved this
	case string:
		rs, ok := right.(string)
		if !ok {
			return false
		}
		return x == rs
	case bool:
		rb, ok := right.(bool)
		if !ok {
			return false
		}
		return x == rb
	case []any:
		rs, ok := right.([]any)
		if !ok {
			return false
		}
		if len(x) != len(rs) {
			return false
		}
		for i := range x {
			if !e.Equal(x[i], rs[i]) {
				return false
			}
		}
		return true
	case nil:
		if right == nil {
			return true
		}
		return false

		// TODO: regexp
	default:
		// TODO: Should be return false, but keep this to detect unexpected precedence.
		panic(catch.Error("== LHS is non equatable value"))
	}
}

func (e *evaluator) Match(left, right any) bool {
	switch matcher := right.(type) {
	case *regexp.Regexp:
		// LHS must be a string
		ls, ok := left.(string)
		if !ok {
			panic(catch.Error("=~ LHS must be a string when RHS is a Regexp"))
		}
		// match against regexp
		return matcher.MatchString(ls)
	case tc.TypeChecker:
		ok, _ := matcher.Check(left)
		return ok
	default:
		return false
	}
}
func (e *evaluator) MatchRegexp(left, right any) []any {
	switch matcher := right.(type) {
	case *regexp.Regexp:
		// LHS must be a string
		ls, ok := left.(string)
		if !ok {
			panic(catch.Error("=~ LHS must be a string when RHS is a Regexp"))
		}
		// match against regexp
		sm := matcher.FindStringSubmatch(ls)
		// Convert to slice of any, since evaluator otherwise also have to match on
		// []string due to go assignability.
		if sm == nil {
			return nil
		}
		sa := make([]any, len(sm))
		for i := range sm {
			sa[i] = sm[i]
		}
		return sa
	default:
		return nil
	}
}
func (e *evaluator) In(left, right any) bool {
	t := reflect.TypeOf(right)
	if t == nil {
		return false
	}
	if t.Kind() == reflect.Slice {
		val := reflect.ValueOf(right)
		for i := 0; i < val.Len(); i++ {
			if e.Equal(left, val.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}
func (e *evaluator) IsTrue(left any) bool {
	if left == nil {
		return false // nil is "falsey"
	}
	lb, ok := left.(bool)
	if !ok {
		panic("value is not a boolean value")
	}
	return lb
}
func (e *evaluator) Not(left any) bool {
	lb, ok := left.(bool)
	if !ok {
		panic("! on non boolean value")
	}
	return !lb
}

func (e *evaluator) Reduce(left any, params []string, expr Expression, start ...any) any {
	if left == nil || reflect.TypeOf(left).Kind() != reflect.Slice {
		left = []any{left}
	}
	var memoName, nextName string
	switch len(params) {
	case 0:
		memoName = "0"
		nextName = "1"
	case 1:
		memoName = params[0]
		nextName = "1"
	case 2:
		memoName, nextName = params[0], params[1]
	}
	hasStart := len(start) > 0
	isFirst := true
	var previous any
	voLeft := reflect.ValueOf(left)
	for i := 0; i < voLeft.Len(); i++ {
		x := voLeft.Index(i).Interface()
		if isFirst {
			if hasStart {
				e.SetVar(memoName, start[0])
				e.SetVar(nextName, x)
				previous = expr.Eval(e)
			} else {
				previous = x
			}
			isFirst = false
			continue
		}
		e.SetVar(memoName, previous)
		e.SetVar(nextName, x)
		previous = expr.Eval(e)
	}
	return previous
}

func (e *evaluator) Var(name string) any {
	result, ok := e.vars[name]
	if !ok {
		// if variable is numeric it is local to this evaluator, else search up parent chain.
		if numericVarPattern.MatchString(name) {
			return nil
		}
		if e.parent != nil {
			return e.parent.Var(name)
		}
		return nil
	}
	return result
}

func (e *evaluator) Property(name string) any {
	// does not have properties, lookup in parent, or give up.
	if e.props == nil {
		if e.parent != nil {
			return e.parent.Property(name)
		}
		panic(catch.Error("attempt to access property in scope that has none"))
	}
	// has properties, do not look up in parent.
	result, ok := e.props[name]
	if !ok {
		return nil
	}
	return result
}

func (e *evaluator) Then(left any, params []string, expr Expression) any {
	if left == nil {
		return nil
	}
	ev := newParentedEvaluator(e)
	var paramName string
	if len(params) > 0 {
		paramName = params[0]
	} else {
		paramName = "0"
	}
	ev.SetVar(paramName, left)
	return expr.Eval(ev)
}

func (e *evaluator) With(left any, params []string, expr Expression) any {
	ev := newParentedEvaluator(e)
	var paramName string
	if len(params) > 0 {
		paramName = params[0]
	} else {
		paramName = "0"
	}
	ev.SetVar(paramName, left)
	return expr.Eval(ev)
}

func (e *evaluator) Lest(left any, expr Expression) any {
	if left != nil {
		return left
	}
	ev := newParentedEvaluator(e)
	return expr.Eval(ev)
}
func (e *evaluator) Map(left any, params []string, expr Expression) any {
	if left == nil || reflect.TypeOf(left).Kind() != reflect.Slice {
		left = []any{left}
	}
	var paramName string
	if len(params) < 1 {
		paramName = "0"
	} else {
		paramName = params[0]
	}
	ev := newParentedEvaluator(e)
	voLeft := reflect.ValueOf(left)
	result := make([]any, voLeft.Len())
	for i := 0; i < voLeft.Len(); i++ {
		x := voLeft.Index(i).Interface()
		ev.SetVar(paramName, x)
		result[i] = expr.Eval(ev)
	}
	return result
}
func (e *evaluator) Filter(left any, params []string, expr Expression) any {
	if left == nil || reflect.TypeOf(left).Kind() != reflect.Slice {
		left = []any{left}
	}
	var paramName string
	if len(params) < 1 {
		paramName = "0"
	} else {
		paramName = params[0]
	}
	ev := newParentedEvaluator(e)
	voLeft := reflect.ValueOf(left)
	result := make([]any, 0, voLeft.Len())
	for i := 0; i < voLeft.Len(); i++ {
		x := voLeft.Index(i).Interface()
		ev.SetVar(paramName, x)
		y := expr.Eval(ev)
		if b, ok := y.(bool); ok {
			if b {
				result = append(result, x)
			}
		} else {
			panic(catch.Error("filter expression did not return a boolean value"))
		}
	}
	return result
}

func (e *evaluator) Count(left any, params []string, expr Expression) any {
	if left == nil || reflect.TypeOf(left).Kind() != reflect.Slice {
		left = []any{left}
	}
	if expr == nil {
		panic(catch.Error("count function requires a lambda expression"))
	}
	var paramName string
	if len(params) < 1 {
		paramName = "0"
	} else {
		paramName = params[0]
	}
	ev := newParentedEvaluator(e)
	voLeft := reflect.ValueOf(left)
	result := 0
	for i := 0; i < voLeft.Len(); i++ {
		x := voLeft.Index(i).Interface()
		ev.SetVar(paramName, x)
		y := expr.Eval(ev)
		if b, ok := y.(bool); ok {
			if b {
				result++
			}
		} else {
			panic(catch.Error("count expression did not return a boolean value"))
		}
	}
	return result
}
func (e *evaluator) Any(left any, params []string, expr Expression) bool {
	if left == nil || reflect.TypeOf(left).Kind() != reflect.Slice {
		left = []any{left}
	}
	var paramName string
	if len(params) < 1 {
		paramName = "0"
	} else {
		paramName = params[0]
	}
	ev := newParentedEvaluator(e)
	voLeft := reflect.ValueOf(left)
	for i := 0; i < voLeft.Len(); i++ {
		x := voLeft.Index(i).Interface()
		ev.SetVar(paramName, x)
		y := expr.Eval(ev)
		if b, ok := y.(bool); ok {
			if b {
				return true
			}
		} else {
			panic(catch.Error("any expression did not return a boolean value"))
		}
	}
	return false
}
func (e *evaluator) All(left any, params []string, expr Expression) bool {
	if left == nil || reflect.TypeOf(left).Kind() != reflect.Slice {
		left = []any{left}
	}
	var paramName string
	if len(params) < 1 {
		paramName = "0"
	} else {
		paramName = params[0]
	}
	ev := newParentedEvaluator(e)
	voLeft := reflect.ValueOf(left)
	length := voLeft.Len()
	if length == 0 {
		return false // empty slice cannot match any predicate
	}
	for i := 0; i < length; i++ {
		x := voLeft.Index(i).Interface()
		ev.SetVar(paramName, x)
		y := expr.Eval(ev)
		if b, ok := y.(bool); ok {
			if !b {
				return false
			}
		} else {
			panic(catch.Error("any expression did not return a boolean value"))
		}
	}
	return true
}
func (e *evaluator) Compare(left, right any) int {
	return utils.ValueOrder(left, right)
}
func (e *evaluator) Compact(left any) any {
	return e.Filter(
		left,
		[]string{},
		&funcexpr{f: func(ev Evaluator) any { return ev.Var("0") != nil }},
	)
}
func (e *evaluator) AllOrNone(left any, params []string, expr Expression) bool {
	if left == nil || reflect.TypeOf(left).Kind() != reflect.Slice {
		left = []any{left}
	}
	count := e.Count(left, params, expr)
	leftLen := e.Len(left)
	return leftLen == 0 || count == leftLen
}

func (e *evaluator) Len(left any) int {
	switch x := left.(type) {
	case string:
		return len(x)
	case []any:
		return len(x)
	default:
		return -1
	}
}

func (e *evaluator) Unique(left any) []any {
	if left == nil || reflect.TypeOf(left).Kind() != reflect.Slice {
		return []any{left}
	}
	hm := utils.NewHashMap()
	leftVo := reflect.ValueOf(left)
	leftLen := leftVo.Len()
	result := make([]any, 0, leftLen)
	for i := 0; i < leftLen; i++ {
		v := leftVo.Index(i).Interface()
		if hm.PutUnique(v, nil) {
			result = append(result, v)
		}
	}
	return result
}

func (e *evaluator) Abs(val any) any {
	f, ok := utils.GetFloat64(val)
	if ok {
		return math.Abs(f)
	}
	i, ok := utils.GetInt64(val)
	if ok {
		if i < 0 {
			return -i
		}
		return i
	}
	panic(catch.Error("abs() got not numeric argument: %v", val))
}

func (e *evaluator) Floor(val any) any {
	f, ok := utils.GetFloat64(val)
	if ok {
		return math.Floor(f)
	}
	i, ok := utils.GetInt64(val)
	if ok {
		return i
	}
	panic(catch.Error("floor() got not numeric argument: %v", val))
}

func (e *evaluator) Ceil(val any) any {
	f, ok := utils.GetFloat64(val)
	if ok {
		return math.Ceil(f)
	}
	i, ok := utils.GetInt64(val)
	if ok {
		return i
	}
	panic(catch.Error("ceil() got not numeric argument: %v", val))
}

func (e *evaluator) Min(left any, right ...any) any {
	if len(right) == 1 {
		if e.Compare(left, right[0]) == 1 {
			return right[0]
		}
		return left
	}
	return e.Reduce(left, []string{},
		&funcexpr{f: func(ev Evaluator) any {
			return e.Min(e.Var("0"), e.Var("1"))
		},
		},
	)
}

func (e *evaluator) Max(left any, right ...any) any {
	if len(right) == 1 {
		if e.Compare(left, right[0]) >= 0 {
			return left
		}
		return right[0]
	}
	return e.Reduce(left, []string{},
		&funcexpr{f: func(ev Evaluator) any {
			return e.Max(e.Var("0"), e.Var("1"))
		},
		},
	)
}
func (e *evaluator) Round(left any) any {
	lf, ok := utils.GetFloat64(left)
	if ok {
		return math.RoundToEven(lf)
	}
	if li, ok := utils.GetInt64(left); ok {
		return li // no-op
	}
	panic(catch.Error("cannot round() non numerical value, %v", left))
}

var numericVarPattern = regexp.MustCompile(`^\d+$`)

type funcexpr struct {
	f func(e Evaluator) any
}

func (f *funcexpr) Eval(e Evaluator) any {
	return f.f(e)
}
func (f *funcexpr) Op() string             { return "func" }
func (f *funcexpr) Children() []Expression { return []Expression{} }
func (f *funcexpr) Literal() any           { return f.f }
