package utils

// IfTrue is a ternary if-the-else function that returns the yes value if predicate is true
// and no value otherwise.
func IfTrue[T any](pred bool, yes, no T) T {
	if pred {
		return yes
	}
	return no
}
