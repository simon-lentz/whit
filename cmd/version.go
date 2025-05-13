package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the current version.",
	Long: `This command prints the current version.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(whitVersion)
	},

	Args: func(cmd *cobra.Command, args []string) error {
		// validate flags/options here return nil if all is fine else an error
		return nil
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
