// Package cmd contains the example hello CLI logic.
package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/spf13/cobra"
	"github.com/wyrth-io/whit/internal/for/cypher"
	"github.com/wyrth-io/whit/internal/validation"
	"github.com/wyrth-io/whit/internal/yammm"
)

var cypherMergeCmd = &cobra.Command{
	Use:   "merge --db dbname --schema <schemafile>.yammm --file <file.[json|csv]>",
	Short: "Merges schema compliant instance data into the DB.",
	Long: `Merges the data in the given file according to the given schema into the given DB. The DB must
have been initialized with the a compatible version of the schema.

The location of the DB server, user name and password are given via environment variables.
If password is not given via environment variable and the input is a tty, the command will prompt
for the password.

The environment variables and their default values are:
	DBURL   neo4j://localhost:7687" 
	DBUSER  neo4j
	DBPWD

These defaults are suitable for running a local DB.

See the command "cypher merge show" for how to view the Cypher statements that will be executed
by the "merge" command.
`,
	Run: func(cmd *cobra.Command, args []string) {
		red := color.New((color.FgRed)).SprintFunc()
		// Create context and meta
		yctx := mustHaveYammmContext()

		if schemaFile == "" {
			fmt.Println(red("No schema file given!"))
		}
		// Read the json file and create a context for the schema.
		var graph any
		if strings.HasSuffix(file, ".csv") {
			graph = readCSVInstance(file, yctx, genmodelFile)
		} else {
			graph = readJZONInstance(file)
		}
		// Validate the instance
		ic := validation.NewTerminatingIssueCollector()
		validator := yammm.NewValidator(yctx, file, graph)
		validator.Validate(ic)

		if ic.HasErrors() {
			_ = validation.NewColorPresentor().Present(ic, validation.Info, os.Stderr)
			os.Exit(1)
		}

		// Generate the cypher statements
		merger := cypher.NewMergeGenerator()
		statements := merger.Process(yctx, graph)

		ctx := context.Background()
		result, err := executeStatements(ctx, statements)
		if err != nil {
			fmt.Printf("ExecuteWrite ended with error: %s\n", err.Error())
			os.Exit(1)
		}
		for i, r := range result.([]any) {
			rr := r.(*neo4j.Record)
			fmt.Printf("Result [%d]: %v\n", i, rr.Values)
		}
		os.Exit(0)
	},

	Args: func(cmd *cobra.Command, args []string) error {
		// validate flags/options here return nil if all is fine else an error
		if len(args) > 0 {
			return errors.New("no additional argument accepted")
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
	cypherCmd.AddCommand(cypherMergeCmd)

	flags := cypherMergeCmd.PersistentFlags()
	flags.StringVarP(&file, "file", "f", "", "instance file to read in JSON, - for stdin")
	flags.StringVarP(&schemaFile, "schema", "s", "", "Yammm model in YAMMM to validate against")
	flags.StringVarP(&genmodelFile, "genmodel", "g", "", "filename of csv genmodel mappings json file")
}
