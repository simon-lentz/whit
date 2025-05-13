package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wyrth-io/whit/internal/utils"
)

var replaceIDCmd = &cobra.Command{
	Use:   "replaceid <filename>",
	Short: `Expands all local "$$id" to global "$$:UUID:id`,
	Long: `This command replaces $$id strings in the input to $$<UUID>:id string thus turning local
id values and references to global. Note that this kind of expansion is performed automatically
by commands when needed. (Thus there is no requirement to run this command in a standard workflow.
This command is however of value if there is a need to know the resulting UUID values beforehand.
`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, fn := range args {
			fi, err := os.Lstat(fn)
			if err != nil {
				fmt.Printf("Cannot stat %s: %s\n", fn, err.Error())
			}

			data, err := os.ReadFile(fn) //nolint:gosec
			if err != nil {
				_, _ = fmt.Printf("Cannot read %s: %s\n", fn, err.Error())
				return
			}
			newContents, wasChanged := utils.AssignIds(string(data), utils.ShortUUID)
			if !wasChanged {
				continue
			}
			// write back to the same file
			if err = os.WriteFile(fn, []byte(newContents), fi.Mode()); err != nil {
				fmt.Printf("Cannot write %s: %s\n", fn, err)
			}
		}
	},

	Args: func(cmd *cobra.Command, args []string) error {
		// validate flags/options here return nil if all is fine else an error
		return nil
	},
}

func init() {
	RootCmd.AddCommand(replaceIDCmd)

	// Nothing here for this umbrella command, if there are any common flags for all subcommands they can
	// be added here, for example:
	// flags := replaceIdCmd.PersistentFlags()
	// flags.IntVarP(&Port, "port", "p", 8088, "The port the server is listening on")
}
