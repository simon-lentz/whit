package pizza

import "gopkg.in/yaml.v3"

// Pizza is a kind of pizza.
type Pizza struct {
	ID    string `yaml:"$id"`
	Size  string `yaml:"size"`
	Type  string `yaml:"type"`
	Links Links  `yaml:"links"`
	line  int
	col   int
}

// SetLine implments SettableLineAndCol.
func (p *Pizza) SetLine(line int) { p.line = line }

// SetCol implments SettableLineAndCol.
func (p *Pizza) SetCol(col int) { p.col = col }

// Validate validates a Pizza.
func (p *Pizza) Validate(v *Validator) {
	// A pizza must have 'ID', 'size' and 'type'.
	if p.ID == "" {
		v.Issue("a Pizza must have an ID", p.line, p.col)
	}
	if p.Size == "" {
		v.Issue("a Pizza must have Size", p.line, p.col)
	}
	if p.Type == "" {
		v.Issue("a Pizza must have Type", p.line, p.col)
	}
	// Size must be S, M, or L
	switch p.Size {
	case "S", "M", "L": // do nothing
	default:
		{
			v.Issue("a pizza.size must be one of S, M, or L", p.line, p.col)
		}
	}
	switch p.Type {
	case "pan", "thin", "cheeze crust": // do nothing
	default:
		v.Issue("a pizza.type must be one of pan, thin, cheeze crust", p.line, p.col)
	}

	// A pizza can have pizza topping relationships only.
	if len(p.Links.LikedPizza) > 0 {
		v.Issue("a pizza.links cannot contain likedPizza links", p.line, p.col)
	}
	// Validate all valid types of links (there is only one kind possible)
	for _, pt := range p.Links.PizzaTopping {
		pt.Validate(v)
	}
}

// UnmarshalYAML provides custom unmarshalling - see genericUnmarsal function.
func (p *Pizza) UnmarshalYAML(value *yaml.Node) error {
	type rawPizza Pizza
	rp := (*rawPizza)(p)
	return genericUnmarsal(rp, p, value)
}
