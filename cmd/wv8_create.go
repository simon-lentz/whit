// Package cmd contains the example hello CLI logic.
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tada/catch"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/wyrth-io/whit/internal/for/wv8"
	"github.com/wyrth-io/whit/internal/validation"
	"github.com/wyrth-io/whit/internal/yammm"
)

var wv8CreateCmd = &cobra.Command{
	Use:   "wv8 create --schema <schema.yammm> --cluster <weaviate cluster address> --file data.json",
	Short: `Stores data in a Weaviate database.`,
	Long: `Stores data in a Weaviate database endpoint according to a schema. The given schema and file are
validated to be free from errors before updating the Weaviate store. There is no check that the given
schema matches the schema of the store.

Note that this command is not idempotent. There is no update support.
`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := mustHaveYammmContext()

		// Create the wv8 client
		cfg := weaviate.Config{
			Host:   wv8StoreURL,
			Scheme: "https",
		}

		client, err := weaviate.NewClient(cfg)
		if err != nil {
			panic(err)
		}
		exitOnError(err, "Could not create Weaviate client: %s", err)

		// Read the json file and validate it.
		graph := readJZONInstance(file)

		validator := yammm.NewValidator(ctx, file, graph)
		ic := validation.NewIssueCollector()
		validator.Validate(ic)
		presentor := validation.NewColorPresentor()
		err = presentor.Present(ic, validation.Info, os.Stderr)
		if ic.HasErrors() || ic.HasFatal() {
			if err != nil {
				fmt.Printf("could not present errors from validation due to: %s", err.Error())
			}
			os.Exit(1)
		}

		// Create instances in Weaviate DB.
		err = catch.Do(func() {
			creator := wv8.NewCreator(client)
			creator.Process(ctx, graph)
		})
		exitOnError(err, "Error while storing instance graph: %s", err)
	},

	Args: func(cmd *cobra.Command, args []string) error {
		// validate flags/options here return nil if all is fine else an error
		if wv8StoreURL == "" {
			return fmt.Errorf("no --cluster specified")
		}
		if schemaFile == "" {
			return fmt.Errorf("no --schema file given")
		}
		if !strings.HasSuffix(schemaFile, ".yammm") {
			return fmt.Errorf("schema file must be a .yammm file")
		}
		if file == "" {
			return fmt.Errorf("no --file given")
		}
		return nil
	},
}

func init() {
	wv8Cmd.AddCommand(wv8CreateCmd)

	flags := wv8CreateCmd.PersistentFlags()
	flags.StringVarP(&file, "file", "f", "", "instance file to read in JSON")
	flags.StringVarP(&schemaFile, "schema", "s", "", "Yammm model in Yammm")
	flags.StringVarP(&wv8StoreURL, "cluster", "c", "", "Weaviate cluster URL")
}
