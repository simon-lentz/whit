package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tada/catch"
	"github.com/wyrth-io/whit/internal/utils/yammel"
)

var yamlValidateCmd = &cobra.Command{
	Use:   "validate -s <schema.json> -f <doc.yaml>",
	Short: "Validates a yaml document against a JSON schema.",
	Long: `This validates a yaml document againt the given jsonschema"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		err := catch.Do(func() {
			yammel.ValidateYamlFile(yamlFile, jsonSchema)
		})
		if err != nil {
			fmt.Println(err)
		}
	},

	Args: func(cmd *cobra.Command, args []string) error {
		// validate flags/options here return nil if all is fine else an error
		if len(args) > 0 {
			return errors.New("no arguments accepted")
		}
		return nil
	},
}
var jsonSchema string
var yamlFile string

func init() {
	yamlCmd.AddCommand(yamlValidateCmd)

	flags := yamlValidateCmd.PersistentFlags()
	flags.StringVarP(&yamlFile, "file", "f", "-", "file to validate or - for stdin")
	flags.StringVarP(&jsonSchema, "schema", "s", "", "schema to use for validation")
}
