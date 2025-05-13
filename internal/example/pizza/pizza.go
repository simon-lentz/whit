// Package pizza is an example of YAML parsing and validation based on the "pizza experiment" in repo "rdata".
package pizza

import (
	"io"
	"os"
	"path/filepath"

	"github.com/tada/catch"
	"gopkg.in/yaml.v3"
)

// Links is contained in Node type structs to describe relationships to other nodes.
// It is generic, not all types of relationships are applicable to all nodes.
// To extend with more relationship types, add them here.
type Links struct {
	PizzaTopping []RelationshipPizzaTopping `yaml:"pizzaTopping"`
	LikedPizza   []LikedPizza               `yaml:"likedPizza"`
}

// KbSnippet is a Knowledge Base Snippet in the pizza example.
// It is the top level object representing the data in a yaml file.
// To ectend with more node types, add them here.
type KbSnippet struct {
	Pizzas   []Pizza   `yaml:"pizzas"`
	Toppings []Topping `yaml:"toppings"`
	Persons  []Person  `yaml:"persons"`
}

// Parse parses bytes (typically read from a yaml file) and unmarshals the data into the given kbSnippet argument.
func (kb *KbSnippet) Parse(data []byte) error {
	return yaml.Unmarshal(data, kb)
}

// Read reads the knowledgebase snippet found in given file and Parses the read data into the receiver.
// If the file name is "-", stdin is read.
func (kb *KbSnippet) Read(file string) {
	var data []byte
	var err error
	switch file {
	case "-":
		{
			data, err = io.ReadAll(os.Stdin)
		}
	case "":
		{
			panic(catch.Error("no file given"))
		}
	default:
		{
			data, err = os.ReadFile(filepath.Clean(file))
		}
	}
	if err != nil {
		panic(catch.Error(err))
	}
	if err := kb.Parse(data); err != nil {
		panic(catch.Error(err))
	}
}

// SettableLineAndCol is an interface for setting line and column information.
// Used in custom unmarshalling to write into unmarshalled node.
type SettableLineAndCol interface {
	SetLine(int)
	SetCol(int)
}

// genericUnmarsal reduces the boilerplate code in every type for unmarshaling line and column.
// This function is called with a typecast to a derived type plus a reference to an SettableLineAndCol.
// The only boilerplate needed for each node is to create this variable. For example see the
// imlementation for Topping.
func genericUnmarsal[T any](t *T, located SettableLineAndCol, value *yaml.Node) error {
	err := value.Decode(t)
	if err != nil {
		return err
	}
	located.SetLine(value.Line)
	located.SetCol(value.Column)
	return nil
}
