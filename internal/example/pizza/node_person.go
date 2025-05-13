package pizza

import (
	"gopkg.in/yaml.v3"
)

// Person is a person (who likes crtain types of pizza).
type Person struct {
	Name  string `yaml:"name"`
	Links Links  `yaml:"links"`
	line  int
	col   int
}

// SetLine implments SettableLineAndCol.
func (p *Person) SetLine(line int) { p.line = line }

// SetCol implments SettableLineAndCol.
func (p *Person) SetCol(col int) { p.col = col }

// UnmarshalYAML provides custom unmarshalling - see genericUnmarsal function.
func (p *Person) UnmarshalYAML(value *yaml.Node) error {
	type rawPerson Person
	rp := (*rawPerson)(p)
	return genericUnmarsal(rp, p, value)
}

// Validate validates a Person.
func (p *Person) Validate(v *Validator) {
	// Person must have a name
	if p.Name == "" {
		v.Issue("a person must have a name", p.line, p.col)
	}
	if len(p.Links.PizzaTopping) > 0 {
		location := p.Links.PizzaTopping[0]
		v.Issue("a person.links cannot contain pizzaTopping links", location.line, location.col)
	}
	// Validate all valid types of links (there is only one kind possible)
	for _, lp := range p.Links.LikedPizza {
		lp.Validate(v)
	}
}
