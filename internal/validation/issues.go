package validation

import (
	"fmt"

	"github.com/wyrth-io/whit/internal/utils"
)

// IssueCollector is used to collect issues (errors, warnings, info, fatal) and where it is desired to
// keep on processing and report back to a user at the end of a process. This way, the user facing portion of
// the logic can be kept separate, contain formatting, decisions of color or no color etc.
// It is adviceable to avoid using colors for parts of the messages in issues and leave all such decisions
// to the presentation logic of the collective result.
type IssueCollector interface {
	// HasErrors returns true if any issues at Error or Fatal level has been reported
	HasErrors() bool

	// HasInfo returns true if any issues at Info level has been reported
	HasWarnings() bool

	// HasInfo returns true if any issues at Info level has been reported
	HasInfo() bool

	// HasFatal returns true if one of the reported errors was considered fatal (further validation not
	// meaningful.
	HasFatal() bool

	// EachIssue iterates over all reported issues and calls the given function with each issue.
	// See [validation.MapEachIssueAtLevel] and [validation.MapEachIssueGteLevel] for conventient
	// maping/filtering and collection of issues in addition to [EachIssueAtLevel] and [EachIssueLteLevel] in
	// this interface.
	EachIssue(f func(issue Issue))

	// EachIssueAtLevel iterates over all issues having the given level.
	EachIssueAtLevel(level IssueLevel, f func(issue Issue))

	// EachIssueLteLevel iterates over all issues less than or equal to the given level.
	EachIssueLteLevel(level IssueLevel, f func(issue Issue))

	// Collect an issue with formatted message of the given IssueLevel.
	Collectf(level IssueLevel, format string, args ...any)

	// CollectFatalIfErrorf collects a fatal error if any given argument is an error
	CollectFatalIfErrorf(format string, args ...any)

	// CollectFatalIfError collects the error if not nil as a fatal error message
	CollectFatalIfError(err error)

	// Collect a prepared issue.
	Collect(issue Issue)

	// Count returns the total count of all reported issues (at all levels).
	Count() int

	// IssueLimit returns the set limit after which the collector will raise a panic.
	// A specific issue collector may have a default limit. A limit < 0 is unlimited.
	IssueLimit() int
}

// IssueLevel is an Enum for available error levels in an IssueCollector.
type IssueLevel int64

const (
	// Fatal is an error after which it is not meaningful to continue.
	Fatal IssueLevel = iota

	// Error is an error after the which it is still meaningful to continue collecting issues but the
	// operation relying on an "ok" should not expect the validated target to be valid.
	Error

	// Warning is an issue that should be corrected but the validated target is technically still possible
	// to use for further processing.
	Warning

	// Info is an issue that is informational and does not need to be addressed.
	Info
)

// MapEachIssueAtLevel calls the given function with each collected issue of the given level and builds a
// slice of all the returned values.
func MapEachIssueAtLevel[T any](collector IssueCollector, level IssueLevel, f func(issue Issue) T) []T {
	var result []T
	collector.EachIssueAtLevel(level, func(issue Issue) {
		result = append(result, f(issue))
	})
	return result
}

// MapEachIssueLteLevel calls the given function with each collected less or equal to the given level and builds a
// slice[T] of all the returned values where T is the return type of the given function.
func MapEachIssueLteLevel[T any](collector IssueCollector, level IssueLevel, f func(issue Issue) T) []T {
	var result []T
	collector.EachIssueLteLevel(level, func(issue Issue) {
		result = append(result, f(issue))
	})
	return result
}

// String returns the IssueLevel in string form.
func (level IssueLevel) String() string {
	switch level {
	case Fatal:
		return "fatal"
	case Error:
		return "error"
	case Warning:
		return "warning"
	case Info:
		return "info"
	default:
		// Cannot really happen unless someone does erroneus type casting.
		return "unknown issue level"
	}
}

// issueCollector implements IssueCollector.
type issueCollector struct {
	issues       []Issue
	fatalCount   int
	errorCount   int
	warningCount int
	infoCount    int
	// panic On Fatal, indicates if a collected fatal message should lead to a panic.
	panicOnFatal bool
	count        int
	issueLimit   int
}

// Issue is an interface for an issue message at an IssueLevel (fatal, error, warning, info) typically
// used with an IssueCollector.
type Issue interface {
	Message() string
	Level() IssueLevel
	String() string
}

type issue struct {
	message string
	level   IssueLevel
}

func (i *issue) Message() string {
	return i.message
}
func (i *issue) Level() IssueLevel {
	return i.level
}

// String returns a default formatted string on the form "level: message".
func (i *issue) String() string {
	return fmt.Sprintf("%s: %s", i.level, i.message)
}

