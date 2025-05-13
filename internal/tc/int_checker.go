package tc

import (
	"fmt"
	"math"

	"github.com/wyrth-io/whit/internal/utils"
)

// DefaultIntChecker is the TypeChecker for any int.
var DefaultIntChecker = &intChecker{min: math.MinInt64, max: math.MaxInt64}

type intChecker struct {
	baseChecker
	min int64
	max int64
}

func (c *intChecker) BaseType() BaseType {
	return valueType(IntKind)
}
func (c *intChecker) SyntaxString() string {
	hasMin := c.min != math.MinInt64
	hasMax := c.max != math.MaxInt64
	switch {
	case hasMin && hasMax:
		return fmt.Sprintf("Integer[%d, %d]", c.min, c.max)
	case hasMin && !hasMax:
		return fmt.Sprintf("Integer[%d, _]", c.min)
	case !hasMin && hasMax:
		return fmt.Sprintf("Integer[_, %d]", c.max)
	default:
		return IntegerS
	}
}
func (c *intChecker) TypeString() string {
	return IntegerS
}

func (c *intChecker) Check(v any) (bool, string) {
	if v == nil {
		return false, "nil value is not an Integer"
	}
	var val int64
	switch i := v.(type) {
	case int64:
		val = i
	case int32:
		val = int64(i)
	case int:
		val = int64(i)
	case int16:
		val = int64(i)
	case int8:
		val = int64(i)
	default:
		return false, "value is not an Integer"
	}
	if val >= c.min && val <= c.max {
		return true, ""
	}
	if val < c.min {
		return false, fmt.Sprintf("Integer value %d is < min %d", val, c.min)
	}
	return false, fmt.Sprintf("Integer value %d is > max %d", val, c.max)
}

func (c *intChecker) Refine(instructions []any) TypeChecker {
	var ok bool
	template := *DefaultIntChecker
	switch len(instructions) {
	case 2:
		if x := instructions[1]; x != nil {
			template.max, ok = utils.GetInt64(x)
			if !ok {
				return nil
			}
		}
		fallthrough
	case 1:
		if x := instructions[0]; x != nil {
			template.min, ok = utils.GetInt64(x)
			if !ok {
				return nil
			}
		}
	case 0:
		return DefaultIntChecker
	default:
		return nil
	}
	return &template
}
