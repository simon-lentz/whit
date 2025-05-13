package yammm

// ModelListener is an interface for listening to events as they occur when walking a yammm model.
// For elements that may have children, there is an Enter/Exit pair of calls. For child elements there
// is an OnXxx call. See ModelWalker for how to walk a graph and get callbacks to this listener.
type ModelListener interface {
	// EnterModel - is called first
	EnterModel(ctx Context)

	// EnterType - before any processing is done for this type. Is called once per type.
	EnterType(ctx Context, t *Type)

	// OnProperty - called once per property for the modeled type.
	OnProperty(ctx Context, t *Type, property *Property)

	// OnAssociation - called for each association for the modeled type.
	OnAssociation(ctx Context, t *Type, a *Association)

	// OnAssociationProperty - called for each association's property for the modeled type.
	OnAssociationProperty(ctx Context, t *Type, a *Association, p *Property)

	// OnComposition - called once for each composition in the modeled type.
	OnComposition(ctx Context, t *Type, c *Composition)

	// ExitType - after all processing is done for this type. Is called once per type.
	ExitType(ctx Context, t *Type)

	// ExitModel - is called last.
	ExitModel(ctx Context)
}

// BaseModelListener is an implementation of ModelListener that does nothing. It is intended to
// be mixed into specialized listeners to avoid having to implement all methods.
type BaseModelListener struct {
	ModelListener
}

// EnterModel implements this GraphListenerMethod and does nothing.
func (b *BaseModelListener) EnterModel(_ Context) {}

// EnterType implements this GraphListenerMethod and does nothing.
func (b *BaseModelListener) EnterType(_ Context, _ *Type) {}

// OnProperty implements this GraphListenerMethod and does nothing.
func (b *BaseModelListener) OnProperty(_ Context, _ *Type, _ *Property) {}

// OnAssociation implements this GraphListenerMethod and does nothing.
func (b *BaseModelListener) OnAssociation(_ Context, _ *Type, _ *Association) {}

// OnAssociationProperty implements this GraphListenerMethod and does nothing.
func (b *BaseModelListener) OnAssociationProperty(_ Context, _ *Type, _ *Association, _ *Property) {}

// OnComposition implements this GraphListenerMethod and does nothing.
func (b *BaseModelListener) OnComposition(_ Context, _ *Type, _ *Composition) {}

// ExitType implements this GraphListenerMethod and does nothing.
func (b *BaseModelListener) ExitType(_ Context, _ *Type) {}

// ExitModel implements this GraphListenerMethod and does nothing.
func (b *BaseModelListener) ExitModel(_ Context) {}

// PluggableModelListener is a GraphListener where each callback delegates to a
// pluggable function. This is useful for unit testing as new types does not have to be
// created.
type PluggableModelListener struct {
	BaseModelListener
	// FEnterModel is called for the corresponding listener method.
	FEnterModel func(ctx Context)
	// FExitModel is called for the corresponding listener method.
	FExitModel func(ctx Context)
	// FEnterType is called for the corresponding listener method.
	FEnterType func(ctx Context, t *Type)
	// FExitType is called for the corresponding listener method.
	FExitType func(ctx Context, t *Type)
	// FOnProperty is called for the corresponding listener method.
	FOnProperty func(ctx Context, t *Type, property *Property)
	// FOnAssociation is called for the corresponding listener method.
	FOnAssociation func(ctx Context, t *Type, a *Association)
	// FOnAssociationProperty is called for the corresponding listener method.
	FOnAssociationProperty func(ctx Context, t *Type, a *Association, p *Property)
	// FOnComposition is called for the corresponding listener method.
	FOnComposition func(ctx Context, t *Type, c *Composition)
}

// EnterModel implements this GraphListenerMethod and delegates to a pluggable function if set.
func (b *PluggableModelListener) EnterModel(ctx Context) {
	if b.FEnterModel != nil {
		b.FEnterModel(ctx)
	}
}

// ExitModel implements this GraphListenerMethod and delegates to a pluggable function if set.
func (b *PluggableModelListener) ExitModel(ctx Context) {
	if b.FExitModel != nil {
		b.FExitModel(ctx)
	}
}

// EnterType implements this GraphListenerMethod and delegates to a pluggable function if set.
func (b *PluggableModelListener) EnterType(ctx Context, t *Type) {
	if b.FEnterType != nil {
		b.FEnterType(ctx, t)
	}
}

// ExitType implements this GraphListenerMethod and delegates to a pluggable function if set.
func (b *PluggableModelListener) ExitType(ctx Context, t *Type) {
	if b.FExitType != nil {
		b.FExitType(ctx, t)
	}
}

// OnProperty implements this GraphListenerMethod and delegates to a pluggable function if set.
func (b *PluggableModelListener) OnProperty(ctx Context, t *Type, property *Property) {
	if b.FOnProperty != nil {
		b.FOnProperty(ctx, t, property)
	}
}

// OnAssociation implements this GraphListenerMethod and delegates to a pluggable function if set.
func (b *PluggableModelListener) OnAssociation(ctx Context, t *Type, a *Association) {
	if b.FOnAssociation != nil {
		b.FOnAssociation(ctx, t, a)
	}
}

// OnAssociationProperty implements this GraphListenerMethod and delegates to a pluggable function if set.
func (b *PluggableModelListener) OnAssociationProperty(ctx Context, t *Type, a *Association, p *Property) {
	if b.FOnAssociationProperty != nil {
		b.FOnAssociationProperty(ctx, t, a, p)
	}
}

// OnComposition implements this GraphListenerMethod and delegates to a pluggable function if set.
func (b *PluggableModelListener) OnComposition(ctx Context, t *Type, c *Composition) {
	if b.FOnComposition != nil {
		b.FOnComposition(ctx, t, c)
	}
}
