package cmd

import (
	"bufio"
	"errors"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wyrth-io/whit/internal/validation"
	"github.com/wyrth-io/whit/parser"
)

// TODO: Add support for --no-color and --warnings-are-errors.

var yammmParseCmd = &cobra.Command{
	Use:   "parse [--out <filename>] <schema.yammm>",
	Short: "Parses and validates a Yammm DSL model.",
	Long: `Parses and validates a Yammm model expressed in Yammm DSL form. Optionally if the --out option
is used the command will generate the model in JSON format to that file. All commands that need a schema
as input will accept the DSL directly. Producing a JSON file is mostly for debugging purposes but could
be used as an interchange format to an application where a DSL parser is not available.
`,
	Run: func(cmd *cobra.Command, args []string) {
		yammmCtx, ic := parser.ParseFile(args[0])
		if yammmCtx == nil {
			// Present outcome from validation (which could error depending of where it was redirected),
			// but still highly unlikely - panic in this unlikely case.
			err := validation.NewColorPresentor().Present(ic, validation.Info, os.Stdout)
			if err != nil {
				panic(err)
			}
			switch {
			case ic.HasFatal():
				log.Fatalf("Fatal error(s) occurred - no output produced")
			case ic.HasErrors():
				log.Fatalf("Error(s) occurred - no output produced")
			default:
				log.Fatalf("internal error: parser did not produce a context (reason unknown)")
			}
		}
		// Present outcome from validation (which could error depending of where it was redirected),
		// but still highly unlikely - panic in this unlikely case.
		err := validation.NewColorPresentor().Present(ic, validation.Info, os.Stdout)
		if err != nil {
			panic(err)
		}

		// var f *os.File

		if outFile != "" {
			f, err := os.Create(outFile)
			if err != nil {
				log.Fatalf("Cannot write to '%s: %s\n", outFile, err.Error())
			}
			defer f.Close() //nolint:all
			w := bufio.NewWriter(f)
			err = yammmCtx.WriteModelAsJSON(w)
			if err != nil {
				log.Fatalf("Writing of JSON outtput failed: %s\n", err)
			}
			err = w.Flush()
			if err != nil {
				log.Fatalf("Flushing of JSON outtput failed: %s\n", err)
			}
		}
	},

	Args: func(cmd *cobra.Command, args []string) error {
		red := color.New((color.FgRed)).SprintFunc()

		// validate flags/options here return nil if all is fine else an error
		if len(args) > 1 {
			return errors.New(red("at most one argument accepted"))
		}
		if len(args) != 1 {
			return errors.New(red("no yammm file given"))
		}
		if outFile == args[0] {
			return errors.New(red("input and output files are the same"))
		}
		return nil
	},
}

func init() {
	yammmCmd.AddCommand(yammmParseCmd)

	flags := yammmParseCmd.PersistentFlags()
	flags.StringVarP(&outFile, "out", "o", "", "the file to write json output to")
}
