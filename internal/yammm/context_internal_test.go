package yammm

import (
	"testing"

	"github.com/hlindberg/testutils"
)

func Test_NewContextReturns_contextImpl(t *testing.T) {
	tt := testutils.NewTester(t)
	ctx := NewContext()
	tt.CheckNotNil(contextImpl(ctx))
}

func contextImpl(ctx Context) *context {
	switch t := ctx.(type) {
	case *context:
		return t
	default:
		return nil
	}
}
