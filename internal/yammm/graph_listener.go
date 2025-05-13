package yammm

import (
	"github.com/wyrth-io/whit/internal/xray"
)

// GraphListener is an interface for listening to events as they occur when walking a yammm instance graph.
// For elements that may have children, there is an Enter/Exit pair of calls. For child elements there
// is an OnXxx call. See GraphWalker for how to walk a graph and get callbacks to this listener.
type GraphListener interface {
	// EnterGraph - is called first
	EnterGraph(ctx Context, topLevel xray.Wrapper)

	// EnterType - before any processing is done for this type. Is called once per type.
	EnterType(ctx Context, t *Type)

	// EnterInstance - before processing an instance of the type.
	EnterInstance(ctx Context, t *Type, data xray.Wrapper)

	// OnProperties - called once per instance with a map of all properties with
	// value or default value.
	OnProperties(ctx Context, t *Type, propMap map[string]any)

	// EnterAssociation - called for each association of the instance.
	EnterAssociation(ctx Context, t *Type, a *Association, data xray.Wrapper)

	// OnEdge - called for each association edge from type t to type tt. The properties of
	// the edge are passed in the assocPropMap, and the primary keys for to/from types are
	// also given.
	OnEdge(ctx Context, a *Association, assocPropMap map[string]any, t *Type,
		fromPkMap map[string]any, tt *Type, toPkMap map[string]any)

	// ExitAssociation - called once for each association from the instance.
	ExitAssociation(ctx Context, t *Type, a *Association, data xray.Wrapper)

	// EnterComposition - called once for each composition contained in the instance before each instance of
	// this composition
	EnterComposition(ctx Context, t *Type, c *Composition, data xray.Wrapper)

	// OnCompositionEdge - called for each composition edge from type t to type tt.
	// The primary keys for to/from types are given.
	OnCompositionEdge(ctx Context, c *Composition, t *Type,
		fromPkMap map[string]any, tt *Type, toPkMap map[string]any)

	// ExitComposition - called once for each composition of the instance after all instances of this
	// composition.
	ExitComposition(ctx Context, t *Type, c *Composition, data xray.Wrapper)

	// ExitInstance - called once per instance after processing all traits of an instance of the type.
	ExitInstance(ctx Context, t *Type, data xray.Wrapper)

	// ExitType - after all processing is done for this type. Is called once per type.
	// Call sequence 2.
	ExitType(ctx Context, t *Type)

	// ExitGraph - is called last.
	ExitGraph(ctx Context, topLevel xray.Wrapper)
}

// BaseGraphListener is an implementation of GraphListener that does nothing. It is intended to
// be mixed into specialized listeners to avoid having to implement all methods.
type BaseGraphListener struct {
	GraphListener
}

// EnterGraph implements this GraphListenerMethod and does nothing.
func (b *BaseGraphListener) EnterGraph(_ Context, _ xray.Wrapper) {}

// EnterType implements this GraphListenerMethod and does nothing.
func (b *BaseGraphListener) EnterType(_ Context, _ *Type) {}

// EnterInstance implements this GraphListenerMethod and does nothing.
func (b *BaseGraphListener) EnterInstance(_ Context, _ *Type, _ xray.Wrapper) {}

// OnProperties implements this GraphListenerMethod and does nothing.
func (b *BaseGraphListener) OnProperties(_ Context, _ *Type, _ map[string]any) {}

// EnterAssociation implements this GraphListenerMethod and does nothing.
func (b *BaseGraphListener) EnterAssociation(_ Context, _ *Type, _ *Association, _ xray.Wrapper) {
}

// OnEdge implements this GraphListenerMethod and does nothing.
func (b *BaseGraphListener) OnEdge(_ Context, _ *Association, _ map[string]any, _ *Type,
	_ map[string]any, _ *Type, _ map[string]any) {
}

// ExitAssociation implements this GraphListenerMethod and does nothing.
func (b *BaseGraphListener) ExitAssociation(_ Context, _ *Type, _ *Association, _ xray.Wrapper) {}

// EnterComposition implements this GraphListenerMethod and does nothing.
func (b *BaseGraphListener) EnterComposition(_ Context, _ *Type, _ *Composition, _ xray.Wrapper) {}

// OnCompositionEdge implements this GraphListenerMethod and does nothing.
func (b *BaseGraphListener) OnCompositionEdge(_ Context, _ *Composition, _ *Type,
	_ map[string]any, _ *Type, _ map[string]any) {
}

// ExitComposition implements this GraphListenerMethod and does nothing.
func (b *BaseGraphListener) ExitComposition(_ Context, _ *Type, _ *Composition, _ xray.Wrapper) {}

// ExitInstance implements this GraphListenerMethod and does nothing.
func (b *BaseGraphListener) ExitInstance(_ Context, _ *Type, _ xray.Wrapper) {}

// ExitType implements this GraphListenerMethod and does nothing.
func (b *BaseGraphListener) ExitType(_ Context, _ *Type) {}

// ExitGraph implements this GraphListenerMethod and does nothing.
func (b *BaseGraphListener) ExitGraph(_ Context, _ xray.Wrapper) {}

