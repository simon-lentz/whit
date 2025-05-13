package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wyrth-io/whit/internal/pio"
	"github.com/wyrth-io/whit/internal/validation"
	"github.com/wyrth-io/whit/internal/yammm"
	"github.com/wyrth-io/whit/parser"
)

var yammmCmd = &cobra.Command{
	Use:   "yammm <subcommand>",
	Short: "Performs a Yammm schema/model related subcommand.",
	Long: `This is an umbrella command for actions on Yammm models.
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
	RootCmd.AddCommand(yammmCmd)

	// Nothing here for this umbrella command, if there are any common flags for all subcommands they can
	// be added here, for example:
	// flags := arrowsCmd.PersistentFlags()
	// flags.IntVarP(&Port, "port", "p", 8088, "The port the server is listening on")
}

func getYammmCtx(fileName string) (ctx yammm.Context, ic validation.IssueCollector, ok bool) {
	ok = true
	if strings.HasSuffix(fileName, ".yammm") {
		ctx, ic = parser.ParseFile(fileName)
		if ctx == nil {
			switch {
			case ic.HasFatal():
				// Present outcome from validation (which could error depending of where it was redirected),
				// but still highly unlikely - panic in this unlikely case.
				_ = validation.NewColorPresentor().Present(ic, validation.Info, os.Stdout)
				log.Fatalf("Fatal error(s) occurred - no output produced")
			case ic.HasErrors():
				// Present outcome from validation (which could error depending of where it was redirected),
				// but still highly unlikely - panic in this unlikely case.
				_ = validation.NewColorPresentor().Present(ic, validation.Info, os.Stdout)
				log.Fatalf("Error(s) occurred - no output produced")
			default:
				log.Fatalf("internal error: parser did not produce a context (reason unknown)")
			}
		}
	} else {
		ic = validation.NewTerminatingIssueCollector()
		ctx = yammm.NewContext()
		reader, err := os.Open(fileName) //nolint:gosec
		ic.CollectFatalIfErrorf("Error opening yammm json file: %s", err)
		defer pio.Close(reader)
		err = ctx.SetModelFromJSON(reader)
		ic.CollectFatalIfErrorf("Error parsing yammm json file: %s", err)
		validation.Do(func() {
			ok = ctx.Complete(ic)
		})
	}
	return ctx, ic, ok
}
