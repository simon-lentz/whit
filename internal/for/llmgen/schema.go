package llmgen

import (
	"encoding/json"

	"github.com/wyrth-io/whit/internal/pio"
	"github.com/wyrth-io/whit/internal/yammm"
)

// Produce produces LLM prompt containing the (simplified) schema in json form.
func Produce(ctx yammm.Context, writer *pio.Writer) error {
	text := `The schema in JSON below describes the node types, links and properties of a Neo4j database.")
A schema has an array named 'types' describing the existing node types.
A type has properties described in the array named 'properties'.
A property has a 'data_type', which is 'int', 'string', 'float' or 'bool'.
A property can be a primary key to the node.
All types have a primary key named 'uid'. This property is a UUID string value.
The 'uid' property uniquely identifies an instance of a node type.
When a type has primary keys in addition to 'uid' the combination of the additional keys also uniquely identifies an instance of a type.
A type may have links to other types described in the array named 'links'.
All elements (schema, type, property and link) have a 'description' property that describes the element in English.
A link describes a link from a node type to another. A link may have properties.

For example, with a schema like this:
{ "types": [
	{
		"name": "Person,
	 	"properties": [
			{ "name": "uid", "data_type": "string", "primary": true},
			{ "name": "name", "data_type": "string", "primary": true}
		],
		"links": [
			{
				"name": "OWNS_Cars",
		  		"links_to": "Car"
			}
		]
	},
	{
		"name": "Car,
	 	"properties": [
			{ "name": "uid", "data_type": "string", "primary": true},
			{ "name": "regNbr", "data_type": "string", "primary": true, "description": "registration number"},
			{ "name": "color", "data_type": "string", "primary": false}
		],
		"links": [
			{
				"name": "OWNS_Cars",
		  		"links_to": "Car"
			}
		]
	}
]}
This can be queried using the Cypher query language as follows:
    MATCH (n:Person)-[r:OWNS_Cars]->[c:Car]
    WHERE c.color == "red"
	RETURN n.name, c.regNbr, c.color

Which would produce a list of People owning red cars showing the name of the person, the registration number
and color of the car.

When you are asked to produce a Cypher query this is the schema to use:
`

	writer.FormatLn(text)
	listener := &Listener{ctx: ctx, writer: writer}
	walker := yammm.NewModelWalker(ctx, listener)
	walker.Walk()
	data, err := json.MarshalIndent(listener.model, "", "  ")
	if err != nil {
		return err
	}
	writer.Println(string(data))
	return nil
}

// SModel describes a complete model.
type SModel struct {
	Description string   `json:"description"`
	Types       []*SType `json:"types"`
	Name        string   `json:"name"`
}

// SType describes a type.
type SType struct {
	Description string       `json:"description"`
	Properties  []*SProperty `json:"properties"`
	Links       []*SLink     `json:"links,omitempty"`
	Name        string       `json:"name"`
}

// SProperty describes a property (for type or association).
type SProperty struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	PrimaryKey  bool   `json:"primary_key"`
	DataType    string `json:"data_type"`
}

// SLink describes a relationship (association on composition).
type SLink struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Properties  []*SProperty `json:"properties,omitempty"`
	TargetType  string       `json:"links_to"`
}

// Listener is a model listener that builds up a SModel (simplified model).
type Listener struct {
	yammm.BaseModelListener
	ctx         yammm.Context
	writer      *pio.Writer
	model       *SModel
	currentType *SType
	currentLink *SLink
	*SLink
}

// EnterModel is an implementation of this yammm.ModelListener.
func (l *Listener) EnterModel(ctx yammm.Context) {
	m := ctx.Model()
	l.model = &SModel{Name: m.Name, Description: m.Documentation}
}

// EnterType is an implementation of this yammm.ModelListener.
func (l *Listener) EnterType(_ yammm.Context, t *yammm.Type) {
	st := &SType{
		Name:        t.Name,
		Description: t.Documentation,
	}
	l.currentType = st
	l.model.Types = append(l.model.Types, st)
}

// ExitType is an implementation of this yammm.ModelListener.
func (l *Listener) ExitType(_ yammm.Context, _ *yammm.Type) {
	l.currentType = nil
}

// OnProperty is an implementation of this yammm.ModelListener.
func (l *Listener) OnProperty(_ yammm.Context, _ *yammm.Type, p *yammm.Property) {
	name := p.Name
	if name == "id" {
		name = "uid"
	}
	sp := &SProperty{
		Name:        name,
		Description: p.Documentation,
		PrimaryKey:  p.IsPrimaryKey,
		DataType:    p.BaseType().Kind().String(),
	}
	l.currentType.Properties = append(l.currentType.Properties, sp)
}

// OnAssociation is an implementation of this yammm.ModelListener.
func (l *Listener) OnAssociation(ctx yammm.Context, _ *yammm.Type, a *yammm.Association) {
	sl := &SLink{
		Name:        a.PropertyName(ctx),
		TargetType:  a.To,
		Description: a.Documentation,
	}
	l.currentLink = sl
	l.currentType.Links = append(l.currentType.Links, sl)
}

// OnAssociationProperty is an implementation of this yammm.ModelListener.
func (l *Listener) OnAssociationProperty(_ yammm.Context, _ *yammm.Type, _ *yammm.Association, p *yammm.Property) {
	sp := &SProperty{
		Name:        p.Name,
		Description: p.Documentation,
		DataType:    p.BaseType().Kind().String(),
	}
	l.currentLink.Properties = append(l.currentLink.Properties, sp)
}

// OnComposition is an implementation of this yammm.ModelListener.
func (l *Listener) OnComposition(ctx yammm.Context, _ *yammm.Type, c *yammm.Composition) {
	sl := &SLink{
		Name:        c.PropertyName(ctx),
		TargetType:  c.To,
		Description: c.Documentation,
	}
	l.currentLink = sl
	l.currentType.Links = append(l.currentType.Links, sl)
}
