// Package cmd contains whit's CLI logic.
package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

const (
	cueFormat        = "cue"
	yammmFormat      = "yammm"
	goFormat         = "go"
	jschemaFormat    = "jschema"
	jsonschemaFormat = "jsonschema"
	markdownFormat   = "markdown"
	mdFormat         = "md"
	wv8Format        = "wv8"
	weaviateFormat   = "weaviate"
	llmFormat        = "llm"
)

// Command flag and argument variables that are common.
var file string
var schemaFile string

func exitOnError(err error, format string, args ...interface{}) {
	fatalPrinter := color.New(color.FgRed, color.Bold).SprintFunc()
	if err != nil {
		fmt.Println(fatalPrinter(fmt.Sprintf(format, args...)))
		os.Exit(1)
	}
}
