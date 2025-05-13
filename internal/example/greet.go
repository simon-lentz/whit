// Package example contains example logic.
package example

import "strings"

// Greet returns the basic type of greeting.
func Greet(s string) string {
	out := "hello world"
	if s != "" {
		return out + " " + s
	}
	return out
}

// GreetWonderful returns the "wonderful" typoe of greeing.
func GreetWonderful(s string) string {
	out := "hello wonderful world"
	if s != "" {
		return out + " " + s
	}
	return out
}

// GreetUpper returns the a greeing in upper case.
func GreetUpper(s string) string {
	return strings.ToUpper(Greet(s))
}

// GreetWonderfulUpper returns the "wonderful" type of greeing in upper case.
func GreetWonderfulUpper(s string) string {
	return strings.ToUpper(GreetWonderful(s))
}
