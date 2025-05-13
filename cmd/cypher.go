package cmd

import (
	"context"
	"fmt"
	"os"
	"syscall"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/spf13/cobra"
	"github.com/wyrth-io/whit/internal/for/cypher"
	"golang.org/x/term"
)

var cypherCmd = &cobra.Command{
	Use:   "cypher <subcommand>",
	Short: "Performs a Cypher related subcommand",
	Long: `This is an umbrella command for actions on Cypher compliant DB (such as Neo4j).
The typical usage is to first create a DB, then it is initialized with a schema and then
data is merged into it. See the subcommands "create", "init", and "merge" for more information.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},

	Args: func(cmd *cobra.Command, args []string) error {
		// validate flags/options here return nil if all is fine else an error
		return nil
	},
}

var dbName string

func init() {
	RootCmd.AddCommand(cypherCmd)
	// Nothing here for this umbrella command, if there are any common flags for all subcommands they can
	// be added here, for example:
	flags := cypherCmd.PersistentFlags()
	flags.StringVarP(&dbName, "db", "d", "", "The name of the DB to operate on")
}

// eceuteWriteOperation takes a function that will be given to neo4j API for execution.
func executeWriteOperation(
	ctx context.Context,
	useDB string,
	f func(transaction neo4j.ManagedTransaction) (any, error),
) (any, error) {
	dbURL, ok := os.LookupEnv("DBURL")
	if !ok {
		dbURL = "neo4j://localhost:7687"
	}
	dbUser, ok := os.LookupEnv("DBUSER")
	if !ok {
		dbUser = "neo4j" // default
	}
	dbPwd, ok := os.LookupEnv("DBPWD")
	if !ok {
		// Ask for password if stdin is a tty
		if !term.IsTerminal(int(syscall.Stdin)) { //nolint:unconvert
			fmt.Println("input is not a terminal and no DBPWD is set")
			os.Exit(1)
		}
		fmt.Print("Password: ")
		bytepw, err := term.ReadPassword(int(syscall.Stdin)) //nolint:unconvert
		if err != nil {
			os.Exit(1)
		}
		dbPwd = string(bytepw)
	}

	driver, err := neo4j.NewDriverWithContext(dbURL, neo4j.BasicAuth(dbUser, dbPwd, ""))
	if err != nil {
		fmt.Printf("could not create neo4j driver: %s", err.Error())
		os.Exit(1)
	}
	defer driver.Close(ctx) //nolint:errcheck
	config := neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}
	if useDB != "" {
		config.DatabaseName = useDB
	}
	session := driver.NewSession(ctx, config)
	defer session.Close(ctx) //nolint:errcheck

	result, err := session.ExecuteWrite(ctx, f)
	if err != nil {
		fmt.Printf("ExecuteWrite ended with error: %s\n", err.Error())
	}
	return result, err
}

// executeStatements executes the given statements in a transaction.
func executeStatements(ctx context.Context, statements []cypher.Statement) (any, error) {
	workLoad := func(transaction neo4j.ManagedTransaction) (any, error) {
		records := []any{}
		for i := range statements {
			result, err := transaction.Run(ctx,
				statements[i].Source,
				statements[i].Parameters,
			)
			if err != nil {
				return records, fmt.Errorf("error while executing statement[%d]=%s: %s",
					i, statements[i].Source, err.Error())
			}
			for result.Next(ctx) {
				records = append(records, result.Record().Values[0])
			}
		}
		return records, nil
	}
	return executeWriteOperation(ctx, dbName, workLoad)
}
