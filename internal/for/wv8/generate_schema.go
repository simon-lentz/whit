package wv8

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/tada/catch"
	"github.com/weaviate/weaviate/entities/models"
	"github.com/wyrth-io/whit/internal/tc"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/yammm"
)

type generator struct {
	ctx    yammm.Context
	model  *yammm.Model
	schema *models.Schema
	gm     *yammm.Genmodel
}

// GenerateSchemaText generates a Weaviate schema written in JSON to the given writer.
func GenerateSchemaText(ctx yammm.Context, gm *yammm.Genmodel, out io.Writer) (err error) {
	var schema *models.Schema
	if schema, err = GenerateSchema(ctx, gm); err == nil {
		var bytes []byte
		if bytes, err = json.MarshalIndent(schema, "", "  "); err == nil {
			_, err = out.Write(bytes)
		}
	}
	return err
}

// GenerateSchema returns a Weaviate model Schema. This schema can be used in a call to weaviate
// to create the schema in a Weaviate store.
func GenerateSchema(ctx yammm.Context, gm *yammm.Genmodel) (*models.Schema, error) {
	if !ctx.IsCompleted() {
		return nil, fmt.Errorf("context not completed")
	}
	if gm.Generator != "wv8" {
		return nil, fmt.Errorf("given genmodel is not for 'wv8': got '%s'", gm.Generator)
	}
	generator := &generator{ctx: ctx, model: ctx.Model(), gm: gm}
	err := catch.Do(func() {
		generator.generate()
	})
	if err == nil {
		return generator.schema, nil
	}
	return nil, err
}

func (gen *generator) generate() {
	gen.schema = &models.Schema{
		Name:    gen.model.Name,
		Classes: make([]*models.Class, 0, len(gen.model.Types)),
	}
	// All types except abstract have concrete representation in Weaviate
	for _, t := range utils.Filter(gen.model.Types, func(t *yammm.Type) bool { return !t.IsAbstract }) {
		hasVectorProperty := false
		wv8props := utils.Map(
			// Filter out the "id" property since it is added automatically by wv8.
			// Also filter out any Spacevector properties (at most one) as it is
			// not a property in Wv8.
			utils.Filter(t.AllProperties(), func(p *yammm.Property) bool {
				isVec := p.BaseType().Kind() != tc.SpacevectorKind
				if isVec {
					hasVectorProperty = true
				}
				return p.Name != "id" && !isVec
			}),
			func(p *yammm.Property) *models.Property {
				w8Prop := &models.Property{
					Name:        p.Name,
					Description: p.Documentation,
					DataType:    makeWv8DataType(gen.ctx, p),
				}
				label := fmt.Sprintf("%s.%s", t.Name, p.Name)
				err := SetOptionalPropertyProperties(w8Prop, gen.gm, label)
				if err != nil {
					panic(catch.Error("error while setting optional weaviate property properties", err))
				}

				return w8Prop
			})

		for _, assoc := range t.AllAssociations() {
			toName := assoc.To
			toType := gen.ctx.LookupType(assoc.To)
			allowedTargets := append(toType.AllSubTypes(), toName)

			if len(assoc.Properties) > 0 {
				// An EDGE class is required EDGE_From_RELNAME_TO
				toName = EdgeName(t, assoc)
				assocProperties := utils.Map(assoc.Properties, func(p *yammm.Property) *models.Property {
					w8Prop := &models.Property{
						Name:        p.Name,
						Description: p.Documentation,
						DataType:    makeWv8DataType(gen.ctx, p),
					}
					label := fmt.Sprintf("%s.%s", toName, p.Name)
					err := SetOptionalPropertyProperties(w8Prop, gen.gm, label)
					if err != nil {
						panic(catch.Error("error while setting optional weaviate property properties", err))
					}
					return w8Prop
				})
				// Add allowedTargets reference property to the edge properties
				assocProperties = append(assocProperties,
					&models.Property{
						Name:        assoc.Name,
						Description: assoc.Documentation,
						DataType:    allowedTargets,
					})
				edgeClass := &models.Class{
					Class:       toName,
					Description: assoc.Documentation,
					Properties:  assocProperties,
				}
				// TODO: What schema properties should apply to this edge class?
				// The edge comes from the Association, but so does the a property
				// in the "fromClass". Where are these properties
				gen.schema.Classes = append(gen.schema.Classes, edgeClass)
				// Add property for the reference to the EDGE (Edge in turn allows a reference
				// to any subtype of toType (inclusive).
				wv8props = append(wv8props,
					&models.Property{
						Name:        assoc.PropertyName(gen.ctx),
						Description: assoc.Documentation,
						DataType:    []string{toName}, // And all subtypes of toName TODO!!
					})
			} else {
				// Add a reference property for the association. The reference may be to any subtype
				// of toType (inclusive).
				// TODO: optional properties - this is for the property that is a reference.
				// Where are its optional properties described? If any???
				wv8props = append(wv8props,
					&models.Property{
						Name:        assoc.PropertyName(gen.ctx),
						Description: assoc.Documentation,
						DataType:    allowedTargets,
					})
			}
		}
		// Compositions
		for _, comp := range t.AllCompositions() {
			// Add a reference property for the composition
			toType := gen.ctx.LookupType(comp.To)
			allowedTargets := append(toType.AllSubTypes(), comp.To)
			wv8props = append(wv8props,
				&models.Property{
					Name:        comp.PropertyName(gen.ctx),
					Description: comp.Documentation,
					DataType:    allowedTargets,
				})
		}

		wv8class := &models.Class{
			Description: t.Documentation,
			Class:       t.Name,
			Properties:  wv8props,
		}
		err := SetOptionalClassProperties(wv8class, gen.gm, t.Name)
		if err != nil {
			panic(catch.Error("error while setting optional weaviate class properties", err))
		}
		// If has vector property, turn off the vectorizer(even if it was set in options).
		if hasVectorProperty {
			wv8class.Vectorizer = "none"
		}
		gen.schema.Classes = append(gen.schema.Classes, wv8class)
	}
}

