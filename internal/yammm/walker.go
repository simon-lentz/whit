package yammm

import (
	"reflect"

	"github.com/wyrth-io/whit/internal/xray"
)

// GraphWalker is a depth first walker of a Yammm instance graph compliant with a Yammm model.
type GraphWalker interface {
	// Walks the given context compliant graph. The graph must have been validated without errors, or this
	// will most likely panic.
	Walk(graph any)
}

type graphWalker struct {
	visitor GraphListener
	ctx     Context
	meta    Meta
}

// NewGraphWalker creates a GraphWalker with a given visitor. The visitor will receive callbacks
// when the GraphWalker.Walk(<<graph>>) is later called to walk the graph.
func NewGraphWalker(ctx Context, visitor GraphListener) GraphWalker {
	meta := NewMeta(ctx)
	return &graphWalker{visitor: visitor, ctx: ctx, meta: meta}
}

// Walk walks the entire graph in depth first order making callbacks to the GraphListener with events
// about walked elements.
func (walker *graphWalker) Walk(graph any) {
	// Since instance model is valid, the generation logic can be rely on data having the correct form.
	topLevel := xray.NewWrapper(graph)

	// The top level Graph fields are all array's of instances. Process those first, and
	// remember all that has associations or compositions. Process those relationships last
	// since both sides must exist. Do compositions before associations (depth first) since there
	// may be associations to/from composed parts.
	ctx := walker.ctx
	visitor := walker.visitor
	visitor.EnterGraph(ctx, topLevel)

	for _, plural := range topLevel.FeatureNames() {
		t := walker.ctx.LookupPluralType(plural)

		visitor.EnterType(ctx, t)
		instances := topLevel.Feature(plural)
		for i := 0; i < instances.Len(); i++ {
			data := instances.FeatureAtIndex(i)
			walker.walkInstance(t, data)
		}
		visitor.ExitType(walker.ctx, t)
	}
	visitor.ExitGraph(ctx, topLevel)
}
func (walker *graphWalker) walkInstance(t *Type, data xray.Wrapper) {
	visitor := walker.visitor
	ctx := walker.ctx

	visitor.EnterInstance(ctx, t, data)
	propMap := walker.propmap(data, t.AllProperties())
	visitor.OnProperties(ctx, t, propMap)
	for _, a := range t.AllAssociations() {
		walker.walkAssociation(t, a, data)
	}
	for _, c := range t.AllCompositions() {
		walker.walkComposition(t, c, data)
	}
	visitor.ExitInstance(walker.ctx, t, data)
}

func (walker *graphWalker) walkAssociation(t *Type, a *Association, data xray.Wrapper) {
	ctx := walker.ctx
	visitor := walker.visitor

	// the property name is the name of the relationship and the type in singular or plural depending on "many"
	fieldName := a.PropertyName(ctx)

	// This could be an optional or required association, but does not really matter as it is too late to do anything
	// if a required association is missing.
	edge := data.Feature(fieldName) // instance or slice of EDGE
	if edge != nil && edge.IsSlice() {
		sliceLen := edge.Len()
		if sliceLen == 0 {
			return
		}
		visitor.EnterAssociation(ctx, t, a, data)
		for i := 0; i < sliceLen; i++ {
			elem := edge.FeatureAtIndex(i)
			walker.walkEdge(t, data, a, elem) // TODO: Needs the overall data as well as the elem
		}
		visitor.ExitAssociation(ctx, t, a, data)
	} else {
		if edge == nil || (reflect.ValueOf(edge).Kind() == reflect.Ptr && reflect.ValueOf(edge).IsNil()) {
			return
		}
		visitor.EnterAssociation(ctx, t, a, data)
		walker.walkEdge(t, data, a, edge)
		visitor.ExitAssociation(ctx, t, a, data)
	}
}
func (walker *graphWalker) walkEdge(t *Type, tData xray.Wrapper, a *Association, data xray.Wrapper) {
	ctx := walker.ctx
	tt := ctx.LookupType(a.To)
	assocPropMap := walker.propmap(data, a.Properties)
	fromPkMap := walker.propmap(tData, t.AllPrimaryKeys())
	toPkMap := walker.propmap(data.Feature("Where"), tt.AllPrimaryKeys())
	walker.visitor.OnEdge(ctx, a, assocPropMap, t, fromPkMap, tt, toPkMap)
}

func (walker *graphWalker) walkPartEdge(t *Type, tData xray.Wrapper, c *Composition, data xray.Wrapper) {
	ctx := walker.ctx
	tt := ctx.LookupType(c.To)
	fromPkMap := walker.propmap(tData, t.AllPrimaryKeys())

	toPkMap := walker.propmap(data, tt.AllPrimaryKeys())
	walker.visitor.OnCompositionEdge(ctx, c, t, fromPkMap, tt, toPkMap)
}

func (walker *graphWalker) walkComposition(t *Type, c *Composition, data xray.Wrapper) {
	ctx := walker.ctx
	visitor := walker.visitor
	fieldName := c.PropertyName(ctx)
	// This could be an optional or required part, but does not really matter as it is too late to do anything
	// if a required composition is missing.
	part := data.Feature(fieldName)
	if part == nil || (reflect.ValueOf(part).Kind() == reflect.Ptr && reflect.ValueOf(part).IsNil()) {
		return
	}
	partType := walker.meta.TypeByName(c.To)
	if part.IsSlice() {
		sliceLen := part.Len()
		if sliceLen == 0 {
			return
		}
		visitor.EnterComposition(ctx, t, c, data)
		for i := 0; i < sliceLen; i++ {
			elem := part.FeatureAtIndex(i) // The part data for one composition
			walker.walkInstance(partType, elem)
			walker.walkPartEdge(t, data, c, elem)
		}
		visitor.ExitComposition(ctx, t, c, data)
	} else {
		// if part == nil || (reflect.ValueOf(part).Kind() == reflect.Ptr && reflect.ValueOf(part).IsNil()) {
		// 	return
		// }
		visitor.EnterComposition(ctx, t, c, data)
		walker.walkInstance(partType, part) // TODO: Walk Part
		walker.walkPartEdge(t, data, c, part)
		visitor.ExitComposition(ctx, t, c, data)
	}
}

func (walker *graphWalker) propmap(w xray.Wrapper, properties []*Property) map[string]any {
	result := make(map[string]any, len(properties))
	for _, p := range properties {
		propVal := w.Value(p.Name)
		// Omit missing property unless it has a Default value.
		if propVal == nil {
			if p.HasDefault() {
				valToUse, _ := p.DefaultStringValue()
				result[p.Name] = valToUse
			}
			continue
		}
		// Set values that are required or optional with a given value that is not go-zero.
		if !p.Optional || p.Optional && !reflect.ValueOf(propVal).IsZero() {
			result[p.Name] = propVal
		}
	}
	return result
}
