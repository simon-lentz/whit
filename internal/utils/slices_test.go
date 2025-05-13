package utils_test

import (
	"strings"
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/utils"
)

func Test_Filter(t *testing.T) {
	tt := testutils.NewTester(t)
	slices := []string{"a", "b", "c", "b"}
	actual := utils.Filter(slices, func(s string) bool {
		return s == "b"
	})
	tt.CheckEqual([]string{"b", "b"}, actual)
}

func Test_All(t *testing.T) {
	tt := testutils.NewTester(t)
	slices := []string{"a", "b"}
	actual := utils.All(slices, func(s string) bool {
		return s == "b" || s == "a"
	})
	tt.CheckTrue(actual)

	actual = utils.All(slices, func(s string) bool {
		return s == "b"
	})
	tt.CheckFalse(actual)

	slices = []string{}
	actual = utils.All(slices, func(s string) bool { return true })
	tt.CheckFalse(actual)
}

func Test_Any(t *testing.T) {
	tt := testutils.NewTester(t)
	slices := []string{"a", "b"}
	actual := utils.Any(slices, func(s string) bool {
		return s == "b"
	})
	tt.CheckTrue(actual)

	actual = utils.Any(slices, func(s string) bool {
		return s == "c"
	})
	tt.CheckFalse(actual)

	slices = []string{}
	actual = utils.Any(slices, func(s string) bool { return true })
	tt.CheckFalse(actual)
}

func Test_Map(t *testing.T) {
	tt := testutils.NewTester(t)
	slices := []string{"a", "b"}
	actual := utils.Map(slices, strings.ToUpper)
	tt.CheckEqual([]string{"A", "B"}, actual)

	slices = []string{}
	actual = utils.Map(slices, strings.ToUpper)
	tt.CheckEqual([]string{}, actual)
}

func Test_Reduce(t *testing.T) {
	tt := testutils.NewTester(t)
	slices := []int{1, 2, 3}
	sum := func(i, prev int) int { return i + prev }
	actual := utils.Reduce(slices, 0, sum)
	tt.CheckEqual(6, actual)

	slices = []int{}
	actual = utils.Reduce(slices, 0, sum)
	tt.CheckEqual(0, actual)
}
func Test_Index(t *testing.T) {
	tt := testutils.NewTester(t)
	slices := []int{1, 2, 3}
	idx := utils.Index(slices, func(x int) bool { return x == 2 })
	tt.CheckEqual(1, idx)

	idx = utils.Index(slices, func(x int) bool { return x == 100 })
	tt.CheckEqual(-1, idx)
}

func Test_Find(t *testing.T) {
	tt := testutils.NewTester(t)
	slices := []int{1, 2, 3}
	found := utils.Find(slices, func(x int) bool { return x == 2 })
	tt.CheckEqual(2, *found)
	*found = 20
	found = utils.Find(slices, func(x int) bool { return x == 20 })
	tt.CheckNotNil(found)
	tt.CheckEqual(20, *found)
	found = utils.Find(slices, func(x int) bool { return x == 100 })
	tt.CheckNil(found)
}

func Test_PtrToLast(t *testing.T) {
	tt := testutils.NewTester(t)
	slices := []int{1, 2, 3}
	var y *int
	x := utils.PtrToLastSlice(slices)
	y = x // this looks strange, but is here to assert that x is actually a pointer to int.
	tt.CheckEqual(3, *x)
	tt.CheckEqual(3, *y)
	*x = 4
	tt.CheckEqual(4, slices[2])
}
