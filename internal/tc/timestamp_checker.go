package tc

import (
	"fmt"

	"cuelang.org/go/pkg/time"
)

// DefaultTimestampChecker is a checker for a timestamp in tim.RFC3339 format.
var DefaultTimestampChecker = &timestampChecker{format: time.RFC3339}

type timestampChecker struct {
	baseChecker
	format string
}

func (c *timestampChecker) SyntaxString() string {
	if c.format == "" {
		return TimestampS
	}
	return fmt.Sprintf("Timestamp[%q]", c.format)
}
func (c *timestampChecker) TypeString() string {
	return TimestampS
}

func (c *timestampChecker) BaseType() BaseType {
	return valueType(StringKind)
}
func (c *timestampChecker) Check(v any) (bool, string) {
	if v == nil {
		return false, "nil value is not a Timestamp"
	}
	if s, ok := v.(string); ok {
		if _, err := time.Parse(c.format, s); err != nil {
			return false, fmt.Sprintf("value does not match Timestamp format : %s", err)
		}
		// TODO: support min/max
		return true, ""
	}
	return false, "value is not a Timestamp"
}
func (c *timestampChecker) Refine(instr []any) TypeChecker {
	if len(instr) == 0 {
		return c
	}
	if s, ok := instr[0].(string); ok {
		return &timestampChecker{format: s}
	}
	return nil
}
