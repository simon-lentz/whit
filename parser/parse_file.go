package parser

import (
	"log"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/wyrth-io/whit/internal/validation"
	"github.com/wyrth-io/whit/internal/yammm"
)

// ErrorCounterListener listens for ANTLR errors and counts them.
type ErrorCounterListener struct {
	antlr.DefaultErrorListener
	count int
}

// SyntaxError counts one syntax error from ANTLR.
func (counter *ErrorCounterListener) SyntaxError(
	recognizer antlr.Recognizer, //nolint:revive
	offendingSymbol interface{}, //nolint:revive
	line, //nolint:revive
	column int, //nolint:revive
	msg string, //nolint:revive
	e antlr.RecognitionException) { //nolint:revive
	counter.count++
}

// Count returrns the current number of collected syntax errors.
func (counter *ErrorCounterListener) Count() int {
	return counter.count
}

// ParseString parses and validates yammm source in string form and returns a yammm.Context and the validation.IssueCollector
// used to validate the model. A nil context is returned if there are errors.
func ParseString(sourceRef string, source string) (yammm.Context, validation.IssueCollector) {
	return parseInput(sourceRef, antlr.NewInputStream(source))
}

// ParseFile parses and validates a yammm file and returns a yammm.Context and the validation.IssueCollector
// used to validate the model. A nil context is returned if there are errors.
func ParseFile(fileName string) (yammm.Context, validation.IssueCollector) {
	input, err := antlr.NewFileStream(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return parseInput(fileName, input)
}

func parseInput(sourceRef string, input antlr.CharStream) (yammm.Context, validation.IssueCollector) {
	lexer := NewYammmGrammarLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewYammmGrammarParser(stream)
	errorListener := antlr.NewDiagnosticErrorListener(true)
	errorCounter := &ErrorCounterListener{}
	p.AddErrorListener(errorListener)
	p.AddErrorListener(errorCounter)
	p.BuildParseTrees = true
	tree := p.Schema()
	listener := NewParseListener(sourceRef)
	var ok bool
	validation.Do(func() {
		// generate the model by listening to parse events
		antlr.ParseTreeWalkerDefault.Walk(listener, tree)
		if errorCounter.Count() > 0 {
			listener.ic.Collectf(validation.Fatal, "Syntax errors in yammm source prevents further processing.")
		}
		// complete the generated yammm model (triggers validation)
		ok = listener.ctx.Complete(listener.ic)
	})

	if listener.ic.HasFatal() {
		return nil, listener.ic
	}
	if listener.ic.HasErrors() {
		return nil, listener.ic
	}
	if !ok {
		// Should not really happen, ok == false only if there were fatals or errors.
		// But if there are bugs in the validation pkg this may go wrong.
		panic("internal error: fatal errors occurred while parsing/validating yamm model, but was not handled")
	}
	return listener.ctx, listener.ic
}
