// Package cmd contains the example hello CLI logic.
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
)

var wv8SchemaDeleteCmd = &cobra.Command{
	Use:   "delete --cluster <weaviate cluster address>",
	Short: "Deletes the schema at a Weaviate endpoint",
	Long: `Deletes the schema at a Weaviate endpoint.
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create the wv8 client
		cfg := weaviate.Config{
			Host:   wv8StoreURL,
			Scheme: "https",
		}
		client, err := weaviate.NewClient(cfg)
		exitOnError(err, "Cannot create weaviate client: %s", err)

		// Delete the schema.
		err = client.Schema().AllDeleter().Do(context.Background())
		exitOnError(err, "Could not delete schema: %s", err)

		// TODO: Only if verbose flag (not yet implemented)
		fmt.Println("Schema deleted.")
	},

	Args: func(cmd *cobra.Command, args []string) error {
		// validate flags/options here return nil if all is fine else an error
		if wv8StoreURL == "" {
			return fmt.Errorf("no --cluster specified")
		}
		return nil
	},
}

func init() {
	wv8SchemaCmd.AddCommand(wv8SchemaDeleteCmd)

	flags := wv8SchemaDeleteCmd.PersistentFlags()
	flags.StringVarP(&wv8StoreURL, "cluster", "c", "", "Weaviate cluster URL")
}
