package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wyrth-io/whit/internal/scrapers"
)

var collectCmd = &cobra.Command{
	Use:   "collect <subcommand>",
	Short: "Performs a data collection/scraper routine.",
	Long:  `Sample command: whit collect nha`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},

	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func addSubcommand() {
	collectCmd.AddCommand(scrapers.NhaCmd)
}

func init() {
	addSubcommand()
	RootCmd.AddCommand(collectCmd)
}
