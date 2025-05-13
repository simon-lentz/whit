// Package gogen has logic for generating go code from a yammm Model.
package gogen

import "fmt"

// tagJ returns a yaml tag string on the form: `yaml:"name"`.
// The opt argment accepts a boolean that if true will include ",omitempty".
func tagJ(s string, opt bool) string {
	if opt {
		return fmt.Sprintf(`json:"%s,omitempty"`, s)
	}
	return fmt.Sprintf(`json:"%s"`, s)
}

// tagY returns a yaml tag string on the form: `yaml:"name"`.
// The opt argment accepts a boolean that if true will include ",omitempty".
func tagY(s string, opt bool) string {
	if opt {
		return fmt.Sprintf(`yaml:"%s,omitempty"`, s)
	}
	return fmt.Sprintf(`yaml:"%s"`, s)
}

// TagJ returns a json tag string on the form: `json:"name"`.
// The opt argment accepts a boolean that if true will include ",omitempty".
func TagJ(s string, opt ...bool) string {
	return fmt.Sprintf("`%s`", tagJ(s, len(opt) > 0 && opt[0]))
}

// TagY returns a yaml tag string on the form: `yaml:"name"`.
// The opt argment accepts a boolean that if true will include ",omitempty".
func TagY(s string, opt ...bool) string {
	return fmt.Sprintf("`%s`", tagY(s, len(opt) > 0 && opt[0]))
}

// TagJY returns a tag string with both json and yaml tags on the form: `json:"name" yaml:"name"`.
// The opt argment accepts a boolean that if true will include ",omitempty" for both tags.
func TagJY(s string, opt ...bool) string {
	optional := len(opt) > 0 && opt[0]
	return fmt.Sprintf("`%s %s`", tagJ(s, optional), tagY(s, optional))
}
