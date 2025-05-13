package utils

import (
	"reflect"
	"strings"
)

// TypeOrder orders types canonically and returns 1 if left has a type greater than right, 0 if
// types are of the same strata, and -1 type left is less than type right.
// The stratas are nil < numbers < strings < slices.
func TypeOrder(left, right any) int {
	leftStrata := TypeStrata(left)
	rightStrata := TypeStrata(right)

	switch {
	case leftStrata > rightStrata:
		return 1
	case leftStrata == rightStrata:
		return 0
	default:
		return -1
	}
}

// NilStrata is the lowest strata followed by
// BoolStrata, NumericStrata, StringStrata, and SliceStrata.
const (
	InvalidStrata = iota
	NilStrata
	BoolStrata
	NumericStrata
	StringStrata
	SliceStrata
)

// TypeStrata returns a strata depending on type. NilStrata is the lowest strata followed by
// BoolStrata, NumericStrata, StringStrata, and SliceStrata.
func TypeStrata(a any) int {
	if a == nil {
		return NilStrata
	}
	switch reflect.TypeOf(a).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return NumericStrata
	case reflect.Float32, reflect.Float64:
		return NumericStrata
	case reflect.String:
		return StringStrata
	case reflect.Slice:
		return SliceStrata
	case reflect.Bool:
		return BoolStrata
	}
	return InvalidStrata
}

// ValueOrder returns the canonical order of the two values, using TypeOrder to first
// determine the order of the type of the values. If values have the same type order they
// are compared taking data type into account.
func ValueOrder(left, right any) int {
	// If type order is different, no need to compare values.
	if to := TypeOrder(left, right); to != 0 {
		return to
	}
	switch TypeStrata(left) {
	case NilStrata:
		return 0 // the other value must also be nil
	case BoolStrata:
		lb, rb := left.(bool), right.(bool)
		if lb == rb {
			return 0
		}
		if lb {
			return 1
		}
		return -1
	case NumericStrata:
		li, liok := GetInt64(left)
		lf, lfok := GetFloat64(left)
		ri, riok := GetInt64(right)
		rf, rfok := GetFloat64(right)
		switch {
		case liok && riok:
			return Int64Compare(li, ri)
		case lfok && rfok:
			return Float64Compare(lf, rf)
		case lfok && riok:
			return Float64Compare(lf, float64(ri))
		case liok && rfok:
			return Float64Compare(float64(li), rf)
		}

	case StringStrata:
		ls := left.(string)
		rs := right.(string)
		return strings.Compare(ls, rs)

	case SliceStrata:
		leftVo := reflect.ValueOf(left)
		leftLen := leftVo.Len()
		rightVo := reflect.ValueOf(right)
		rightLen := rightVo.Len()
		minLen := Min(leftLen, rightLen)
		for i := 0; i < minLen; i++ {
			leftVal := leftVo.Index(i).Interface()
			rightVal := rightVo.Index(i).Interface()
			which := ValueOrder(leftVal, rightVal)
			if which == 0 {
				continue
			}
			return which
		}
		// Equal up to same length, the shorter is smaller
		if leftLen == rightLen {
			return 0
		}
		if leftLen > rightLen {
			return 1
		}
		return -1
	}
	panic("Unknown value strata")
}

// GetInt64 from any. Returns bool ok true if value was any integer and the value as an Int64.
func GetInt64(val any) (int64, bool) {
	switch x := val.(type) {
	case int:
		return int64(x), true
	case int8:
		return int64(x), true
	case int16:
		return int64(x), true
	case int32:
		return int64(x), true
	case int64:
		return x, true
	}
	return 0, false
}

// GetFloat64 from any. Returns bool ok true if value was any float and the value as a Float64.
func GetFloat64(val any) (float64, bool) {
	switch x := val.(type) {
	case float32:
		return float64(x), true
	case float64:
		return x, true
	}
	return 0, false
}

// Int64Compare compares two int64 and returns 1 if left is > right, 0 if equal and -1
// if right > left.
func Int64Compare(left, right int64) int {
	if left == right {
		return 0
	}
	if left > right {
		return 1
	}
	return -1
}

// Float64Compare compares two float64 and returns 1 if left is > right, 0 if equal and -1
// if right > left.
func Float64Compare(left, right float64) int {
	if left == right {
		return 0
	}
	if left > right {
		return 1
	}
	return -1
}
