// Package yammel contains YAML related utilities. (Pardon the cute name).
package yammel

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"

	"github.com/fatih/color"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/tada/catch"
	"github.com/wyrth-io/whit/internal/pio"
	"gopkg.in/yaml.v3"
)

// ValidateYamlFile validates a Json file as per the given schema (or no schema if schema is "").
// This will panic with a catch.Error, and the caller should then exit with 1 status.
func ValidateYamlFile(file string, schemaFile string) {
	doc := yamlDocumentLoader(file)
	schema := jsonSchemaLoader(schemaFile)
	validate(doc, schema)
}

func yamlDocumentLoader(file string) (doc interface{}) {
	var data []byte
	var err error
	if file == "-" {
		data, err = io.ReadAll(os.Stdin)
	} else {
		data, err = os.ReadFile(filepath.Clean(file))
	}
	if err != nil {
		panic(catch.Error(err))
	}

	var m interface{}
	err = yaml.Unmarshal(data, &m)
	if err != nil {
		panic(catch.Error(err))
	}
	m, err = toStringKeys(m)
	if err != nil {
		panic(catch.Error(err))
	}
	return m
}

func jsonSchemaLoader(schemaFile string) *jsonschema.Schema {
	schemaFile = filepath.Clean(schemaFile)
	// use this as the name of the schema in the registry
	// will be strange if user gave strange schema file name, but probably does not matter
	baseName := filepath.Base(schemaFile)

	compiler := jsonschema.NewCompiler()
	f, err := os.Open(schemaFile)
	if err != nil {
		panic(catch.Error(err))
	}
	if err := compiler.AddResource(baseName, f); err != nil {
		panic(err)
	}
	defer pio.Close(f)

	schema, err := compiler.Compile(baseName)
	if err != nil {
		panic(catch.Error(err))
	}
	return schema
}

func validate(doc interface{}, schema *jsonschema.Schema) {
	err := schema.Validate(doc)
	if err != nil {
		message := fmt.Sprintf("%#v", err)
		panic(catch.Error(colorizeValidationError(message)))
	}
	fmt.Println("validation successful")
}

// colorizeValidationError colorizes the error message from jsonschema.
// The jsonschema package standard errors are on the form "I[ref] S[ref] error text"
// Where I stands for instance document and S for schema - the refs are JSON pointers to the
// position in the document.
// This colorization makes refs blue and error messages red
// Since jsonschema handles json documents and the yaml content has already be parsed into
// a go structure it is not possible to produce line/column references in error messages.
// To do that, a YAML parser needs to be added to jsonschema package.
func colorizeValidationError(message string) string {
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()

	r := regexp.MustCompile(`\[I(.*)\] \[S(.*)\] (.*)`)
	return r.ReplaceAllString(message, "[I"+blue("$1)"+"] [S"+blue("$2")+"] "+red("$3")))
}
func toStringKeys(val interface{}) (interface{}, error) {
	var err error
	switch val := val.(type) {
	case map[interface{}]interface{}:
		m := make(map[string]interface{})
		for k, v := range val {
			k, ok := k.(string)
			if !ok {
				return nil, errors.New("found non-string key")
			}
			m[k], err = toStringKeys(v)
			if err != nil {
				return nil, err
			}
		}
		return m, nil
	case []interface{}:
		var l = make([]interface{}, len(val))
		for i, v := range val {
			l[i], err = toStringKeys(v)
			if err != nil {
				return nil, err
			}
		}
		return l, nil
	default:
		return val, nil
	}
}
