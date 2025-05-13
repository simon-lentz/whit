package tc

import (
	"fmt"
	"sort"
	"strings"

	"github.com/wyrth-io/whit/internal/utils"
)

type enumChecker struct {
	baseChecker
	members map[string]any
}

// DefaultEnumChecker is an Enum checker that matches nothing.
var DefaultEnumChecker = &enumChecker{members: map[string]any{}}

func (c *enumChecker) BaseType() BaseType {
	return valueType(StringKind)
}
func (c *enumChecker) SyntaxString() string {
	quoted := utils.Map(utils.Keys(c.members), func(s string) string { return fmt.Sprintf("%q", s) })
	return fmt.Sprintf("Enum[%s]", strings.Join(quoted, ", "))
}
func (c *enumChecker) TypeString() string {
	return "Enum"
}

func (c *enumChecker) Check(v any) (bool, string) {
	if v == nil {
		return false, "nil value is not an Enum"
	}
	if s, ok := v.(string); ok {
		if _, ok := c.members[s]; ok {
			return true, ""
		}
		keys := []string{}
		for k := range c.members {
			keys = append(keys, `"`+k+`"`)
		}
		sort.Strings(keys)
		return false, fmt.Sprintf("String value '%s' does not match Enum[%s]", s, strings.Join(keys, ", "))
	}
	return false, "value is not a String, cannot match an Enum"
}
func (c *enumChecker) Refine(instr []any) TypeChecker {
	if len(instr) == 0 {
		return c
	}
	ec := &enumChecker{members: make(map[string]any, len(instr))}
	for i := range instr {
		if s, ok := instr[i].(string); ok {
			ec.members[s] = true
		} else {
			return nil
		}
	}
	return ec
}
