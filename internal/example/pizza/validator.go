package pizza

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// Validator holds context during a validation of a KbSnippet.
// Usage, create a blank Validator and call On().
type Validator struct {
	// file is the file being validated
	file   string
	issues []Issue
	data   *KbSnippet
}

// Issue describes an issue with a message and location in source text.
type Issue struct {
	Message string
	Line    int
	Column  int
}

// HasPizzaNamed return true if there is a Pizza with the given name (ID).
func (v *Validator) HasPizzaNamed(name string) bool {
	for _, p := range v.data.Pizzas {
		if p.ID == name {
			return true
		}
	}
	return false
}

// HasToppingNamed return true if there is a Topping with the given name.
func (v *Validator) HasToppingNamed(name string) bool {
	for _, t := range v.data.Toppings {
		if t.Name == name {
			return true
		}
	}
	return false
}

// On sets the file and Kb to operate when validating.
func (v *Validator) On(file string, data *KbSnippet) {
	v.file = file
	v.data = data
}

// Issue adds an issue from a string message.
func (v *Validator) Issue(message string, line, column int) {
	v.issues = append(v.issues, Issue{Message: message, Line: line, Column: column})
}

// String produces a string containing all validation messages.
func (v *Validator) String() string {
	// prepend with filename to produce lines like:
	// "someplace/foo/bar.yaml: A Pizza must have a name"
	// error text in red
	red := color.New(color.FgRed).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	var outlines []string //nolint:golint,prealloc

	// suppress filename if it is stdin (represented as "-")
	var filename string
	if v.file == "-" {
		filename = ""
	} else {
		filename = v.file
	}
	for _, issue := range v.issues {
		var extra string
		switch {
		case issue.Line > 0 && issue.Column > 0:
			extra = fmt.Sprintf(":%d:%d", issue.Line, issue.Column)
		case issue.Line > 0:
			extra = fmt.Sprintf(":%d", issue.Line)
		}
		outlines = append(outlines, bold(filename+extra+":")+" "+red(issue.Message))
	}
	return strings.Join(outlines, "\n")
}

// Validate validates the KbSninppet given in the call to "On", and returns the count of the reported issues.
func (v *Validator) Validate() (issueCount int) {
	for _, p := range v.data.Pizzas {
		p.Validate(v)
	}
	for _, t := range v.data.Toppings {
		t.Validate(v)
	}
	for _, p := range v.data.Persons {
		p.Validate(v)
	}
	return len(v.issues)
}

// ValidatePizzas reads and validates pizzas. Returns the validated structure.
func ValidatePizzas(file string) *KbSnippet {
	var snippet KbSnippet
	snippet.Read(file)
	var validator Validator
	validator.On(file, &snippet)
	if validator.Validate() == 0 {
		fmt.Println("ok")
	} else {
		// TODO: Windows, must use print to color.Output to get color
		fmt.Println(validator.String())
	}
	return &snippet
}
