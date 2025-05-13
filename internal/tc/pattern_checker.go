package tc

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/wyrth-io/whit/internal/utils"
)

type patternChecker struct {
	baseChecker
	patterns []*regexp.Regexp
}

// DefaultPatternChecker is a  Pattern checker that matches nothing.
var DefaultPatternChecker = &patternChecker{patterns: []*regexp.Regexp{}}

func (c *patternChecker) BaseType() BaseType {
	return valueType(StringKind)
}

func (c *patternChecker) SyntaxString() string {
	quoted := utils.Map(c.patterns, func(r *regexp.Regexp) string { return fmt.Sprintf("%q", r.String()) })
	return fmt.Sprintf("Pattern[%s]", strings.Join(quoted, ", "))
}
func (c *patternChecker) TypeString() string {
	return "Pattern"
}

func (c *patternChecker) Check(v any) (bool, string) {
	if v == nil {
		return false, "nil value is not a String, cannot match a Pattern"
	}
	if s, ok := v.(string); ok {
		for i := range c.patterns {
			if c.patterns[i].Match([]byte(s)) {
				return true, ""
			}
		}
		keys := []string{}
		for _, k := range c.patterns {
			keys = append(keys, `"`+k.String()+`"`)
		}
		sort.Strings(keys)
		return false, fmt.Sprintf("String value '%s' does not match Pattern[%s]", s, strings.Join(keys, ", "))
	}
	return false, "value is not a String, cannot match a Pattern"
}

func (c *patternChecker) Refine(instr []any) TypeChecker {
	if len(instr) == 0 {
		return c
	}
	ec := &patternChecker{patterns: make([]*regexp.Regexp, len(instr))}
	for i := range instr {
		if s, ok := instr[i].(string); ok {
			ec.patterns[i] = regexp.MustCompile(s)
		} else if r, ok := instr[i].(*regexp.Regexp); ok {
			ec.patterns[i] = r
		} else {
			return nil
		}
	}
	return ec
}
