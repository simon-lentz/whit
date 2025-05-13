// Package cmd contains the example hello CLI logic.
package cmd

import (
	"github.com/spf13/cobra"
)

var yamlCmd = &cobra.Command{
	Use:   "yaml <subcommand>",
	Short: "Performs a yaml related <subcommand>",
	Long: `This is an umbrella command for yaml related commands.
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},

	Args: func(cmd *cobra.Command, args []string) error {
		// validate flags/options here return nil if all is fine else an error
		return nil
	},
}

func init() {
	RootCmd.AddCommand(yamlCmd)

	// Nothing here for this umbrella command, if there are any common flags for all subcommands they can
	// be added here, for example:
	// flags := yamlCmd.PersistentFlags()
	// flags.IntVarP(&Port, "port", "p", 8088, "The port the server is listening on")
}
