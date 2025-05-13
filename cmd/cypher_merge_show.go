// Package cmd contains the example hello CLI logic.
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wyrth-io/whit/internal/for/cypher"
)

var cypherDumpCmd = &cobra.Command{
	Use:   "show --schema <schema.yammm> --file <data.[json|csv]>",
	Short: "Shows the Cypher statements that 'cypher merge' would execute.",
	Long: `Shows the generated cypher statements as they would be executed by the "cypher merge" command.
This command is mostly useful for debugging and prints out the statements and a map of the parameters
that accompany them. The statements are numbered and in color if output is to a tty.
`,
	Run: func(cmd *cobra.Command, args []string) {
		blue := color.New((color.FgBlue)).SprintFunc()
		green := color.New((color.FgGreen)).SprintFunc()

		ctx := mustHaveYammmContext()

		// Read the json file and create a context for the schema.
		var graph any
		if strings.HasSuffix(file, ".csv") {
			graph = readCSVInstance(file, ctx, genmodelFile)
		} else {
			graph = readJZONInstance(file)
		}
		// Generate the cypher statements
		merger := cypher.NewMergeGenerator()
		statements := merger.Process(ctx, graph)

		for i := range statements {
			// Print statement and pretty print the parameters map
			bytes, _ := json.MarshalIndent(statements[i].Parameters, "    ", "  ") //nolint:all
			fmt.Printf("%s %s %s\n",
				blue(fmt.Sprintf("[%.3d]", i)), // index
				green(statements[i].Source),    // cypher text
				string(bytes))                  // the map of properties
		}
		os.Exit(0)
	},

	Args: func(cmd *cobra.Command, args []string) error {
		// validate flags/options here return nil if all is fine else an error
		if len(args) > 1 {
			return errors.New("at most one argument accepted")
		}
		if file == "" {
			return errors.New("no file given")
		}
		if strings.HasSuffix(file, ".csv") {
			if genmodelFile == "" {
				return errors.New("no --genmodel file given")
			}
		} else {
			if genmodelFile != "" {
				return errors.New("--genmodel only supported for CSV files")
			}
		}

		if schemaFile == "" {
			return errors.New("no schema file given")
		}

		return nil
	},
}

func init() {
	cypherMergeCmd.AddCommand(cypherDumpCmd)

	flags := cypherDumpCmd.PersistentFlags()
	flags.StringVarP(&file, "file", "f", "", "instance file to read in JSON or CSV, - for stdin")
	flags.StringVarP(&schemaFile, "schema", "s", "", "schema in Yamm to validate against")
	flags.StringVarP(&genmodelFile, "genmodel", "g", "", "filename of csv genmodel mappings json file")
}
