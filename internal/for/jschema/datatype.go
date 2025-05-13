package jschema

import (
	"github.com/wyrth-io/whit/internal/tc"
	"github.com/wyrth-io/whit/internal/yammm"
)

// BaseTypeName returns the JSONSchema string form of the base type. For example "string".
// NOTE: This will return "spacevector" for the SpacevectorType which cannot be directly
// translated to the correct schema.
func BaseTypeName(p *yammm.Property) string {
	return tc.LongName(p.TypeChecker().BaseType())
}
