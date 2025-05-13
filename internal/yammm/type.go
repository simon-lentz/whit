package yammm

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/tada/catch"
	"github.com/wyrth-io/whit/internal/tc"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/validation"
)

// Type represents entities other than relationships and data types.
type Type struct {
	Located
	Name            string         `json:"name"`
	PluralName      string         `json:"plural_name"`
	Properties      []*Property    `json:"properties,omitempty"`
	Invariants      []*Invariant   `json:"invariants,omitempty"`
	Inherits        []string       `json:"inherits,omitempty"`
	Compositions    []*Composition `json:"compositions,omitempty"`
	Associations    []*Association `json:"associations,omitempty"`
	IsPart          bool           `json:"is_part,omitempty"`
	IsAbstract      bool           `json:"is_abstract,omitempty"`
	Documentation   string         `json:"documentation,omitempty"`
	allPrimaryKeys  []*Property
	allProperties   []*Property
	allSuperTypes   *utils.Set[string]
	allSubTypes     *utils.Set[string]
	allAssociations []*Association
	allCompositions []*Composition
	completed       bool
	idPreamble      string
	idAdded         bool
	propMap         map[string]*Property // cache for lookup
}

func (t *Type) complete(ctx Context, ic validation.IssueCollector) (ok bool) { //nolint:gocognit
	ok = true
	if t.completed {
		return ok
	}
	// must set this first as this is what stops possible endless recursion
	t.completed = true

	// SUPER - computing all super types (transitively)
	t.allSuperTypes = utils.NewSet[string]()
	for _, st := range t.Inherits {
		stT := ctx.LookupType(st)
		if stT == nil {
			ok = false
			ic.Collectf(validation.Error, "%stype '%s' inherits from '%s' - super type not found", t.Label(), t.Name, st)
		} else {
			t.allSuperTypes.Add(stT.Name)
			stT.addSubtype(t.Name)
			if !stT.complete(ctx, ic) {
				ok = false
			}
			t.allSuperTypes = t.allSuperTypes.Union(stT.allSuperTypes)
		}
	}
	t.allSuperTypes.Each(func(s string) {
		ctx.LookupType(s).addSubtype(t.Name)
	})

	// PROPERTIES AND PRIMARY KEYS
	addPks := func(tt *Type) {
		for i, p := range tt.Properties {
			// "id" is special, only add it once
			if p.Name == "id" {
				if t.idAdded {
					continue
				}
				t.idAdded = true
			}
			if p.IsPrimaryKey {
				t.allPrimaryKeys = append(t.allPrimaryKeys, tt.Properties[i])
			}
			t.allProperties = append(t.allProperties, tt.Properties[i])
		}
	}
	// Properties and Primary keys from this type
	// The type itself is not allowed to have duplicate definitions (even if type compatible)
	if _, noDups := processDuplicates(ic, t.Properties, false); !noDups {
		ok = false
	}

	addPks(t)
	// Primary Keys from Inherited types
	t.allSuperTypes.Each(func(name string) { addPks(ctx.LookupType(name)) })

	// Check for overall duplicates and drop compatible inherited duplicates. Only check this if
	// not already failed since it would lead to same problem potentially reported multiple times.
	if ok {
		if uniqueProps, noDups := processDuplicates(ic, t.allProperties, true); !noDups {
			ok = false
		} else {
			t.allProperties = uniqueProps
			// Since all properties include all primary keys, there will be no additional
			// errors from the primary key set, but duplicates must be dropped.
			if uniqueProps, noDups := processDuplicates(ic, t.allPrimaryKeys, true); !noDups {
				t.allPrimaryKeys = uniqueProps
			}
		}
	}
	if t.IsPart && len(t.allPrimaryKeys) == 0 {
		ic.Collectf(validation.Error, "%spart type '%s' does not have any primary keys - required for parts", t.Label(), t.Name)
	}
	for _, p := range t.allProperties {
		if !p.complete(ctx, ic) {
			ok = false
		}
	}
	// id preamle is "package:typename:pk1:...:pkn:"
	t.idPreamble = fmt.Sprintf("%s:%s:%s", ctx.Model().Name, t.Name,
		strings.Join(
			utils.Filter(
				utils.Map(t.allPrimaryKeys,
					func(p *Property) string { return p.Name },
				),
				func(s string) bool { return s != "id" },
			), ":"),
	)

	// ASSOCIATIONS and COMPOSITIONS
	assocSet := utils.NewSet[string]()
	compSet := utils.NewSet[string]()
	for i := range t.Associations {
		if assocSet.Add(t.Associations[i].Key()) {
			t.allAssociations = append(t.allAssociations, t.Associations[i])
		}
		for ip := range t.Associations[i].Properties {
			if !t.Associations[i].Properties[ip].complete(ctx, ic) {
				ok = false
			}
		}
	}
	for i := range t.Compositions {
		if compSet.Add(t.Compositions[i].Key()) {
			t.allCompositions = append(t.allCompositions, t.Compositions[i])
		}
	}
	addRels := func(tt *Type) {
		for _, tta := range tt.allAssociations {
			if assocSet.Add(tta.Key()) {
				t.allAssociations = append(t.allAssociations, tta)
			}
		}
		for _, ttc := range tt.allCompositions {
			if compSet.Add(ttc.Key()) {
				t.allCompositions = append(t.allCompositions, ttc)
			}
		}
	}
	t.allSuperTypes.Each(func(name string) { addRels(ctx.LookupType(name)) })

	return ok
}
func processDuplicates(ic validation.IssueCollector, props []*Property, skipDups bool) (result []*Property, ok bool) {
	propSet := make(map[string]*Property, len(props))
	result = make([]*Property, 0, len(props))
	ok = true
	for pi := range props {
		p := props[pi]
		if p2, found := propSet[p.Name]; found {
			// already defined - ok if matching type, but do not add
			if skipDups && p.HasSameType(p2) {
				continue // ok, just skip it in result
			}
			ic.Collectf(validation.Error, "%sproperty '%s' definition clashes with definition at %s",
				p.Label(), p.Name, strings.TrimRight(p2.Label(), " "),
			)
			ok = false
		}
		propSet[p.Name] = p
		result = append(result, p)
	}
	return result, ok
}

