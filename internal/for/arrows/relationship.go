// Package arrows contains utilities for operating on neo4j Arrows graphs.
package arrows

import (
	"fmt"
	"strings"

	"github.com/wyrth-io/whit/internal/validation"
)

// Relationship describes a relationship between two nodes in arrows.
type Relationship struct {
	ID              string                 `json:"id"`
	Type            string                 `json:"type"`
	Style           map[string]interface{} `json:"style"`      // string key and mix of string and int value
	Properties      map[string]string      `json:"properties"` // TODO: don't know what they look like yet
	FromID          string                 `json:"fromId"`
	ToID            string                 `json:"toId"`
	isInherits      bool
	isComposition   bool // normalized so that from is the Composition and to is the Part
	isAssociation   bool // normal relationship
	reversedEdge    bool // for composition expressed as COMPOSES_INTO in diagram and is reversed to COMPOSED_OF in normalization
	toMin           Multiplicity
	toMax           Multiplicity
	fromMin         Multiplicity
	fromMax         Multiplicity
	needsDefinition bool        // if it must have a separate named output entry
	nProperties     []*Property // normalized properties

}

// Multiplicity is an Enum for valid multiplicities (0, 1, Many).
type Multiplicity int

const (
	// Zero meaning a relation is optional.
	Zero Multiplicity = iota

	// One meaning that one relation is required.
	One

	// Many meaning 1 or more.
	Many = -1
)

// String returns "0", "1" or "M".
func (m Multiplicity) String() (s string) {
	switch m {
	case Zero:
		s = "0"
	case One:
		s = "1"
	case Many:
		s = "M"
	default:
		s = "???"
	}
	return
}

// normalize sets flags and multiplicity from the relationship Type and #to, #from propertoes. For meta relations (#) the # is stripped
// and IsAssociation() returns false. The #to and #from properties are deleted.
func (r *Relationship) normalize(ic validation.IssueCollector) {
	if len(r.Type) < 1 {
		ic.Collectf(validation.Fatal, "Relationship with id '%s' does not have a relationship type", r.ID)
	}
	if r.Type[0:1] == "#" {
		r.Type = r.Type[1:]
	} else {
		r.isAssociation = true
	}
	switch r.Type {
	case "INHERITS", "IS_A":
		if r.isAssociation {
			ic.Collectf(validation.Warning, "relationship id '%s' is named '%s' but does not start with # and does not specify inheritance", r.ID, r.Type) //nolint:lll
		} else {
			r.isInherits = true
		}
	case "COMPOSES_INTO":
		if r.isAssociation {
			ic.Collectf(validation.Warning, "relationship id '%s' is named '%s' but does not start with # and does not specify composition", r.ID, r.Type) //nolint:lll
		} else {
			r.isComposition = true
			r.reversedEdge = true
			r.ToID, r.FromID = r.FromID, r.ToID
		}
	case "COMPOSED_OF":
		if r.isAssociation {
			ic.Collectf(validation.Warning, "relationship id '%s' is named '%s' but does not start with # and does not specify inheritance", r.ID, r.Type) //nolint:lll
		} else {
			r.isComposition = true
		}
	default:
		if !r.isAssociation {
			ic.Collectf(validation.Error, "relationship id '%s' - unknown special relationship '#%s'", r.ID, r.Type)
		}
	}

	// Multiplicity defaults - one optional at other end, both ways
	r.fromMin = Zero
	r.fromMax = One
	r.toMin = Zero
	r.toMax = One
	toKey := "#to"
	fromKey := "#from"
	if r.reversedEdge {
		toKey = "#from"
		fromKey = "#to"
	}
	if m, ok := r.Properties[toKey]; ok {
		m = strings.TrimSpace(m)
		switch m {
		case "01": // the default, do nothing
		case "1", "11":
			r.toMin = One
			r.toMax = One
		case "M", "0M", "m", "0m":
			r.toMin = Zero
			r.toMax = Many
		case "1M", "1m":
			r.toMin = One
			r.toMax = Many
		default:
			ic.Collectf(validation.Error, "illegal multiplicity '%s' in property %s of relationship id '%s'", m, toKey, r.ID)
		}
	}
	if m, ok := r.Properties[fromKey]; ok {
		m = strings.TrimSpace(m)
		switch m {
		case "01": // the default, do nothing
		case "1", "11":
			r.fromMin = One
			r.fromMax = One
		case "M", "0M", "m", "0m":
			r.fromMin = Zero
			r.fromMax = Many
		case "1M", "1m":
			r.fromMin = One
			r.fromMax = Many
		default:
			ic.Collectf(validation.Error, "illegal multiplicity '%s' in property %s of relationship id '%s'", m, fromKey, r.ID)
		}
		// Delete the meta properties to not have to deal with them later.
	}
	delete(r.Properties, "#to")
	delete(r.Properties, "#from")
	if r.isComposition && len(r.Properties) > 0 {
		ic.Collectf(validation.Error, "relationship properties other than #to/#from not allowed in COMPOSED type relationship - id '%s' from node id '%s'", //nolint:lll
			r.ID, r.FromID)
	}
	if r.isAssociation && len(r.Properties) > 0 && r.toMax == Many {
		// helper flag
		r.needsDefinition = true
	}
	r.normalizeProperties(ic, true)
}

func (r *Relationship) normalizeProperties(ic validation.IssueCollector, forMeta bool) {
	if !r.isAssociation {
		return
	}
	np := make([]*Property, 0, len(r.Properties))
	for key, typeString := range r.Properties {
		nname, primary, optional := normalizePropertyName(key)
		if nname == "" {
			ic.Collectf(validation.Error, "association '%s' property '%s' - illegal property name\n", r.Type, key)
		}
		if forMeta && primary && optional {
			ic.Collectf(validation.Error, "association '%s' property '%s' cannot be both primary (+) and optional (?)\n", r.Type, key)
		}
		if !forMeta && (primary || optional) {
			ic.Collectf(validation.Error, "association '%s' property '%s' has illegal ending of '+ or '?'\n", r.Type, key)
		}
		np = append(np, &Property{
			Name:     nname,
			Type:     typeString,
			primary:  primary,
			optional: optional,
		})
		r.nProperties = np
	}
}

// IsComposition returns true if this is a Composition.
func (r *Relationship) IsComposition() bool {
	return r.isComposition
}

// IsInheritance returns true if this is an Inheritance.
func (r *Relationship) IsInheritance() bool {
	return r.isInherits
}

// IsAssociation returns true if this is an Association.
func (r *Relationship) IsAssociation() bool {
	return r.isAssociation
}

// String returns the relationship in a human readable form, showing to,from, relationship type and the
// multiplicity.
func (r *Relationship) String() string {
	var special string
	if !r.isAssociation {
		special = "#"
	}
	if r.reversedEdge {
		return fmt.Sprintf("(%s)[%s..%s]--%s%s-->[%s..%s](%s)", r.ToID, r.toMin, r.toMax, special, r.Type, r.fromMin, r.fromMax, r.FromID)
	}
	return fmt.Sprintf("(%s)[%s..%s]--%s%s-->[%s..%s](%s)", r.FromID, r.fromMin, r.fromMax, special, r.Type, r.toMin, r.toMax, r.ToID)
}

// ToMultiplicity returns 01, 11, 0M or 1M.
func (r *Relationship) ToMultiplicity() string {
	return r.toMin.String() + r.toMax.String()
}

// FromMultiplicity returns 01, 11, 0M or 1M.
func (r *Relationship) FromMultiplicity() string {
	return r.fromMin.String() + r.fromMax.String()
}
