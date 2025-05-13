package csvgen_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/jzon"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/validation"
	"github.com/wyrth-io/whit/internal/yammm"
	"github.com/wyrth-io/whit/parser"
)

// validatedModelMessages validates the model (no assumption it is correct) together with
// instance validation. Thus messages can also come from validation when model is broken.
// The created context and the parsed graph are returned if there were no failures.
func ValidateModelMessages(t *testing.T, model, instance string, expected ...string) (yammm.Context, any) {
	t.Helper()
	tt := testutils.NewTester(t)
	// use false to stop panic on bad model from makeContext and use ic from model validation
	ctx, ic := makeContext(t.Name(), model, false)
	graph := makeJzonInstance(t.Name(), instance)
	v := yammm.NewValidator(ctx, t.Name(), graph)
	result := v.Validate(ic)
	if len(expected) == 0 {
		tt.CheckTrue(result)
	} else {
		tt.CheckFalse(result)
	}
	actual := utils.NewSet[string]()
	ic.EachIssue(func(issue validation.Issue) { actual.Add(issue.Message()) })
	expectedMessages := utils.NewSet(expected...)
	diffA := actual.Diff(expectedMessages)
	diffE := expectedMessages.Diff(actual)
	diff := diffA.Union(diffE)
	if diff.Size() != 0 {
		fmt.Printf("actual<->expected:\n%v\n", diff)
	}
	tt.CheckEqual(0, diffE.Size())
	return ctx, graph
}

// makeJzonInstance makes an interface for json input using jzon parser.
func makeJzonInstance(sourceName string, jsonString string) any {
	ctx := jzon.NewContext(sourceName, jsonString)
	node, err := ctx.UnmarshalNode()
	if err != nil {
		panic(fmt.Sprintf("Bad json given to makeJzonInstance: %s", err))
	}
	return node
}

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

// makeTmpFile creates a tmp file with the given contents. It is the caller's responsibility to remove it.
func makeTmpFile(content string) *os.File {
	// Create a temporary file
	file, errs := os.CreateTemp("", "temp-*.csv")
	if errs != nil {
		panic(errs)
	}

	_, errs = file.WriteString(content)
	if errs != nil {
		panic(errs)
	}

	// Close the file
	errs = file.Close()
	if errs != nil {
		panic(errs)
	}

	return file
}
