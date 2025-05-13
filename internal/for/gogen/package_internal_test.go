package gogen_test

import (
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/for/gogen"
)

func Test_TagJ(t *testing.T) {
	tt := testutils.NewTester(t)
	s := gogen.TagJ("foo")
	tt.CheckEqual("`json:\"foo\"`", s)

	s = gogen.TagJ("foo", false)
	tt.CheckEqual("`json:\"foo\"`", s)
	s = gogen.TagJ("foo", true)
	tt.CheckEqual("`json:\"foo,omitempty\"`", s)
}

func Test_TagY(t *testing.T) {
	tt := testutils.NewTester(t)
	s := gogen.TagY("foo")
	tt.CheckEqual("`yaml:\"foo\"`", s)

	s = gogen.TagY("foo", false)
	tt.CheckEqual("`yaml:\"foo\"`", s)
	s = gogen.TagY("foo", true)
	tt.CheckEqual("`yaml:\"foo,omitempty\"`", s)
}

func Test_TagJY(t *testing.T) {
	tt := testutils.NewTester(t)
	s := gogen.TagJY("foo")
	tt.CheckEqual("`json:\"foo\" yaml:\"foo\"`", s)

	s = gogen.TagJY("foo", false)
	tt.CheckEqual("`json:\"foo\" yaml:\"foo\"`", s)
	s = gogen.TagJY("foo", true)
	tt.CheckEqual("`json:\"foo,omitempty\" yaml:\"foo,omitempty\"`", s)
}
