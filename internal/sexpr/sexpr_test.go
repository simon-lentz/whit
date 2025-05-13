package sexpr_test

import (
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/sexpr"
	"github.com/wyrth-io/whit/internal/utils"
)

func TestList(t *testing.T) {
	tt := testutils.NewTester(t)
	li := sexpr.NewList(1, 2, 3)
	var results []int64
	for x := li; !x.Empty(); x = x.Rest() {
		v := x.First()
		results = append(results, v.Literal().(int64))
	}
	tt.CheckEqual([]int64{1, 2, 3}, results)
}

func TestEvalSExpression(t *testing.T) {
	tt := testutils.NewTester(t)
	li := sexpr.NewList(sexpr.NewSymbol("+"), 1, 2)
	e := sexpr.PluggableEvaluator{}
	e.AddOp("+", func(_ string, args []any) any {
		lhs, _ := utils.GetInt64(args[0])
		rhs, _ := utils.GetInt64(args[1])
		return lhs + rhs
	})
	actual := li.Eval(&e)
	tt.CheckEqual(3, actual)
}
