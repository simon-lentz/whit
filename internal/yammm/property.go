package yammm

import (
	"github.com/wyrth-io/whit/internal/tc"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/validation"
)

// Property describes a property in a Type or Association. The DataType is expressed as
// a slice of string where the first entry is the name of a yammm data type with initial upper case
// letter - for example "String", "Integer", "Float". Depending on data type additional arguments
// can describe a value range - for example []string{"Integer", "0", "100"}. See [NewTypeChecker] for available
// types and constraints.
// An instance of the property is always a single value, lists or maps are not supported.
type Property struct {
	Located

	// Name is the name of the property, it must have an initial lower case letter.
	Name string `json:"name"`

	// DataType is a value constraint expressed using an array of string values. These values
	// encode the constraints.
	DataType []string `json:"datatype"`

	// Optional is a flag that indicates if this property is optional or not.
	Optional bool `json:"optional,omitempty"`

	// IsPrimary indicates if this propery is a primary key or not. Possibly in combination
	// with other properties also marked as being a primary key.
	IsPrimaryKey bool `json:"primary,omitempty"`

	// Documentation is a string in markdown describing the property.
	Documentation string `json:"documentation,omitempty"`

	// completed is true if the property has been completed. This flag is important to avoid completing
	// an inherited property each time it is included in a subtype.
	completed bool

	// typeChecker is a cashed created TypeChecker created from the DataType on demand.
	typeChecker tc.TypeChecker
}

// TypeChecker returns a type checker for the DataType of the property.
func (p *Property) TypeChecker() tc.TypeChecker {
	return p.typeChecker
}

// DataTypeString returns a string in the yammm concrete grammar syntax for the data type of this propery.
func (p *Property) DataTypeString() string {
	return p.typeChecker.SyntaxString()
}

// ShortDataTypeString returns a string in the yammm concrete grammar syntax for the data type of this propery
// without the detailed constraints.
func (p *Property) ShortDataTypeString() string {
	return p.typeChecker.TypeString()
}

func (p *Property) complete(ctx Context, ic validation.IssueCollector) (ok bool) {
	if p.completed {
		return true
	}
	p.completed = true
	var effectiveConstraints []string
	if len(p.DataType) == 2 && p.DataType[0] == "Alias" {
		dt := ctx.LookupDataType(p.DataType[1])
		if dt == nil {
			ic.Collectf(validation.Fatal,
				"%scannot compute type checker for property '%s': referenced custom data type not found '%s'",
				p.Label(),
				p.Name,
				p.DataType[1],
			)
			return false
		}
		effectiveConstraints = dt.Constraint
	} else {
		effectiveConstraints = p.DataType
	}
	p.typeChecker = tc.NewTypeChecker(effectiveConstraints)
	if p.typeChecker == nil {
		ic.Collectf(validation.Fatal,
			"%scannot compute type checker for property '%s' from its constraint '%v'",
			p.Label(),
			p.Name,
			effectiveConstraints, // Note, cannot use the SyntaxString method on typechecker here...
		)
		return false
	}
	return true
}

// BaseType returns the property's tc.BaseType.
func (p Property) BaseType() tc.BaseType {
	return p.TypeChecker().BaseType() // TODO: WAS: .String()
}

// Validate validates the Property and returns true if it did not raise any errors.
func (p Property) validate(ctx Context, ic validation.IssueCollector) (ok bool) {
	ok = p.complete(ctx, ic) // make sure property is completed.
	if len(p.Name) < 1 {
		ic.Collectf(validation.Error, "%sa property with empty name having constraint '%s' - must have a name",
			p.Label(), p.DataTypeString())
		ok = false
	} else {
		firstLetter := p.Name[0:1]
		if firstLetter < "a" || firstLetter > "z" {
			ok = false
			ic.Collectf(validation.Error, "%sinvalid property name '%s' - property name must start with lower case 'a'-'z'",
				p.Label(), p.Name)
		} else if !IsLCName(p.Name) {
			ok = false
			ic.Collectf(validation.Error, "%sinvalid property name '%s' - contains illegal characters", p.Label(), p.Name)
		}
	}
	// Not meaningful to validate if property name is wrong
	if ok {
		if p.TypeChecker() == nil {
			ok = false
			ic.Collectf(validation.Error, "%sinvalid data type for property '%s'", p.Label(), p.Name)
		}
	}
	return
}

// GoName returns the name of the property in generated go code. (An initial upper cased name).
func (p Property) GoName() string {
	return utils.Capitalize(p.Name)
}

// HasSameType returns true if this property has the same type as the given property and
// optionality and primary key are set the same way.
func (p *Property) HasSameType(p2 *Property) bool {
	dt1 := p.DataType
	dt2 := p2.DataType
	if len(dt1) != len(dt2) {
		return false
	}
	for i := range dt1 {
		if dt1[i] != dt2[i] {
			return false
		}
	}
	return p.Optional == p2.Optional && p.IsPrimaryKey == p2.IsPrimaryKey
}

// HasDefault returns true if this has a default value.
func (p *Property) HasDefault() bool {
	// This can be extended later allowing different type of default functions.
	return p.DataType[0] == tc.UUIDS
}

// DefaultStringValue returns a default string. For properties not having a default value
// the return is an empty string and false. For UUID an empty string and true are returned.
// The actual UUID default value is handled separately since it is either a deterministic
// v5 UUID generated from the set of primary keys, a local UUID to resolve, or an actual UUID.
func (p *Property) DefaultStringValue() (value string, ok bool) {
	if p.DataType[0] == tc.UUIDS {
		return "", true // UUID is produced later
	}
	return "", false
}
