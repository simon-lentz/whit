package csvgen

import (
	"regexp"
	"strconv"
	"strings"
)

var scientificPattern = regexp.MustCompile(`[Ee][+-]?[0-9]+$`)

// IsFloat returns true if the given string most likely represents a Float value.
// It returns true if:
//   - The value can be converted to signed Float64
//   - Contains a single "." or ends with [Ee][+-]digits
//   - If it starts with "0", the next char must be "."
//
// Thereby not recognizing "01" or "1" as being Float.
func IsFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return false
	}
	periodIdx := strings.Index(s, ".")
	periodLast := strings.LastIndex(s, ".")
	if periodIdx != periodLast {
		return false
	}
	hasScientific := scientificPattern.MatchString(s)
	if periodIdx != -1 || hasScientific {
		if s[0] == '-' {
			s = s[0:]
		}
		if s[0] == '0' && s[1] != '.' {
			return false
		}
		return true
	}
	return false
}

// IsInteger returns true if the given string most likely represents an Integer value.
// It returns true if:
//   - The value can be converted to signed Int64
//   - Contains only digits and an optional leading -
//   - Is exactly "0" or digit sequence not starting with "0"
//
// Thereby not recognizing "01" as being Integer.
func IsInteger(s string) bool {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return false
	}
	if i < 0 {
		return s[1] != '0'
	}
	if i > 0 {
		return s[0] != '0'
	}
	return true
}

// IsBoolean returns true if the given string most likely represents a boolean value.
// It returns true if:
//   - it is case independently "true" or "false"
//
// While "0" and "1" are also convertible it would be wrong to take those integer values to mean boolean
// by default.
func IsBoolean(s string) bool {
	lowered := strings.ToLower(s)
	return lowered == "true" || lowered == "false"
}
