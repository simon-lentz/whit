package utils_test

import (
	"fmt"
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/utils"
)

func Test_AssignIds(t *testing.T) {
	tt := testutils.NewTester(t)
	s := "nothing to replace here"
	s, changed := utils.AssignIds(s, utils.ShortUUID)
	tt.CheckFalse(changed)
	tt.CheckEqual("nothing to replace here", s)

	startval := 0
	mockUUID := func() string {
		startval++
		return fmt.Sprintf("%d", startval)
	}
	s = "abc$$foo def"
	s, changed = utils.AssignIds(s, mockUUID)
	tt.CheckTrue(changed)
	tt.CheckEqual("abc$$:1:foo def", s)
	s, changed = utils.AssignIds(s, mockUUID)
	tt.CheckFalse(changed)
	tt.CheckEqual("abc$$:1:foo def", s)

	startval = 100
	s = "$$a $$b $$a $$b"
	s, changed = utils.AssignIds(s, mockUUID)
	tt.CheckTrue(changed)
	tt.CheckEqual("$$:101:a $$:102:b $$:101:a $$:102:b", s)
}

func Test_ReAssign(t *testing.T) {
	tt := testutils.NewTester(t)
	startval := 0
	mockUUID := func() string {
		startval++
		return fmt.Sprintf("%d", startval)
	}
	// local ids can be used after they have been turned global.
	startval = 0
	s := "$$:101:a $$a"
	s, changed := utils.AssignIds(s, mockUUID)
	tt.CheckTrue(changed)
	tt.CheckEqual("$$:101:a $$:101:a", s)
}
