package tc

import (
	"fmt"
	"math"

	"github.com/wyrth-io/whit/internal/utils"
)

// DefaultStringChecker is a TypeChecker for any string.
var DefaultStringChecker = &stringChecker{minLen: 0, maxLen: math.MaxInt64}

type stringChecker struct {
	baseChecker
	minLen int
	maxLen int
}

func (c *stringChecker) BaseType() BaseType {
	return valueType(StringKind)
}
func (c *stringChecker) SyntaxString() string {
	hasMin := c.minLen != 0
	hasMax := c.maxLen != math.MaxInt
	switch {
	case hasMin && hasMax:
		return fmt.Sprintf("String[%d, %d]", c.minLen, c.maxLen)
	case hasMin && !hasMax:
		return fmt.Sprintf("String[%d, _]", c.minLen)
	case !hasMin && hasMax:
		return fmt.Sprintf("String[_, %d]", c.maxLen)
	default:
		return StringS
	}
}
func (c *stringChecker) TypeString() string {
	return "String"
}

func (c *stringChecker) Check(v any) (ok bool, message string) {
	if v == nil {
		return false, "nil value is not a String"
	}
	if s, ok := v.(string); ok {
		i := len(s)
		if i <= c.maxLen && i >= c.minLen {
			return true, ""
		}
		if i > c.maxLen {
			return false, fmt.Sprintf("String length %d is longer than max allowed %d", i, c.maxLen)
		}
		if i < c.minLen {
			return false, fmt.Sprintf("String length %d is shorter than min allowed %d", i, c.minLen)
		}
	}
	return false, "value is not a String"
}
func (c *stringChecker) Refine(instructions []any) TypeChecker {
	var ok bool
	template := *DefaultStringChecker
	switch len(instructions) {
	case 2:
		if x := instructions[1]; x != nil {
			var max int64
			max, ok = utils.GetInt64(x)
			if !ok {
				return nil
			}
			template.maxLen = int(max)
		}
		fallthrough
	case 1:
		if x := instructions[0]; x != nil {
			var min int64
			min, ok = utils.GetInt64(x)
			if !ok {
				return nil
			}
			template.minLen = int(min)
		}
	case 0:
		return DefaultFloatChecker
	default:
		return nil
	}
	return &template
}
