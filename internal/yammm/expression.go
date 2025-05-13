package yammm

import (
	"fmt"
	"reflect"

	"github.com/tada/catch"
	"github.com/wyrth-io/whit/internal/tc"
	"github.com/wyrth-io/whit/internal/utils"
)

// Expression is an interface for expressions. See SExpr, Literal, and Op.
type Expression interface {
	// Children returns all children except the Operation.
	Children() []Expression

	// Op returns the operation name.
	Op() string

	// Eval evaluates the operation by calling the given evaluator.
	Eval(evaluator Evaluator) any

	// Literal returns a literal value (for literal value expressions).
	Literal() any
}

// SExpr is an Expression.
type SExpr []Expression

// Literal is an Expression.
type literalp struct {
	Val any
}

// NewLiteral returns a new Expression for a literal value.
func NewLiteral(val any) Expression {
	if _, ok := val.(*literalp); ok {
		panic("literal in literal")
	}
	return &literalp{Val: val}
}

// Op is an Expression used to describe the name of an operation.
type Op string

// DatatypeLiteral is an Expression used to describe a datatype (unconstrained).
type DatatypeLiteral string

// Op implements this Expression function.
func (o Op) Op() string { return string(o) }

// Children implements this Expression function.
func (Op) Children() []Expression { return []Expression{} }

// Literal implements this Expression function.
func (o Op) Literal() any { return string(o) }

// Eval implements this Expression function.
func (o Op) Eval(_ Evaluator) any { return nil }

// Op implements this Expression function.
func (*literalp) Op() string { return "lit" }

// Children implements this Expression function.
func (*literalp) Children() []Expression { return []Expression{} }

// Literal implements this Expression function.
func (lit *literalp) Literal() any { return lit.Val }

// Eval implements this Expression function.
func (lit *literalp) Eval(_ Evaluator) any {
	return lit.Val
}

// Op implements this Expression function.
func (DatatypeLiteral) Op() string { return "dt" }

// Children implements this Expression function.
func (DatatypeLiteral) Children() []Expression { return []Expression{} }

// Literal implements this Expression function.
func (dt DatatypeLiteral) Literal() any { return string(dt) }

// Eval implements this Expression function.
func (dt DatatypeLiteral) Eval(_ Evaluator) any { return tc.NewTypeChecker([]string{string(dt)}) }

// Op implements this Expression function.
func (e SExpr) Op() string { return e[0].Literal().(string) }

// Children implements this Expression function.
func (e SExpr) Children() []Expression { return e[1:] }

// Literal implements this Expression function.
func (e SExpr) Literal() any { return e.Op() }