// makeWv8DataType transforms the yammm data type encoding of a data type to Weaviate.
// This will panic if a yammm data type is not handled (it should not).
func makeWv8DataType(ctx yammm.Context, p *yammm.Property) []string {
	switch p.DataType[0] {
	case "Alias":
		// could in theory fail, but model is validated so cannot happen.
		dt := ctx.LookupDataType(p.DataType[1])
		return makeWv8DataTypeFromString(dt.Constraint[0])
	default:
		return makeWv8DataTypeFromString(p.DataType[0])
	}
}
func makeWv8DataTypeFromString(s string) []string {
	switch s {
	case "String", "Enum", "Pattern":
		return []string{"text"}
	case "Integer":
		return []string{"int"}
	case "Float":
		return []string{"number"}
	case "Boolean":
		return []string{"boolean"}
	case "Timestamp", "Date":
		// Note that Timestamp and Date are both wv8 Date with full RDC3339 format.
		// Instances with different format for the Timestamp will need to be transformed.
		return []string{"date"}
	case "UUID":
		return []string{"uuid"}

	default:
		panic(catch.Error("Internal Error: makeWv8DataType got type %s and had no transformation for it", s))
	}
}

// GetOptionalPropertyProperties returns any set optional Weaviate properties for a Property.
func GetOptionalPropertyProperties(wv8data any) (optionals *PropertyOptionals) {
	if topLevelMap, ok := wv8data.(map[string]any); ok {
		if bytes, err := json.Marshal(topLevelMap); err == nil {
			po := PropertyOptionals{}
			if err := json.Unmarshal(bytes, &po); err == nil {
				return &po
			}
		}
	}
	return nil
}

// PropertyOptionals describes the Weaviate optional properties settable for a property. It is used
// to read a genmodel.
type PropertyOptionals struct {
	IndexFiltrable  *bool  `json:"IndexFiltrable,omitempty"`
	IndexInverted   *bool  `json:"IndexInverted,omitempty"`
	IndexSearchable *bool  `json:"IndexSearchable,omitempty"`
	Tokenization    string `json:"Tokenization,omitempty"`
	ModuleConfig    any    `json:"ModuleConfig,omitempty"`
}

