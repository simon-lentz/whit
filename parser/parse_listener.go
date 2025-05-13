package parser

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/wyrth-io/whit/internal/validation"
	"github.com/wyrth-io/whit/internal/yammm"
)

// ParseListener implements the ANTLR listener interface for the YammmGrammar.
type ParseListener struct {
	*BaseYammmGrammarListener
	ctx               yammm.Context
	currentTypeName   string
	currentDT         []string
	currentProperties []*yammm.Property
	ic                validation.IssueCollector
	source            string
	currentInvariants []*yammm.Invariant
}

// NewParseListener returns an ANTLR parse listener for the Yammm grammar. It holds a context in which the
// Yammm model is built, and an issue collector to be used for reporting semantic errors during validation
// of the parsed model.
func NewParseListener(sourceRef string) *ParseListener {
	return &ParseListener{
		ctx:    yammm.NewContext(),
		ic:     validation.NewIssueCollector(),
		source: sourceRef,
	}
}

// ExitSchema_name takes the given schema name in the parsed source and creates a new Yammm Model with this name.
// This model is set as the main model of the Yammm Context. (This is the first thing that is parsed as it must come first
// in the .yammm file being parsed.
func (pl *ParseListener) ExitSchema_name(ctx *Schema_nameContext) { //nolint:all
	schemaNameToken := ctx.GetToken(YammmGrammarLexerSTRING, 0)
	modelName := schemaNameToken.GetText()
	model := yammm.NewModel(ConvertString(modelName))
	model.Source = pl.source // Pick up the source reference from when the pl was created and set in model
	symbol := schemaNameToken.GetSymbol()
	model.Line = symbol.GetLine()
	model.Column = symbol.GetColumn()
	if ctx.DOC_COMMENT() != nil {
		model.Documentation = stripDelimiters(ctx.DOC_COMMENT().GetText())
	}
	err := pl.ctx.SetMainModel(model)
	pl.ic.CollectFatalIfErrorf("Could not define Yammm data model %s: %s", modelName, err)
}
func locate(t *yammm.Located, source string, line, column int) {
	if t != nil {
		t.Source = source
		t.Line = line
		t.Column = column
	}
}
func (pl *ParseListener) EnterType(ctx *TypeContext) { //nolint:all
	typeName := ctx.Type_name().GetText()
	nameToken := ctx.Type_name().GetStart()
	line := nameToken.GetLine()
	column := nameToken.GetColumn()
	source := pl.source
	isAbstract := ctx.is_abstract != nil
	isPart := ctx.is_part != nil
	var err error
	var t *yammm.Type
	// TODO: The below won't work if an AddXXXType fails with error since it won't return
	// a *Type that can be located and reported!
	switch {
	case isPart:
		t, err = pl.ctx.AddCompositionPartType(typeName, []*yammm.Property{})
		locate(&t.Located, source, line, column)
		pl.ic.CollectFatalIfErrorf("%sCould not create part type %s: %s", t.Label(), typeName, err)
	case isAbstract:
		t, err = pl.ctx.AddAbstractType(typeName, []*yammm.Property{})
		locate(&t.Located, source, line, column)
		pl.ic.CollectFatalIfErrorf("%sCould not create abstract type %s: %s", t.Label(), typeName, err)
	default:
		t, err = pl.ctx.AddType(typeName, []*yammm.Property{})
		locate(&t.Located, source, line, column)
		pl.ic.CollectFatalIfErrorf("%sCould not create type %s: %s", t.Label(), typeName, err)
	}
	if err == nil {
		if ctx.DOC_COMMENT() != nil {
			t.Documentation = stripDelimiters(ctx.DOC_COMMENT().GetText())
		}
	}
	pl.currentTypeName = typeName
}
func (pl *ParseListener) ExitType(ctx *TypeContext) { //nolint:all
	pl.currentTypeName = ""
}

