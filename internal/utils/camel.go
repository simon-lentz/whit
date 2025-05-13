package utils

import (
	"regexp"
	"strings"
)

// ToLowerCamel converts a string to initial lower case camelCase. The given string is broken up
// into segments delimited by non words and where a lower case letter is followed by an upper case
// letters. Each segment except the first (which is decapitalied) are capitalized. Should two numeric segments
// end up next to eachother they will be separated by an underscore ("_"). Empty segments are ignored.
func ToLowerCamel(s string) string {
	return toCamel(s, true)
}

// ToUpperCamel converts a string to initial upper case CamelCase. The given string is broken up
// into segments delimited by non words and where a lower case letter is followed by an upper case
// letters. Each segment except the first (which is decapitalied) are capitalized. Should two numeric segments
// end up next to eachother they will be separated by an underscore ("_"). Empty segments are ignored.
func ToUpperCamel(s string) string {
	return toCamel(s, false)
}

func toCamel(s string, toLower bool) string {
	nonWord := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	result := nonWord.ReplaceAllString(s, "_")
	// No lookahed/lookbehind in go regexp
	// instead replace all occurrences of "lowercase" or digits followed by "uppercase"
	lowerUpper := regexp.MustCompile(`([a-z]|[0-9])([A-Z])`)
	result = lowerUpper.ReplaceAllString(result, "${1}_${2}")

	// split the string on _
	// Then downcase first slice initial char, and upcase first char of all others.
	// If two elements next to each other are numeric insert a separating _.
	slices := strings.Split(result, "_")
	for i := 0; i < len(slices); i++ {
		slices[i] = Capitalize(slices[i])
	}
	if toLower {
		slices[0] = DeCapitalize(slices[0])
	}
	output := make([]string, 0, len(slices))
	numEnd := false
	for _, s := range slices {
		// skip empty
		if len(s) == 0 {
			continue
		}
		// Add _ between two number segments
		first := s[0]
		if numEnd && first < '9' && first > '0' {
			output = append(output, "_")
		}
		last := s[len(s)-1]
		numEnd = (last < '9' && last > '0')
		output = append(output, s)
	}
	return strings.Join(output, "")
}
