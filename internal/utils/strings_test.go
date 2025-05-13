package utils_test

import (
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/utils"
)

func Test_Capitaize(t *testing.T) {
	tt := testutils.NewTester((t))
	tt.CheckEqual("Blah", utils.Capitalize("blah"))
}

func Test_Capitaize_does_nothing_on_empty_string(t *testing.T) {
	tt := testutils.NewTester((t))
	tt.CheckEqual("", utils.Capitalize(""))
}

func Test_CapitalizeAll_returns_array_with_all_entrie_capitalized(t *testing.T) {
	tt := testutils.NewTester((t))
	tt.CheckEqual([]string{"Abc", "Åäö", "ÅÄö", ""}, utils.CapitalizeAll([]string{"abc", "åäö", "ÅÄö", ""}))
}

func Test_DeCapitaize(t *testing.T) {
	tt := testutils.NewTester((t))
	tt.CheckEqual("blah", utils.DeCapitalize("Blah"))
	tt.CheckEqual("blah", utils.DeCapitalize("blah"))
	tt.CheckEqual("", utils.DeCapitalize(""))
}

func Test_UniqueFold_returns_unique_array_of_strings_cast_independent(t *testing.T) {
	tt := testutils.NewTester((t))
	tt.CheckEqual([]string{"abc", "åäö"}, utils.UniqueFold([]string{"abc", "åäö", "ÅÄö"}))
}

func Test_DeleteFold_removes_all_instances_of_given_string_case_independently(t *testing.T) {
	tt := testutils.NewTester((t))
	tt.CheckEqual([]string{"abc"}, utils.DeleteFold([]string{"abc", "åäö", "ÅÄö"}, "åÄö"))
}

func Test_CamelToSnake_turns_camel_to_snake_and_retains_case(t *testing.T) {
	tt := testutils.NewTester(t)
	tt.CheckEqual("camel_Case", utils.CamelToSnake(("camelCase")))
	tt.CheckEqual("Camel_Case", utils.CamelToSnake(("CamelCase")))
	tt.CheckEqual("_camel_Case", utils.CamelToSnake(("_camelCase")))
	tt.CheckEqual("__camel_Case", utils.CamelToSnake(("__camel___Case")))
	tt.CheckEqual("The_FBI_Office", utils.CamelToSnake(("TheFBIOffice")))
	tt.CheckEqual("The_FBI_Office", utils.CamelToSnake(("The__F__B__I__Office")))   // spaced acronym
	tt.CheckEqual("The_f_b_i_Office", utils.CamelToSnake(("The__f__b__i__Office"))) // not acronym
	tt.CheckEqual("The_FBI_Office", utils.CamelToSnake(("The_FBI_Office")))
}