func (pl *ParseListener) ExitDatatype(ctx *DatatypeContext) { //nolint:all
	typeName := toLowerFirstC(ctx.Type_name().GetText())
	dt, err := pl.ctx.AddDataType(typeName, pl.currentDT)
	token := ctx.Type_name().GetStart()
	line := token.GetLine()
	col := token.GetColumn()
	locate(&dt.Located, pl.source, line, col)
	pl.ic.CollectFatalIfErrorf("Could not add datatype %s: %s", typeName, err)
	pl.currentTypeName = ""
	pl.currentDT = nil
}

func (pl *ParseListener) ExitExtends_types(ctx *Extends_typesContext) { //nolint:all
	for _, nameCtx := range ctx.AllType_name() {
		superName := nameCtx.GetText()
		err := pl.ctx.AddInherits(pl.currentTypeName, superName)
		pl.ic.CollectFatalIfErrorf("Could not make %s extend %s: %s", pl.currentTypeName, superName, err)
	}
}

// ExitProperty transforms the parse tree for a Yammm Property.
func (pl *ParseListener) ExitProperty(ctx *PropertyContext) {
	propName := ctx.Property_name().GetText() // given lc name or a lc_keyword
	isPrimary := ctx.is_primary != nil
	isRequired := ctx.is_required != nil
	if isPrimary {
		isRequired = true
	}
	ctx.Data_type_ref()
	prop, err := pl.ctx.AddProperty(pl.currentTypeName, propName, pl.currentDT, !isRequired, isPrimary)
	pl.ic.CollectFatalIfErrorf("Cannot add property %s to %s: %s", propName, pl.currentTypeName, err)
	pl.currentDT = nil
	if err == nil {
		if ctx.DOC_COMMENT() != nil {
			prop.Documentation = stripDelimiters(ctx.DOC_COMMENT().GetText())
		}
		line := ctx.Property_name().GetStart().GetLine()
		col := ctx.Property_name().GetStart().GetColumn()
		locate(&prop.Located, pl.source, line, col)
	}
}
func convertTypeRange(tok antlr.Token) string {
	if tok == nil {
		return ""
	}
	txt := tok.GetText()
	if txt == "_" {
		return ""
	}
	return txt
}

// ExitIntegerT is called when exiting the integerT production.
func (pl *ParseListener) ExitIntegerT(ctx *IntegerTContext) {
	hasMin := !(ctx.min == nil || ctx.min.GetTokenType() == YammmGrammarLexerUSCORE)
	hasMax := !(ctx.max == nil || ctx.max.GetTokenType() == YammmGrammarLexerUSCORE)
	switch {
	case hasMin && hasMax:
		pl.currentDT = []string{"Integer", convertTypeRange(ctx.min), convertTypeRange(ctx.max)}
	case hasMin:
		pl.currentDT = []string{"Integer", convertTypeRange(ctx.min), ""}
	case hasMax:
		pl.currentDT = []string{"Integer", "", convertTypeRange(ctx.min)}
	default:
		pl.currentDT = []string{"Integer", "", ""}
	}
}

// ExitFloatT is called when exiting the floatT production.
func (pl *ParseListener) ExitFloatT(ctx *FloatTContext) {
	hasMin := !(ctx.min == nil || ctx.min.GetTokenType() == YammmGrammarLexerUSCORE)
	hasMax := !(ctx.max == nil || ctx.max.GetTokenType() == YammmGrammarLexerUSCORE)
	switch {
	case hasMin && hasMax:
		pl.currentDT = []string{"Float", convertTypeRange(ctx.min), convertTypeRange(ctx.max)}
	case hasMin:
		pl.currentDT = []string{"Float", convertTypeRange(ctx.min), ""}
	case hasMax:
		pl.currentDT = []string{"Float", "", convertTypeRange(ctx.max)}
	default:
		pl.currentDT = []string{"Float", "", ""}
	}
}

// ExitBoolT is called when exiting the boolT production.
func (pl *ParseListener) ExitBoolT(_ *BoolTContext) {
	pl.currentDT = []string{"Boolean"}
}

