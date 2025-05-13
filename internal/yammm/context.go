package yammm

import (
	"bytes"
	"encoding/json"
	"io"

	"cuelang.org/go/pkg/strings"
	"github.com/gertd/go-pluralize"
	"github.com/pkg/errors"
	"github.com/wyrth-io/whit/internal/tc"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/validation"
)

// Context holds state for a model. It has a cue context for cuelang operations required for data types.
// It holds a cue scope to enable data type referencing, which is required to compute the base data type
// of DataType when generating go code from a meta model.
//
// A Context contains a reference to a main model, either by loading one from JSON, or by programatically
// building one step by step.
type Context interface {
	// Complete validates and indexes everything for operations that require the context to be completed.
	// Will return all found issues in the given IssueCollector. If there were any errors or fatal issues
	// reported false will be returned, otherwise true.
	Complete(ic validation.IssueCollector) (ok bool)

	// IsCompleted returns true if model is completed. In case there were errors during
	// the completion false will be returned.
	IsCompleted() bool

	// SetMainModel sets the main model which is operated on by other context operations.
	// Will return error if main model is already set, or if model does not have a name.
	SetMainModel(model *Model) error

	// LookupType returns a Type pointer for the given name, or nil if this type is not
	// defined in the model.
	LookupType(name string) *Type

	// LookupPLuralType returns a Type pointer for the given name in plural, or nil if this type is not
	// defined in the model.
	LookupPluralType(name string) *Type

	// Model returns the Model in this context or nil if it has not been set.
	Model() *Model

	// Name returns the name of the model. Returns empty string if model is not defined.
	Name() string

	// LookupDataType returns a DataType pointer for the given name, or nil if this data type is not
	// defined in the model.
	LookupDataType(name string) *DataType

	// SetModelFromJSON unmarshals a yammm model in Json form and sets it as the main model.
	SetModelFromJSON(reader io.Reader) error

	// WriteModelAsJSON marshals the main yammm model in Json format.
	WriteModelAsJSON(writer io.Writer) error

	// AddType adds a new empty type. The type will have both Name and plural names set.
	// Use other AddXxx methods to detail the model further.
	AddType(name string, properties []*Property) (*Type, error)

	// AddInherits adds an inheritance from inheritedType to the given type. The inheritedType
	// does not have to be defined yet, but must be when calling Complete. Multiple inheritance is
	// allowed but all inherited traits (properties and relationships) must be unique.
	AddInherits(typeName, inheritedTypeName string) error

	// AddProperty adds a property to the named type.
	AddProperty(typeName, property string, constraint []string, optional, primary bool) (*Property, error)

	// AddAbstractType adds an abstract type.
	AddAbstractType(name string, properties []*Property) (*Type, error)

	// AddCompositionPartType adds a type that is part of a composition.
	AddCompositionPartType(name string, properties []*Property) (*Type, error)

	// AddAssociation adds an association relationship between two types. Setting `optional`
	// to true means that an instance of the from type is not required to have this association.
	// Setting `many` to true means that the from type can have multiple references to the to type,
	// otherwise max one.
	AddAssociation(fromTypeName, name, toTypeName string, optional, many bool, properties []*Property, doc string) (*Association, error)

	// AddComposition adds a Part (to type) to a composition (from type). Setting `optional`
	// to true means that an instance of the composition type is not required to have this part.
	// Setting `many` to true means that the composition type can have multiple references to the Part type,
	// otherwise max one.
	AddComposition(fromType, name, toType string, optional, many bool, doc string) (*Composition, error)

	// AddDataType adds a new named data type that can be used in value constraints
	// in properties and other data types. The constraint string slice must be a valid yammm TypeChecker
	// sequence. The name must start with a letter a-z. The name of the data type may not be one of
	// the preexisting data types "string", "int", "number", "bool".
	AddDataType(name string, constraint []string) (*DataType, error)

	// AddInvariant adds an invariant constraint to the type.
	AddInvariant(typeName string, invariant *Invariant) error

	// EachType calls the given function for each type in the model.
	EachType(f func(t *Type))
}

// thePluralizer is initiaized when a context is created for the first time.
// It holds the state of the pluralizer as it is not meaningful to have several of those throughout
// the code.
var thePluralizer *pluralize.Client

// Pluralize returns the plural of the given string.
func Pluralize(s string) string {
	if thePluralizer == nil {
		thePluralizer = pluralize.NewClient()
	}
	return thePluralizer.Plural(s)
}

// context is the implementation of the Context interface.
type context struct {
	idxOfTypes       map[string]*Type
	idxOfPluralTypes map[string]*Type
	referencedTypes  []string
	main             *Model
	// If the context has been completed or not
	completed bool
}

// NewContext creates a new context in which a meta model is defined.
func NewContext() Context {
	ctx := context{
		idxOfTypes: make(map[string]*Type),
	}
	return &ctx
}

