package yammm

import (
	"github.com/wyrth-io/whit/internal/tc"
	"github.com/wyrth-io/whit/internal/validation"
)

// DataType represents a data type - it has a name (first character in lower case), and a yammm type constraint
// expressed as a slice of string (see [NewTypeChecker] for available constraints).
type DataType struct {
	Located

	// Name is the name of the data type - must start with 'a'-'z'.
	Name string `json:"name"`

	// Constraint is a slice of instructions for a TypeChecker.
	Constraint []string `json:"constraint"`

	// Documentation is a markdown string describing the datatype.
	Documentation string `json:"documentation,omitempty"`
}

// Validate validates the DataType and returns true if it did not raise any errors.
func (dt *DataType) validate(_ Context, ic validation.IssueCollector) (ok bool) {
	ok = true
	if len(dt.Name) < 1 {
		ic.Collectf(validation.Error, "%sdata type has empty name", dt.Label())
		ok = false
	} else {
		firstLetter := dt.Name[0:1]
		if firstLetter < "a" || firstLetter > "z" {
			ok = false
			ic.Collectf(validation.Error, "%sinvalid data type name '%s' - data type name must start with lower case 'a'-'z'",
				dt.Label(), dt.Name)
		}
	}
	if tc.NewTypeChecker(dt.Constraint) == nil {
		ok = false
		ic.Collectf(validation.Error, "%sinvalid constraint for data type '%s': %v", dt.Label(), dt.Name, dt.Constraint)
	}
	return
}
