package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/wyrth-io/whit/internal/for/csvgen"
	"github.com/wyrth-io/whit/internal/jzon"
	"github.com/wyrth-io/whit/internal/yammm"
	"gopkg.in/yaml.v3"
)

// readJSONInstance reads an instance graph in JSON. The JZON reader keeps track of
// the source line of the parsed elements.
func readJZONInstance(file string) any {
	red := color.New((color.FgRed)).SprintFunc()
	bytes, err := os.ReadFile(file) //nolint:gosec
	if err != nil {
		fmt.Println(red("Error while reading file %s: %s", file, err.Error()))
		os.Exit(1)
	}
	ctx := jzon.NewContext(file, string(bytes))
	node, err := ctx.UnmarshalNode()

	if err != nil {
		fmt.Println(red("Error while Unmarshaling JSON from file %s: %s", file, err.Error()))
		os.Exit(1)
	}
	return node
}

// readYamlInstance reads an instance graph in YAML.
// TODO: This does not handle integer values as the standard YAML parser returns float64
// for all numbers.
func readYamlInstance(file string) map[string]any {
	var graph map[string]any
	data, err := os.ReadFile(file) //nolint:gosec
	red := color.New((color.FgRed)).SprintFunc()
	if err != nil {
		fmt.Println(red("Error while reading file %s: %s", file, err))
		os.Exit(1)
	}
	// TODO: This produces float64 for all numbers!!
	if err = yaml.Unmarshal(data, &graph); err != nil {
		fmt.Println(red("Error while Unmarshaling YAML from file %s: %s", file, err))
		os.Exit(1)
	}
	return graph
}

// readCSVInstance reads a CSV file and returns it as an instance graph.
// Note the dependency on the cmd package global genmodelFile.
func readCSVInstance(file string, ctx yammm.Context, genmodelFile string) any {
	var gm csvgen.Genmodel
	reader, err := os.Open(genmodelFile) //nolint:gosec
	exitOnError(err, "Could not read genmodel file %s: %s", genmodelFile, err)
	data, err := io.ReadAll(reader)
	exitOnError(err, "Could not read genmodel file %s: %s", genmodelFile, err)
	err = json.Unmarshal(data, &gm)
	exitOnError(err, "Could not unmarshal JSON in genmodel file %s: %s", genmodelFile, err)

	f, err := os.Open(file) //nolint:gosec
	exitOnError(err, "Could not open CSVfile '%s': %s", genmodelFile, err)

	csvReader := csv.NewReader(f)

	graph, err := csvgen.NewGraph(ctx, &gm, csvReader)
	exitOnError(err, "Error while creating a graph adapter for CSV file '%s': %s", file, err)
	return graph
}
