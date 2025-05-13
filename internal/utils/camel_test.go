package utils_test

import (
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/utils"
)

func TestToLowerCamel(t *testing.T) {
	tt := testutils.NewTester(t)
	// Should replace all non ascii word chars with single _ character.
	// Should then split with _ between flank lower to upper case
	r := utils.ToLowerCamel("St(range)___pCamelCase32_33Foo")
	// tt.CheckEqual("St_range_p_Camel_Case32_33Foo", r)
	tt.CheckEqual("stRangePCamelCase32_33Foo", r)
}
func TestToUpperCamel(t *testing.T) {
	tt := testutils.NewTester(t)
	// Should replace all non ascii word chars with single _ character.
	// Should then split with _ between flank lower to upper case
	r := utils.ToUpperCamel("St(range)___pCamelCase32_33Foo")
	// tt.CheckEqual("St_range_p_Camel_Case32_33Foo", r)
	tt.CheckEqual("StRangePCamelCase32_33Foo", r)
}
