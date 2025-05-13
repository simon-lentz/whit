// Package cmd contains the example hello CLI logic.
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wyrth-io/whit/internal/for/wv8"
	"github.com/wyrth-io/whit/internal/pio"
)

var wv8SchemaGenmodelCmd = &cobra.Command{
	Use:   "wv8 schema genmodel --schema <schema.yammm> --out [w8genmodel.json|-]",
	Short: "Produces a wv8 template genmodel for a schema.",
	Long: `Defines an empty/template Weaviate genmodel for the schema. By default the output is sent to stdout.
Use the --out option to write to a given filename (or explicitly to stdout if the filename is given as "-").
The generated weaviate genmodel has two (unused) definitions for classes and properties to serve as a starting
point/example of what can be defined per class or property. Some of the information in these defaults are
toy samples. Consult Weaviate and used Weaviate modul documentation for the detailed meaning of the
genmodel information. 
`,
	Run: func(cmd *cobra.Command, args []string) {
		fatalPrinter := color.New(color.FgRed, color.Bold).SprintFunc()

		ctx := mustHaveYammmContext()

		genmodel := wv8.ProduceEmptyGenModel(ctx)

		var f *os.File
		var err error
		if outFile == "-" || outFile == "" {
			f = os.Stdout
		} else {
			f, err = os.Create(outFile)
			if err != nil {
				fmt.Println(fatalPrinter("Cannot write to '%s: %s", outFile, err.Error()))
				os.Exit(1)
			}
			defer pio.Close(f) // close and panic on error
		}
		bytes, err := json.MarshalIndent(genmodel, "", "  ")
		if err != nil {
			fmt.Println(fatalPrinter("Cannot marshal genmodel to JSON: %s", err.Error()))
			os.Exit(1)
		}
		_, err = f.Write(bytes)
		if err != nil {
			fmt.Println(fatalPrinter("Error writing genmodel to file: %s", err.Error()))
			os.Exit(1)
		}
	},

	Args: func(cmd *cobra.Command, args []string) error {
		// validate flags/options here return nil if all is fine else an error
		if schemaFile == "" {
			return fmt.Errorf("no schema file given")
		}
		if !strings.HasSuffix(schemaFile, ".yammm") {
			return fmt.Errorf("schema file must be a .yammm file")
		}
		return nil
	},
}

func init() {
	wv8SchemaCmd.AddCommand(wv8SchemaGenmodelCmd)

	flags := wv8SchemaGenmodelCmd.PersistentFlags()
	flags.StringVarP(&schemaFile, "schema", "s", "", "Yammm model in Yammm")
	flags.StringVarP(&outFile, "out", "o", "", "the file to write output to")
}
