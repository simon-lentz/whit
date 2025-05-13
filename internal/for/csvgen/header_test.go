package csvgen_test

import (
	"bytes"
	"encoding/csv"
	"errors"
	"strings"
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/for/csvgen"
	"github.com/wyrth-io/whit/internal/pio"
)

func TestMissingValuesError(t *testing.T) {
	tt := testutils.NewTester(t)
	buf := new(bytes.Buffer)
	writer := pio.WriterOn(buf)

	csv := `A,B,C
1,,3
`
	reader := makeCsvReaderOn(csv)
	err := csvgen.HeaderToYammm(t.Name(), writer, reader, "", "")
	var missingErr *csvgen.MissingValuesError
	tt.CheckTrue(errors.As(err, &missingErr))
	tt.CheckEqual([]string{"B"}, missingErr.Cols)
}

func TestDerivedNames(t *testing.T) {
	tt := testutils.NewTester(t)
	buf := new(bytes.Buffer)
	writer := pio.WriterOn(buf)

	csv := `A,B
1,2
`
	reader := makeCsvReaderOn(csv)
	err := csvgen.HeaderToYammm(t.Name(), writer, reader, "Bwah", "Mukfluk")
	tt.CheckNotError(err)
	expected := `/* schema derived from: TestDerivedNames */
schema "Bwah"
type Mukfluk {
    /* A */
    a Integer
    /* B */
    b Integer
}
`
	tt.CheckStringSlicesEqual(strings.Split(expected, "\n"), strings.Split(buf.String(), "\n"))
}
func TestSearchesForValuesToDetermineType(t *testing.T) {
	tt := testutils.NewTester(t)
	buf := new(bytes.Buffer)
	writer := pio.WriterOn(buf)

	csv := `A,B,C
1,,
,2,
,,3
`
	reader := makeCsvReaderOn(csv)
	err := csvgen.HeaderToYammm(t.Name(), writer, reader, "", "")
	tt.CheckNotError(err)

	expected := `/* schema derived from: TestSearchesForValuesToDetermineType */
schema "TestSearchesForValuesToDetermineType"
type TestSearchesForValuesToDetermineType {
    /* A */
    a Integer
    /* B */
    b Integer
    /* C */
    c Integer
}
`
	tt.CheckStringSlicesEqual(strings.Split(expected, "\n"), strings.Split(buf.String(), "\n"))
}

func TestHeaderMapping(t *testing.T) {
	tt := testutils.NewTester(t)
	csvStr := `A,B
`
	reader := makeCsvReaderOn(csvStr)
	buf := new(bytes.Buffer)
	writer := pio.WriterOn(buf)
	err := csvgen.HeaderToMapping(t.Name(), writer, reader, "")
	tt.CheckNotError(err)
	expected := `{
  "type": "TestHeaderMapping",
  "property_map": {
    "a": "A",
    "b": "B"
  },
  "association_map": {}
}
`
	tt.CheckStringSlicesEqual(strings.Split(expected, "\n"), strings.Split(buf.String(), "\n"))
}

func TestHeaderMappingTypeName(t *testing.T) {
	tt := testutils.NewTester(t)
	csvStr := `A,B
`
	reader := makeCsvReaderOn(csvStr)
	buf := new(bytes.Buffer)
	writer := pio.WriterOn(buf)
	err := csvgen.HeaderToMapping(t.Name(), writer, reader, "Mukfluk")
	tt.CheckNotError(err)
	expected := `{
  "type": "Mukfluk",
  "property_map": {
    "a": "A",
    "b": "B"
  },
  "association_map": {}
}
`
	tt.CheckStringSlicesEqual(strings.Split(expected, "\n"), strings.Split(buf.String(), "\n"))
}

func makeCsvReaderOn(s string) *csv.Reader {
	buf := bytes.NewBufferString(s)
	return csv.NewReader(buf)
}
