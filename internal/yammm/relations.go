package yammm

import (
	"fmt"

	"github.com/wyrth-io/whit/internal/validation"
)

// Relationship is a base type for the two concrete forms of relationship; Association and Composition.
type Relationship struct {
	Located

	// Name is the type/name of the relationship - in all caps
	Name string `json:"name"`

	// To is the name of the to Type of the relation in Name
	To string `json:"to"`

	// Optional indicates a 0? relation
	Optional bool `json:"optional"`

	// Many indicates a ?M relation
	Many bool `json:"many"`

	// Documentation contains markdown text describing the relation.
	Documentation string `json:"documentation,omitempty"`

	targetPkMap map[string]*Property // cache
}

// Association is a Relationship with properties and a Where clause.
type Association struct {
	Relationship
	// Properties of the relationship
	Properties []*Property `json:"properties,omitempty"`
	propMap    map[string]*Property
}

// Composition is a Relationship without properties and where clause (as the Part-of-composition data is embedded).
type Composition struct {
	Relationship
}

func (r Relationship) validateR(forType string, _ Context, ic validation.IssueCollector) (ok bool) {
	ok = true
	if len(r.To) < 1 {
		ok = false
		ic.Collectf(validation.Error, "%s To(Name) must start with UC letter - got '%s'", forType, r.To)
	} else {
		c := r.To[0:1]
		if c < "A" || c > "Z" {
			ok = false
			ic.Collectf(validation.Error, "%s To(Name) must start with UC letter - got '%s'", forType, r.To)
		}
	}
	return
}

// TargetMap returns the target Type and a map of its primary keys. The model must have
// been completed or this will panic.
func (r Relationship) TargetMap(ctx Context) (*Type, map[string]*Property) {
	targetType := ctx.LookupType(r.To)
	if r.targetPkMap == nil {
		pks := targetType.AllPrimaryKeys()
		pkMap := make(map[string]*Property, len(pks))
		for _, p := range pks {
			pkMap[p.Name] = p
		}
		r.targetPkMap = pkMap
	}
	return targetType, r.targetPkMap
}
func (c Composition) validate(ctx Context, ic validation.IssueCollector) (ok bool) {
	ok = true
	if !c.validateR("composition", ctx, ic) {
		ok = false
	}
	if len(c.Name) < 1 {
		ok = false
		ic.Collectf(validation.Error, "composition name is empty - must start with UC letter - got '%s'", c.Name)
	}
	return
}

func (a Association) validate(ctx Context, ic validation.IssueCollector) (ok bool) {
	ok = true
	if !a.validateR("association", ctx, ic) {
		ok = false
	}
	if len(a.Name) < 1 {
		ok = false
		ic.Collectf(validation.Error, "association name is empty - must start with UC letter - got '%s'", a.Name)
	} else {
		c := a.Name[0:1]
		if c < "A" || c > "Z" {
			ok = false
			ic.Collectf(validation.Error, "association name must start with UC letter - got '%s'", a.Name)
		} else if !IsUCName(a.Name) {
			ok = false
			ic.Collectf(validation.Error, "association name has illegal characters - got '%s'", a.Name)
		}
	}

	for _, p := range a.Properties {
		if !p.validate(ctx, ic) {
			ok = false
		}
		// Cannot call BaseType() here, not yet set up.
		if p.DataType != nil && p.DataType[0] == "Spacevector" {
			ic.Collectf(validation.Error, "association '%s' has property '%s' of Spacevector type", a.Name, p.Name)
			ok = false
		}
	}
	return
}

// PropMap returns a map of association property name to Property. The model
// must have been completed or this may panic.
func (a *Association) PropMap() map[string]*Property {
	if a.propMap == nil {
		a.propMap = make(map[string]*Property, len(a.Properties))
		for _, p := range a.Properties {
			a.propMap[p.Name] = p
		}
	}
	return a.propMap
}

// Key returns the unique key for a relationship on the form "name_toType".
func (r Relationship) Key() string {
	return fmt.Sprintf("%s_%s", r.Name, r.To)
}

// PropertyName returns the property name of a relationship on the form '<name>_<type>' where
// <name> is the name of the composition, and <type> is the singular or plural name depending
// on the compositions 'Many' flag.
func (r Relationship) PropertyName(ctx Context) string {
	if r.Many {
		toType := ctx.LookupType(r.To)
		return fmt.Sprintf("%s_%s", r.Name, toType.PluralName)
	}
	return fmt.Sprintf("%s_%s", r.Name, r.To)
}
