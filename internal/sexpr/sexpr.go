package sexpr

import (
	"fmt"
	"regexp"

	"github.com/wyrth-io/whit/internal/tc"
	"github.com/wyrth-io/whit/internal/utils"
)

// Evaluator is an interface for evaluation of calls, getting a literal value or symbol.
type Evaluator interface {
	Call(op string, args []any) any
	Literal(Value) any
	Symbol(Symbol) any
}

// Value is an interface for elements of an S-expression (atom or list).
type Value interface {
	Literal() any
	Eval(evaluator Evaluator) any
}

// Symbol is a symbolic name. (As opposed to a literal string).
type Symbol interface {
	Value
	Literal() any
	String() string
}

// List is a list of Value.
type List interface {
	Value
	First() Value
	Rest() List
	Empty() bool
	Elements() []Value
}
type list []Value
type symbol string

func (s symbol) Literal() any {
	return string(s)
}
func (s symbol) String() string {
	return string(s)
}
func (s symbol) Eval(e Evaluator) any {
	return e.Symbol(s)
}

// NewSymbol returns a new Symbol.
func NewSymbol(s string) Symbol {
	return symbol(s)
}

var emptyList = list{}

// NewList returns a new List containing all the given values. Native values are wrapped
// in specialized types implementing the expected interfaces.
func NewList(elements ...any) List {
	result := make([]Value, len(elements))
	for i := range elements {
		element := elements[i]
		if element == nil {
			result[i] = nilValue{}
		}
		if x, ok := utils.GetInt64(element); ok {
			result[i] = intValue(x)
			continue
		}
		if x, ok := utils.GetFloat64(element); ok {
			result[i] = floatValue(x)
			continue
		}
		switch x := element.(type) {
		case tc.TypeChecker:
		case string:
			result[i] = stringValue(x)
		case bool:
			result[i] = boolValue(x)
		case Value:
			result[i] = x
		case *regexp.Regexp:
			result[i] = regexpValue{x}
		default:
			result[i] = anyValue{x}
		}
	}
	return list(result)
}
func (li list) First() Value {
	return li[0]
}

func (li list) Rest() List {
	if len(li) < 2 {
		return emptyList
	}
	return li[1:]
}
func (li list) Literal() any {
	return li
}
func (li list) Empty() bool {
	return len(li) == 0
}

// Eval evaluates the list if the first element is a Symbol, otherwise the
// list is returned verbatim.
func (li list) Eval(e Evaluator) any {
	if !li.Empty() {
		switch x := li[0].(type) {
		case Symbol:
			args := make([]any, len(li)-1)
			rest := li.Rest().Elements()
			for i := range rest {
				args[i] = rest[i].Eval(e)
			}
			return e.Call(x.String(), args)
		default:
			return e.Literal(li)
		}
	}
	return li
}

// Returns this as a slice of Value.
func (li list) Elements() []Value {
	return li
}

type stringValue string
type floatValue float64
type intValue int64
type boolValue bool
type regexpValue [1]*regexp.Regexp
type anyValue [1]any
type nilValue [0]any

func (i intValue) Literal() any            { return int64(i) }
func (i intValue) Eval(e Evaluator) any    { return e.Literal(i) }
func (f floatValue) Literal() any          { return float64(f) }
func (f floatValue) Eval(e Evaluator) any  { return e.Literal(f) }
func (s stringValue) Literal() any         { return string(s) }
func (s stringValue) Eval(e Evaluator) any { return e.Literal(s) }
func (b boolValue) Literal() any           { return bool(b) }
func (b boolValue) Eval(e Evaluator) any   { return e.Literal(b) }
func (r regexpValue) Literal() any         { return r[0] }
func (r regexpValue) Eval(e Evaluator) any { return e.Literal(r) }
func (a anyValue) Literal() any            { return a[0] }
func (a anyValue) Eval(e Evaluator) any    { return e.Literal(a) }
func (n nilValue) Literal() any            { return nil }
func (n nilValue) Eval(e Evaluator) any    { return e.Literal(nil) }

// BaseEvaluator implements the Evaluator interface and can be included in
// implementations. The BaseEvaluator returns literals to their literal value, but
// nothing else.
type BaseEvaluator struct{}

// Call implements this Evaluator method.
func (e *BaseEvaluator) Call(_ string, _ []any) any {
	return nil
}

// Literal returns the result of calling Literal() on the given value.
func (e *BaseEvaluator) Literal(x Value) any {
	return x.Literal()
}

// Symbol returns the symbol as a literal data type value.
func (e *BaseEvaluator) Symbol(x Symbol) any {
	return x.Literal()
}

// PluggableEvaluator allows adding operators to the evaluator.
type PluggableEvaluator struct {
	BaseEvaluator
	funcs map[string]func(string, []any) any
}

// Call calls the function (added with AddOp) for the given operator and passes the
// operator and the arguments to the function.
func (e *PluggableEvaluator) Call(op string, args []any) any {
	if e.funcs != nil {
		if f, ok := e.funcs[op]; ok {
			return f(op, args)
		}
	}
	panic(fmt.Errorf("undefined operation: '%s'", op))
}

// AddOp adds an operator function to this evaluator.
func (e *PluggableEvaluator) AddOp(op string, f func(string, []any) any) {
	if e.funcs == nil {
		e.funcs = map[string]func(string, []any) any{}
	}
	e.funcs[op] = f
}
