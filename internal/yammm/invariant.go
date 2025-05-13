package yammm

import (
	"github.com/wyrth-io/whit/internal/validation"
)

// Invariant describes one invariant for a Type. It consists of a Name (for human identification), and
// a Constraint expression.
type Invariant struct {
	// Name is the human readable version of the constraint expression. It is typically used as a message
	// when the constraint is asserted and the assertion fails.
	Name string

	// Constraint is an expression that must evaluate to true for an instance of the type to be valid.
	Constraint Expression
}

// InvariantValidator is used to validate instances of types with invariant constraints.
// Use NewInvariantValidator to create one.
type InvariantValidator struct {
	BaseGraphListener
	evaluator Evaluator
	ic        validation.IssueCollector
}

// NewInvariantValidator returns a new invariant validator.
func NewInvariantValidator(ic validation.IssueCollector) *InvariantValidator {
	return &InvariantValidator{
		evaluator: NewEvaluator(),
		ic:        ic,
	}
}

// Validate validates the instance graph against the given context model's invariant constraints.
func (iv *InvariantValidator) Validate(ctx Context, graph any) {
	walker := NewGraphWalker(ctx, iv)
	walker.Walk(graph)
}

// OnProperties implements this GraphListener interface.
func (iv *InvariantValidator) OnProperties(_ Context, t *Type, props map[string]any) {
	for i := range t.Invariants {
		inv := t.Invariants[i]
		anyResult := inv.Constraint.Eval(iv.evaluator.NewWithProperties(props))
		switch y := anyResult.(type) {
		case bool:
			if !y {
				// TODO: Invariant location
				iv.ic.Collectf(validation.Error, "%s failed invariant assertion: [%s]", t.Name, inv.Name)
			}
			return
		default:
			iv.ic.Collectf(validation.Error,
				"invariant %s.[%s] expected a boolean result: got '%t', '%v'",
				t.Name, inv.Name, anyResult, anyResult,
			)
		}
	}
}
