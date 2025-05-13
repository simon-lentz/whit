package tc

import (
	"fmt"
	"math"

	"github.com/wyrth-io/whit/internal/utils"
)

// DefaultFloatChecker is a TypeChecker for any float.
var DefaultFloatChecker = &floatChecker{min: -math.MaxFloat64, max: math.MaxFloat64}

type floatChecker struct {
	baseChecker
	min float64
	max float64
}

func (c *floatChecker) BaseType() BaseType {
	return valueType(FloatKind)
}
func (c *floatChecker) SyntaxString() string {
	hasMin := c.min != -math.MaxFloat64
	hasMax := c.max != math.MaxFloat64
	switch {
	case hasMin && hasMax:
		return fmt.Sprintf("Float[%g, %g]", c.min, c.max)
	case hasMin && !hasMax:
		return fmt.Sprintf("Float[%g, _]", c.min)
	case !hasMin && hasMax:
		return fmt.Sprintf("Float[_, %g]", c.max)
	default:
		return FloatS
	}
}
func (c *floatChecker) TypeString() string {
	return FloatS
}

func (c *floatChecker) Check(v any) (bool, string) {
	if v == nil {
		return false, "nil value is not a Float"
	}
	var val float64
	switch i := v.(type) {
	case float64:
		val = i
	case float32:
		val = float64(i)
	case int64:
		val = float64(i)
	case int32:
		val = float64(i)
	case int:
		val = float64(i)
	case int16:
		val = float64(i)
	case int8:
		val = float64(i)
	default:
		return false, "value is not a Float"
	}
	if val >= c.min && val <= c.max {
		return true, ""
	}
	if val < c.min {
		return false, fmt.Sprintf("Float value %g is < min %g", val, c.min)
	}
	return false, fmt.Sprintf("Float value %g is > max %g", val, c.max)
}

func (c *floatChecker) Refine(instructions []any) TypeChecker {
	var ok bool
	template := *DefaultFloatChecker
	switch len(instructions) {
	case 2:
		if x := instructions[1]; x != nil {
			template.max, ok = utils.GetFloat64(x)
			if !ok {
				if max, ok := utils.GetInt64(x); ok {
					template.max = float64(max)
				} else {
					return nil
				}
			}
		}
		fallthrough
	case 1:
		if x := instructions[0]; x != nil {
			template.min, ok = utils.GetFloat64(x)
			if !ok {
				if min, ok := utils.GetInt64(x); ok {
					template.min = float64(min)
				} else {
					return nil
				}
			}
		}
	case 0:
		return DefaultFloatChecker
	default:
		return nil
	}
	return &template
}
