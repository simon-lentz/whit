// Package cmd contains the example hello CLI logic.
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate/entities/models"
	"github.com/wyrth-io/whit/internal/for/wv8"
	"github.com/wyrth-io/whit/internal/yammm"
)

var wv8SchemaInitCmd = &cobra.Command{
	Use:   "wv8 schema init --schema <schema.yammm> --cluster <weaviate cluster address> [--genmodel genmodelfile.json]",
	Short: "Defines a Weaviate schema in a Weaviate endpoint.",
	Long: `Defines a Weaviate schema in a Weaviate endpoint. An optional genmodel for Weaviate can be given
as a  separate file with --genmodel/-g. See the command "whit wv8 schema genmodel" for more information.
`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := mustHaveYammmContext()

		// Create the wv8 client
		cfg := weaviate.Config{
			Host:   wv8StoreURL, // Replace with your endpoint
			Scheme: "https",
		}
		genmodel := yammm.NewGenmodel()
		if genmodelFile != "" {
			reader, err := os.Open(genmodelFile)
			exitOnError(err, "Could not read genmodel file %s: %s", genmodelFile, err)
			data, err := io.ReadAll(reader)
			exitOnError(err, "Could not read genmodel file %s: %s", genmodelFile, err)
			err = json.Unmarshal(data, &genmodel)
			exitOnError(err, "Could not unmarshal JSON in genmodel file %s: %s", genmodelFile, err)
		}
		client, err := weaviate.NewClient(cfg)
		if err != nil {
			panic(err)
		}
		exitOnError(err, "Could not create Weaviate client: %s", err)

		// Create schema
		schema, err := wv8.GenerateSchema(ctx, genmodel)
		exitOnError(err, "Could not generate a Weaviate schema from yammm: %s", err)

		for _, cl := range schema.Classes {
			emptyClass := &models.Class{Class: cl.Class}
			err = client.Schema().ClassCreator().WithClass(emptyClass).Do(context.Background())
			exitOnError(err, "Could not create Weaviate class: %s: %s", cl.Class, err)
		}
		for _, cl := range schema.Classes {
			for _, p := range cl.Properties {
				err = client.Schema().PropertyCreator().WithClassName(cl.Class).WithProperty(p).Do(context.Background())
				exitOnError(err, "Could not create property: %s for Weaviate class %s: %s", p.Name, cl.Class, err.Error())
			}
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

// Cluster example: sandbox1-gfbp9s3p.weaviate.network  (sigh for lint).

func init() {
	wv8SchemaCmd.AddCommand(wv8SchemaInitCmd)

	flags := wv8SchemaInitCmd.PersistentFlags()
	//	flags.StringVarP(&file, "file", "f", "", "instance file to read in JSON, - for stdin")
	flags.StringVarP(&schemaFile, "schema", "s", "", "Yammm model in Yammm")
	flags.StringVarP(&wv8StoreURL, "cluster", "c", "", "Weaviate cluster URL")
	flags.StringVarP(&genmodelFile, "genmodel", "g", "", "filename of Weaviate genmodel")
}
