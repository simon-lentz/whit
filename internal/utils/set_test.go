package utils_test

import (
	"sort"
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/utils"
)

func Test_Set_Add(t *testing.T) {
	tt := testutils.NewTester((t))
	set := utils.NewSet[string]()
	tt.CheckTrue(set.Add("a"))
	tt.CheckTrue(set.Add("b"))
	tt.CheckFalse(set.Add("b"))
}
func Test_Set_Remove(t *testing.T) {
	tt := testutils.NewTester((t))
	set := utils.NewSet[string]()
	tt.CheckTrue(set.Add("a"))
	tt.CheckTrue(set.Add("b"))
	set2 := set.Remove("b")
	tt.CheckTrue(set.Contains("b"))
	tt.CheckFalse(set2.Contains("b"))
	tt.CheckTrue(set.Contains("a"))
}
func Test_Set_New_with_values(t *testing.T) {
	tt := testutils.NewTester((t))
	set := utils.NewSet(1, 2, 3, 4)
	tt.CheckTrue(set.Contains(3))
}

func Test_Set_Contains(t *testing.T) {
	tt := testutils.NewTester((t))
	set := utils.NewSet[string]()
	set.Add("a")
	set.Add("b")
	tt.CheckTrue(set.Contains("a"))
	tt.CheckTrue(set.Contains("b"))
	tt.CheckFalse(set.Contains("c"))
}

func Test_Set_Each(t *testing.T) {
	tt := testutils.NewTester((t))
	set := utils.NewSet[string]()
	set.Add("a")
	set.Add("b")
	set.Add("b")
	var count int
	seen := make(map[string]bool, 2)
	set.Each(func(val string) {
		count++
		seen[val] = true
	})
	tt.CheckEqual(2, count)               // called twice
	tt.CheckTrue(utils.HasKey(seen, "a")) // once with a
	tt.CheckTrue(utils.HasKey(seen, "b")) // once with b
}

func Test_Set_Union(t *testing.T) {
	tt := testutils.NewTester((t))
	set := utils.NewSet[string]()
	set.Add("a")
	set.Add("b")
	set2 := utils.NewSet[string]()
	set2.Add("b")
	set2.Add("c")
	var count int
	actual := set.Union(set2)
	seen := make(map[string]bool, 2)
	actual.Each(func(val string) {
		count++
		seen[val] = true
	})
	tt.CheckEqual(3, count)               // called twice
	tt.CheckTrue(utils.HasKey(seen, "a")) // once with a
	tt.CheckTrue(utils.HasKey(seen, "b")) // once with b
	tt.CheckTrue(utils.HasKey(seen, "b")) // once with b
}

func Test_NewSetFrom(t *testing.T) {
	tt := testutils.NewTester(t)
	slice := []string{"a", "b"}
	set := utils.NewSetFrom(slice, func(s string) string {
		return s
	})
	tt.CheckTrue(set.Contains(("a")))
	tt.CheckTrue(set.Contains(("b")))
}

func Test_Set_Size(t *testing.T) {
	tt := testutils.NewTester(t)
	set := utils.NewSet[string]()
	set.Add("a")
	set.Add("b")
	set.Add("b")
	tt.CheckEqual(2, set.Size())
}

func Test_Set_Slices(t *testing.T) {
	tt := testutils.NewTester(t)
	set := utils.NewSet[string]()
	set.Add("a")
	set.Add("b")
	set.Add("b")
	slices := set.Slices()
	sort.Slice(slices, func(i, j int) bool { return slices[i] < slices[j] })
	tt.CheckEqual([]string{"a", "b"}, slices)
}

func Test_Set_Intersection(t *testing.T) {
	tt := testutils.NewTester(t)
	set := utils.NewSet[string]()
	set2 := utils.NewSet[string]()
	set.Add("a")
	set.Add("b")
	set.Add("c")
	set2.Add("x")
	set2.Add("y")
	set2.Add("a")
	set2.Add("b")
	result := set.Intersection(set2)
	tt.CheckEqual(2, result.Size())
	tt.CheckTrue(result.Contains("a"))
	tt.CheckTrue(result.Contains("b"))
}
func Test_Set_Empty(t *testing.T) {
	tt := testutils.NewTester(t)
	set := utils.NewSet[string]()
	tt.CheckTrue(set.Empty())
	set.Add("a")
	tt.CheckFalse(set.Empty())
}

func Test_Set_Equal(t *testing.T) {
	tt := testutils.NewTester(t)
	set := utils.NewSet[string]()
	set.Add("a")
	set.Add("b")

	set2 := utils.NewSet[string]()
	set2.Add("a")
	set2.Add("b")
	tt.CheckTrue(set.Equal(set2))
}

func Test_Set_String(t *testing.T) {
	tt := testutils.NewTester(t)
	set := utils.NewSet[string]()
	set.Add("a")
	set.Add("b")
	actual := set.String()
	// Order is unspecified
	tt.CheckTrue(actual == "Set[a,b,]" || actual == "Set[b,a,]")
}

func Test_Set_Diff(t *testing.T) {
	tt := testutils.NewTester(t)
	set := utils.NewSet(1, 2, 3)
	set2 := utils.NewSet(2, 3)
	tt.CheckTrue(utils.NewSet(1).Equal(set.Diff(set2)))
}
