package pio

import (
	"fmt"
	"io"

	"github.com/tada/catch"
)

// Writer wraps an io.Writer and provides Printf() and Println() with error checking.
// If an error occurs, the functions will panic with a catch.Error.
type Writer struct {
	w      io.Writer
	indent int
}

// IndentedWriter returns a new Writer with 4 spaces increased indentation.
// A call to Indent() outputs the indentation amount of spaces. This will panic if
// nested more than 100 indentations deep.
func IndentedWriter(w *Writer) *Writer {
	if w.indent > 100-4 {
		panic("Too much indentation - not supported")
	}
	return &Writer{w: w.w, indent: w.indent + 4}
}

// OutdentedWriter returns a new Writer with 4 spaces decreased indentation.
// A call to Indent() outputs the indentation amount of spaces. This will panic if
// outdented when indent is 0.
func OutdentedWriter(w *Writer) *Writer {
	if w.indent == 0 {
		panic("Not enough indentation - negative indentation not supported")
	}
	return &Writer{w: w.w, indent: w.indent - 4}
}

// Indent output the amount of indentation (a number of spaces) set for this Writer.
func (w *Writer) Indent() {
	if w.indent == 0 {
		return
	}
	spacesFormat := "%" + fmt.Sprintf(`%d`, w.indent) + "s"
	w.Printf(spacesFormat, " ")
}

// Spaces returns a string with spaces for the current amount of indentation.
func (w *Writer) Spaces() string {
	if w.indent == 0 {
		return ""
	}
	spacesFormat := "%" + fmt.Sprintf(`%d`, w.indent) + "s"
	return fmt.Sprintf(spacesFormat, " ")
}

// WriterOn creates a Writer for an io.Writer.
func WriterOn(w io.Writer) *Writer {
	return &Writer{w: w}
}

// Println behaves as fmt.Println but to this Writer's configured io.Writer.
// Any error from the underlying Fprintln will result in a panic(catch.Error(err)).
func (w *Writer) Println(a ...any) {
	_, err := fmt.Fprintln(w.w, a...)
	if err != nil {
		panic(catch.Error(err))
	}
}

// Printf behaves as fmt.Printf but to this Writer's configured io.Writer.
// Any error from the underlying Fprintf will result in a panic(catch.Error(err)).
func (w *Writer) Printf(s string, a ...any) {
	_, err := fmt.Fprintf(w.w, s, a...)
	if err != nil {
		panic(catch.Error(err))
	}
}

// Indentedf is the same as first calling w.Indent() and then w.Printf().
func (w *Writer) Indentedf(s string, a ...any) *Writer {
	w.Indent()
	w.Printf(s, a...)
	return w
}

// Format is the same as first calling w.Indent() and then w.Printf().
func (w *Writer) Format(s string, a ...any) *Writer {
	w.Indent()
	w.Printf(s, a...)
	return w
}

// FormatLn is the same as first calling w.Indent() and then w.Printf() with a final new line.
func (w *Writer) FormatLn(s string, a ...any) *Writer {
	w.Indent()
	w.Printf(s, a...)
	w.Printf("\n")
	return w
}

// FormatJoinLn outputs each line in the given lines slice indented, joining lines with the given
// join string if there is one more element in lines. Each line, including the last is ended with a new line.
func (w *Writer) FormatJoinLn(lines []string, join string) *Writer {
	for i, txt := range lines {
		if i > 0 {
			w.Printf("%s\n", join)
		}
		w.Indent()
		w.Printf("%s", txt)
	}
	w.Printf("\n")
	return w
}

// Indented returns a writer indented more than this writer.
func (w *Writer) Indented() *Writer {
	return IndentedWriter(w)
}

// Outdented returns a writer with less indentation that this writer.
func (w *Writer) Outdented() *Writer {
	return OutdentedWriter(w)
}
