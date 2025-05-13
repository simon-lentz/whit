package utils_test

import (
	"sort"
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/utils"
)

func Test_HasKey(t *testing.T) {
	tt := testutils.NewTester(t)
	tt.CheckTrue(utils.HasKey(map[string]int{"a": 1, "b": 2}, "a"))
	tt.CheckFalse(utils.HasKey(map[string]int{"a": 1, "b": 2}, "c"))
}
func Test_Keys(t *testing.T) {
	type Something interface {
		Blah() string
	}
	tt := testutils.NewTester(t)
	h := map[int]Something{1: nil, 2: nil, 3: nil}
	keys := utils.Keys(h)
	sort.Ints(keys)
	tt.CheckEqual([]int{1, 2, 3}, keys)
}

func Test_Values(t *testing.T) {
	tt := testutils.NewTester(t)
	h := map[int]string{1: "a", 2: "b", 3: "c"}
	vals := utils.Values(h)
	sort.Strings(vals)
	tt.CheckEqual([]string{"a", "b", "c"}, vals)
}

func Test_FilterMap(t *testing.T) {
	tt := testutils.NewTester(t)
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := utils.FilterMap(m, func(_ string, _ int) bool { return true })
	tt.CheckEqual(m, m2)

	m2 = utils.FilterMap(m, func(k string, _ int) bool { return k == "b" })
	tt.CheckEqual(1, len(m2))
	tt.CheckEqual(2, m2["b"])
}
