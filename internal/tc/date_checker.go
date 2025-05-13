package tc

import (
	"fmt"
	"time"
)

// DefaultDateChecker is a TypeChecker for a Date in time.RFC3339Date format.
var DefaultDateChecker = &dateChecker{}

type dateChecker struct {
	baseChecker
}

func (c *dateChecker) SyntaxString() string {
	return "Date"
}
func (c *dateChecker) TypeString() string {
	return c.SyntaxString()
}
func (c *dateChecker) BaseType() BaseType {
	return valueType(StringKind)
}
func (c *dateChecker) Check(v any) (bool, string) {
	if v == nil {
		return false, "nil value is not a Date"
	}
	if s, ok := v.(string); ok {
		if _, err := time.Parse(time.DateOnly, s); err != nil {
			return false, fmt.Sprintf("value '%s' does not match Date format '%s'", s, time.DateOnly)
		}
		// TODO: support min/max
		return true, ""
	}
	return false, "value is not a Date"
}
func (c *dateChecker) Refine(instr []any) TypeChecker {
	if len(instr) == 0 {
		return c
	}
	return nil
}
