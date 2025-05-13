// Package cmd contains the example hello CLI logic.
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/data/replication"
	"github.com/wyrth-io/whit/internal/validation"
	"github.com/wyrth-io/whit/internal/yammm"
	"github.com/wyrth-io/whit/parser"
)

var wv8ExperimentCmd = &cobra.Command{
	Use:   "wv8 experiment --schema <schemafile.yammm> --cluster <weaviate cluster address>",
	Short: "Runs a weaviate experiment.",
	Long: `Runs a weaviate experiment/example. This is a temporary command used for testing some
implementation detail.
`,
	Run: func(cmd *cobra.Command, args []string) {
		red := color.New((color.FgRed)).SprintFunc()
		blue := color.New((color.FgBlue)).SprintFunc()

		ctx := mustHaveYammmContext()

		// Create the wv8 client
		cfg := weaviate.Config{
			Host:   wv8StoreURL, // Replace with your endpoint
			Scheme: "https",
		}

		client, err := weaviate.NewClient(cfg)
		exitOnError(err, "%s", err)

		switch wv8Experiment {
		case 1:
			fmt.Println(blue("Experiment: Create a Car with v5 deterministic UUID"))

			// Insert a car with a UUID v5 id for registration number ABC123
			carType := ctx.LookupType("Car")
			v5UUID, err := carType.InstanceID(map[string]any{"regNbr": "ABC123"}, map[string]uuid.UUID{})
			if err != nil {
				panic(err)
			}

			propMap := map[string]interface{}{
				"regNbr": "ABC123",
			}

			created, err := client.Data().Creator().
				WithClassName("Car").
				WithID(v5UUID.String()).
				WithProperties(propMap).
				WithConsistencyLevel(replication.ConsistencyLevel.ALL). // default QUORUM
				Do(context.Background())
			exitOnError(err, "Could not create Car ABC123 instance: %s", err)
			bytes, err := json.MarshalIndent(created, "", "  ")
			exitOnError(err, "Could not marshal response to JSON: %s", err)

			fmt.Printf("Got object back after creation: %s\n", string(bytes))
		case 2:
			fmt.Println(blue("Experiment: Create a Car XYZ987 with v5 deterministic UUID"))

			// Insert a car with a UUID v5 id for registration number ABC123
			carType := ctx.LookupType("Car")
			carUUID, err := carType.InstanceID(map[string]any{"regNbr": "XYZ987"}, map[string]uuid.UUID{})
			exitOnError(err, "Coult not produce instance id for Car XYZ987: %s", err)

			propMap := map[string]interface{}{
				"regNbr": "XYZ987",
			}

			created, err := client.Data().Creator().
				WithClassName("Car").
				WithID(carUUID.String()).
				WithProperties(propMap).
				WithConsistencyLevel(replication.ConsistencyLevel.ALL). // default QUORUM
				Do(context.Background())
			exitOnError(err, "Could not create Car instance: %s", err)
			bytes, err := json.MarshalIndent(created, "", "  ")
			exitOnError(err, "Could not marshal response to JSON: %s", err)

			fmt.Printf("Got object back after creation: %s\n", string(bytes))

			// Create a part
			propMap = map[string]interface{}{
				"kind": "Engine",
			}
			autoPartType := ctx.LookupType("AutoPart")
			partUUID, err := autoPartType.InstanceID(map[string]any{"kind": "Engine"}, map[string]uuid.UUID{})
			exitOnError(err, "Coult not produce instance id for AutoPart Engine: %s", err)
			created, err = client.Data().Creator().
				WithClassName("AutoPart").
				WithID(partUUID.String()).
				WithProperties(propMap).
				WithConsistencyLevel(replication.ConsistencyLevel.ALL). // default QUORUM
				Do(context.Background())
			exitOnError(err, "Could not create AutoPart instance: %s", err)
			bytes, err = json.MarshalIndent(created, "", "  ")
			exitOnError(err, "Could not marshal response to JSON: %s", err)
			fmt.Printf("Got object back after creation: %s\n", string(bytes))

			// Link engine to Car
			payload := client.Data().ReferencePayloadBuilder().
				WithClassName("AutoPart").
				WithID(partUUID.String()).
				Payload()

			err = client.Data().ReferenceCreator().
				WithClassName("Car").
				WithID(carUUID.String()).
				WithReferenceProperty("hAS_AutoParts").
				WithReference(payload).
				WithConsistencyLevel(replication.ConsistencyLevel.ALL).
				Do(context.Background())
			exitOnError(err, "Could not create Link between car and auto part : %s", err)
			bytes, err = json.MarshalIndent(created, "", "  ")
			exitOnError(err, "Could not marshal response to JSON: %s", err)
			fmt.Printf("Got object back after creation %d bytes: %s\n", len(bytes), string(bytes))

		case 3:
			fmt.Println(blue("Experiment: Link from Car XYZ987 to Engine"))

			// Insert a car with a UUID v5 id for registration number ABC123
			carType := ctx.LookupType("Car")
			carUUID, err := carType.InstanceID(map[string]any{"regNbr": "XYZ987"}, map[string]uuid.UUID{})
			exitOnError(err, "Coult not produce instance id for Car XYZ987: %s", err)

			autoPartType := ctx.LookupType("AutoPart")
			partUUID, err := autoPartType.InstanceID(map[string]any{"kind": "Engine"}, map[string]uuid.UUID{})
			exitOnError(err, "Coult not produce instance id for AutoPart Engine: %s", err)

			// Link engine to Car
			payload := client.Data().ReferencePayloadBuilder().
				WithClassName("AutoPart").
				WithID(partUUID.String()).
				Payload()

			err = client.Data().ReferenceCreator().
				WithClassName("Car").
				WithID(carUUID.String()).
				WithReferenceProperty("hAS_AutoParts").
				WithReference(payload).
				WithConsistencyLevel(replication.ConsistencyLevel.ALL).
				Do(context.Background())
			exitOnError(err, "Could not create Link between car and auto part : %s", err)

		default:
			fmt.Println(red("unknown experiment %d", wv8Experiment))
		}
	},

	Args: func(cmd *cobra.Command, args []string) error {
		// validate flags/options here return nil if all is fine else an error
		if wv8StoreURL == "" {
			return fmt.Errorf("no --cluster specified")
		}
		if schemaFile == "" {
			return fmt.Errorf("no schema file given")
		}
		if !strings.HasSuffix(schemaFile, ".yammm") {
			return fmt.Errorf("schema file must be a .yammm file")
		}
		return nil
	},
}

// For example: https://sandbox1-gfbp9s3p.weaviate.network
var wv8Experiment int

func init() {
	wv8Cmd.AddCommand(wv8ExperimentCmd)

	flags := wv8ExperimentCmd.PersistentFlags()
	//	flags.StringVarP(&file, "file", "f", "", "instance file to read in JSON, - for stdin")
	flags.StringVarP(&schemaFile, "schema", "s", "", "Yammm model in Yammm")
	flags.StringVarP(&wv8StoreURL, "cluster", "c", "", "Weaviate cluster URL")
	flags.IntVarP(&wv8Experiment, "xperiment", "x", 1, "Run experiment 1 to n")
}

func mustHaveYammmContext() yammm.Context {
	var ctx yammm.Context
	var ic validation.IssueCollector
	ctx, ic = parser.ParseFile(schemaFile)
	if ctx == nil {
		switch {
		case ic.HasFatal():
			log.Fatalf("fatal error(s) occurred - no output produced")
		case ic.HasErrors():
			log.Fatalf("error(s) occurred - no output produced")
		default:
			log.Fatalf("internal error: parser did not produce a context (reason unknown)")
		}
	}
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
	return ctx
}