func (t Type) validate(ctx Context, ic validation.IssueCollector) (ok bool) {
	ok = true
	// Names in Name and Plural must be initial upper case
	if len(t.PluralName) < 1 {
		ok = false
		ic.Collectf(validation.Error, "%stype '%s' Plural is empty, must start with UC letter - got '%s'", t.Label(), t.Name, t.PluralName)
	} else {
		c := t.PluralName[0:1]
		if c < "A" || c > "Z" {
			ok = false
			ic.Collectf(validation.Error, "%stype '%s' Plural must start with UC letter - got '%s'", t.Label(), t.Name, t.PluralName)
		} else if !IsUCName(t.PluralName) {
			ok = false
			ic.Collectf(validation.Error, "%stype '%s' Plural is an invalid identifier - got '%s'", t.Label(), t.Name, t.PluralName)
		}
	}

	if len(t.Name) < 1 {
		ok = false
		ic.Collectf(validation.Error, "%stype '%s' name must start with UC letter", t.Label(), t.Name)
	} else {
		c := t.Name[0:1]
		if c < "A" || c > "Z" {
			ok = false
			ic.Collectf(validation.Error, "%stype '%s' name must start with UC letter", t.Label(), t.Name)
		} else if !IsUCName(t.Name) {
			ok = false
			ic.Collectf(validation.Error, "%stype '%s' name is an invalid identifier", t.Label(), t.Name)
		}
	}

	// Rest of structure
	for _, p := range t.Properties {
		if !p.validate(ctx, ic) {
			ok = false
		}
	}
	for _, a := range t.Associations {
		if !a.validate(ctx, ic) {
			ok = false
		}
	}
	for _, c := range t.Compositions {
		if !c.validate(ctx, ic) {
			ok = false
		}
		// TO must be a part
		toType := ctx.LookupType(c.To)
		if toType != nil && (!toType.IsPart || toType.IsAbstract) { // missing to type checked elsewhere.
			ic.Collectf(validation.Error, "%scomposition '%s.%s' must reference a concrete part type", t.Label(), t.Name, c.PropertyName(ctx))
		}
	}
	return ok
}

