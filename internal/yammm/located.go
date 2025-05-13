package yammm

import "fmt"

// Located keeps track of source (e.g. file name), line and column of a modelled element.
type Located struct {
	// Line is the line where the schema name is defined.
	Line int `json:"line,omitempty"`

	// Column is the column of where the schema name starts.
	Column int `json:"column,omitempty"`

	// Source is the name of the data source. Typically only used for error messages.
	Source string `json:"source,omitempty"`
}

// Label formats a string describing source, plus line if it is set, plus column if it is set.
// The format is `[file:line:col]`
// Returns empty string if name is not set, and only file if not line is set.
func (loc *Located) Label() (result string) {
	if loc == nil {
		return ""
	}
	source := loc.Source
	line := loc.Line
	col := loc.Column
	if source == "" && line == 0 {
		return ""
	}
	switch {
	case line != 0 && col != 0:
		result = fmt.Sprintf("[%s:%d:%d] ", source, line, col)
	case line != 0 && col == 0:
		result = fmt.Sprintf("[%s:%d] ", source, line)
	default:
		result = fmt.Sprintf("[%s] ", source)
	}

	return result
}