// Eval implements this Expression function.
func (e SExpr) Eval(evaluator Evaluator) any { //nolint:gocyclo
	children := e.Children()
	switch e.Op() {
	case "+":
		return evaluator.Add(children[0].Eval(evaluator), children[1].Eval(evaluator))
	case "-":
		return evaluator.Sub(children[0].Eval(evaluator), children[1].Eval(evaluator))
	case "-x": // unary minus
		return evaluator.ChangeSign(children[0].Eval(evaluator))
	case "*":
		return evaluator.Mul(children[0].Eval(evaluator), children[1].Eval(evaluator))
	case "/":
		return evaluator.Div(children[0].Eval(evaluator), children[1].Eval(evaluator))
	case "%":
		return evaluator.Mod(children[0].Eval(evaluator), children[1].Eval(evaluator))
	case "^":
		lhs := evaluator.IsTrue(children[0].Eval(evaluator))
		rhs := evaluator.IsTrue(children[1].Eval(evaluator))
		return lhs != rhs
	case "?":
		if evaluator.IsTrue(children[0].Eval(evaluator)) {
			return children[1].Eval(evaluator)
		} else if len(children) == 3 {
			return children[2].Eval(evaluator)
		}
		return nil // no else part and was not true
	case "!":
		return !evaluator.IsTrue(children[0].Eval(evaluator))
	case "&&":
		if evaluator.IsTrue(children[0].Eval(evaluator)) {
			return evaluator.IsTrue(children[1].Eval(evaluator))
		}
		return false
	case "||":
		if evaluator.IsTrue(children[0].Eval(evaluator)) {
			return true
		}
		return evaluator.IsTrue(children[1].Eval(evaluator))

	case "==":
		return evaluator.Equal(children[0].Eval(evaluator), children[1].Eval(evaluator))
	case "!=":
		return !evaluator.Equal(children[0].Eval(evaluator), children[1].Eval(evaluator))
	case "=~":
		return evaluator.Match(children[0].Eval(evaluator), children[1].Eval(evaluator))
	case "!~":
		return !evaluator.Match(children[0].Eval(evaluator), children[1].Eval(evaluator))
	case "<":
		return evaluator.Compare(children[0].Eval(evaluator), children[1].Eval(evaluator)) == -1
	case "<=":
		return evaluator.Compare(children[0].Eval(evaluator), children[1].Eval(evaluator)) <= 0
	case ">":
		return evaluator.Compare(children[0].Eval(evaluator), children[1].Eval(evaluator)) == 1
	case ">=":
		return evaluator.Compare(children[0].Eval(evaluator), children[1].Eval(evaluator)) >= 0
	case "in":
		return evaluator.In(children[0].Eval(evaluator), children[1].Eval(evaluator))
	case "$":
		name := children[0].Literal()
		return evaluator.Var(name.(string))
	case "p":
		// property name of object set in the evaluator
		name := children[0].Literal()
		return evaluator.Property(name.(string))
	case ".":
		// lhs should be an object with properties (for now a map[string]any).
		lhs := children[0].Eval(evaluator)
		name := children[1].Literal().(string)
		if props, ok := lhs.(map[string]any); ok {
			return evaluator.NewWithProperties(props).Property(name)
		}
		panic(catch.Error("property canot be extracted from lhs value, got: %v", lhs))

	case "reduce":
		lhs, args, params, body := e.getLHSArgParams(evaluator, children, 0, 1, 2, true)
		return evaluator.Reduce(lhs, params, body, args...)

	case "then":
		lhs, _, params, body := e.getLHSArgParams(evaluator, children, 0, 0, 1, true)
		return evaluator.Then(lhs, params, body)
	case "lest":
		lhs, _, _, body := e.getLHSArgParams(evaluator, children, 0, 0, 1, true)
		return evaluator.Lest(lhs, body)
	case "with":
		lhs, _, params, body := e.getLHSArgParams(evaluator, children, 0, 0, 1, true)
		return evaluator.With(lhs, params, body)

	case "filter":
		lhs, _, params, body := e.getLHSArgParams(evaluator, children, 0, 0, 1, true)
		return evaluator.Filter(lhs, params, body)

	case "map":
		lhs, _, params, body := e.getLHSArgParams(evaluator, children, 0, 0, 1, true)
		return evaluator.Map(lhs, params, body)

	case "compact":
		return evaluator.Compact(children[0].Eval(evaluator))
	case "unique":
		return evaluator.Unique(children[0].Eval(evaluator))
	case "len":
		return evaluator.Len(children[0].Eval(evaluator))
	case "count":
		lhs, _, params, body := e.getLHSArgParams(evaluator, children, 0, 0, 1, true)
		return evaluator.Count(lhs, params, body)
	case "all":
		lhs, _, params, body := e.getLHSArgParams(evaluator, children, 0, 0, 1, true)
		return evaluator.All(lhs, params, body)
	case "all_or_none":
		lhs, _, params, body := e.getLHSArgParams(evaluator, children, 0, 0, 1, true)
		return evaluator.AllOrNone(lhs, params, body)
	case "any":
		lhs, _, params, body := e.getLHSArgParams(evaluator, children, 0, 0, 1, true)
		return evaluator.Any(lhs, params, body)
	case "compare":
		lhs, args, _, _ := e.getLHSArgParams(evaluator, children, 1, 1, 0, false)
		return evaluator.Compare(lhs, args[0])
	case "[]":
		result := make([]any, len(children))
		for i := range children {
			result[i] = children[i].Eval(evaluator)
		}
		return result
	case "@":
		// Slice is more complex since it handles values, relationship and type name as LHS.
		return e.Slice(evaluator)
	case "abs":
		lhs, _, _, _ := e.getLHSArgParams(evaluator, children, 0, 0, 0, false)
		return evaluator.Abs(lhs)
	case "floor":
		lhs, _, _, _ := e.getLHSArgParams(evaluator, children, 0, 0, 0, false)
		return evaluator.Floor(lhs)
	case "ceil":
		lhs, _, _, _ := e.getLHSArgParams(evaluator, children, 0, 0, 0, false)
		return evaluator.Ceil(lhs)
	case "round":
		lhs, _, _, _ := e.getLHSArgParams(evaluator, children, 0, 0, 0, false)
		return evaluator.Round(lhs)
	case "min":
		lhs, args, _, _ := e.getLHSArgParams(evaluator, children, 0, 1, 0, false)
		return evaluator.Min(lhs, args...)
	case "max":
		lhs, args, _, _ := e.getLHSArgParams(evaluator, children, 0, 1, 0, false)
		return evaluator.Max(lhs, args...)
	case "match":
		lhs, args, _, _ := e.getLHSArgParams(evaluator, children, 1, 1, 0, false)
		return evaluator.MatchRegexp(lhs, args[0])
	}
	panic(catch.Error("unknown operand %s", e.Op()))
}

