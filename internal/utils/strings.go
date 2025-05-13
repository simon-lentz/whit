// Package utils contains various utility functions.
package utils

import (
	"regexp"
	"strings"
	"unicode"
)

// Capitalize returns a new string with initial capital letter.
func Capitalize(s string) string {
	if len(s) < 1 {
		return s
	}
	a := []rune(s)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}

// DeCapitalize returns a new string with initial capital letter.
func DeCapitalize(s string) string {
	if len(s) < 1 {
		return s
	}
	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

// CapitalizeAll returns a new slice where each entry is Capitalized.
func CapitalizeAll(list []string) []string {
	result := make([]string, len(list))
	for i, s := range list {
		result[i] = Capitalize(s)
	}
	return result
}

// UniqueFold returns a unique list of strings where equality is unicode case independent.
// The first found unique value is retained in the output with original case.
func UniqueFold(list []string) []string {
	var result []string
	set := make(map[string]bool)
	for _, s := range list {
		sl := strings.ToLower(s)
		if !set[sl] {
			result = append(result, s)
			set[sl] = true
		}
	}
	return result
}

// DeleteFold deletes all instances of matching given s in list. The match is case independent.
func DeleteFold(list []string, s string) []string {
	result := make([]string, 0, len(list))
	for _, entry := range list {
		if !strings.EqualFold(entry, s) {
			result = append(result, entry)
		}
	}
	return result
}

// CamelToSnake transforms a "camelCase" or "CamelCase" to snake form; "camel_Case", "Camel_Case" respectively.
// Case of letters are not changed. Multiple consecutive "_" are changed to single "_".
func CamelToSnake(s string) (s2 string) {
	if isAllUpperCase((s)) {
		return s
	}
	re := regexp.MustCompile(`^[a-z_][^A-Z]*|[A-Z][^A-Z]*`)
	submatchall := re.FindAllString(s, -1)
	s2 = strings.Join(submatchall, "_")
	// Complicated regexp, "_ or more at start, capture all in $1" or "one or more anywhere, capture the first".
	// Replace with "$1$2" either replaces with captured $1, if the match was at the beginning (in which case $2 is empty),
	// or if the match is anywhere else, capture $1 is empty and $2 is the capture of one _ while the match is multiple
	// underscores. Thus "___a___b___" becomes "___a_b_"
	matcher := regexp.MustCompile(`^(_+)|(_)+`)
	s2 = matcher.ReplaceAllString(s2, `$1$2`)
	s2 = fixAcronyms(s2)
	return
}
func isAllUpperCase(s string) bool {
	for _, r := range s {
		if r < 'A' || r > 'Z' {
			return false
		}
	}
	return true
}
func fixAcronyms(s string) string {
	isUpperCase := func(s string) bool { return s[0] >= 'A' && s[0] <= 'Z' }
	safeSlice := func(s string, start int) string {
		if start < 1 || len(s) <= start {
			return ""
		}
		return s[start : start+1]
	}
	var result string
	for i := range s {
		lb1 := safeSlice(s, i-1)
		lb2 := safeSlice(s, i-2)
		lh1 := safeSlice(s, i+1)
		lh2 := safeSlice(s, i+2)
		lh3 := safeSlice(s, i+3)
		lh4 := safeSlice(s, i+4)
		this := s[i : i+1]
		if lb2 == "_" && isUpperCase(lb1) && this == "_" && isUpperCase(lh1) && lh2 == "_" && isUpperCase(lh3) && !isUpperCase(lh4) {
			continue
		}
		result += this
	}
	return result
}
