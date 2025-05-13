// Package cmd contains the example hello CLI logic.
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wyrth-io/whit/internal/for/cypher"
)

var cypherInitShowCmd = &cobra.Command{
	Use:   "show --schema <schema.yammm>",
	Short: "Shows the Cypher statements init will execute.",
	Long: `Shows the generated cypher statements as they would be executed by the "cypher init" command.
This command is mostly useful for debugging and prints out the statements and a map of the parameters
that accompany them. The statements are numbered and in color if output is to a tty.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		red := color.New((color.FgRed)).SprintFunc()
		blue := color.New((color.FgBlue)).SprintFunc()
		green := color.New((color.FgGreen)).SprintFunc()

		if schemaFile == "" {
			fmt.Println(red("No schema file given!"))
		}
		ctx := mustHaveYammmContext()
		// Generate the cypher statements
		statements := cypher.GenerateInit(ctx)

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
		if len(args) > 0 {
			return errors.New("no arguments accepted")
		}
		return nil
	},
}

func init() {
	cypherInitCmd.AddCommand(cypherInitShowCmd)

	flags := cypherInitShowCmd.PersistentFlags()
	flags.StringVarP(&schemaFile, "schema", "s", "", "schema in Yamm to validate against")
}
