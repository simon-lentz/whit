// Package cmd contains the example hello CLI logic.
package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
)

var wv8MetaCmd = &cobra.Command{
	Use:   "wv8 meta --cluster <weaviate cluster address>",
	Short: "Gets Weaviate meta information for a Weaviate endpoint",
	Long: `Gets Weaviate meta information for a Weaviate endpoint for the purpose of checking
current settings, for example to verify that the init command produced the expected result.
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create the wv8 client
		cfg := weaviate.Config{
			Host:   wv8StoreURL, // Replace with your endpoint
			Scheme: "https",
		}

		client, err := weaviate.NewClient(cfg)
		exitOnError(err, "%s", err)

		metaGetter := client.Misc().MetaGetter()
		meta, err := metaGetter.Do(context.Background())
		exitOnError(err, "Error while getting meta information: %s", err)

		fmt.Printf("Weaviate meta information:\n")
		fmt.Printf("Hostname: %s version: %s\n", meta.Hostname, meta.Version)
		bytes, err := json.MarshalIndent(meta.Modules, "", "  ")
		exitOnError(err, "Error while creating JSON output of enabled modules: %s", err)
		fmt.Printf("Enabled modules:%s\n", string(bytes))
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
	wv8Cmd.AddCommand(wv8MetaCmd)

	flags := wv8MetaCmd.PersistentFlags()
	flags.StringVarP(&wv8StoreURL, "cluster", "c", "", "Weaviate cluster URL")
}
