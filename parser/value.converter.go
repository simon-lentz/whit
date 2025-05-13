package parser

import "strings"

const (
	dq = `"`
	sq = `'`
)

// ConvertString converts a parsed double or single quoted string to a Go string.
// At the moment this function only removes the enclosing quotes but should also translate
// escape sequences to control characters which is left TBD.
func ConvertString(s string) string {
	if strings.HasPrefix(s, dq) {
		s = strings.TrimPrefix(s, dq)
		s = strings.TrimSuffix(s, dq)
	} else if strings.HasPrefix(s, sq) {
		s = strings.TrimPrefix(s, sq)
		s = strings.TrimSuffix(s, sq)
	}
	return s
}

// toLowerFirstC returns a string with first character in lower case. Only works for strings where
// the first character is in ASCII range.
func toLowerFirstC(s string) string {
	if len(s) != 0 && (s[0] <= 'Z' && s[0] >= 'A') {
		return string(s[0]+' ') + s[1:]
	}
	return s
}

// stripDelimiters strips leading /* and trailing */.
func stripDelimiters(doc string) string {
	first := strings.Index(doc, "/*")
	last := strings.LastIndex(doc, "*/")
	delimiterSize := 2
	doc = doc[first+delimiterSize : last]
	return strings.Trim(doc, " \t")
}
