package xray

import "fmt"

// Position is an interface for getting Line and Col information from an element describing source.
type Position interface {
	// Line returns the linenumber.
	Line() int
	// Col returns the column on the line.
	Col() int
}

// Label returns a "[sourceName:line:col]" label for a Position. If line is 0 the line and col parts are
// omitted.
func Label(p Position, sourceName string) string {
	line := p.Line()
	col := p.Col()
	switch {
	case sourceName != "" && line > 0 && col > 1:
		return fmt.Sprintf("[%s:%d:%d] ", sourceName, line, col)
	case sourceName != "" && line > 0:
		return fmt.Sprintf("[%s:%d] ", sourceName, line)
	case sourceName != "":
		return fmt.Sprintf("[%s] ", sourceName)
	case sourceName == "" && line > 0 && col > 0:
		return fmt.Sprintf("[%d:%d] ", line, col)
	case sourceName == "" && line > 0:
		return fmt.Sprintf("[%d] ", line)
	}
	return ""
}
