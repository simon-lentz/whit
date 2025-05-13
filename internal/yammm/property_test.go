package yammm_test

import (
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/yammm"
)

func Test_PropertyHasSameType(t *testing.T) {
	tt := testutils.NewTester(t)
	p1 := yammm.Property{Name: "x", DataType: []string{"String"}}
	p2 := yammm.Property{Name: "x", DataType: []string{"String", "0"}}
	tt.CheckTrue(p1.HasSameType(&p1))
	tt.CheckFalse(p1.HasSameType(&p2))
}

func Test_PropertyHasDefault(t *testing.T) {
	tt := testutils.NewTester(t)
	p1 := yammm.Property{Name: "x", DataType: []string{"UUID"}}
	p2 := yammm.Property{Name: "x", DataType: []string{"String", "0"}}
	tt.CheckTrue(p1.HasDefault())
	tt.CheckFalse(p2.HasDefault())
}
