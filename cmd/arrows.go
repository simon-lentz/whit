package cmd

import (
	"github.com/spf13/cobra"
)

var arrowsCmd = &cobra.Command{
	Use:   "arrows <subcommand>",
	Short: "Performs an arrows related sub command",
	Long: `This is an umbrella command for actions on neo4j arrows files
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
	RootCmd.AddCommand(arrowsCmd)

	// Nothing here for this umbrella command, if there are any common flags for all subcommands they can
	// be added here, for example:
	// flags := arrowsCmd.PersistentFlags()
	// flags.IntVarP(&Port, "port", "p", 8088, "The port the server is listening on")
}
