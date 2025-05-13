package validation

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hlindberg/testutils"
)

func Test_NewIssueCollector(t *testing.T) {
	// Test that expected implementation is returned
	// (configured to panic on fatal, and default, and custom issue limit)
	tt := testutils.NewTester(t)
	x := NewIssueCollector()
	tt.CheckNotNil(x)
	tt.CheckEqual("*validation.issueCollector", fmt.Sprintf("%v", reflect.TypeOf(x)))
	var ic *issueCollector
	var ok bool
	if ic, ok = x.(*issueCollector); ok {
		tt.CheckFalse(ic.panicOnFatal)
	}
	tt.CheckTrue(ok)
	tt.CheckEqual(200, x.IssueLimit())

	x = NewIssueCollector(123)
	tt.CheckEqual(123, x.IssueLimit())
	if ic, ok = x.(*issueCollector); ok {
		tt.CheckFalse(ic.panicOnFatal)
	}
	tt.CheckTrue(ok)
}

func Test_NewTerminatingIssueCollector(t *testing.T) {
	// Test that expected implementation is returned
	// (configured to panic on fatal)
	tt := testutils.NewTester(t)
	x := NewTerminatingIssueCollector()
	tt.CheckNotNil(x)
	tt.CheckEqual("*validation.issueCollector", fmt.Sprintf("%v", reflect.TypeOf(x)))
	var ok bool
	var ic *issueCollector
	if ic, ok = x.(*issueCollector); ok {
		tt.CheckTrue(ic.panicOnFatal)
	}
	tt.CheckTrue(ok)
	tt.CheckEqual(200, x.IssueLimit())

	x = NewTerminatingIssueCollector(123)
	tt.CheckEqual(123, x.IssueLimit())
	if ic, ok = x.(*issueCollector); ok {
		tt.CheckTrue(ic.panicOnFatal)
	}
	tt.CheckTrue(ok)
}

func Test_IssueCollector_counts_collected_issues(t *testing.T) {
	// Test that expected implementation is returned
	// (configured to panic on fatal)
	x := NewIssueCollector()
	invariant := func(ic IssueCollector, count int, errors, warnings, info, fatal bool) {
		t.Helper()
		tt := testutils.NewTester(t) // Don't want the other tt to change At
		tt.At(0).CheckEqual(count, x.Count())
		tt.At(1).CheckEqual(errors, x.HasErrors())
		tt.At(2).CheckEqual(warnings, x.HasWarnings())
		tt.At(3).CheckEqual(info, x.HasInfo())
		tt.At(4).CheckEqual(fatal, x.HasFatal())
	}
	// Check there are no messages
	tt := testutils.NewTester(t)
	messages := []string{}
	messageAppender := func(issue Issue) { messages = append(messages, issue.Message()) }
	messageBuilder := func(issue Issue) string { return issue.Message() }
	x.EachIssue(messageAppender)
	tt.CheckEqual([]string{}, messages)

	invariant(x, 0, false, false, false, false)
	x.Collectf(Error, "error 1")
	invariant(x, 1, true, false, false, false)
	x.Collectf(Error, "error 2")
	invariant(x, 2, true, false, false, false)
	messages = []string{}

	// Only the two errors present
	x.EachIssue(messageAppender)
	testutils.CheckEqualElements([]string{"error 1", "error 2"}, messages, t)
	tt.CheckEqual([]string{"error 1", "error 2"}, MapEachIssueAtLevel(x, Error, messageBuilder))

	x.Collectf(Warning, "warning 1")
	invariant(x, 3, true, true, false, false)
	x.Collectf(Warning, "warning 2")
	invariant(x, 4, true, true, false, false)
	tt.CheckEqual([]string{"warning 1", "warning 2"}, MapEachIssueAtLevel(x, Warning, messageBuilder))

	x.Collectf(Info, "info 1")
	invariant(x, 5, true, true, true, false)
	x.Collectf(Info, "info 2")
	invariant(x, 6, true, true, true, false)
	tt.CheckEqual([]string{"info 1", "info 2"}, MapEachIssueAtLevel(x, Info, messageBuilder))

	x.Collectf(Fatal, "fatal 1")
	invariant(x, 7, true, true, true, true)
	x.Collectf(Fatal, "fatal 2")
	invariant(x, 8, true, true, true, true)
	tt.CheckEqual([]string{"fatal 1", "fatal 2"}, MapEachIssueAtLevel(x, Fatal, messageBuilder))

	tt.CheckEqual([]string{
		"error 1", "error 2",
		"warning 1", "warning 2",
		"fatal 1", "fatal 2",
	}, MapEachIssueLteLevel(x, Warning, messageBuilder))

	messages = []string{}
	x.EachIssue(messageAppender)
	testutils.CheckEqualElements([]string{ // unordered
		"error 1", "error 2",
		"warning 1", "warning 2",
		"info 1", "info 2",
		"fatal 1", "fatal 2",
	}, messages, t)
}

func Test_Issue_String(t *testing.T) {
	tt := testutils.NewTester(t)
	i := &issue{level: Error, message: "meh"}
	tt.CheckEqual("error: meh", i.String())

	i = &issue{level: Warning, message: "meh"}
	tt.CheckEqual("warning: meh", i.String())

	i = &issue{level: Fatal, message: "meh"}
	tt.CheckEqual("fatal: meh", i.String())

	i = &issue{level: Info, message: "meh"}
	tt.CheckEqual("info: meh", i.String())
}

// Test being over the report limit.

func TestIssueCollector_panics_on_fatal(t *testing.T) {
	defer testutils.ShouldPanic(t)
	x := NewTerminatingIssueCollector()
	x.Collectf(Fatal, "boom")
}

func TestIssueCollector_panics_when_over_limit(t *testing.T) {
	defer testutils.ShouldNotPanic(t)
	x := NewIssueCollector(2)
	c := Do(func() {
		x.Collectf(Fatal, "boom1")
		x.Collectf(Fatal, "boom2")
	})
	messageBuilder := func(issue Issue) string { return issue.Message() }
	tt := testutils.NewTester(t)
	tt.CheckEqual([]string{
		"boom1", "boom2",
		"stopping - too many issues!",
	}, MapEachIssueLteLevel(x, Warning, messageBuilder))

	tt.CheckEqual(3, c.Count())
}

func TestDo_catches_collector_panic(t *testing.T) {
	tt := testutils.NewTester(t)
	defer testutils.ShouldNotPanic(t)
	c := Do(func() {
		x := NewTerminatingIssueCollector()
		x.Collectf(Fatal, "boom")
	})
	tt.CheckTrue(c.HasFatal())
}

func TestDo_panics_if_not_a_collector_value(t *testing.T) {
	defer testutils.ShouldPanic(t)
	Do(func() {
		panic("Aaargh!")
	})
}

func Test_MessageSet(t *testing.T) {
	tt := testutils.NewTester(t)
	ic := NewIssueCollector()
	ic.Collectf(Error, "a")
	ic.Collectf(Error, "b")
	ic.Collectf(Error, "b")
	ms := MessageSet(ic)
	tt.CheckEqual(2, ms.Size())
	tt.CheckTrue(ms.Contains("a"))
	tt.CheckTrue(ms.Contains("b"))
}

// TODO: Test new methods in IssueCollector CollectFatalIf...
