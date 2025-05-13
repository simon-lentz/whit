package yammm

import (
	"reflect"
	"strings"
)

// Meta is the meta API for a generated go yammm model. The generator will generate a package specific
// impementation of this interface.
type Meta interface {
	// PropertyByName returns the value of a property given as its go field name. This function returns
	// nil if the property is not a property of the given value, or if the value is not part of the model
	// described by this meta.
	PropertyByName(value any, propName string) any

	// PropertiesMap returns a map of property name to value for all properties of the given value and type.
	// Properties that are optional and have "go zero value" are not included in the returned map.
	PropertiesMap(value any, yammmType *Type) map[string]any

	// PropertiesMapForMap returns a map of property name to value for all properties of the given value and type.
	// Properties that are optional and have "go zero value" are not included in the returned map.
	PropertiesMapForMap(value map[string]any, yammmType *Type) map[string]any

	// Propmap returns a map of property name to value for a given slice of properties.
	Propmap(value any, properties []*Property) map[string]any

	// PropmapForMap returns a map of property name to value for a given slice of properties.
	PropmapForMap(value map[string]any, properties []*Property) map[string]any

	// PrimaryKeysMap returns a map of all primary key properties with their values from the given value and type.
	PrimaryKeysMap(value any, yammmType *Type) map[string]any

	// PrimaryKeysMapForMap returns a map of all primary key properties with their values from the given value and type.
	PrimaryKeysMapForMap(value map[string]any, yammmType *Type) map[string]any

	// TypeModel returns the yammm Type of the given value. If the value is not an instance of a yammm modeled value in
	// this package, nil is returned.
	TypeModel(value any) *Type

	// TypeByName return the yammm Type of the given name. If the type is not found nil is returned.
	TypeByName(name string) *Type

	// TypeByPluralName return the yammm Type of the given name in plural. If the type is not found nil is returned.
	TypeByPluralName(name string) *Type

	// Context returns a yammm.Context for the package specific model.
	Context() Context
}

type meta struct {
	ctx Context
}

// NewMeta returns a Meta interface for reflection of an instance of a Yammm model.
func NewMeta(ctx Context) Meta {
	return &meta{ctx: ctx}
}

// Context returns the model Context.
func (m *meta) Context() Context {
	return m.ctx
}

// TypeModel return the yammm modeled Type of the given value, or nil if the given value is not a type in the model
// described by this Meta.
func (m *meta) TypeModel(v any) *Type {
	t := reflect.TypeOf(v)
	pkgAndName := t.String()
	parts := strings.Split(pkgAndName, ".")
	pkgName := parts[0]
	if len(pkgName) > 0 && pkgName[0:1] == "*" {
		pkgName = pkgName[1:]
	}
	typeName := parts[1]
	if pkgName != m.ctx.Name() {
		return nil
	}
	return m.ctx.LookupType(typeName)
}

// PropertyByName returns the value of a property given the Go property name in string form.
// Returns nil for uknown properties.
// TODO: Bad name of method since it is really for a Go Struct Field, not just yammm Properties.
func (m *meta) PropertyByName(v any, name string) any {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(name)
	if !f.IsValid() {
		return nil
	}
	if !f.CanInterface() {
		return nil
	}
	return f.Interface()
}

// Propmap returns the given properties as a map of property name to value.
func (m *meta) Propmap(value any, properties []*Property) map[string]any {
	result := make(map[string]any, len(properties))
	for _, p := range properties {
		propVal := m.PropertyByName(value, p.GoName())
		// Omit empty values if value is "go zero value"
		if !p.Optional || p.Optional && !reflect.ValueOf(propVal).IsZero() {
			result[p.Name] = propVal
		}
	}
	return result
}

// PropmapForMap returns the given properties as a map of property name to value.
func (m *meta) PropmapForMap(value map[string]any, properties []*Property) map[string]any {
	result := make(map[string]any, len(properties))
	for _, p := range properties {
		if propVal, ok := value[p.Name]; ok {
			// Omit empty values if value is "go zero value"
			if !p.Optional || p.Optional && !reflect.ValueOf(propVal).IsZero() {
				result[p.Name] = propVal
			}
		}
	}
	return result
}

// PrimaryKeysMap returns the primary keys for the given value and type as a map of name to value.
func (m *meta) PrimaryKeysMap(value any, yammmType *Type) map[string]any {
	return m.Propmap(value, yammmType.AllPrimaryKeys())
}

// PrimaryKeysMap returns the primary keys for the given value and type as a map of name to value.
func (m *meta) PrimaryKeysMapForMap(value map[string]any, yammmType *Type) map[string]any {
	return m.PropmapForMap(value, yammmType.AllPrimaryKeys())
}

// PropertiesMap returns all properties for the given value as a map of name to value.
func (m *meta) PropertiesMap(value any, yammmType *Type) map[string]any {
	return m.Propmap(value, yammmType.AllProperties())
}

// PropertiesMapForMap returns all properties for the given value as a map of name to value.
func (m *meta) PropertiesMapForMap(value map[string]any, yammmType *Type) map[string]any {
	return m.PropmapForMap(value, yammmType.AllProperties())
}

func (m *meta) TypeByName(name string) *Type {
	return m.ctx.LookupType(name)
}

func (m *meta) TypeByPluralName(name string) *Type {
	return m.ctx.LookupPluralType(name)
}

// TODO: Method for getting DataType in a similar way. Maybe change TypeOf to TypeModel, and consequently DataTypeModel.
// TODO: Method for Properties, Associations, and Compositions.