// getLhsArgsParams extracts the values to feed into an Evaluator call.
func (e *SExpr) getLHSArgParams(
	evaluator Evaluator,
	children []Expression,
	minArgs, maxArgs,
	maxParams int,
	acceptBody bool,
) (lhs any, args []any, params []string, body Expression) {
	// LHS is 0
	lhs = children[0].Eval(evaluator)

	// Args is 1 - may be nil or empty slice. Check min/max allowed and evaluate them.
	argCount := 0
	argsChild := children[1]
	if argsChild != nil {
		argExprsAny := argsChild.Literal().([]any)
		argCount = len(argExprsAny)
		args = make([]any, argCount)
		for i := range argExprsAny {
			args[i] = argExprsAny[i].(Expression).Eval(evaluator)
		}
	}
	if argCount < minArgs || argCount > maxArgs {
		if minArgs == maxArgs {
			panic(catch.Error("function accepts %d argument(s), got %d", minArgs, argCount))
		}
		panic(catch.Error("function takes %d to %d argument(s), got %d", minArgs, maxArgs, argCount))
	}

	// Params is 2 - may be nil or empty slice
	paramsChild := children[2]
	if paramsChild != nil {
		voParams := reflect.ValueOf(paramsChild.Literal())
		for i := 0; i < voParams.Len(); i++ {
			params = append(params, voParams.Index(i).Interface().(string))
		}
		if len(params) > maxParams {
			panic(catch.Error("number of parameters exceeds accepted %d: got %d", maxParams, len(params)))
		}
	}

	// Body is 3 - may be nil or Expression
	bodyChild := children[3]
	if bodyChild != nil {
		if !acceptBody {
			panic(catch.Error("function does not accept a lambda expression"))
		}
		body = bodyChild
	}
	return lhs, args, params, body
}

// ResolveIdentifier returns a type checker instance for all UC names recognized as
// the name of a data type, otherwise the name is the name of a relation (which is TODO:).
func (e *SExpr) ResolveIdentifier(evaluator Evaluator) any {
	children := e.Children()
	idLiteral := children[0]
	// identifier, if it is a type name then slice the type, else it is a relationship property
	identifier := idLiteral.Children()[0].Literal().(string)
	if e.isDataType(identifier) {
		var val any
		// TypeChecker is the closest thing, but acts on strings.
		// TODO: Make this better. For now transform the data to string form instructions
		instructions := make([]string, len(children))
		for i := range children {
			if i == 0 {
				instructions[i] = identifier
			} else {
				elem := children[i].Eval(evaluator)
				// if nil use "_", otherwise format strings and numbers
				if elem == nil {
					instructions[i] = "_"
					continue
				}
				switch data := elem.(type) {
				case int, int8, int16, int32, int64:
					instructions[i] = fmt.Sprintf("%d", data)
				case float32, float64:
					instructions[i] = fmt.Sprintf("%f", data)
				case string:
					instructions[i] = data
				default:
					panic("illegal value in type constraint")
				}
			}
		}
		val = tc.NewTypeChecker(instructions)
		if val == nil {
			panic(fmt.Sprintf("cannot create a type checker from given instructions %v", instructions))
		}
		return val // done, no further slicing required.
	}
	// A relation property
	// TODO: get relationship from context instance
	return []any{"this", "should", "be", "replaced"}
}

// Slice evaluates the s expression (@ lhs at1 at2 ...).
func (e *SExpr) Slice(evaluator Evaluator) any {
	children := e.Children()
	values := make([]any, len(children))
	for i := range children {
		values[i] = children[i].Eval(evaluator)
	}
	// type checker has its own "slicer"
	if tc, ok := values[0].(tc.TypeChecker); ok {
		return tc.Refine(values[1:])
	}
	// LHS value is now sliceable. Must have one value (get one element), but accepts two.
	// to specify a range. If second value is -1 means (all of the lhs).
	var from, to int64
	var isRange bool
	var ok bool
	if len(children) == 3 {
		isRange = true
		from, ok = utils.GetInt64(values[1])
		if !ok {
			panic("not an integer")
		}
		to, ok = utils.GetInt64(values[2])
		if !ok {
			to = -1
		}
	} else {
		from, ok = utils.GetInt64(values[1])
		if !ok {
			panic("not an integer")
		}
	}

	switch lhsVal := values[0].(type) {
	case string:
		// expect 1 or two expressions (start, end)
		switch {
		case isRange && to != -1:
			return lhsVal[from:to]
		case isRange:
			return lhsVal[from:]
		default:
			return lhsVal[from]
		}
	case []any:
		switch {
		case isRange && to != -1:
			return lhsVal[from:to]
		case isRange:
			return lhsVal[from:]
		default:
			return lhsVal[from]
		}
	default:
		panic(fmt.Sprintf("don't know how to slice lhs value of type %t", lhsVal))
	}
}

// // Extracts the UC identifier from an "I" SExpr
//
//	func (e *SExpr) extractIdentifier() string {
//		x := e.Children()[1].Literal()
//		return x.(string)
//	}
func (e *SExpr) isDataType(name string) bool {
	switch name {
	case "Integer", "Float", "Boolean", "String", "Enum", "Date", "Timestamp", "Pattern", "UUID":
		return true
	default:
		return false
	}
}
