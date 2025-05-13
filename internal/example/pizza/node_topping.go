package pizza

import "gopkg.in/yaml.v3"

// Topping is a kind of topping.
type Topping struct {
	Name string `yaml:"name"`
	// Not allowed to have any yet, but this is allowed in the schema
	Links Links `yaml:"links"`
	line  int
	col   int
}

// SetLine implments SettableLineAndCol.
func (t *Topping) SetLine(line int) { t.line = line }

// SetCol implments SettableLineAndCol.
func (t *Topping) SetCol(col int) { t.col = col }

// UnmarshalYAML provides custom unmarshalling - see genericUnmarsal function.
func (t *Topping) UnmarshalYAML(value *yaml.Node) error {
	type rawTopping Topping
	rt := (*rawTopping)(t)
	return genericUnmarsal(rt, t, value)
}

// Validate validates a Topping.
func (t *Topping) Validate(v *Validator) {
	// Topping must have a name
	if t.Name == "" {
		v.Issue("a topping must have a name", t.line, t.col)
	}
	// A Topping does not have any outgoing relationships (at the moment, they are allowed in the schema though)
	if len(t.Links.LikedPizza) > 0 {
		location := t.Links.LikedPizza[0]
		v.Issue("a topping.links cannot contain likedPizza links", location.line, location.col)
	}
	if len(t.Links.PizzaTopping) > 0 {
		location := t.Links.LikedPizza[0]
		v.Issue("a topping.links cannot contain pizzaTopping links", location.line, location.col)
	}
}
