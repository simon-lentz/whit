package utils

// IntType is a type constraint for generics dscribing all signed integer types.
type IntType interface {
	int | int8 | int16 | int32 | int64
}

// Min returns the smaller of the two given integer values.
func Min[T IntType](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the greater of the two given integer values.
func Max[T IntType](a, b T) T {
	if a > b {
		return a
	}
	return b
}
