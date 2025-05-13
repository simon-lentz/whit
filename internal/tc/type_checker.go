package tc

import (
	"regexp"
	"strconv"

	"github.com/google/uuid"
)

// TypeChecker is an interface for checking the runtime type of a go any (specifically the base types
// supported by Yammm type constraints. A new type checker is constructed using the function NewTypeChecker.
//
// TODO: Aliases are defined in a context, can introduce an AliasLookup interface for easier testing.
type TypeChecker interface {
	// Check checks if the given value is in compliance with the type constraints.
	Check(any) (ok bool, errorMessage string)

	// BaseType returns a BaseType for the base/underlying type of the value this is a checker for.
	BaseType() BaseType

	// SyntaxString returns a string on the form of the yammm conrete grammar data type representation.
	SyntaxString() string

	// TypeString returns a string on the form of the yammm concrete grammar data type representation
	// without any details of its additional constraints.
	TypeString() string

	// GetUUID returns a uuid.UUID value of the given value (or an error). The special error
	// RequiresResolution is returned for a $$local uuid reference that requires resolution.
	GetUUID(any) (uuid.UUID, error)

	// Refines returns a new type checker of the same kind but with constrained refined by
	// the given detail instructions. Different data types have different number of
	// instructions. Nil is returned if the instructions are not correct
	Refine(instructions []any) TypeChecker
}

type baseChecker struct{}

func (b *baseChecker) GetUUID(_ any) (uuid.UUID, error) { return nullUUID, ErrNotUUIDTypeChecker }

// NewTypeChecker returns a new type checker based on the given string instructions.
// ["Integer", "min", "max"], ["Float", "min", "max"], ["String", "min", "max"], ["Boolean"],
// ["Enum", "choice", ...], ["Pattern", "pattern", ...], ["Timestamp", ""], ["Date"].
// The types Integer, Float and String can be specified with up to two extra arguments (min/max),
// and if given one extra argument it is expected to mean the min. Further, if an empty string
// is passed as an argument it means to take the default value for the min or max.
//
// Nil is returned if the instructions are not well formed.
func NewTypeChecker(instructions []string) (tc TypeChecker) { //nolint:gocognit,gocyclo
	if len(instructions) == 0 {
		return nil
	}
	switch instructions[0] {
	case "String":
		min := ""
		max := ""
		switch len(instructions) {
		case 1:
		case 2:
			min = instructions[1]
		case 3:
			min = instructions[1]
			max = instructions[2]
		}

		if min == "" && max == "" {
			return DefaultStringChecker
		}
		template := *DefaultStringChecker
		if min != "" {
			i, err := strconv.Atoi(min)
			if err != nil {
				return nil
			}
			if 1 < 0 {
				return nil
			}
			template.minLen = i
		}
		if max != "" {
			i, err := strconv.Atoi(max)
			if err != nil {
				return nil
			}
			if i < 0 {
				return nil
			}
			template.maxLen = i
		}
		tc = &template

	case "Integer":
		min := ""
		max := ""
		switch len(instructions) {
		case 1:
		case 2:
			min = instructions[1]
		case 3:
			min = instructions[1]
			max = instructions[2]
		}
		if min == "" && max == "" {
			return DefaultIntChecker
		}
		template := *DefaultIntChecker
		if min != "" {
			i, err := strconv.ParseInt(min, 10, 64)
			if err != nil {
				return nil
			}
			template.min = i
		}
		if max != "" {
			i, err := strconv.ParseInt(max, 10, 64)
			if err != nil {
				return nil
			}
			template.max = i
		}
		tc = &template

	case "Float":
		min := ""
		max := ""
		switch len(instructions) {
		case 1:
		case 2:
			min = instructions[1]
		case 3:
			min = instructions[1]
			max = instructions[2]
		}

		if min == "" && max == "" {
			return DefaultFloatChecker
		}
		template := *DefaultFloatChecker
		if min != "" {
			i, err := strconv.ParseFloat(min, 64)
			if err != nil {
				return nil
			}
			template.min = i
		}
		if max != "" {
			i, err := strconv.ParseFloat(max, 64)
			if err != nil {
				return nil
			}
			template.max = i
		}
		tc = &template

	case "Boolean":
		if len(instructions) != 1 {
			return nil
		}
		tc = DefaultBoolChecker

	case "Enum":
		if len(instructions) < 2 {
			return DefaultEnumChecker // matches nothing.
		}
		ec := &enumChecker{members: map[string]any{}}
		for i := range instructions[1:] {
			ec.members[instructions[i+1]] = true
		}
		tc = ec

	case "Pattern":
		if len(instructions) < 2 {
			return DefaultPatternChecker // matches nothing
		}
		pc := &patternChecker{}
		for i := range instructions[1:] {
			pc.patterns = append(pc.patterns, regexp.MustCompile(instructions[i+1]))
		}
		tc = pc

	case "Date":
		if len(instructions) != 1 {
			return nil
		}
		tc = DefaultDateChecker

	case "Timestamp":
		switch len(instructions) {
		case 1:
			tc = DefaultTimestampChecker
		case 2:
			tc = &timestampChecker{format: instructions[1]}
		default:
			return nil
		}
	case "UUID":
		if len(instructions) != 1 {
			return nil
		}
		tc = DefaultUUIDChecker

	case "Spacevector":
		if len(instructions) != 2 {
			return nil
		}
		i, err := strconv.ParseInt(instructions[1], 10, 64)
		if err != nil {
			return nil
		}
		return &spacevectorChecker{dimensions: int(i)}
	}
	return tc
}
