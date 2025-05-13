// Package pio (Protected IO) contains various io functons that perform an underlying
// io call (like in the package catch/pio).
package pio

import (
	"os"

	"github.com/tada/catch"
)

// Close closes the given file and panics with a check.Error if the f.Close() errors.
func Close(f *os.File) {
	err := f.Close()
	if err != nil {
		panic(catch.Error(err))
	}
}
