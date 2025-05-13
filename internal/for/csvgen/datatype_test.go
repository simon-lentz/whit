package csvgen_test

import (
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/for/csvgen"
)

func TestIsFloat(t *testing.T) {
	tt := testutils.NewTester(t)
	for input, expected := range map[string]bool{
		"0":       false,
		"1":       false,
		"01.2":    false,
		"0.1.2":   false,
		"0e":      false,
		"0E":      false,
		"E1":      false,
		"0.1":     true,
		"1.1":     true,
		"123.1":   true,
		"1E-1":    true,
		"1E+1":    true,
		"1E10":    true,
		"0.1E-1":  true,
		"0.1E+1":  true,
		"0.11E10": true,
		".1":      true,
		".1e10":   true,
	} {
		tt.CheckEqual(expected, csvgen.IsFloat(input))
	}
}

func TestIsInteger(t *testing.T) {
	tt := testutils.NewTester(t)
	for input, expected := range map[string]bool{
		"0":     true,
		"1":     true,
		"01":    false,
		"-01":   false,
		"01.2":  false,
		"0.1.2": false,
		"0e":    false,
		"0E":    false,
		"E1":    false,
		"0.1":   false,
	} {
		tt.CheckEqual(expected, csvgen.IsInteger(input))
	}
}

func TestIsBoolean(t *testing.T) {
	tt := testutils.NewTester(t)
	tt.CheckTrue(csvgen.IsBoolean("true"))
	tt.CheckTrue(csvgen.IsBoolean("false"))
	tt.CheckTrue(csvgen.IsBoolean("True"))
	tt.CheckTrue(csvgen.IsBoolean("False"))
	tt.CheckFalse(csvgen.IsBoolean("T"))
	tt.CheckFalse(csvgen.IsBoolean("0"))
}