// ExitStringT is called when exiting the stringT production.
func (pl *ParseListener) ExitStringT(ctx *StringTContext) {
	hasMin := ctx.min != nil && ctx.min.GetTokenType() == YammmGrammarLexerINTEGER
	hasMax := ctx.max != nil && ctx.max.GetTokenType() == YammmGrammarLexerINTEGER
	switch {
	case hasMin && hasMax:
		pl.currentDT = []string{"String", convertTypeRange(ctx.min), convertTypeRange(ctx.max)}
	case hasMin:
		pl.currentDT = []string{"String", convertTypeRange(ctx.min), ""}
	case hasMax:
		pl.currentDT = []string{"String", "", convertTypeRange(ctx.max)}
	default:
		pl.currentDT = []string{"String", "", ""}
	}
}

// ExitEnumT is called when exiting the enumT production.
func (pl *ParseListener) ExitEnumT(ctx *EnumTContext) {
	enumTokens := ctx.AllSTRING()
	enumStrings := make([]string, len(enumTokens))
	pl.currentDT = []string{"Enum"}
	// convert tokens ('val' or "val") to normalized strings ("val") for cue (not needed for DT instructions)
	for i, e := range enumTokens {
		text := e.GetText()
		pl.currentDT = append(pl.currentDT, ConvertString(text))
		enumStrings[i] = fmt.Sprintf(`"%s"`, ConvertString(text))
	}
}

// ExitPatternT is called when exiting the patternT production.
func (pl *ParseListener) ExitPatternT(ctx *PatternTContext) {
	patternTokens := ctx.AllSTRING()
	patternStrings := make([]string, len(patternTokens))
	pl.currentDT = []string{"Pattern"}
	// convert tokens ('val' or "val") to normalized strings ("val")
	// regexp match is done with unary =~ operator
	for i, e := range patternTokens {
		text := e.GetText()
		pl.currentDT = append(pl.currentDT, ConvertString(text))
		patternStrings[i] = fmt.Sprintf(`=~ "%s"`, ConvertString(text))
	}
}

// ExitTimestampT is called when exiting the timestampT production.
// When used without a pattern the default pattern is "2006-01-02T15:04:05Z07:00" (RFC339).
func (pl *ParseListener) ExitTimestampT(ctx *TimestampTContext) {
	// #date: string & time.Format(time.RFC3339Date)
	formatToken := ctx.format
	if formatToken == nil {
		pl.currentDT = []string{"Timestamp"}
	} else {
		// given as input, while it would be nie to validate the format itself there is no
		// such method as format allows arbitrary text and then ends up simply not matching.
		text := formatToken.GetText()
		pl.currentDT = []string{"Timestamp", ConvertString(text)}
	}
}

// ExitDateT is called when exiting the dateT production.
// A date is on the form "2006-01-02" (Date part of RFC3339).
func (pl *ParseListener) ExitDateT(_ *DateTContext) {
	pl.currentDT = []string{"Date"}
}

// ExitUuidT is called when exiting the dateT production.
func (pl *ParseListener) ExitUuidT(_ *UuidTContext) { //nolint
	pl.currentDT = []string{"UUID"}
}

// ExitSpacevectorT is called when exiting the spacevectorT production.
// The type always specifies the dimensionality of the vector.
func (pl *ParseListener) ExitSpacevectorT(ctx *SpacevectorTContext) {
	// One of space name (string) or dimensions (integer) will be set
	dimText := ctx.dimensions.GetText()
	pl.currentDT = []string{"Spacevector", dimText}
}

// ExitAlias is called when exiting the alias production. (A reference to a user defined data type).
func (pl *ParseListener) ExitAlias(ctx *AliasContext) {
	pl.currentDT = []string{"Alias", toLowerFirstC(ctx.GetText())}
}

