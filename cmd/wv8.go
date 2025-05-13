package cmd

import (
	"github.com/spf13/cobra"
)

var wv8Cmd = &cobra.Command{
	Use:   "wv8 <subcommand>",
	Short: "Performs a Weaviate related subcommand.",
	Long: `This is an umbrella command for various Weaviate commands.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},

	Args: func(cmd *cobra.Command, args []string) error {
		// validate flags/options here return nil if all is fine else an error
		return nil
	},
}

// Common flags.

var wv8StoreURL string

func init() {
	RootCmd.AddCommand(wv8Cmd)

	// Nothing here for this umbrella command, if there are any common flags for all subcommands they can
	// be added here, for example:
	// flags := helloCmd.PersistentFlags()
	// flags.IntVarP(&Port, "port", "p", 8088, "The port the server is listening on")
}
