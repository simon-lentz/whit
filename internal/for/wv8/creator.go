package wv8

import (
	"context"

	"github.com/google/uuid"
	"github.com/tada/catch"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/data/replication"
	"github.com/wyrth-io/whit/internal/tc"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/yammm"
)

// Creator creates instances in a wv8 cluster.
type Creator interface {
	// Process a graph of data that has been validated against the model described by the context.
	// If the graph was not validated prior to this call the behaviour is undefined and will most likely panic.
	Process(ctx yammm.Context, graph any)
}

type creator struct {
	yammm.BaseGraphListener
	client         *weaviate.Client
	replacementMap map[string]uuid.UUID
	links          []link
}

type link struct {
	propName string // property name in fromType
	fromName string
	fromID   string
	toName   string
	toID     string
}

// NewCreator returns a creator for a weaviate cluster as determined by the given client.
func NewCreator(client *weaviate.Client) Creator {
	return &creator{client: client}
}

func (c *creator) Process(ctx yammm.Context, graph any) {
	// Get the ID replacements for the graph
	idMapper := yammm.NewIDMapper()
	var err error
	if c.replacementMap, err = idMapper.Map(ctx, graph); err != nil {
		panic(catch.Error("could not create UUID replacement map: %s", err))
	}

	walker := yammm.NewGraphWalker(ctx, c)
	walker.Walk(graph)

	c.storeLinks()
}
func (c *creator) OnProperties(_ yammm.Context, t *yammm.Type, propMap map[string]any) {
	var uid uuid.UUID
	var err error
	if uid, err = t.InstanceID(propMap, c.replacementMap); err != nil {
		panic(catch.Error("Cannot create id for from-type %s: %s", t.Name, err))
	}
	// TODO: Get vector properties from the type and create separate propMap for those.
	vecMap := map[string]any{}
	for _, p := range t.AllProperties() {
		if p.BaseType().Kind() == tc.SpacevectorKind {
			if sv, ok := propMap[p.Name]; ok {
				vecMap[p.Name] = sv
				delete(propMap, p.Name)
			}
		}
	}
	c.storeInstance(t.Name, uid, propMap, vecMap)
}

func (c *creator) OnEdge(
	ctx yammm.Context,
	a *yammm.Association,
	assocPropMap map[string]any,
	fromType *yammm.Type,
	fromPks map[string]any,
	toType *yammm.Type,
	toPks map[string]any,
) {
	var fromID, toID uuid.UUID
	var err error
	if fromID, err = fromType.InstanceID(fromPks, c.replacementMap); err != nil {
		panic(catch.Error("Cannot create id for from-type %s: %s", fromType.Name, err))
	}
	if toID, err = toType.InstanceID(toPks, c.replacementMap); err != nil {
		panic(catch.Error("Cannot create id for to-type %s: %s", toType.Name, err))
	}
	// If association has properties then an edge class is needed between from/to instances.
	if len(a.Properties) > 0 {
		// Requires an edge
		edgeName := EdgeName(fromType, a)
		edgeID, err := uuid.NewRandom()
		if err != nil {
			panic(catch.Error("Cannot generate random UUID for edge: %s", edgeName, err))
		}
		c.storeInstance(edgeName, edgeID, assocPropMap, map[string]any{})
		// Reference from toType to edge
		c.addLink(link{
			propName: a.PropertyName(ctx),
			fromName: fromType.Name, fromID: fromID.String(),
			toName: edgeName, toID: edgeID.String(),
		})
		// Reference from edgeType to toType
		c.addLink(link{
			propName: a.Name,
			fromName: edgeName, fromID: edgeID.String(),
			toName: toType.Name, toID: toID.String(),
		})
	} else {
		c.addLink(link{
			propName: a.PropertyName(ctx),
			fromName: fromType.Name, fromID: fromID.String(),
			toName: toType.Name, toID: toID.String(),
		})
	}
}

func (c *creator) OnCompositionEdge(ctx yammm.Context, comp *yammm.Composition,
	fromType *yammm.Type, fromPkMap map[string]any,
	toType *yammm.Type, toPkMap map[string]any) {
	var fromID, toID uuid.UUID
	var err error
	if fromID, err = fromType.InstanceID(fromPkMap, c.replacementMap); err != nil {
		panic(catch.Error("Cannot create id for from-type %s: %s", fromType.Name, err))
	}
	if toID, err = toType.InstanceID(toPkMap, c.replacementMap); err != nil {
		panic(catch.Error("Cannot create id for to-type %s: %s", toType.Name, err))
	}

	c.addLink(link{
		propName: comp.PropertyName(ctx),
		fromName: fromType.Name,
		fromID:   fromID.String(),
		toName:   toType.Name,
		toID:     toID.String(),
	})
}
func (c *creator) addLink(aLink link) {
	c.links = append(c.links, aLink)
}
func (c *creator) storeInstance(typeName string, uid uuid.UUID, propMap, vecMap map[string]any) {
	propsWithoutID := utils.FilterMap(propMap, func(k string, _ any) bool { return k != "id" })
	// Create instance. TODO: keep the "created" returned and log it in debug logging mode.
	dataCreator := c.client.Data().Creator().
		WithClassName(typeName).
		WithID(uid.String()).
		WithProperties(propsWithoutID).
		WithConsistencyLevel(replication.ConsistencyLevel.ALL) // default QUORUM

	if len(vecMap) > 0 {
		// for now only supports a single anonymous user supplied spacevector.
		keys := utils.Keys(vecMap)
		vec := vecMap[keys[0]].([]float32)
		dataCreator.WithVector(vec)
	}
	_, err := dataCreator.Do(context.Background())

	if err != nil {
		panic(catch.Error("could not create instance of type %s: %s",
			typeName, err,
		))
	}
}

func (c *creator) storeLinks() {
	var err error
	for i := range c.links {
		theLink := c.links[i]

		payload := c.client.Data().ReferencePayloadBuilder().
			WithClassName(theLink.toName).
			WithID(theLink.toID).
			Payload()

		err = c.client.Data().ReferenceCreator().
			WithClassName(theLink.fromName).
			WithID(theLink.fromID).
			WithReferenceProperty(FormatPropName(theLink.propName)).
			WithReference(payload).
			WithConsistencyLevel(replication.ConsistencyLevel.ALL).
			Do(context.Background())

		if err != nil {
			panic(catch.Error("could not create link from %s(%s) to %s(%s): %s",
				theLink.fromName, theLink.fromID, theLink.toName, theLink.toID, err,
			))
		}
	}
}