// EnterAssociation sets an empty set of properties in the listeners context to be picked up
// in the ExitAssociation callback.
func (pl *ParseListener) EnterAssociation(_ *AssociationContext) {
	pl.currentProperties = []*yammm.Property{}
	pl.currentInvariants = []*yammm.Invariant{}
}

// ExitAssociation is called when production association is exited.
func (pl *ParseListener) ExitAssociation(ctx *AssociationContext) {
	optional, many := pl.HandleMultiplicity(ctx.thisMp)
	relName := ctx.thisName.GetText()
	toTypeName := ctx.toType.GetText()
	doc := ""
	if ctx.DOC_COMMENT() != nil {
		doc = stripDelimiters(ctx.DOC_COMMENT().GetText())
	}
	_, err := pl.ctx.AddAssociation(pl.currentTypeName, relName, toTypeName, optional, many, pl.currentProperties, doc)
	pl.ic.CollectFatalIfErrorf("Could not add association %s from %s to %s: %s", relName, pl.currentTypeName, toTypeName, err)
}

// EnterComposition sets an empty set of properties in the listeners context to be picked up
// in the ExitComposition callback.
func (pl *ParseListener) EnterComposition(_ *CompositionContext) {
	pl.currentProperties = []*yammm.Property{}
}

// ExitComposition is called when production association is exited.
func (pl *ParseListener) ExitComposition(ctx *CompositionContext) {
	optional, many := pl.HandleMultiplicity(ctx.thisMp)
	toTypeName := ctx.toType.GetText()
	doc := ""
	if ctx.DOC_COMMENT() != nil {
		doc = stripDelimiters(ctx.DOC_COMMENT().GetText())
	}
	var name string
	if ctx.thisName != nil {
		name = ctx.thisName.GetText()
	}
	_, err := pl.ctx.AddComposition(pl.currentTypeName, name, toTypeName, optional, many, doc)
	pl.ic.CollectFatalIfErrorf("Could not add composition %s from %s to %s: %s", name, pl.currentTypeName, toTypeName, err)
}

// ExitRel_property adds a relationship property to the listeners context to be picked up later.
func (pl *ParseListener) ExitRel_property(ctx *Rel_propertyContext) { //nolint:all
	propName := ctx.Property_name().GetText() // given lc name or a lc_keyword
	isRequired := ctx.is_required != nil
	// ctx.Data_type_ref()
	doc := ""
	if ctx.DOC_COMMENT() != nil {
		doc = stripDelimiters(ctx.DOC_COMMENT().GetText())
	}
	pl.currentProperties = append(pl.currentProperties, &yammm.Property{
		Name:          propName,
		DataType:      pl.currentDT,
		Optional:      !isRequired,
		IsPrimaryKey:  false,
		Documentation: doc,
	})
}

// ExitInvariant is called when exiting the boolT production.
func (pl *ParseListener) ExitInvariant(ctx *InvariantContext) {
	visitor := &ExprVisitor{}
	name, err := strconv.Unquote(ctx.message.GetText())
	if err != nil {
		panic(fmt.Errorf("invariant name as illegal string name: %s", err.Error()))
	}
	_ = pl.ctx.AddInvariant(pl.currentTypeName, &yammm.Invariant{Name: name, Constraint: visitor.Visit(ctx.constraint)})
}

// HandleMultiplicity handles the various combinations of one:many that are possible.
func (pl *ParseListener) HandleMultiplicity(multiplicity IMultiplicityContext) (optional, many bool) {
	if multiplicity == nil {
		return true, false
	}
	switch multiplicity.GetText() {
	case "(_)", "(_:one)", "(one)":
		optional = true
		many = false
	case "(many)", "(_:many)":
		optional = true
		many = true
	case "(one:one)":
		optional = false
		many = false
	case "(one:many)":
		optional = false
		many = true
	default:
		pl.ic.Collectf(validation.Fatal, "internal error: unknown multiplicty parsed by grammar: %s", multiplicity.GetText())
	}
	return
}
