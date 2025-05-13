package validation

import (
	"bufio"
	"io"

	"github.com/fatih/color"
)

// IssuePresentor is an interface for presenting issues collected by an
// IssueCollector.
type IssuePresentor interface {
	// Present the set of issues specified by the issuesLevel. Different implementations
	// may interpret this argument differently; matching, <, <=, etc.
	Present(collector IssueCollector, issuesLevel IssueLevel, out io.Writer) error
}
type colorPresentor struct{}

// NewColorPresentor returns a new IssuePresentor that outputs colored error messages.
func NewColorPresentor() IssuePresentor {
	return &colorPresentor{}
}

// Present outputs the issues in the order they were reported and colors each message
// fatal (bold, red), error (red), warning (yellow), info (green). Each on a separate line
// with an error level label first, for example "error: something wenr wrong\n"
// The presentation is limited to issue levels less than or equal to the given level,
// Thus validation.Warning would output Fatal, Error and Warning issues, but not Info.
// The function returns error if writing to the given Writer got an error.
// TODO: Functions does not filter on IssueLevel (what the _ parameter is for).
func (p *colorPresentor) Present(collector IssueCollector, _ /*issuesLte*/ IssueLevel, out io.Writer) (err error) {
	// present errors and warnings in the order they were reported

	type colorF func(a ...interface{}) string
	colorFuncs := map[IssueLevel]colorF{
		Fatal:   color.New(color.FgRed, color.Bold).SprintFunc(),
		Error:   color.New(color.FgRed).SprintFunc(),
		Warning: color.New(color.FgYellow).SprintFunc(),
		Info:    color.New(color.FgGreen).SprintFunc(),
	}
	output := bufio.NewWriter(out)

	collector.EachIssue(func(issue Issue) {
		_, err = output.WriteString(colorFuncs[issue.Level()](issue.String()) + "\n")
		if err != nil {
			return
		}
	})
	err = output.Flush()
	return
}
