package yammm

import (
	"encoding/json"
	"reflect"
	"sort"

	"github.com/wyrth-io/whit/internal/jzon"
	"github.com/wyrth-io/whit/internal/tc"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/validation"
	"github.com/wyrth-io/whit/internal/xray"
)

// Validator is an interface for validating an instance against a Yammm schema/model.
type Validator interface {
	// Validates the given model and calls the issue collector with found issues.
	// A value of false is returned if there are any issues at warning level or higher.
	// The issue collector can then be checked for details about reported problem severities.
	Validate(ic validation.IssueCollector) bool
}

type validator struct {
	meta       Meta
	graph      any
	sourceName string
	ic         validation.IssueCollector
}

// NewValidator returns a new instance to use for validation of an instance of a Yammmm schema/model.
// The given graph can be a go Struct if the instance is structured or a `map[string]any` if the instance
// is unstructured. The graph can also be an implementation of Wrapper.
// A context with the loaded and completed Yammm model is also given.
func NewValidator(ctx Context, sourceName string, graph any) Validator {
	return &validator{meta: NewMeta(ctx), graph: graph, sourceName: sourceName}
}

func (v *validator) Validate(ic validation.IssueCollector) (result bool) {
	v.ic = ic
	result = true
	ctx := v.meta.Context()
	if ctx == nil {
		ic.Collectf(validation.Fatal, "Cannot validate instance: no context defined")
		return false
	}
	if !v.meta.Context().IsCompleted() {
		model := v.meta.Context().Model()
		if model == nil {
			ic.Collectf(validation.Fatal, "Cannot validate instance: no model defined")
		} else {
			ic.Collectf(validation.Fatal, "%sCannot validate instance: model not completed", model.Label())
		}
		return false
	}
	// Top level must be an object
	topLevel := xray.NewWrapper(v.graph)
	if topLevel == nil || !topLevel.IsObject() {
		ic.Collectf(validation.Fatal, "%sGraph does not have an object at the top level", v.label(topLevel))
		return false
	}
	// All top level properties must be a plural type name present in the model.
	// The value of each property must be a slice.
	names := topLevel.FeatureNames()
	if len(names) == 0 {
		ic.Collectf(validation.Warning, "%sThe graph top level is empty", v.label(topLevel))
		return false
	}
	var t *Type
	for _, name := range names {
		t = v.meta.Context().LookupPluralType(name)
		if t == nil {
			ic.Collectf(validation.Error,
				"%sThe name '%s' is not the plural name of a type in the model",
				v.label(topLevel.FeatureName(name)),
				name)
			result = false
			continue
		}
		instances := topLevel.Feature(name)
		// The value must be a slice of instances.
		if instances == nil || !instances.IsSlice() {
			ic.Collectf(validation.Error,
				"%sThe value of '%s' is not a list of instances",
				v.label(topLevel.FeatureName(name)),
				name)
			result = false
			continue
		}
		// There must be instances of this type.
		if instances.Len() < 1 {
			ic.Collectf(validation.Warning,
				"The list of instances for '%s' is empty",
				v.label(instances),
				name)
			result = false
			continue
		}
		// Must be a concrete type.
		if t.IsAbstract || t.IsPart {
			ic.Collectf(validation.Error,
				"%sThe type '%s' is not a concrete type in the model",
				v.label(topLevel.FeatureName(name)),
				t.Name)
			result = false
			continue
		}
		// Validate all instances in the slice
		if !v.validateInstances(t, instances) {
			result = false
			continue
		}
	}
	if result {
		// Now that structure is ok, invariants can be validated.
		iv := NewInvariantValidator(ic)
		iv.Validate(ctx, v.graph)
		if ic.HasFatal() || ic.HasErrors() {
			result = false
		}
	}
	return result
}

// validateInstances checks that the slice per type has entries that are objects and continues
// by validating each instance.
func (v *validator) validateInstances(t *Type, instances xray.Wrapper) (result bool) {
	result = true
	// Validate each instance
	for i := 0; i < instances.Len(); i++ {
		// Must be an Object.
		obj := instances.FeatureAtIndex(i)
		if obj == nil || !obj.IsObject() {
			v.ic.Collectf(validation.Error,
				"%sEntry %d for type '%s' is not an object representation",
				v.label(obj),
				i, t.Name)
			result = false
			continue
		}
		// And the object must be a valid representation of that type
		if !v.validateInstance(t, obj) {
			result = false
			continue
		}
	}
	return
}