// NewIssueCollector returns [IssueCollector] that does not panic on fatal errors. The caller should
// check [IssueCollector.HasFatalIssue] if a fatal error was reported and then deciding what to do.
// Function accepts an optional limit on the number of collected issues. If exceeding this value
// a panic(collector) is called. The default is 200. If set to a negative value the collection is "unlimited".
func NewIssueCollector(limits ...int) IssueCollector {
	return newIssueCollector(limits...)
}

func newIssueCollector(limits ...int) *issueCollector {
	limit := 200 // default value
	switch len(limits) {
	case 0: // do nothing, default is set
	case 1:
		limit = limits[0]
	default:
		panic(fmt.Sprintf("NewIssueCollector accepts at most one int value, got: %v", limits))
	}
	return &issueCollector{issueLimit: limit}
}

// NewTerminatingIssueCollector returns an [IssueCollector] that panics on the first reported fatal issue.
// The expected usage is to wrap logic that makes use of this IssueCollector in a
// call to [Do] which returns the collector when a fatal error was reported. The caller of [Do] can then
// present the collected issues.
// Function accepts an optional limit on the number of collected issues. If exceeding this value
// a panic(collector) is called. The default is 200. If set to a negative value the collection is "unlimited".
func NewTerminatingIssueCollector(limits ...int) IssueCollector {
	ic := newIssueCollector(limits...)
	ic.panicOnFatal = true
	return ic
}

// HasErrors returns true if an [Error] was collected.
func (collector *issueCollector) HasErrors() bool {
	return collector.errorCount > 0
}

// HasWarnings returns true if a [Warning] was collected.
func (collector *issueCollector) HasWarnings() bool {
	return collector.warningCount > 0
}

// HasInfo returns true if an [Info] was collected.
func (collector *issueCollector) HasInfo() bool {
	return collector.infoCount > 0
}

// HasFatal returns true if during collection, the work stopped because of
// too many errors, or some error condition that prevented continued validation/work.
func (collector *issueCollector) HasFatal() bool {
	return collector.fatalCount > 0
}

// EachIssue calls the given function for all reported issues of any level.
func (collector *issueCollector) EachIssue(f func(issue Issue)) {
	for _, i := range collector.issues {
		f(i)
	}
}

// EachIssueAtLevel calls the given function only for issues with level equal to the given level.
func (collector *issueCollector) EachIssueAtLevel(level IssueLevel, f func(issue Issue)) {
	for _, i := range collector.issues {
		if i.Level() != level {
			continue
		}
		f(i)
	}
}

// EachIssueLteLevel calls the given function for issues with level Less Than or Equal (Lte) than the given level.
// (Those that are "equal or more severe").
func (collector *issueCollector) EachIssueLteLevel(level IssueLevel, f func(issue Issue)) {
	for _, i := range collector.issues {
		if i.Level() > level {
			continue
		}
		f(i)
	}
}

// Collect adds a prepared Issue.
func (collector *issueCollector) Collect(i Issue) {
	collector.count++
	switch i.Level() {
	case Fatal:
		collector.fatalCount++
	case Error:
		collector.errorCount++
	case Warning:
		collector.warningCount++
	case Info:
		collector.infoCount++
	}

	collector.issues = append(collector.issues, i)
	if collector.panicOnFatal && collector.fatalCount > 0 {
		panic(collector)
	}
	if collector.issueLimit > 0 && collector.count >= collector.issueLimit {
		collector.count++
		collector.fatalCount++
		collector.issues = append(collector.issues, &issue{level: Fatal, message: "stopping - too many issues!"})
		panic(collector)
	}
}

// Collectf collects (adds) an Issue with a formatted message from the given format and arguments
// in the style of fmt.Sprintf.
func (collector *issueCollector) Collectf(level IssueLevel, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	collector.Collect(&issue{level: level, message: message})
}

func (collector *issueCollector) Count() int {
	return collector.count
}

func (collector *issueCollector) IssueLimit() int {
	return collector.issueLimit
}

func (collector *issueCollector) CollectFatalIfError(err error) {
	if err != nil {
		collector.CollectFatalIfErrorf("%s", err)
	}
}

func (collector *issueCollector) CollectFatalIfErrorf(format string, args ...any) {
	if len(args) > 0 {
		for _, a := range args {
			if e, ok := a.(error); ok && e != nil {
				// if one non nil error is found it is taken as an error
				collector.Collectf(Fatal, format, args...)
				return
			}
		}
	}
}

// MessageSet creates a Set[string] for all collected messages.
func MessageSet(ic IssueCollector) *utils.Set[string] {
	m := utils.NewSet[string]()
	ic.EachIssue(func(issue Issue) {
		m.Add(issue.Message())
	})
	return m
}
