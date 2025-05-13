package csvgen_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/for/csvgen"
	"github.com/wyrth-io/whit/internal/pio"
)

func TestProcessor(t *testing.T) {
	tt := testutils.NewTester(t)

	model := `schema "test"
	type Person {
		name String primary
		age Integer required
		popularity Float required
		isImaginary Boolean required
		-->OWNS (many) Car { since Date }
	}
	type Car {
		regNbr String primary
	}
	`
	ctx, _ := makeContext(t.Name(), model, true)

	genModel := `{
	"type": "Person",
	"property_map": {"name": "NAME", "age": "AGE", "popularity": "RATING", "isImaginary": "IMAGINARY"},
	"association_map": {
		"OWNS_Cars": { 
			"property_map": { "since": "OWNERSHIP_SINCE"},
			"where": { "regNbr": "REGISTRATION_NUMBER"}
		}
	}
}`
	csvData := `NAME,AGE,REGISTRATION_NUMBER,OWNERSHIP_SINCE,RATING,IMAGINARY
Fred Flintstone,42,ABC123,2023-01-02,1,true
Henry Ford,63,T1,1908-01-02,0.2,false
`
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	writer := pio.WriterOn(w)

	file := makeTmpFile(csvData)
	defer os.Remove(file.Name()) //nolint:errcheck

	var gm csvgen.Genmodel
	err := json.Unmarshal([]byte(genModel), &gm)
	tt.CheckNotError(err)
	err = csvgen.Process(file.Name(), ctx, gm, writer)
	tt.CheckNotError(err)
	err = w.Flush()
	tt.CheckNotError(err)
	generated := buf.String()

	// validate the generated json (to keep the generator honest).
	x, _ := ValidateModelMessages(t, model, generated)
	tt.CheckNotNil(x)

	fmt.Println(generated)
	// tt.CheckTrue(false) // uncomment to see output
	// TODO: Assert generated.
}

// TODO: Test missing values.
