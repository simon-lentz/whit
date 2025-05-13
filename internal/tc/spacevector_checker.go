package tc

import (
	"fmt"
	"math"

	"github.com/wyrth-io/whit/internal/utils"
)

// DefaultSpaceVectorChecker is the (singleton) TypeChecker for Spacevector (dimension 0) space
// vectors.
var DefaultSpaceVectorChecker = &spacevectorChecker{}

type spacevectorChecker struct {
	baseChecker
	dimensions int
}

func (c *spacevectorChecker) SyntaxString() string {
	return fmt.Sprintf("Spacevector[%d]", c.dimensions)
}
func (c *spacevectorChecker) TypeString() string {
	return "Spacevector"
}

func (c *spacevectorChecker) BaseType() BaseType {
	return spaceType(c.dimensions)
}
func (c *spacevectorChecker) Check(v any) (bool, string) {
	if v == nil {
		return false, "nil value is not a Spacevector"
	}
	var vectorLength int
	switch vals := v.(type) {
	case []float32:
		vectorLength = len(vals)
	case []float64:
		vectorLength = len(vals)
		for i := range vals {
			f := vals[i]
			if f < -math.MaxFloat32 || f > math.MaxFloat32 {
				return false, "value is outside range of allowe spacevector float32 values"
			}
		}
	default:
		return false, "value is not a Spacevector value"
	}
	if vectorLength != c.dimensions {
		return false, fmt.Sprintf("length of given spacevector must be %d: got %d",
			c.dimensions, vectorLength)
	}
	return true, ""
}
func (c *spacevectorChecker) Refine(instr []any) TypeChecker {
	if len(instr) == 0 {
		return c
	}
	if len(instr) != 1 {
		return nil
	}
	dim, ok := utils.GetInt64(instr[0])
	if !ok {
		return nil
	}
	return &spacevectorChecker{dimensions: int(dim)}
}
