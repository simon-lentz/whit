package utils

// Ptr returns a pointer to the given (literal) value.
// This is used since `&42` is not valid go code.
func Ptr[T any](v T) *T {
	return &v
}