// validateInstance validates one instance against its type.
func (v *validator) validateInstance(t *Type, obj xray.Wrapper) (result bool) { //nolint:gocognit,gocyclo
	result = true
	features := utils.NewSet[string]() // collect all feature names
	if len(t.allPrimaryKeys) > 1 {
		// The id property can not be set manually when there are other primary keys.
		theID := obj.Value("id")
		if !(theID == nil || reflect.ValueOf(theID).IsZero()) {
			v.ic.Collectf(validation.Error,
				"%sProperty value of '%s.id' should not be set: not allowed when there are other primary keys",
				v.label(obj), t.Name)
			result = false
		}
	}
	for _, p := range t.allProperties {
		features.Add(p.Name)

		// Required properties must be set.
		// TODO: Special rule for the id property; it cannot be set if there are other primary
		// keys unless it is set to the deterministically UUID derived from the other primary keys.
		prop := obj.Value(p.Name)
		if p.Optional || p.HasDefault() {
			// Can be missing (or the zero value which is taken as missing)
			if prop == nil || reflect.ValueOf(prop).IsZero() {
				continue
			}
		} else {
			// Must be set (Zero value ignored since property is required the value must be compliant with data type)
			if prop == nil {
				v.ic.Collectf(validation.Error,
					"%sProperty value of '%s.%s' is required and is missing",
					v.label(obj), t.Name, p.Name)
				result = false
				continue
			}
		}
		// Set values must comply with the data type for the property
		if !v.validateDataType(obj.Feature(p.Name), t.Name, p, prop) {
			result = false
			continue
		}
	}
	for _, r := range t.allAssociations {
		// If association is Many, the Feature for the Property is a Slice otherwise an Object
		// The to type determines the content of the object(s) (all primary keys of that type)
		// If the association is Optional it may be missing.
		// The association may also have property values that must be compliant.
		assocPropName := r.PropertyName(v.meta.Context())
		features.Add(assocPropName)
		allLinks := []xray.Wrapper{} // list built up of documents to validate

		// Must have required Association (at least one that is specified, the target could be msising though)
		nested := obj.Feature(assocPropName)
		if nested == nil {
			if !r.Optional {
				result = false
				v.ic.Collectf(validation.Error,
					"%sAssociation '%s.%s' is required and is missing",
					v.label(obj),
					t.Name, assocPropName)
			}
			continue
		}
		// Must have Object with Where property or slice of objects with Where property.
		if r.Many {
			if !nested.IsSlice() {
				result = false
				v.ic.Collectf(validation.Error,
					"%sAssociation '%s.%s' must be a list since relation is to many",
					v.label(nested),
					t.Name, assocPropName)
				continue
			}
			for i := 0; i < nested.Len(); i++ {
				allLinks = append(allLinks, nested.FeatureAtIndex(i))
			}
		} else {
			if !nested.IsObject() {
				result = false
				v.ic.Collectf(validation.Error,
					"%sAssociation '%s.%s' must be an object since relation is to one",
					v.label(nested),
					t.Name, assocPropName)
				continue
			}
			allLinks = append(allLinks, nested)
		}
		// Each object must have all primary keys of the toType set in the Where object
		for _, link := range allLinks {
			where := link.Feature("Where")
			if where == nil {
				result = false
				v.ic.Collectf(validation.Error,
					"%sAssociation '%s.%s' is not an object with a Where property",
					v.label(link),
					t.Name, assocPropName)
				continue
			}
			if !where.IsObject() {
				result = false
				v.ic.Collectf(validation.Error,
					"%sAssociation '%s.%s' the Where property is not an object",
					v.label(where),
					t.Name, assocPropName)
				continue
			}
			// Validate the properties of the Where clause (the primary keys of the toType must be set).
			toType := v.meta.Context().LookupType(r.To)
			pks := toType.allPrimaryKeys
			whereExpected := utils.NewSetFrom[string](pks, func(p *Property) string { return p.Name })
			whereActual := utils.NewSet(where.FeatureNames()...)
			if where.HasCapitalizedFeatureNames() {
				// All features have initial UC - modify the feature names
				whereExpected = utils.NewSet(utils.Map(whereExpected.Slices(), utils.Capitalize)...)
			}

			// missing properties
			// If not having "id" set, it must have all primary keys.
			if !whereActual.Contains("id") {
				withoutID := whereExpected.Remove("id")
				if diff := withoutID.Diff(whereActual); diff.Size() != 0 {
					result = false
					v.ic.Collectf(validation.Error,
						"%sAssociation '%s.%s' is missing primary key 'id' or all of %v",
						v.label(where),
						t.Name, assocPropName, diff.Slices())
				}
			}
			// excess properties
			if diff := whereActual.Diff(whereExpected); diff.Size() != 0 {
				result = false
				v.ic.Collectf(validation.Error,
					"%sAssociation '%s.%s.Where' has excess primary key %v",
					v.label(where),
					t.Name, assocPropName, diff.Slices())
			}
			// Validate the data type for present and known properties.
			propsToValidate := whereActual.Intersection(whereExpected)
			for _, p := range pks {
				if !propsToValidate.Contains(p.Name) {
					continue
				}
				if !v.validateDataType(where.FeatureName(p.Name), t.Name+".Where", p, where.Value(p.Name)) {
					result = false
				}
			}

			// Properties of link must comply with properties of Association
			requiredProps := utils.Filter(r.Properties, func(p *Property) bool { return !p.Optional })
			nameOfProperty := func(p *Property) string { return p.Name }

			relPropsExpected := utils.NewSetFrom(r.Properties, nameOfProperty)
			relPropsRequired := utils.NewSetFrom(requiredProps, nameOfProperty)
			relPropsActual := utils.NewSet(link.FeatureNames()...)
			relPropsActual = relPropsActual.Diff(utils.NewSet("Where")) // Drop "Where" property
			if link.HasCapitalizedFeatureNames() {
				// All features have initial UC - modify the feature names
				relPropsExpected = utils.NewSet(utils.Map(relPropsExpected.Slices(), utils.Capitalize)...)
				relPropsRequired = utils.NewSet(utils.Map(relPropsRequired.Slices(), utils.Capitalize)...)
			}

			if diff := relPropsRequired.Diff(relPropsActual); diff.Size() != 0 {
				result = false
				v.ic.Collectf(validation.Error,
					"%sAssociation '%s.%s' is missing required properties %v",
					v.label(link),
					t.Name, assocPropName /*r.Key()*/, diff.Slices())
			}
			if diff := relPropsActual.Diff(relPropsExpected); diff.Size() != 0 {
				result = false
				v.ic.Collectf(validation.Error,
					"%sAssociation '%s.%s' has excess properties %v",
					v.label(link),
					t.Name, assocPropName /*r.Key()*/, diff.Slices())
			}
			// Validate the data type of the properties
			propsToValidate = relPropsActual.Intersection(relPropsExpected)
			for i := range r.Properties {
				if !propsToValidate.Contains(r.Properties[i].Name) {
					continue
				}
				if !v.validateDataType(link.FeatureName(r.Properties[i].Name),
					t.Name+"."+r.Key(), r.Properties[i], link.Value(r.Properties[i].Name)) {
					result = false
				}
			}
		}
	}
	for _, r := range t.allCompositions {
		// If composition is Many, the Feature for the Property is a Slice otherwise an Object
		// The to type determines the content of the object(s) (all primary keys of that type)
		// If the association is Optional it may be missing.
		// The association may also have property values that must be compliant.
		compPropName := r.PropertyName(v.meta.Context())
		features.Add(compPropName)

		// Must have required composition
		nested := obj.Feature(compPropName)
		if nested == nil {
			if !r.Optional {
				result = false
				v.ic.Collectf(validation.Error,
					"%sComposition '%s.%s' is required and is missing",
					v.label(obj),
					t.Name, compPropName)
			}
			continue
		}
		instancesToValidate := []xray.Wrapper{}
		if r.Many {
			if nested.Kind() != reflect.Slice {
				result = false
				v.ic.Collectf(validation.Error,
					"%sComposition '%s.%s' must be a list since relation is to 'many'",
					v.label(nested),
					t.Name, r.Key())
				continue
			}
			for i := 0; i < nested.Len(); i++ {
				instancesToValidate = append(instancesToValidate, nested.FeatureAtIndex(i))
			}
		} else {
			if nested.Kind() == reflect.Slice {
				result = false
				v.ic.Collectf(validation.Error,
					"%sComposition '%s.%s' cannot be a list since relation is to 'one'",
					v.label(nested),
					t.Name, r.Key())
				continue
			}
			instancesToValidate = append(instancesToValidate, nested)
		}

		// Validate each nested instance
		toType := v.meta.Context().LookupType(r.To)
		for _, d := range instancesToValidate {
			if !d.IsObject() {
				result = false
				v.ic.Collectf(validation.Error,
					"%sComposition '%s.%s' value must be an Object",
					v.label(d),
					t.Name, r.Key())
				continue
			}
			if !v.validateInstance(toType, d) {
				result = false
			}
		}
	}
	// Get all features of the instances as a set
	actualFeatures := utils.NewSet[string](obj.FeatureNames()...)
	if obj.HasCapitalizedFeatureNames() {
		// All features have initial UC - modify the feature names
		features = utils.NewSet(utils.Map(features.Slices(), utils.Capitalize)...)
	}
	diff := actualFeatures.Diff(features)
	if diff.Size() > 0 {
		result = false
		diffSlices := diff.Slices()
		sort.Strings(diffSlices)
		v.ic.Collectf(validation.Error,
			"%sObject of type '%s' has excess properties: %v",
			v.label(obj),
			t.Name, diffSlices,
		)
	}
	return result
}

