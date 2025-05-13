package cmd

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wyrth-io/whit/internal/for/csvgen"
	"github.com/wyrth-io/whit/internal/pio"
	"github.com/wyrth-io/whit/internal/utils"
)

var csvDebugCmd = &cobra.Command{
	Use:   "debug --out <filename> ??? <csvfile>",
	Short: "Performs debugging operation on csv file..",
	Long: `This command is for debugging csv related problems.
`,
	Run: func(cmd *cobra.Command, args []string) {
		var w *bufio.Writer
		var writer *pio.Writer

		fatalPrinter := color.New(color.FgRed, color.Bold).SprintFunc()

		// Prepare the output
		var fo *os.File
		var err error
		if outFile == "-" || outFile == "" {
			fo = os.Stdout
		} else {
			fo, err = os.Create(outFile)
			if err != nil {
				fmt.Println(fmt.Sprintf(fatalPrinter("Cannot write to '%s': %s"), outFile, err.Error()))
				os.Exit(1)
			}
		}
		defer pio.Close(fo) // close and panic on error
		w = bufio.NewWriter(fo)
		writer = pio.WriterOn(w)

		// Open and read the file line by line
		f, err := os.Open(args[0])
		if err != nil {
			fmt.Println(fmt.Sprintf(fatalPrinter("Could not read '%s': %s"), args[0], err.Error()))
			os.Exit(1)
		}
		reader := csv.NewReader(f)
		// basePath := filepath.Base(args[0])
		// derivedName := strings.TrimSuffix(basePath, filepath.Ext(basePath))
		// derivedTypeName := utils.ToUpperCamel(derivedName)

		// Read and process the header
		firstRow, err := reader.Read()
		if err != nil {
			fmt.Println(fmt.Sprintf(fatalPrinter("Could not read header line of '%s': %s"), args[0], err.Error()))
			os.Exit(1)
		}
		colMap := make(map[string]int, len(firstRow))
		for i := range firstRow {
			colMap[firstRow[i]] = i
		}
		camels := make([]string, len(firstRow))
		for i := range firstRow {
			camels[i] = utils.ToLowerCamel(firstRow[i])
		}
		colNbr := colMap[columnName]
		writer.FormatLn("Column: '%s', is number: %d, mapped to: '%s'", columnName, colNbr, camels[colNbr])
		// TODO: hack here - show this column
		lineNbr := 1
		for {
			row, err := reader.Read()
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error reading csv line %d: %s\n", lineNbr, err.Error())
				os.Exit(1)
			}
			val := row[colNbr]
			typeStr := ""
			if csvgen.IsBoolean(val) || strings.HasSuffix(firstRow[colNbr], "?") {
				typeStr = "Boolean"
			} else if csvgen.IsFloat(val) {
				typeStr = "Float"
			} else if csvgen.IsInteger(val) {
				typeStr = "Integer"
			} else {
				typeStr = "String"
			}
			writer.FormatLn("[%d] %s %s", lineNbr, val, typeStr)
			lineNbr++
		}
		err = w.Flush()
		if err != nil {
			fmt.Println(fatalPrinter("Could not flush generated otput: %s", err.Error()))
		}
	},

	Args: func(cmd *cobra.Command, args []string) error {
		red := color.New((color.FgRed)).SprintFunc()

		if len(args) > 1 {
			return errors.New(red("at most one csv file accepted"))
		}
		if len(args) != 1 {
			return errors.New(red("no csv file given"))
		}
		if len(columnName) < 1 {
			return errors.New(red("--col <column name> must be given"))
		}
		return nil
	},
}
var columnName string

func init() {
	csvCmd.AddCommand(csvDebugCmd)

	flags := csvDebugCmd.PersistentFlags()
	flags.StringVarP(&outFile, "out", "o", "", "the file to write output to")
	flags.StringVarP(&columnName, "col", "c", "", "the name of a column")
}