// PluggableListener is a GraphListener where each callback delegates to a
// pluggable function. This is useful for unit testing as new types does not have to be
// created.
type PluggableListener struct {
	BaseGraphListener
	// FEnterGraph is called for the corresponding listener method.
	FEnterGraph func(ctx Context, topLevel xray.Wrapper)
	// FExitGraph is called for the corresponding listener method.
	FExitGraph func(ctx Context, topLevel xray.Wrapper)
	// FEnterType is called for the corresponding listener method.
	FEnterType func(ctx Context, t *Type)
	// FExitType is called for the corresponding listener method.
	FExitType func(ctx Context, t *Type)
	// FEnterInstance is called for the corresponding listener method.
	FEnterInstance func(ctx Context, t *Type, data xray.Wrapper)
	// FExitInstance is called for the corresponding listener method.
	FExitInstance func(ctx Context, t *Type, data xray.Wrapper)
	// FOnProperties is called for the corresponding listener method.
	FOnProperties func(ctx Context, t *Type, propMap map[string]any)
	// FEnterAssociation is called for the corresponding listener method.
	FEnterAssociation func(ctx Context, t *Type, a *Association, data xray.Wrapper)
	// FExitAssociation is called for the corresponding listener method.
	FExitAssociation func(ctx Context, t *Type, a *Association, data xray.Wrapper)
	// FEnterComposition is called for the corresponding listener method.
	FEnterComposition func(ctx Context, t *Type, c *Composition, data xray.Wrapper)
	// FExitComposition is called for the corresponding listener method.
	FExitComposition func(ctx Context, t *Type, c *Composition, data xray.Wrapper)
	// FOnEdge is called for the corresponding listener method.
	FOnEdge func(ctx Context, a *Association, assocPropMap map[string]any, t *Type,
		fromPkMap map[string]any, tt *Type, toPkMap map[string]any)
	// FOnCompositionEdge is called for the corresponding listener method.
	FOnCompositionEdge func(ctx Context, c *Composition, t *Type,
		fromPkMap map[string]any, tt *Type, toPkMap map[string]any)
}

// EnterGraph implements this GraphListenerMethod and delegates to a pluggable function if set.
func (b *PluggableListener) EnterGraph(ctx Context, topLevel xray.Wrapper) {
	if b.FEnterGraph != nil {
		b.FEnterGraph(ctx, topLevel)
	}
}

// ExitGraph implements this GraphListenerMethod and delegates to a pluggable function if set.
func (b *PluggableListener) ExitGraph(ctx Context, topLevel xray.Wrapper) {
	if b.FExitGraph != nil {
		b.FExitGraph(ctx, topLevel)
	}
}

// EnterType implements this GraphListenerMethod and delegates to a pluggable function if set.
func (b *PluggableListener) EnterType(ctx Context, t *Type) {
	if b.FEnterType != nil {
		b.FEnterType(ctx, t)
	}
}

// ExitType implements this GraphListenerMethod and delegates to a pluggable function if set.
func (b *PluggableListener) ExitType(ctx Context, t *Type) {
	if b.FExitType != nil {
		b.FExitType(ctx, t)
	}
}

// EnterInstance implements this GraphListenerMethod and delegates to a pluggable function if set.
func (b *PluggableListener) EnterInstance(ctx Context, t *Type, data xray.Wrapper) {
	if b.FEnterInstance != nil {
		b.FEnterInstance(ctx, t, data)
	}
}

// ExitInstance implements this GraphListenerMethod and delegates to a pluggable function if set.
func (b *PluggableListener) ExitInstance(ctx Context, t *Type, data xray.Wrapper) {
	if b.FExitInstance != nil {
		b.FExitInstance(ctx, t, data)
	}
}

// OnProperties implements this GraphListenerMethod and delegates to a pluggable function if set.
func (b *PluggableListener) OnProperties(ctx Context, t *Type, propMap map[string]any) {
	if b.FOnProperties != nil {
		b.FOnProperties(ctx, t, propMap)
	}
}

// EnterAssociation implements this GraphListenerMethod and delegates to a pluggable function if set.
func (b *PluggableListener) EnterAssociation(ctx Context, t *Type, a *Association, data xray.Wrapper) {
	if b.FEnterAssociation != nil {
		b.FEnterAssociation(ctx, t, a, data)
	}
}

// OnEdge implements this GraphListenerMethod and delegates to a pluggable function if set.
func (b *PluggableListener) OnEdge(ctx Context, a *Association, assocPropMap map[string]any, t *Type,
	fromPkMap map[string]any, tt *Type, toPkMap map[string]any) {
	if b.FOnEdge != nil {
		b.FOnEdge(ctx, a, assocPropMap, t, fromPkMap, tt, toPkMap)
	}
}

// OnCompositionEdge implements this GraphListenerMethod and delegates to a pluggable function if set.
func (b *PluggableListener) OnCompositionEdge(ctx Context, c *Composition, t *Type,
	fromPkMap map[string]any, tt *Type, toPkMap map[string]any) {
	if b.FOnCompositionEdge != nil {
		b.FOnCompositionEdge(ctx, c, t, fromPkMap, tt, toPkMap)
	}
}

// ExitAssociation implements this GraphListenerMethod and delegates to a pluggable function if set.
func (b *PluggableListener) ExitAssociation(ctx Context, t *Type, a *Association, data xray.Wrapper) {
	if b.FExitAssociation != nil {
		b.FExitAssociation(ctx, t, a, data)
	}
}

// EnterComposition implements this GraphListenerMethod and delegates to a pluggable function if set.
func (b *PluggableListener) EnterComposition(ctx Context, t *Type, c *Composition, data xray.Wrapper) {
	if b.FEnterComposition != nil {
		b.FEnterComposition(ctx, t, c, data)
	}
}

// ExitComposition implements this GraphListenerMethod and delegates to a pluggable function if set.
func (b *PluggableListener) ExitComposition(ctx Context, t *Type, c *Composition, data xray.Wrapper) {
	if b.FExitComposition != nil {
		b.FExitComposition(ctx, t, c, data)
	}
}