// AllSubTypes returns the names of all types extending this type requrively.
func (t Type) AllSubTypes() []string {
	if t.allSubTypes == nil {
		return []string{}
	}
	return t.allSubTypes.Slices()
}

// AllSuperTypes returns the names of all types this type extends. The name of the type itself
// is not included.
func (t Type) AllSuperTypes() []string {
	if t.allSuperTypes == nil {
		return []string{}
	}
	return t.allSuperTypes.Slices()
}

// AllProperties returns a slice of pointers to each property (in the type itself, or inherited).
func (t Type) AllProperties() []*Property {
	return t.allProperties
}

// AllAssociations returns all associations from the type and all inherited.
func (t Type) AllAssociations() []*Association {
	return t.allAssociations
}

// AllCompositions returns all compositions from the type and all inherited.
func (t Type) AllCompositions() []*Composition {
	return t.allCompositions
}

// AllPrimaryKeys returns all primary keys for the type (direct and inherited).
func (t Type) AllPrimaryKeys() []*Property {
	return t.allPrimaryKeys
}

// InstanceID returns a UUID for an instance of this type. For a type with more primary keys
// than the always present 'id' key a deterministic v5 UUID will be generated.
// For types with only a single primary key ('id') a UUID will be generated if the 'id' is not set.
// This function also makes use of resolved `$$local` and `$$:uuid:local` references
// (as obtained from the `IDMapper.Map()` method and passed to this function.
// Note that the instance properties are given as a map of property name to value.
func (t Type) InstanceID(
	propMap map[string]any,
	replacements map[string]uuid.UUID,
) (uid uuid.UUID, err error) {
	pks := t.allPrimaryKeys
	if len(pks) <= 1 {
		idValue, ok := (propMap["id"]).(string)
		if ok && idValue != "" {
			// has a set id, replace it if needed
			replacement, hasReplacement := replacements[idValue]
			if hasReplacement {
				return replacement, nil
			}
			// no replacement, must be an UUID in string form, transform it to UUID.
			// (can possibly error if validator is faulty).
			return tc.DefaultUUIDChecker.GetUUID(idValue)
		}
		// has no set id or replacement, generate a random UUID
		return uuid.NewRandom()
	}
	// Multiple primary keys - generate a deterministic UUID.
	// The value to SHA1 is "package:type:propname:propname...:propvalue🔗...🔗...".
	parts := make([]string, 0, len(pks)-1)
	// Validation has ensured all Pks have values in the propmap.
	for _, p := range pks {
		if p.Name == "id" {
			continue
		} // is empty, and cannot be input to itself...
		parts = append(parts, fmt.Sprintf("%v", propMap[p.Name]))
	}
	payload := strings.Join(parts, "\U0001F517") // join with 🔗 char.
	return uuid.NewSHA1(WhitNamespace, []byte(t.idPreamble+payload)), nil
}

// IDPreamble returns the preamble to be combined with the joined property values to form a
// deterministic v5 UUID.
func (t Type) IDPreamble() string {
	return t.idPreamble
}

// addSubtype adds a subtype by name to this type.
func (t *Type) addSubtype(subName string) {
	if t.allSubTypes == nil {
		t.allSubTypes = utils.NewSet[string]()
	}
	t.allSubTypes.Add(subName)
}

// PropByName returns the *Property for the name. If the property does not exist
// nil is returned. This method can only be called on a completed model.
func (t *Type) PropByName(propName string) *Property {
	m := t.PropMap()
	if p, ok := m[propName]; ok {
		return p
	}
	return nil
}

// PropMap returns a map of property name to property for all properties.
// The model must have been completed or this will panic.
func (t *Type) PropMap() map[string]*Property {
	if !t.completed {
		panic(catch.Error("Model must be completed"))
	}
	if t.propMap == nil {
		t.propMap = make(map[string]*Property, len(t.allProperties))
		for i := range t.allProperties {
			p := t.allProperties[i]
			t.propMap[p.Name] = p
		}
	}
	return t.propMap
}
