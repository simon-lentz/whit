package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/lithammer/shortuuid/v4"
)

// ShortUUID returns a short Base57 encoded verion 4 UUID suitable to use as ID string.
func ShortUUID() string {
	return shortuuid.New()
}

// AssignIds will replace all occurrences of a particular $$([w])+ pattern with the same
// $$:<UUID>:\1 pattern where UUID is unique. True is returned if there were any changes.
// The UUID is obtained from the given uuidGen function (for example utils.ShortUUID).
func AssignIds(s string, uuidGen func() string) (result string, changed bool) {
	r := regexp.MustCompile(`\$\$\w+`)
	matches := r.FindAllString(s, -1)

	if matches == nil {
		return s, false
	}
	rGlobal := regexp.MustCompile(`\$\$:\w+:(\w+)`)
	globalMatches := rGlobal.FindAllString(s, -1)
	localToGlob := map[string]string{}
	for _, g := range globalMatches {
		parts := strings.Split(g, ":")
		localToGlob[parts[2]] = g
	}
	for _, id := range matches {
		local := id[2:]
		if g, ok := localToGlob[local]; ok {
			s = strings.ReplaceAll(s, fmt.Sprintf("$$%s", local), g)
		}
		short := uuidGen()
		replacement := fmt.Sprintf("$$:%s:%s", short, local)
		s = strings.ReplaceAll(s, id, replacement)
	}

	return s, true
}