// Complete completes the contained main model by indexing and cross referencing everything.
func (ctx *context) Complete(ic validation.IssueCollector) bool {
	// Error if there is no model
	if ctx.main == nil {
		ic.Collectf(validation.Error, "no model set in context - cannot proceed")
		return false
	}
	if ic.HasFatal() {
		return false
	}

	// Build an index of all type names to *Type
	for i, t := range ctx.main.Types {
		if _, ok := ctx.idxOfTypes[t.Name]; ok {
			ic.Collectf(validation.Error, "%stype '%s' is defined multiple times", t.Label(), t.Name)
			return false
		}
		ctx.idxOfTypes[t.Name] = ctx.main.Types[i]
	}
	// Validate the model
	if !ctx.main.validate(ctx, ic) {
		return false
	}
	for _, t := range ctx.main.Types {
		referencesOk := true
		for _, x := range t.Associations {
			if _, ok := ctx.idxOfTypes[x.To]; !ok {
				ic.Collectf(validation.Error, "%stype '%s' referenced in association '%s' does not exist", t.Label(), x.To, x.Name)
				referencesOk = false
			}
		}
		for _, x := range t.Compositions {
			if _, ok := ctx.idxOfTypes[x.To]; !ok {
				ic.Collectf(validation.Error, "%stype '%s' referenced in composition '%s' does not exist", t.Label(), x.To, x.Name)
				referencesOk = false
			}
		}
		if !referencesOk {
			return false
		}
	}
	for i := range ctx.main.Types {
		if !ctx.main.Types[i].complete(ctx, ic) {
			return false
		}
	}
	// Collect the names of Types for which there needs to be a Cue REF_TO_TYPE (i.e. associations
	// To).
	ctx.referencedTypes = utils.Reduce(ctx.main.Types, utils.NewSet[string](),
		func(t *Type, memo *utils.Set[string]) *utils.Set[string] {
			for _, a := range t.Associations {
				memo.Add(a.To)
			}
			return memo
		}).Slices()

	ctx.completed = true
	return true
}

func (ctx *context) LookupType(name string) *Type {
	t, ok := ctx.idxOfTypes[name]
	if !ok {
		return nil
	}
	return t
}

func (ctx *context) LookupPluralType(name string) *Type {
	if !ctx.IsCompleted() {
		panic("LookupPluralType() can only be invoked when context have been completed")
	}
	if ctx.idxOfPluralTypes == nil {
		ctx.idxOfPluralTypes = make(map[string]*Type, len(ctx.idxOfTypes))
		for _, t := range ctx.idxOfTypes {
			ctx.idxOfPluralTypes[t.PluralName] = t
		}
	}
	t, ok := ctx.idxOfPluralTypes[name]
	if !ok {
		return nil
	}
	return t
}

func (ctx *context) SetMainModel(m *Model) error {
	if ctx.main != nil {
		return errors.Errorf("context already has a main model")
	}
	if len(m.Name) < 1 {
		return errors.Errorf("a model must have a name")
	}
	ctx.main = m
	return nil
}

func (ctx *context) IsCompleted() bool { return ctx.completed }

func (ctx *context) SetModelFromJSON(reader io.Reader) (err error) {
	var m Model
	var data []byte
	if data, err = io.ReadAll(reader); err == nil {
		if err = json.Unmarshal(data, &m); err == nil {
			err = ctx.SetMainModel(&m)
		}
	}
	return err
}

func (ctx *context) WriteModelAsJSON(writer io.Writer) (err error) {
	data, err := json.Marshal(ctx.main)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = json.Indent(buf, data, "", "    "); err != nil {
		return err
	}
	_, err = io.WriteString(writer, buf.String())
	if err != nil {
		return err
	}
	return nil
}
func (ctx *context) addType(name string, abstract, part bool, properties []*Property) (*Type, error) {
	if err := ctx.errorIfCompleted(); err != nil {
		return nil, err
	}
	amendedProperties := append(properties, &Property{ //nolint:gocritic
		Located:      Located{Line: 1, Column: 1, Source: "<yammm built in 'id UUID primary'>"},
		Name:         "id",
		DataType:     []string{tc.UUIDS},
		IsPrimaryKey: true,
		Optional:     false,
	})
	t := &Type{Name: name, PluralName: Pluralize(name), Properties: amendedProperties, IsAbstract: abstract, IsPart: part}
	ctx.main.Types = append(ctx.main.Types, t)
	return ctx.main.Types[len(ctx.main.Types)-1], nil
}

func (ctx *context) AddType(name string, properties []*Property) (*Type, error) {
	return ctx.addType(name, false, false, properties)
}

func (ctx *context) AddAbstractType(name string, properties []*Property) (*Type, error) {
	return ctx.addType(name, true, false, properties)
}
func (ctx *context) AddCompositionPartType(name string, properties []*Property) (*Type, error) {
	return ctx.addType(name, false, true, properties)
}
func (ctx *context) errorIfCompleted() error {
	if ctx.completed {
		return errors.New("cannot add type - context already completed")
	}
	if ctx.main == nil {
		return errors.New("cannot add type - context has no main model")
	}
	return nil
}

