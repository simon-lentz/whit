package cmd

import (
	"github.com/spf13/cobra"
)

var wv8SchemaCmd = &cobra.Command{
	Use:   "wv8 schema <subcommand>",
	Short: "Runs a Weaviate schema related <subcommand>",
	Long: `This is an umbrella command for various Weaviate schema related commands.
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},

	Args: func(cmd *cobra.Command, args []string) error {
		// validate flags/options here return nil if all is fine else an error
		return nil
	},
}

// The name of the genmodel file to use.
var genmodelFile string

func init() {
	wv8Cmd.AddCommand(wv8SchemaCmd)

	// Nothing here for this umbrella command, if there are any common flags for all subcommands they can
	// be added here, for example:
	// flags := helloCmd.PersistentFlags()
	// flags.IntVarP(&Port, "port", "p", 8088, "The port the server is listening on")
}
