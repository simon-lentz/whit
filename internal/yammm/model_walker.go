package yammm

import "fmt"

// ModelWalker is a depth first walker of a Yammm model. It walks all concrete types and flattens
// all inherited traits.
type ModelWalker interface {
	// Walks the given context which must have been completed without errors.
	Walk()
}
type modelWalker struct {
	visitor ModelListener
	ctx     Context
	meta    Meta
}

// NewModelWalker creates a ModelWalker with a given visitor. The visitor will receive callbacks
// when the ModelWalker.Walk(<<ctx>>) is later called to walk the graph.
func NewModelWalker(ctx Context, visitor ModelListener) ModelWalker {
	if !ctx.IsCompleted() {
		panic(fmt.Errorf("context given to NewModelWalker is not completed"))
	}
	meta := NewMeta(ctx)
	return &modelWalker{visitor: visitor, ctx: ctx, meta: meta}
}

// Walk walks the model making callbacks to the ModelListener with events about walked elements.
// Only concrete types result in callbacks. All (i.e. including inherited) properties, associations and
// compositions are included in the walk.
func (walker *modelWalker) Walk() {
	walker.visitor.EnterModel(walker.ctx)
	for _, t := range walker.ctx.Model().Types {
		walker.walkType(t)
	}
	walker.visitor.ExitModel(walker.ctx)
}
func (walker *modelWalker) walkType(t *Type) {
	if t.IsAbstract {
		return
	}
	walker.visitor.EnterType(walker.ctx, t)
	for _, p := range t.AllProperties() {
		walker.walkProperty(walker.ctx, t, p)
	}
	for _, a := range t.AllAssociations() {
		walker.walkAssociation(walker.ctx, t, a)
	}
	for _, c := range t.AllCompositions() {
		walker.walkComposition(walker.ctx, t, c)
	}
	walker.visitor.ExitType(walker.ctx, t)
}

func (walker *modelWalker) walkProperty(ctx Context, t *Type, p *Property) {
	walker.visitor.OnProperty(ctx, t, p)
}

func (walker *modelWalker) walkAssociation(ctx Context, t *Type, a *Association) {
	walker.visitor.OnAssociation(ctx, t, a)
	for _, p := range a.Properties {
		walker.visitor.OnAssociationProperty(walker.ctx, t, a, p)
	}
}

func (walker *modelWalker) walkComposition(ctx Context, t *Type, c *Composition) {
	walker.visitor.OnComposition(ctx, t, c)
}
