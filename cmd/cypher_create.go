package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/spf13/cobra"
)

var cypherCreateCmd = &cobra.Command{
	Use:   "create --db name",
	Short: "Creates a new cypher database.",
	Long: `Creates a new cypher database with a given name. The location of the DB server,
user name and password are given via environment variables. If password is not given via environment
variable and the input is a tty, the command will prompt for the password.

The environment variables and their default values are:
    DBURL   neo4j://localhost:7687" 
    DBUSER  neo4j
    DBPWD

These defaults are suitable for running a local DB.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		workLoad := func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				`CREATE DATABASE $dbName`,
				map[string]any{"dbName": dbName},
			)
			if err != nil {
				fmt.Printf("could not create database '%s': %s", dbName, err.Error())
			}
			records := []any{}
			for result.Next(ctx) {
				records = append(records, result.Record().Values[0])
			}
			return records, err
		}
		result, err := executeWriteOperation(ctx, "", workLoad)

		if err != nil {
			fmt.Printf("ExecuteWrite ended with error: %s\n", err.Error())
			os.Exit(1)
		}
		for i, r := range result.([]any) {
			rr := r.(*neo4j.Record)
			fmt.Printf("Result [%d]: %v\n", i, rr.Values)
		}
	},

	Args: func(cmd *cobra.Command, args []string) error {
		// validate flags/options here return nil if all is fine else an error
		return nil
	},
}

func init() {
	cypherCmd.AddCommand(cypherCreateCmd)
	// Nothing here as this uses --db to get name of db inherited from `cypherCmd`
}
