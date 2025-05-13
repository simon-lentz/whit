package pizza

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// LikedPizza is a relationship from a Person to a Pizza.
type LikedPizza struct {
	Pizza string `yaml:"pizza"`
	line  int
	col   int
}

// SetLine implments SettableLineAndCol.
func (lp *LikedPizza) SetLine(line int) { lp.line = line }

// SetCol implments SettableLineAndCol.
func (lp *LikedPizza) SetCol(col int) { lp.col = col }

// Validate validates a LinkedPizza.
func (lp *LikedPizza) Validate(v *Validator) {
	// "pizza" must be a reference to a Pizza ID
	if !v.HasPizzaNamed(lp.Pizza) {
		v.Issue(fmt.Sprintf("the likedPizza '%s' is unknown", lp.Pizza), lp.line, lp.col)
	}
}

// UnmarshalYAML provides custom unmarshalling - see genericUnmarsal function.
func (lp *LikedPizza) UnmarshalYAML(value *yaml.Node) error {
	type rawLikedPizza LikedPizza
	rlp := (*rawLikedPizza)(lp)
	return genericUnmarsal(rlp, lp, value)
}
