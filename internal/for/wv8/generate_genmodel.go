package wv8

import (
	"fmt"

	"github.com/weaviate/weaviate/entities/models"
	"github.com/wyrth-io/whit/internal/yammm"
)

// ProduceEmptyGenModel produces a Genmodel for the model in a context (which must have been completed).
// This must be generator specific.
func ProduceEmptyGenModel(ctx yammm.Context) *yammm.Genmodel {
	gmListener := &genmListener{gm: yammm.NewGenmodel()}
	walker := yammm.NewModelWalker(ctx, gmListener)
	walker.Walk()
	return gmListener.gm
}

type genmListener struct {
	yammm.BaseModelListener
	gm *yammm.Genmodel
}

var emptyMap = map[string]any{}

func (gml *genmListener) EnterModel(_ yammm.Context) {
	gml.gm.Generator = "wv8"
}

// DefaultPropertyOptionals is used to define the defaults as output when generating a genmodel.
type DefaultPropertyOptionals struct {
	IndexFiltrable  bool   `json:"IndexFiltrable"`
	IndexInverted   bool   `json:"IndexInverted"`
	IndexSearchable bool   `json:"IndexSearchable"`
	Tokenization    string `json:"Tokenization"`
	ModuleConfig    any    `json:"ModuleConfig"`
}

// DefaultClassOptionals is used to define the defaults as output when generating a genmodel.
type DefaultClassOptionals struct {
	InvertedIndexConfig *models.InvertedIndexConfig `json:"InvertedIndexConfig"`
	ModuleConfig        any                         `json:"ModuleConfig"`
	MultiTenancyConfig  *models.MultiTenancyConfig  `json:"MultiTenancyConfig"`
	ReplicationConfig   *models.ReplicationConfig   `json:"ReplicationConfig"`
	ShardingConfig      any                         `json:"ShardingConfig"`
	VectorIndexConfig   any                         `json:"VectorIndexConfig"`
	VectorIndexType     string                      `json:"VectorIndexType"`
	Vectorizer          string                      `json:"Vectorizer"`
}

func (gml *genmListener) ExitModel(_ yammm.Context) {
	// Output the defaults for property and class as documentation/help.
	defProp := &DefaultPropertyOptionals{
		IndexFiltrable:  true,
		IndexSearchable: true,
		IndexInverted:   true,
		Tokenization:    "word",
		ModuleConfig:    nil,
	}
	gml.gm.Definitions["default_property"] = defProp
	defClass := &DefaultClassOptionals{
		InvertedIndexConfig: &models.InvertedIndexConfig{
			Bm25:                   &models.BM25Config{B: 1.0, K1: 1.0}, // actually not defaults
			CleanupIntervalSeconds: 1,
			IndexNullState:         true,
			IndexTimestamps:        true,
			IndexPropertyLength:    true,
			Stopwords: &models.StopwordConfig{
				Additions: []string{"newstopword"},
				Preset:    "preexisting-list-of-words",
				Removals:  []string{"words-to-remove"},
			},
		},
		ModuleConfig: nil,
		MultiTenancyConfig: &models.MultiTenancyConfig{
			Enabled: true,
		},
		ReplicationConfig: &models.ReplicationConfig{
			Factor: 3,
		},
		ShardingConfig:    nil,
		VectorIndexType:   "some-vector-index-type",
		VectorIndexConfig: nil,
		Vectorizer:        "some-vectorizer",
	}
	gml.gm.Definitions["default_class"] = defClass
}
func (gml *genmListener) EnterType(_ yammm.Context, t *yammm.Type) {
	gml.gm.Genmodels[t.Name] = emptyMap
}
func (gml *genmListener) OnProperty(_ yammm.Context, t *yammm.Type, p *yammm.Property) {
	if p.Name == "id" {
		return // id is built in and cannot be configured.
	}
	name := fmt.Sprintf("%s.%s", t.Name, p.Name)
	gml.gm.Genmodels[name] = emptyMap
}
func (gml *genmListener) OnAssociation(ctx yammm.Context, t *yammm.Type, a *yammm.Association) {
	// The property in from.type
	propName := a.PropertyName(ctx)
	name := fmt.Sprintf("%s.%s", t.Name, propName)
	gml.gm.Genmodels[name] = emptyMap

	// The Edge class if required.
	if len(a.Properties) > 0 {
		gml.gm.Genmodels[EdgeName(t, a)] = emptyMap
	}
}

func (gml *genmListener) OnAssociationProperty(_ yammm.Context, t *yammm.Type, a *yammm.Association, p *yammm.Property) {
	name := fmt.Sprintf("%s.%s", EdgeName(t, a), p.Name)
	gml.gm.Genmodels[name] = emptyMap
}

func (gml *genmListener) OnComposition(ctx yammm.Context, t *yammm.Type, c *yammm.Composition) {
	propName := c.PropertyName(ctx)
	name := fmt.Sprintf("%s.%s", t.Name, propName)
	gml.gm.Genmodels[name] = emptyMap
}
