package yammm

import (
	"github.com/wyrth-io/whit/internal/validation"
)

// Model is the top level structure - it describes an instance of a meta-model.
type Model struct {
	Located

	// Documentation is overall documentation for the schema.
	Documentation string `json:"documentation,omitempty"`

	// Name is the "package name" of the module. It may be referenced from other models
	// as an import.
	Name string `json:"name"`

	// Types are all types in the model; regular, Abstract, and Parts. Types can describe relationships to
	// other types and may inherit from other types.
	Types []*Type `json:"types"`

	// DataTypes defines named value constraints that can be used in other data types and in properties of
	// relations and types.
	DataTypes []*DataType `json:"data_types,omitempty"`
}

// NewModel returns a new empty model with the given name.
func NewModel(name string) *Model {
	return &Model{Name: name, Types: []*Type{}, DataTypes: []*DataType{}}
}

func (m Model) validate(ctx Context, ic validation.IssueCollector) (ok bool) {
	ok = true
	falseIfNotOk := func(result bool) bool {
		if !result {
			return false
		}
		return ok
	}
	for _, t := range m.Types {
		ok = falseIfNotOk(t.validate(ctx, ic))
	}
	for _, dt := range m.DataTypes {
		ok = falseIfNotOk(dt.validate(ctx, ic))
	}
	return
}
