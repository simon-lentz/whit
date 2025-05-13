package cmd

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wyrth-io/whit/internal/for/csvgen"
	"github.com/wyrth-io/whit/internal/pio"
	"github.com/wyrth-io/whit/internal/utils"
)

var csvHeaderCmd = &cobra.Command{
	Use:   "header --out <filename> [--format mapping|yammm] <csvfile>",
	Short: "Converts csv header to Yammm DSL or csv genmodel.",
	Long: `This command converts the header row of a csv file according to the wanted format.
The format 'mapping' will generate a genmodel property map to be used when converting the csv to yammm
instance data. The format 'yammm' will produce a yammm schema with a single type having properties
derived from the csv content (header and sampled data).
For generation of Yammm, it is important that all columns have values in row 1. If that is not the
case in the dataset a copy of the data should be made of the two first rows, and typical values
added where they are missing, then using this copy to compute the data types of the columns.
`,
	Run: func(cmd *cobra.Command, args []string) {

		fatalPrinter := color.New(color.FgRed, color.Bold).SprintFunc()

		// Prepare the output
		var fo *os.File
		var err error
		if outFile == "-" || outFile == "" {
			fo = os.Stdout
		} else {
			fo, err = os.Create(outFile)
			if err != nil {
				fmt.Println(fmt.Sprintf(fatalPrinter("Cannot write to '%s': %s"), outFile, err.Error()))
				os.Exit(1)
			}
		}
		defer pio.Close(fo) // close and panic on error
		w := bufio.NewWriter(fo)
		writer := pio.WriterOn(w)

		// Create csv Reader
		f, err := os.Open(args[0])
		if err != nil {
			fmt.Println(fmt.Sprintf(fatalPrinter("Could not read '%s': %s"), args[0], err.Error()))
			os.Exit(1)
		}
		reader := csv.NewReader(f)

		switch format {
		case "mapping":
			err := csvgen.HeaderToMapping(args[0], writer, reader, typeName)
			if err != nil {
				fmt.Println(fmt.Sprintf(fatalPrinter("Could not produce mapping for '%s': %s"),
					args[0], err.Error()))
				os.Exit(1)
			}

		case "yammm":
			// Yamm, outputs a yammm model.
			err = csvgen.HeaderToYammm(args[0], writer, reader, schemaName, typeName)
			var missingErr *csvgen.MissingValuesError
			if err != nil {
				if errors.As(err, &missingErr) {
					output := strings.Join(
						utils.Map(missingErr.Cols,
							func(s string) string { return fmt.Sprintf("%q", s) }),
						", ")
					fmt.Println(fatalPrinter(
						"The following columns are empty - the generated data type may therefore not be correct: %s",
						output))
				} else {
					fmt.Println(fatalPrinter("Could generate yammm output: %s", err.Error()))
				}
				os.Exit(1)
			}
		}
		err = w.Flush()
		if err != nil {
			fmt.Println(fatalPrinter("Could not flush generated output (may be truncated): %s", err.Error()))
			os.Exit(1)
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
		if !(format == "yammm" || format == "mapping") {
			return errors.New(red("illegal --format, must be either 'yammm' or 'mapping'"))
		}
		return nil
	},
}

var schemaName string
var typeName string

func init() {
	csvCmd.AddCommand(csvHeaderCmd)

	flags := csvHeaderCmd.PersistentFlags()
	flags.StringVarP(&outFile, "out", "o", "", "the file to write output to")
	flags.StringVarP(&format, "format", "f", "yammm", "the format to use for output ('yammm' or 'mapping')")
	flags.StringVarP(&schemaName, "schemaname", "s", "", "the name of the schema in yammm output")
	flags.StringVarP(&typeName, "typename", "t", "", "the name of the type in yamm output")
}
