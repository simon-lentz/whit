package tc

import "reflect"

// BaseType is an interface for valid base types.
type BaseType interface {
	// Kind is the kind of base type.
	Kind() Kind

	// Dim is the dimension of a SpaceVector, 0 for all non vectors.
	Dim() int
}

type valueType int
type spaceType int

// Kind is the kind of the basetype. See constants.
type Kind int

func (v valueType) Kind() Kind { return Kind(v) } // represents the kind directly
func (v valueType) Dim() int   { return 0 }

func (v spaceType) Kind() Kind { return SpacevectorKind }
func (v spaceType) Dim() int   { return int(v) } // represents the dimension

// Constants for Kind of BaseType.
const (
	// UnspecifiedT is the unspecified data type kind - this means the base type must be determined by the base type of the constraint.
	// The constraint must then be given.
	UnspecifiedKind Kind = iota

	// StringKind is the kind for any length Unicode string.
	StringKind

	// IntT is the base type kind for all kinds of integers - 8, 16, 32, 64 bit, signed/unsigned is expressed as additional constraints.
	IntKind

	// FloatKind is the base type kind for all kinds of floats. Bitsize etc. is expressed as additional constraints.
	FloatKind

	// BoolKind is the base type kind for Boolean values.
	BoolKind

	// SpacevectorKind is the base type kind for a vector space (typically float32 array).
	// This kind has a dimension.
	SpacevectorKind
)

// Lower case string consonants for type names.
const (
	LunspecifiedS  = "unspecified"
	LintS          = "int"
	LintJS         = "integer"
	LfloatS        = "float"
	LfloatJS       = "float"
	LstringS       = "string"
	LstringJS      = "string"
	LboolS         = "bool"
	LboolJS        = "boolean"
	LspacevectorS  = "spacevector"
	LspacevectorJS = "spacevector"
)

// Upper case commonly used string values in yammm types.
// TODO: Rename on form IntegerConstraintName.
const (
	IntegerS    = "Integer"
	FloatS      = "Float"
	StringS     = "String"
	BooleanS    = "Boolean"
	EnumS       = "Enum"
	PatternS    = "Pattern"
	DateS       = "Date"
	TimestampS  = "Timestamp"
	AliasS      = "Alias"
	UUIDS       = "UUID"
	Spacevector = "Spacevector"
)

func (k Kind) String() string {
	switch k {
	case IntKind:
		return LintS
	case FloatKind:
		return LfloatS
	case StringKind:
		return LstringS
	case BoolKind:
		return LboolS
	case SpacevectorKind:
		return LspacevectorS
	default:
		return LunspecifiedS
	}
}

// LongName returns the base type in string form with the type named spelled out
// in long form - e.g. like in JSON schema. One of:
//
//	"integer", "float", "string", "boolean", and (non JSON compatible) "spacevector".
func LongName(b BaseType) (s string) {
	switch b.Kind() {
	case UnspecifiedKind:
		s = LunspecifiedS
	case IntKind:
		s = LintJS
	case FloatKind:
		s = LfloatJS
	case StringKind:
		s = LstringJS
	case BoolKind:
		s = LboolJS
	case SpacevectorKind:
		s = LspacevectorJS
	default:
		s = "Unknown base type - internal error"
	}
	return
}

// GoName returns the base type in (short "go") string form, one of "int", "float", "string",
// "bool", and (the non go compatible) "spacevector".
func GoName(b BaseType) (s string) {
	switch b.Kind() {
	case UnspecifiedKind:
		s = LunspecifiedS
	case IntKind:
		s = LintS
	case FloatKind:
		s = LfloatS
	case StringKind:
		s = LstringS
	case BoolKind:
		s = LboolS
	case SpacevectorKind:
		s = "[]float32"
	default:
		s = "Unknown base type - internal error"
	}
	return
}

// BaseKindOfReflectKind returns a Kind from reflection's Kind.
// TODO: Does not work for spacevec since Kind only denotes that it is a Slice, the type is
// needed to get Elem and get kind from it.
func BaseKindOfReflectKind(rk reflect.Kind) Kind {
	switch rk {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return IntKind
	case reflect.Float32, reflect.Float64:
		return FloatKind
	case reflect.Bool:
		return BoolKind
	case reflect.String:
		return StringKind
	default:
		return UnspecifiedKind
	}
}
