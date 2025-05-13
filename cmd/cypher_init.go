// Package cmd contains the example hello CLI logic.
package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/spf13/cobra"
	"github.com/wyrth-io/whit/internal/for/cypher"
	"github.com/wyrth-io/whit/internal/utils"
)

var cypherInitCmd = &cobra.Command{
	Use:   "init --db <dbname> --schema  <schemaname.yammm>",
	Short: "Initializes a db with a schema.",
	Long: `Initializes the given db with the given schema. The db must have been created first.
This will create indexes and constraints in the DB to ensure as far as it is possile to
have the DB uphold the schema integrity of the DB.

The schema can be given as a Yammm DSL file or as a JSON serialization of a Yammm model. The schema
will be validated before the initialization takes place and any errors will stop the execution.

The location of the DB server, user name and password are given via environment variables.
If password is not given via environment variable and the input is a tty, the command will prompt
for the password.

The environment variables and their default values are:
	DBURL   neo4j://localhost:7687" 
	DBUSER  neo4j
	DBPWD

These defaults are suitable for running a local DB.

Use the sub command "cypher init show" to see the statements that init will execute.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		red := color.New((color.FgRed)).SprintFunc()

		if schemaFile == "" {
			fmt.Println(red("No schema file given!"))
		}
		// Create context and meta
		yctx := mustHaveYammmContext()

		// Generate cyphercode for schema
		statements := cypher.GenerateInit(yctx)

		ctx := context.Background()
		idxops, stmnts := utils.Partition(statements, func(s cypher.Statement) bool { return s.NeedsSepTx })
		makeWorkLoad := func(s []cypher.Statement) func(neo4j.ManagedTransaction) (any, error) {
			return func(transaction neo4j.ManagedTransaction) (any, error) {
				records := []any{}
				for i := range s {
					result, err := transaction.Run(ctx,
						s[i].Source,
						s[i].Parameters,
					)
					if err != nil {
						return records, fmt.Errorf("error while executing statement[%d]=%s: %s",
							i, s[i].Source, err.Error())
					}
					for result.Next(ctx) {
						records = append(records, result.Record().Values[0])
					}
				}
				return records, nil
			}
		}
		workLoad := makeWorkLoad(stmnts)
		result, err := executeWriteOperation(ctx, dbName, workLoad)
		if err != nil {
			fmt.Printf("ExecuteWrite ended with error: %s\n", err.Error())
			os.Exit(1)
		}
		for i, r := range result.([]any) {
			rr := r.(*neo4j.Record)
			fmt.Printf("Result [%d]: %v\n", i, rr.Values)
		}
		workLoad = makeWorkLoad(idxops)

		result, err = executeWriteOperation(ctx, dbName, workLoad)
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
		if len(args) > 1 {
			return errors.New("at most one argument accepted")
		}
		return nil
	},
}

func init() {
	cypherCmd.AddCommand(cypherInitCmd)

	flags := cypherInitCmd.PersistentFlags()
	flags.StringVarP(&schemaFile, "schema", "s", "", "Yammm model in YAMMM to validate against")
}
