package utils_test

import (
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/utils"
)

func TestHashMap(t *testing.T) {
	tt := testutils.NewTester(t)
	hm := utils.NewHashMap()
	hm.Put([]any{1, 2, 3}, "yay")
	hm.Put([]any{3, 2, 1}, "aya")
	tt.CheckEqual("yay", hm.Get([]any{1, 2, 3}))
	tt.CheckEqual("aya", hm.Get([]any{3, 2, 1}))
}
func TestHashMap_PutUnique(t *testing.T) {
	tt := testutils.NewTester(t)
	hm := utils.NewHashMap()
	wasUnique := hm.PutUnique("key", "value")
	tt.CheckTrue(wasUnique)
	wasUnique = hm.PutUnique("key", "othervalue")
	tt.CheckFalse(wasUnique)
	tt.CheckEqual("value", hm.Get("key")) // value did not change
}

func TestHashMap_GetOk(t *testing.T) {
	tt := testutils.NewTester(t)
	hm := utils.NewHashMap()
	hm.Put("key", "value")
	v, ok := hm.GetOk("key")
	tt.CheckTrue(ok)
	tt.CheckEqual("value", v)
	v, ok = hm.GetOk("notkey")
	tt.CheckFalse(ok)
	tt.CheckNil(v)
}
func TestHashMap_Get_NonExistingIsNil(t *testing.T) {
	tt := testutils.NewTester(t)
	hm := utils.NewHashMap()
	v := hm.Get("key")
	tt.CheckNil(v)
}
