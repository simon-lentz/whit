// Package cmd contains the example hello CLI logic.
package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
)

var wv8SchemaViewCmd = &cobra.Command{
	Use:   "wv8 schema view --cluster <weaviate cluster address",
	Short: "Obtains and outputs a Weaviate schema from a Weaviate endpoint.",
	Long: `Obtains and outputs a Weaviate schema from a Weaviate endpoint. The schema is output
in the native Weaviate JSON format.
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create the wv8 client
		cfg := weaviate.Config{
			Host:   wv8StoreURL,
			Scheme: "https",
		}

		client, err := weaviate.NewClient(cfg)
		exitOnError(err, "%s", err)

		schemaDump, err := client.Schema().Getter().Do(context.Background())
		exitOnError(err, "%s", err)

		// Display Schema as JSON.
		bytes, err := json.MarshalIndent(schemaDump, "", "  ")
		exitOnError(err, "%s", err)

		fmt.Printf("%s\n", bytes)
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
	wv8SchemaCmd.AddCommand(wv8SchemaViewCmd)

	flags := wv8SchemaViewCmd.PersistentFlags()
	flags.StringVarP(&wv8StoreURL, "cluster", "c", "", "Weaviate cluster URL")
}
