// Package cmd contains the example hello CLI logic.
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wyrth-io/whit/internal/validation"
	"github.com/wyrth-io/whit/internal/yammm"
	"github.com/wyrth-io/whit/parser"
)

var yammmValidateCmd = &cobra.Command{
	Use:   "validate --schema <schemafile>.[yammm | json] --file instance.[json|yaml|yml|csv]",
	Short: "Validates a data/instance file against a Yammm schema.",
	Long: `Validates the given instance file in one of the supported formats (determined by the file extension)
against the given schemafile in JSON or Yammm DSL format. This command is useful to check the correctness of a data file
before an attempt is used to import it to a DB.
`,
	Run: func(cmd *cobra.Command, args []string) {
		red := color.New((color.FgRed)).SprintFunc()

		if file == "" {
			fmt.Println(red("No file given!"))
			cmd.HelpFunc()(cmd, args)
			os.Exit(1)
		}
		if schemaFile == "" {
			fmt.Println(red("No schema file given!"))
		}

		// Create context and meta
		var ctx yammm.Context
		var ic validation.IssueCollector
		if strings.HasSuffix(schemaFile, ".yammm") {
			ctx, ic = parser.ParseFile(schemaFile)
			if ctx == nil {
				switch {
				case ic.HasFatal():
					log.Fatalf("Fatal error(s) occurred - no output produced")
				case ic.HasErrors():
					log.Fatalf("Error(s) occurred - no output produced")
				default:
					log.Fatalf("internal error: parser did not produce a context (reason unknown)")
				}
			}
		} else {
			reader, err := os.Open(schemaFile)
			if err != nil {
				fmt.Println(red("Could not read schema file %s: %s", schemaFile, err))
				os.Exit(1)
			}
			ctx = yammm.NewContext()
			err = ctx.SetModelFromJSON(reader)
			if err != nil {
				fmt.Println(red("Error when setting model from JSON: %s", err))
				os.Exit(1)
			}
			ic = validation.NewTerminatingIssueCollector()
			ctx.Complete(ic)
			if ic.HasErrors() {
				presentor := validation.NewColorPresentor()
				_ = presentor.Present(ic, validation.Error, os.Stderr)
				os.Exit(1)
			}
		}
		// Read the instance file
		var graph any
		switch {
		case strings.HasSuffix(file, ".json"):
			graph = readJZONInstance(file)
		case strings.HasSuffix(file, ".yaml") || strings.HasSuffix(file, ".yml"):
			graph = readYamlInstance(file)
		case strings.HasSuffix(file, ".csv"):
			graph = readCSVInstance(file, ctx, genmodelFile)
		default:
			fmt.Println(red("File must have .json, .yaml, .yml, or .csv suffix"))
			os.Exit(1)
		}

		// Validate the instance
		validator := yammm.NewValidator(ctx, file, graph)
		validator.Validate(ic)

		if ic.HasErrors() {
			_ = validation.NewColorPresentor().Present(ic, validation.Info, os.Stderr)
			os.Exit(1)
		}

		os.Exit(0)
	},

	Args: func(cmd *cobra.Command, args []string) error {
		// validate flags/options here return nil if all is fine else an error
		if len(args) > 1 {
			return errors.New("at most one argument accepted")
		}
		if strings.HasSuffix(file, ".csv") && genmodelFile == "" {
			return errors.New("genmodel mapping filename must be provided")
		}
		return nil
	},
}

func init() {
	yammmCmd.AddCommand(yammmValidateCmd)

	flags := yammmValidateCmd.PersistentFlags()
	flags.StringVarP(&file, "file", "f", "", "instance file to read in JSON, - for stdin")
	flags.StringVarP(&schemaFile, "schema", "s", "", "Yammm model in YAMMM or JSON to validate against")
	flags.StringVarP(&genmodelFile, "genmodel", "g", "", "filename of csv genmodel mappings json file")
}