// AddProperty adds a property to the named type.
func (ctx *context) AddProperty(typeName string, property string, constraint []string, optional bool, primary bool) (*Property, error) {
	if err := ctx.errorIfCompleted(); err != nil {
		return nil, err
	}
	for i := range ctx.main.Types {
		if ctx.main.Types[i].Name != typeName {
			continue
		}
		t := ctx.main.Types[i]
		p := &Property{Name: property, DataType: constraint, Optional: optional, IsPrimaryKey: primary}
		t.Properties = append(t.Properties, p)
		return p, nil
	}
	return nil, errors.Errorf("cannot add property '%s' - type '%s' not found", property, typeName)
}
func (ctx *context) AddInvariant(typeName string, invariant *Invariant) error {
	if err := ctx.errorIfCompleted(); err != nil {
		return err
	}
	for i := range ctx.main.Types {
		if ctx.main.Types[i].Name != typeName {
			continue
		}
		t := ctx.main.Types[i]
		t.Invariants = append(t.Invariants, invariant)
		return nil
	}
	return errors.Errorf("cannot add invariant '%s' - type '%s' not found", invariant.Name, typeName)
}

// AddAssociation adds an association relationship between two types. Setting `optional`
// to true means that an instance of the from type is not required to have this association.
// Setting `many` to true means that the from type can have multiple references to the to type,
// otherwise max one.
func (ctx *context) AddAssociation(
	fromTypeName, name, toTypeName string,
	optional, many bool,
	properties []*Property,
	doc string,
) (*Association, error) {
	if err := ctx.errorIfCompleted(); err != nil {
		return nil, err
	}
	fromType := *utils.Find(ctx.main.Types, func(t *Type) bool { return t.Name == fromTypeName })
	if fromType == nil {
		return nil, errors.Errorf("cannot add association - from type '%s' not found", fromTypeName)
	}
	assoc := &Association{
		Relationship: Relationship{
			Name:          name,
			To:            toTypeName,
			Optional:      optional,
			Many:          many,
			Documentation: doc,
		},
		Properties: properties,
	}
	fromType.Associations = append(fromType.Associations, assoc)
	return assoc, nil
}

// AddComposition adds a Part (to type) to a composition (from type). Setting `optional`
// to true means that an instance of the composition type is not required to have this part.
// Setting `many` to true means that the composition type can have multiple references to the Part type,
// otherwise max one.
func (ctx *context) AddComposition(fromTypeName, name, toTypeName string, optional, many bool, doc string) (*Composition, error) {
	if err := ctx.errorIfCompleted(); err != nil {
		return nil, err
	}
	// Default relationship name is the toType in singlar/plural and upper case.
	if name == "" {
		if many {
			name = strings.ToUpper(Pluralize(toTypeName))
		} else {
			name = strings.ToUpper(toTypeName)
		}
	}
	fromType := *utils.Find(ctx.main.Types, func(t *Type) bool { return t.Name == fromTypeName })
	if fromType == nil {
		return nil, errors.Errorf("cannot add composition - composition type '%s' not found", fromTypeName)
	}
	comp := &Composition{
		Relationship: Relationship{
			Name:          name,
			To:            toTypeName,
			Optional:      optional,
			Many:          many,
			Documentation: doc,
		},
	}
	fromType.Compositions = append(fromType.Compositions, comp)
	return comp, nil
}

// AddDataType adds a new named data type that can be used in value consraints expressed in cue
// in properties and other data types. The cueConstraint string must be a valida Cuelang value constraint
// exression. The name must start with a letter a-z. The name of the data type may not be one of
// the preexisting data types "string", "int", "number", "bool". When a named data type definition is
// used in another constraint it should be preceded with a #.
func (ctx *context) AddDataType(name string, constraint []string) (*DataType, error) {
	if err := ctx.errorIfCompleted(); err != nil {
		return nil, err
	}
	dt := &DataType{
		Name:       name,
		Constraint: constraint,
	}
	ctx.main.DataTypes = append(ctx.main.DataTypes, dt)
	return dt, nil
}

// AddInherits adds an inheritance from inheritedType to the given type. The inheritedType
// does not have to be defined yet, but must be when calling Complete.
func (ctx *context) AddInherits(typeName string, inheritedTypeName string) error {
	if err := ctx.errorIfCompleted(); err != nil {
		return err
	}

	fromType := *utils.Find(ctx.main.Types, func(t *Type) bool { return t.Name == typeName })
	if fromType == nil {
		return errors.Errorf("cannot add inheritance - type '%s' not found", typeName)
	}
	fromType.Inherits = append(fromType.Inherits, inheritedTypeName)
	return nil
}

// LookupDataType returns a DataType pointer for the given name, or nil if this data type is not
// defined in the model.
func (ctx *context) LookupDataType(name string) *DataType {
	if ctx.main == nil {
		return nil
	}
	return *utils.Find(ctx.main.DataTypes, func(dt *DataType) bool {
		return dt.Name == name
	})
}

func (ctx *context) Name() string {
	if ctx.main == nil {
		return ""
	}
	return ctx.main.Name
}

func (ctx *context) Model() *Model {
	return ctx.main
}

func (ctx *context) EachType(f func(t *Type)) {
	for i := range ctx.main.Types {
		f(ctx.main.Types[i])
	}
}
