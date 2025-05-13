package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wyrth-io/whit/internal/for/csvgen"
	"github.com/wyrth-io/whit/internal/pio"
	"github.com/wyrth-io/whit/internal/validation"
)

var csvConvertCmd = &cobra.Command{
	Use:   "convert --out <filename> --schema [<yammmfile.[yammm|json>]] --genmodel <genmodelfile> <csvfile>",
	Short: "Converts a csv file to Yammm compliant JSON.",
	Long: `Converts csv file to yamm compliant JSON format under the direction of a genmodel file for mapping
csv columns to property values and relations. The genmodel is a JSON file describing how columns are mapped
to properties of a type, and how columns can be used to form associations to other types (already loaded in
a DB).

The conversion will convert data types from csv string form to integer, float and boolean values using a best effort.
Booleans are case insensitive and accept "true", "t", "1", "yes", "y" as being boolean true, and "false", "f",
"0", "-1", "no", "n", and "" (empty string) as being boolean false. All data types with a string base type are used
verbatim.

When mapping associations the properties of the association are mapped separately from those properties
that are primary keys (the "where" clause). The where part must either contain either only the mapping of "id"
(a UUID in stringform), or all other primary keys of the target type. The where part can reference any instance
of the target type as there is no check that they exist at the time of conversion to JSON. They must however
exist when importing/merging to a DB.

Note that this version cannot convert compositions (i.e. extraction of some columns to an instance of a separate
type). To achieve this, model it as an association and make a separate genmodel for each type to extract. Then
extract the main part with associations to the other. 

Example
=======
Given a model like this:

schema "example"
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

And csv file format like this:

NAME,AGE,REGISTRATION_NUMBER,OWNERSHIP_SINCE,RATING,IMAGINARY
Fred Flintstone,42,ABC123,2023-01-02,0.8,true
Henry Ford,63,T1,1908-01-02,0.2,false

Then this genmodel like this:

{   "type": "Person",
	"property_map": {"name": "NAME", "age": "AGE", "popularity": "RATING", "isImaginary": "IMAGINARY"},
	"association_map": {
		"OWNS_Cars": { 
			"property_map": { "since": "OWNERSHIP_SINCE"},
			"where": { "regNbr": "REGISTRATION_NUMBER"}
		}
	}
}

Would map the csv to a YAMM model instance in JSON, like this:

{
    "People":
[
      {
        "OWNS_Cars": [
          {
            "Where": {
              "regNbr": "ABC123"
            },
            "since": "2023-01-02"
          }
        ],
        "age": 42,
        "isImaginary": true,
        "name": "Fred Flintstone",
        "popularity": 0.8
      },
      {
        "OWNS_Cars": [
          {
            "Where": {
              "regNbr": "T1"
            },
            "since": "1908-01-02"
          }
        ],
        "age": 63,
        "isImaginary": false,
        "name": "Henry Ford",
        "popularity": 0.2
      }
    ]
}
`,
	Run: func(cmd *cobra.Command, args []string) {
		var w *bufio.Writer
		var writer *pio.Writer

		ctx, ic, ok := getYammmCtx(schemaFile)
		fatalPrinter := color.New(color.FgRed, color.Bold).SprintFunc()

		// Present outcome from validation (which could error depending of where it was redirected),
		// but still highly unlikely.
		err := validation.NewColorPresentor().Present(ic, validation.Info, os.Stdout)
		if err != nil {
			fmt.Println(fatalPrinter("fatal: %s", err))
			os.Exit(1)
		}
		if ic.HasFatal() {
			os.Exit(1)
		}
		if ic.HasErrors() {
			os.Exit(2)
		}
		if !ok {
			// Should not really happen, ok == false only if there were fatals or errors.
			// But if there are bugs in the validation pkg this may go wrong.
			return
		}

		// Prepare the output
		var f *os.File
		if outFile == "-" || outFile == "" {
			f = os.Stdout
		} else {
			f, err = os.Create(outFile)
			if err != nil {
				fmt.Println(fmt.Sprintf(fatalPrinter("Cannot write to '%s': %s"), outFile, err.Error()))
				os.Exit(1)
			}
		}
		defer pio.Close(f) // close and panic on error
		w = bufio.NewWriter(f)
		writer = pio.WriterOn(w)

		// Load Genmodel
		var gm csvgen.Genmodel
		reader, err := os.Open(genmodelFile)
		exitOnError(err, "Could not read genmodel file %s: %s", genmodelFile, err)
		data, err := io.ReadAll(reader)
		exitOnError(err, "Could not read genmodel file %s: %s", genmodelFile, err)
		err = json.Unmarshal(data, &gm)
		exitOnError(err, "Could not unmarshal JSON in genmodel file %s: %s", genmodelFile, err)

		err = csvgen.Process(args[0], ctx, gm, writer)
		if err != nil {
			fmt.Println(fatalPrinter("Could not process csv: %s", err.Error()))
		}

		// Make sure buffered io on output file is flushed to disc
		err = w.Flush()
		if err != nil {
			fmt.Println(fatalPrinter("Could not flush generated otput: %s", err.Error()))
		}
	},

	Args: func(cmd *cobra.Command, args []string) error {
		red := color.New((color.FgRed)).SprintFunc()

		if len(args) > 1 {
			return errors.New(red("at most one csv file accepted"))
		}
		if len(args) != 1 {
			return errors.New(red("no csv file given"))
		}
		if len(genmodelFile) < 1 {
			return errors.New(red("no genmodel file given"))
		}
		return nil
	},
}

func init() {
	csvCmd.AddCommand(csvConvertCmd)

	flags := csvConvertCmd.PersistentFlags()
	flags.StringVarP(&outFile, "out", "o", "", "the file to write output to")
	flags.StringVarP(&genmodelFile, "genmodel", "g", "", "filename of csv genmodel")
	flags.StringVarP(&schemaFile, "schema", "s", "", ".yammm or .json schema file")
}
