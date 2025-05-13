package utils_test

import (
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/utils"
)

func TestValueStrata(t *testing.T) {
	tt := testutils.NewTester(t)
	tt.CheckEqual(utils.NilStrata, utils.TypeStrata(nil))
	tt.CheckEqual(utils.BoolStrata, utils.TypeStrata(true))
	tt.CheckEqual(utils.BoolStrata, utils.TypeStrata(false))
	tt.CheckEqual(utils.NumericStrata, utils.TypeStrata(1))
	tt.CheckEqual(utils.NumericStrata, utils.TypeStrata(3.14))
	tt.CheckEqual(utils.StringStrata, utils.TypeStrata("hola"))
	tt.CheckEqual(utils.SliceStrata, utils.TypeStrata([]string{"hola"}))
}
func TestTypeOrder(t *testing.T) {
	tt := testutils.NewTester(t)
	tt.CheckEqual(0, utils.TypeOrder(nil, nil))
	tt.CheckEqual(1, utils.TypeOrder(false, nil))
	tt.CheckEqual(-1, utils.TypeOrder(nil, true))

	tt.CheckEqual(0, utils.TypeOrder(false, false))
	tt.CheckEqual(0, utils.TypeOrder(true, true))
	tt.CheckEqual(0, utils.TypeOrder(false, true))
	tt.CheckEqual(0, utils.TypeOrder(true, false))

	tt.CheckEqual(0, utils.TypeOrder(1, 1))
	tt.CheckEqual(0, utils.TypeOrder(1, 3.14))
	tt.CheckEqual(0, utils.TypeOrder(3.14, 1))
	tt.CheckEqual(0, utils.TypeOrder(1.12, 3.14))
	tt.CheckEqual(-1, utils.TypeOrder(false, 1))
	tt.CheckEqual(-1, utils.TypeOrder(true, 1.2))
	tt.CheckEqual(1, utils.TypeOrder(1, false))
	tt.CheckEqual(1, utils.TypeOrder(1, true))

	tt.CheckEqual(0, utils.TypeOrder("a", "b"))
	tt.CheckEqual(1, utils.TypeOrder("a", 1))
	tt.CheckEqual(-1, utils.TypeOrder(1, "a"))

	tt.CheckEqual(0, utils.TypeOrder([]string{"a"}, []int{1}))
	tt.CheckEqual(1, utils.TypeOrder([]int{1}, 1))
	tt.CheckEqual(-1, utils.TypeOrder(1, []int{1}))
}
func TestValueOrder(t *testing.T) {
	tt := testutils.NewTester(t)
	tt.CheckEqual(0, utils.ValueOrder(nil, nil))
	tt.CheckEqual(1, utils.ValueOrder(false, nil))
	tt.CheckEqual(-1, utils.ValueOrder(nil, true))

	tt.CheckEqual(0, utils.ValueOrder(false, false))
	tt.CheckEqual(0, utils.ValueOrder(true, true))
	tt.CheckEqual(-1, utils.ValueOrder(false, true))
	tt.CheckEqual(1, utils.ValueOrder(true, false))

	tt.CheckEqual(0, utils.ValueOrder(1, 1))
	tt.CheckEqual(-1, utils.ValueOrder(1, 3.14))
	tt.CheckEqual(1, utils.ValueOrder(3.14, 1))
	tt.CheckEqual(-1, utils.ValueOrder(1.12, 3.14))
	tt.CheckEqual(-1, utils.ValueOrder(false, 1))
	tt.CheckEqual(-1, utils.ValueOrder(true, 1.2))
	tt.CheckEqual(1, utils.ValueOrder(1, false))
	tt.CheckEqual(1, utils.ValueOrder(1, true))

	tt.CheckEqual(-1, utils.ValueOrder("a", "b"))
	tt.CheckEqual(1, utils.ValueOrder("a", 1))
	tt.CheckEqual(-1, utils.ValueOrder(1, "a"))

	tt.CheckEqual(1, utils.ValueOrder([]string{"a"}, []int{1}))
	tt.CheckEqual(1, utils.ValueOrder([]int{1}, 1))
	tt.CheckEqual(-1, utils.ValueOrder(1, []int{1}))

	tt.CheckEqual(1, utils.ValueOrder([]any{"a", "b", "c"}, []any{"a", "b"}))
	tt.CheckEqual(1, utils.ValueOrder([]any{"a", "b", 1}, []any{"a", "b"}))
	tt.CheckEqual(-1, utils.ValueOrder([]any{"a", "b", 1}, []any{"a", "b", "c"}))
}
