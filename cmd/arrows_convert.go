package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wyrth-io/whit/internal/for/arrows"
	"github.com/wyrth-io/whit/internal/pio"
	"github.com/wyrth-io/whit/internal/validation"
)

// TODO: Add support for --no-color and --warnings-are-errors.

var arrowsConvertCmd = &cobra.Command{
	Use:   "convert [<arrowsfile.json>]",
	Short: "convert",
	Long: `Converts an arrows file to other formats. Supported formats are "cue" and "yammm". The yammm output
is in JSON format."
	`,
	Run: func(cmd *cobra.Command, args []string) {
		name := "" // TODO: Change this - this is the filename not model name
		if len(args) > 0 {
			name = args[0]
		}
		issueCollector := validation.NewTerminatingIssueCollector()

		if list {
			arrowsNodes := func(arrows *arrows.Arrows) []string {
				var lines []string
				for _, n := range arrows.Parsed.Nodes {
					lines = append(lines, "("+n.Caption+")")
				}
				return lines
			}

			arrowsRelations := func(arrows *arrows.Arrows) []string {
				var lines []string
				for _, n := range arrows.Parsed.Relationships {
					lines = append(lines, n.Type)
				}
				return lines
			}

			graph := arrows.New(name)
			graph.Parse(issueCollector)
			fmt.Printf("Nodes: %v\n", strings.Join(arrowsNodes(&graph), ", "))
			fmt.Printf("Relations: %v\n", strings.Join(arrowsRelations(&graph), ", "))
			return
		}
		// show presentation error (cannot write, cannot flush, etc. in the same style as the ColorPresentor)
		// TODO: Maybe implement functions on Presentor interface `SprintFunc(Level, format, ...args)`
		fatalPrinter := color.New(color.FgRed, color.Bold).SprintFunc()
		if outFile == "" {
			fmt.Println(fatalPrinter("no output file given"))
			os.Exit(1)
		}
		var f *os.File
		var err error
		if outFile == "-" {
			f = os.Stdout
		} else {
			f, err = os.Create(outFile)
			if err != nil {
				fmt.Println(fatalPrinter("Cannot write to '%s: %s", outFile, err.Error()))
				os.Exit(1)
			}
			defer pio.Close(f) // close and panic on error
		}
		w := bufio.NewWriter(f)

		graph := arrows.New(name)
		_ = validation.Do(func() {
			graph.Parse(issueCollector)
			switch format {
			case cueFormat:
				graph.CueMarshalMeta(w, issueCollector)
			case yammmFormat:
				modelName := packageName
				if modelName == "" {
					modelName = strings.TrimSuffix(name, filepath.Ext(name))
				}
				graph.MarshalYammm(modelName, w, issueCollector)
			default:
				panic("should not happen - illegal format not filtered out")
			}
		})
		// Present outcome from validation
		err = validation.NewColorPresentor().Present(issueCollector, validation.Info, os.Stdout)
		if err != nil {
			fmt.Println(fatalPrinter("fatal: %s", err))
			os.Exit(1)
		}
		// Make sure buffered io on output file is flushed to disc
		err = w.Flush()
		if err != nil {
			fmt.Println(fatalPrinter("Could not flush generated otput: %s", err.Error()))
		}

		// TODO: Also send to log, or just log? (no color in log)
		if issueCollector.HasErrors() {
			os.Exit(2)
		}
		if issueCollector.HasFatal() {
			os.Exit(1)
		}
	},

	Args: func(cmd *cobra.Command, args []string) error {
		red := color.New((color.FgRed)).SprintFunc()

		// validate flags/options here return nil if all is fine else an error.
		if len(args) > 1 {
			return errors.New(red("at most one argument accepted"))
		}
		if len(args) != 1 {
			return errors.New(red("no arrows file given"))
		}
		if !(format == cueFormat || format == yammmFormat) {
			return errors.New(red("unknown format given to --format"))
		}
		return nil
	},
}

var list bool
var packageName string

func init() {
	arrowsCmd.AddCommand(arrowsConvertCmd)

	flags := arrowsConvertCmd.PersistentFlags()
	flags.BoolVarP(&list, "list", "l", false, "presents a simple list of nodes and relations")
	flags.StringVarP(&format, "format", "f", "", "the output file format")
	flags.StringVarP(&outFile, "out", "o", "", "the file to write output to")
	flags.StringVarP(&packageName, "package", "p", "", "the name of the generated package (defaults to base name of arrows file)")
	arrowsConvertCmd.MarkFlagsMutuallyExclusive("list", "format")

	// TODO: --nowerror, precense of warnings does not cause exit != 0
}