// SetOptionalPropertyProperties sets the Weaviate optional properties for a property from
// the data map (if present, is a map with a wv8 key).
func SetOptionalPropertyProperties(w8Prop *models.Property, gm *yammm.Genmodel, label string) error {
	// Check if there is a genmodel, if it is a definition do lookup
	data, ok := gm.Genmodels[label]
	if !ok {
		return nil
	} // had no genmodel
	if definition, ok := data.(string); ok {
		data, ok = gm.Definitions[definition]
		if !ok {
			return fmt.Errorf("genmodel for %s references non existing definition '%s'", label, definition)
		}
	}
	// Data is now supposed to be for a property. It is howeveer in map[string]any form
	po := GetOptionalPropertyProperties(data)
	if po != nil {
		w8Prop.IndexFilterable = po.IndexFiltrable
		w8Prop.IndexInverted = po.IndexInverted
		w8Prop.IndexSearchable = po.IndexSearchable
		if len(po.Tokenization) > 0 {
			w8Prop.Tokenization = po.Tokenization
		}
		if po.ModuleConfig != nil {
			w8Prop.ModuleConfig = po.ModuleConfig
		}
	}
	return nil
}

// ClassOptionals describes the Weaviate optional properties settable for a class. It is used
// to read a genmodel.
type ClassOptionals struct {
	InvertedIndexConfig *models.InvertedIndexConfig `json:"InvertedIndexConfig,omitempty"`
	ModuleConfig        any                         `json:"ModuleConfig,omitempty"`
	MultiTenancyConfig  *models.MultiTenancyConfig  `json:"MultiTenancyConfig,omitempty"`
	ReplicationConfig   *models.ReplicationConfig   `json:"ReplicationConfig,omitempty"`
	ShardingConfig      any                         `json:"ShardingConfig,omitempty"`
	VectorIndexConfig   any                         `json:"VectorIndexConfig,omitempty"`
	VectorIndexType     string                      `json:"VectorIndexType,omitempty"`
	Vectorizer          string                      `json:"Vectorizer,omitempty"`
}

// GetOptionalClassProperties returns ClassOptionals or nil from the given data.
func GetOptionalClassProperties(wv8Data any) (optionals *ClassOptionals) {
	if topLevelMap, ok := wv8Data.(map[string]any); ok {
		if bytes, err := json.Marshal(topLevelMap); err == nil {
			co := ClassOptionals{}
			if err := json.Unmarshal(bytes, &co); err == nil {
				return &co
			}
		}
	}
	return nil
}

// SetOptionalClassProperties sets the Weaviate optional properties for a class from
// the data map (if present, is a map with a wv8 key).
func SetOptionalClassProperties(w8Class *models.Class, gm *yammm.Genmodel, label string) error {
	// Check if there is a genmodel, if it is a definition do lookup
	data, ok := gm.Genmodels[label]
	if !ok {
		return nil
	} // had no genmodel
	if definition, ok := data.(string); ok {
		data, ok = gm.Definitions[definition]
		if !ok {
			return fmt.Errorf("genmodel for %s references non existing definition '%s'", label, definition)
		}
	}
	co := GetOptionalClassProperties(data)
	if co != nil {
		w8Class.InvertedIndexConfig = co.InvertedIndexConfig
		w8Class.ModuleConfig = co.ModuleConfig
		w8Class.MultiTenancyConfig = co.MultiTenancyConfig
		w8Class.ReplicationConfig = co.ReplicationConfig
		w8Class.ShardingConfig = co.ShardingConfig
		w8Class.VectorIndexConfig = co.VectorIndexConfig
		w8Class.VectorIndexType = co.VectorIndexType
		w8Class.Vectorizer = co.Vectorizer // TODO: Should be set to "none" if vector property is present
	}
	return nil
}
