package tc

// DefaultBoolChecker is the (singleton) TypeChecker for Boolan values.
var DefaultBoolChecker = &boolChecker{}

type boolChecker struct {
	baseChecker
}

func (c *boolChecker) SyntaxString() string {
	return "Boolean"
}
func (c *boolChecker) TypeString() string {
	return c.SyntaxString()
}
func (c *boolChecker) BaseType() BaseType {
	return valueType(BoolKind)
}
func (c *boolChecker) Check(v any) (bool, string) {
	if v == nil {
		return false, "nil value is not a Boolean"
	}
	if _, ok := v.(bool); ok {
		return true, ""
	}
	return false, "value is not a Boolean"
}
func (c *boolChecker) Refine(instr []any) TypeChecker {
	if len(instr) == 0 {
		return c
	}
	return nil
}
