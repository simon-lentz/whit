package gogen_test

import (
	"os"

	"github.com/wyrth-io/whit/internal/validation"
	"github.com/wyrth-io/whit/internal/yammm"
	"github.com/wyrth-io/whit/parser"
)

// makeContext makes a context from yammm model source.
func makeContext(sourceRef string, source string, assertModel bool) (yammm.Context, validation.IssueCollector) {
	yammmCtx, ic := parser.ParseString(sourceRef, source)
	if yammmCtx == nil {
		if assertModel {
			err := validation.NewColorPresentor().Present(ic, validation.Info, os.Stderr)
			if err != nil {
				panic(err)
			}
			panic("Bad input to makeContext - error in test itself.")
		}
	}
	return yammmCtx, ic
}
