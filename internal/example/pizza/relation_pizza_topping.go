package pizza

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// RelationshipPizzaTopping is a relationship from a Pizza to a Topping.
type RelationshipPizzaTopping struct {
	Topping string `yaml:"topping"`
	Extra   bool   `yaml:"extra"`
	line    int
	col     int
}

// SetLine implments SettableLineAndCol.
func (pt *RelationshipPizzaTopping) SetLine(line int) { pt.line = line }

// SetCol implments SettableLineAndCol.
func (pt *RelationshipPizzaTopping) SetCol(col int) { pt.col = col }

// Validate validates a PizzaTopping.
func (pt *RelationshipPizzaTopping) Validate(v *Validator) {
	// "topping" must be a reference to a Topping in the kb snippet
	if !v.HasToppingNamed(pt.Topping) {
		v.Issue(fmt.Sprintf("the pizzaTopping '%s' is unknown", pt.Topping), pt.line, pt.col)
	}
}

// UnmarshalYAML provides custom unmarshalling - see genericUnmarsal function.
func (pt *RelationshipPizzaTopping) UnmarshalYAML(value *yaml.Node) error {
	type rawPizzaTopping RelationshipPizzaTopping
	rpt := (*rawPizzaTopping)(pt)
	return genericUnmarsal(rpt, pt, value)
}
