package parser //nolint:all

import (
	"testing"

	"github.com/hlindberg/testutils"
)

func Test_stripDelimiters(t *testing.T) {
	tt := testutils.NewTester(t)
	result := stripDelimiters("/* abc */")
	tt.CheckEqual("abc", result)
}