// validateDataType checks the value of a property against the modeled data type. The
// `typeName` argument is only used as a string in error messages.
func (v *validator) validateDataType(wrapper xray.Wrapper, typeName string, p *Property, val any) (result bool) {
	result = true
	actualKind := tc.UnspecifiedKind
	expectedKind := p.BaseType().Kind()
	// If value is parsed from json using json package "numbers" feature,
	// and the value is a json.Number, then decide kind manually.
	if v, ok := val.(json.Number); ok {
		if n, err := v.Int64(); err == nil {
			actualKind = tc.IntKind
			val = n
		} else if n, err := v.Float64(); err == nil {
			actualKind = tc.FloatKind
			val = n
		}
	} else {
		valKind := reflect.ValueOf(val).Kind()
		switch valKind {
		case reflect.Int64, reflect.Int:
			actualKind = tc.IntKind
		case reflect.Float64, reflect.Float32:
			actualKind = tc.FloatKind
		case reflect.Slice:
			kv := reflect.ValueOf(val)
			// the elements are interface as array could hold mixed types of values in JSON
			// TODO: Need to check that they are all either float 32 or 64.
			// TODO: Transform value to []float32 as this is needed for checking
			transformed := []float32{}
			actualKind = tc.SpacevectorKind // until proven wrong
			for i := 0; i < kv.Len(); i++ {
				ki := kv.Index(i).Interface()
			again:
				switch n := ki.(type) {
				case float32:
					transformed = append(transformed, n)
					continue
				case float64:
					transformed = append(transformed, float32(n))
					continue
				case jzon.Node:
					ki = n.RawValue()
					goto again
				default:
					actualKind = tc.UnspecifiedKind
				}
			}
			val = transformed
		default:
			actualKind = tc.BaseKindOfReflectKind(valKind)
		}
	}
	if expectedKind != actualKind && !(expectedKind == tc.FloatKind && actualKind == tc.IntKind) {
		v.ic.Collectf(validation.Error,
			"%sProperty value of '%s.%s' must have %s base type: got %s",
			v.label(wrapper),
			typeName, p.Name, expectedKind.String(), actualKind.String())
		return false
	}
	if ok, msg := p.TypeChecker().Check(val); !ok {
		v.ic.Collectf(validation.Error, v.label(wrapper)+msg)
		return false
	}
	return result
}

func (v *validator) label(w xray.Wrapper) string {
	// Not all wrappers have Location
	if pos, ok := w.(xray.Position); ok {
		return xray.Label(pos, v.sourceName)
	}
	return ""
}
