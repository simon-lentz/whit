package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	yammmcue "github.com/wyrth-io/whit/internal/for/cue"
	"github.com/wyrth-io/whit/internal/for/gogen"
	"github.com/wyrth-io/whit/internal/for/jschema"
	"github.com/wyrth-io/whit/internal/for/llmgen"
	"github.com/wyrth-io/whit/internal/for/markdown"
	"github.com/wyrth-io/whit/internal/for/wv8"
	"github.com/wyrth-io/whit/internal/pio"
	"github.com/wyrth-io/whit/internal/validation"
	"github.com/wyrth-io/whit/internal/yammm"
)

// TODO: Add support for --no-color and --warnings-are-errors.

var yammmConvertCmd = &cobra.Command{
	Use:   "convert --format [cue|go|jschema|md] --out <filename> [<yammmfile.[yammm|json>]]",
	Short: "Converts a Yammm schema/model to other formats.",
	Long: `Converts a Yammm model file in Yamm DSL or JSON format to other formats. Any errors appear on
stderr if output is to stdout.

The available formats are "go", "cue", "jschema" (JSONschema), "wv8" (weaviate), "llm", and "md" (Markdown).
`,
	Run: func(cmd *cobra.Command, args []string) {
		var w *bufio.Writer
		var writer *pio.Writer

		ctx, ic, ok := getYammmCtx(args[0])

		fatalPrinter := color.New(color.FgRed, color.Bold).SprintFunc()

		// Present outcome from validation (which could error depending of where it was redirected),
		// but still highly unlikely.
		err := validation.NewColorPresentor().Present(ic, validation.Info, os.Stdout)
		if err != nil {
			fmt.Println(fatalPrinter("fatal: %s", err))
			os.Exit(1)
		}
		if ic.HasFatal() {
			os.Exit(1)
		}
		if ic.HasErrors() {
			os.Exit(2)
		}
		if !ok {
			// Should not really happen, ok == false only if there were fatals or errors.
			// But if there are bugs in the validation pkg this may go wrong.
			return
		}

		var f *os.File
		if outFile == "-" || outFile == "" {
			f = os.Stdout
		} else {
			f, err = os.Create(outFile)
			if err != nil {
				fmt.Println(fmt.Sprintf(fatalPrinter("Cannot write to '%s': %s"), outFile, err.Error()))
				os.Exit(1)
			}
			defer pio.Close(f) // close and panic on error
		}
		w = bufio.NewWriter(f)
		writer = pio.WriterOn(w)

		switch format {
		case goFormat:
			gogen.Marshal(ctx, writer)

		case cueFormat:
			yammmcue.Marshal(ctx, writer)

		case jschemaFormat, jsonschemaFormat:
			jschema.Marshal(ctx, writer)

		case markdownFormat, mdFormat:
			markdown.Marshal(ctx, writer)

		case weaviateFormat, wv8Format:
			gm := yammm.NewGenmodel()
			gm.Generator = "wv8"
			if genmodelFile != "" {
				reader, err := os.Open(genmodelFile)
				exitOnError(err, "Could not read genmodel file %s: %s", genmodelFile, err)
				data, err := io.ReadAll(reader)
				exitOnError(err, "Could not read genmodel file %s: %s", genmodelFile, err)
				err = json.Unmarshal(data, gm)
				exitOnError(err, "Could not unmarshal JSON in genmodel file %s: %s", genmodelFile, err)
				if gm.Generator != "wv8" {
					fmt.Println(fatalPrinter("The genmodel file is not for weaviate: got %s", gm.Generator))
					os.Exit(1)
				}
			}

			err := wv8.GenerateSchemaText(ctx, gm, w)
			if err != nil {
				panic(err)
			}

		case llmFormat:
			err := llmgen.Produce(ctx, writer)
			if err != nil {
				panic(err)
			}

		default:
			panic("should not happen - argument format should have been validated - got unrecognized format")
		}
		// Make sure buffered io on output file is flushed to disc
		err = w.Flush()
		if err != nil {
			fmt.Println(fatalPrinter("Could not flush generated otput: %s", err.Error()))
		}
	},

	Args: func(cmd *cobra.Command, args []string) error {
		red := color.New((color.FgRed)).SprintFunc()

		// validate flags/options here return nil if all is fine else an error
		if len(args) > 1 {
			return errors.New(red("at most one argument accepted"))
		}
		if len(args) != 1 {
			return errors.New(red("no yammm file given"))
		}
		if len(format) < 2 {
			return errors.New(red("no --format specified"))
		}
		switch format {
		case goFormat, cueFormat, jschemaFormat, jsonschemaFormat,
			mdFormat, markdownFormat, llmFormat:
			if genmodelFile != "" {
				return fmt.Errorf("--genmodel option only applies to the weaviate format")
			}
		case wv8Format, weaviateFormat:
		default:
			return fmt.Errorf("the format '%s' is not recognized, use --help to see valid formats", format)
		}
		return nil
	},
}

func init() {
	yammmCmd.AddCommand(yammmConvertCmd)

	flags := yammmConvertCmd.PersistentFlags()
	flags.StringVarP(&outFile, "out", "o", "", "the file to write output to")
	flags.StringVarP(&format, "format", "f", "", "the output file format")
	flags.StringVarP(&genmodelFile, "genmodel", "g", "", "filename of Weaviate genmodel")

	// TODO: --nowerror, precense of warnings does not cause exit != 0
}
