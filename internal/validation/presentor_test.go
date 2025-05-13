package validation_test

import (
	"os"

	"github.com/wyrth-io/whit/internal/validation"
)

func ExampleIssuePresentor() {
	// Create a collector and collect some issues
	collector := validation.NewIssueCollector()
	collector.Collectf(validation.Error, "this is an error message")
	collector.Collectf(validation.Warning, "it is not so great today")
	collector.Collectf(validation.Info, "it may rain today")
	collector.Collectf(validation.Fatal, "struck by thunder")

	// Create a ColorPresentor and present all issues at Info level and below (i.e. warning, error, fatal)
	p := validation.NewColorPresentor()
	_ = p.Present(collector, validation.Info, os.Stdout)
	//output:
	// error: this is an error message
	// warning: it is not so great today
	// info: it may rain today
	// fatal: struck by thunder
}
