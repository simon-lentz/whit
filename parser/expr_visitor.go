package parser

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/tada/catch"
	"github.com/wyrth-io/whit/internal/yammm"
)

// ExprVisitor is an ANTLR visitor that produces a yammm.Expression from
// a parse tree for the yammm DSL.
type ExprVisitor struct {
	BaseYammmGrammarVisitor
}

// Visit implements this visitor method.
func (v *ExprVisitor) Visit(tree antlr.ParseTree) yammm.Expression {
	if tree == nil {
		return nil
	}
	switch val := tree.(type) {
	case *LiteralContext:
		return v.VisitLiteral(val)
	case *ValueContext:
		return v.VisitValue(val)
	case *VariableContext:
		return v.VisitVariable(val)
	case *MuldivContext:
		return v.VisitMuldiv(val)
	case *PlusminusContext:
		return v.VisitPlusminus(val)
	case *CompareContext:
		return v.VisitCompare(val)
	case *EqualityContext:
		return v.VisitEquality(val)
	case *MatchContext:
		return v.VisitMatch(val)
	case *InContext:
		return v.VisitIn(val)
	case *IfContext:
		return v.VisitIf(val)
	case *GroupContext:
		return v.VisitGroup(val)
	case *AndContext:
		return v.VisitAnd(val)
	case *OrContext:
		return v.VisitOr(val)
	case *NotContext:
		return v.VisitNot(val)
	case *AtContext:
		return v.VisitAt(val)
	case *ListContext:
		return v.VisitList(val)
	case *DatatypeNameContext:
		return v.VisitDatatypeName(val)
	case *DatatypeKeywordContext:
		panic("wtf")
	case *ArgumentsContext:
		return v.VisitArguments(val)
	case *ParametersContext:
		return v.VisitParameters(val)
	case *FcallContext:
		return v.VisitFcall(val)
	case *UminusContext:
		return v.VisitUminus(val)
	case *LiteralNilContext:
		return yammm.NewLiteral(nil)
	case *NameContext:
		return v.VisitName(val)
	case *PeriodContext:
		return v.VisitPeriod(val)
	// case *RegexpContext:
	// 	regexpText := val.left.GetText()
	// 	return yammm.NewLiteral(regexp.MustCompile(regexpText[1 : len(regexpText)-1]))
	case *ExprContext:
		return v.VisitExpr(val)
	default:
		s := fmt.Sprintf("Unknown parser context: '%t': %s", val, reflect.TypeOf(tree).Name())
		panic(s)
	}
}

// VisitExpr implements this visitor method.
func (v *ExprVisitor) VisitExpr(ctx *ExprContext) yammm.Expression {
	children := ctx.GetChildren()
	fmt.Printf("Children of expr: %d", len(children))
	panic(catch.Error("invalid expression: '%s'", ctx.GetText()))
}

// VisitValue implements this visitor method.
func (v *ExprVisitor) VisitValue(ctx *ValueContext) yammm.Expression {
	return v.Visit(ctx.left)
}

// VisitLiteral implements this visitor method.
func (v *ExprVisitor) VisitLiteral(ctx *LiteralContext) yammm.Expression {
	token := ctx.v
	switch token.GetTokenType() {
	case YammmGrammarLexerSTRING:
		s, err := strconv.Unquote(token.GetText())
		if err != nil {
			panic(err)
		}
		return yammm.NewLiteral(s)
	case YammmGrammarLexerINTEGER:
		i, err := strconv.ParseInt(token.GetText(), 10, 64)
		if err != nil {
			panic(err)
		}
		return yammm.NewLiteral(i)
	case YammmGrammarLexerFLOAT:
		f, err := strconv.ParseFloat(token.GetText(), 64)
		if err != nil {
			panic(err)
		}
		return yammm.NewLiteral(f)
	case YammmGrammarLexerBOOLEAN:
		b, err := strconv.ParseBool(token.GetText())
		if err != nil {
			panic(err)
		}
		return yammm.NewLiteral(b)
	case YammmGrammarLexerREGEXP:
		s := token.GetText()
		re := s[1 : len(s)-1]
		r := regexp.MustCompile(re)
		return yammm.NewLiteral(r)
	}
	return nil
}

// VisitUminus implements this visitor method.
func (v *ExprVisitor) VisitUminus(ctx *UminusContext) yammm.Expression {
	return yammm.SExpr{yammm.Op("-x"), v.Visit(ctx.right)}
}

// VisitVariable implements this visitor method.
func (v *ExprVisitor) VisitVariable(ctx *VariableContext) yammm.Expression {
	s := ctx.left.GetText()
	return yammm.SExpr{yammm.Op("$"), yammm.NewLiteral(s[1:])}
}

// VisitPeriod implements this visitor method.
func (v *ExprVisitor) VisitPeriod(ctx *PeriodContext) yammm.Expression {
	// Change rhs to literal name if it is a "get property of implied $self".
	nameExpr := v.Visit(ctx.name)
	if nameExpr.Op() == "p" {
		nameExpr = nameExpr.Children()[0]
	}
	return yammm.SExpr{yammm.Op("."), v.Visit(ctx.left), nameExpr}
}

// VisitName implements this visitor method.
func (v *ExprVisitor) VisitName(ctx *NameContext) yammm.Expression {
	s := ctx.left.GetText()
	return yammm.SExpr{yammm.Op("p"), yammm.NewLiteral(s)}
}

// VisitMuldiv implements this visitor method.
func (v *ExprVisitor) VisitMuldiv(ctx *MuldivContext) yammm.Expression {
	return v.binaryExpr(ctx.op, ctx.left, ctx.right)
}

// VisitPlusminus implements this visitor method.
func (v *ExprVisitor) VisitPlusminus(ctx *PlusminusContext) yammm.Expression {
	return v.binaryExpr(ctx.op, ctx.left, ctx.right)
}

// VisitCompare implements this visitor method.
func (v *ExprVisitor) VisitCompare(ctx *CompareContext) yammm.Expression {
	return v.binaryExpr(ctx.op, ctx.left, ctx.right)
}

// VisitEquality implements this visitor method.
func (v *ExprVisitor) VisitEquality(ctx *EqualityContext) yammm.Expression {
	return v.binaryExpr(ctx.op, ctx.left, ctx.right)
}

// VisitMatch implements this visitor method.
func (v *ExprVisitor) VisitMatch(ctx *MatchContext) yammm.Expression {
	return v.binaryExpr(ctx.op, ctx.left, ctx.right)
}

// VisitIn implements this visitor method.
func (v *ExprVisitor) VisitIn(ctx *InContext) yammm.Expression {
	return v.binaryExpr(ctx.op, ctx.left, ctx.right)
}

// VisitOr implements this visitor method.
func (v *ExprVisitor) VisitOr(ctx *OrContext) yammm.Expression {
	return v.binaryExpr(ctx.op, ctx.left, ctx.right)
}

// VisitAnd implements this visitor method.
func (v *ExprVisitor) VisitAnd(ctx *AndContext) yammm.Expression {
	return v.binaryExpr(ctx.op, ctx.left, ctx.right)
}

// VisitIf implements this visitor method.
func (v *ExprVisitor) VisitIf(ctx *IfContext) yammm.Expression {
	return yammm.SExpr{yammm.Op(
		ctx.op.GetText()),
		v.Visit(ctx.left),
		v.Visit(ctx.trueBranch),
		v.Visit(ctx.falseBranch),
	}
}

// VisitNot implements this visitor method.
func (v *ExprVisitor) VisitNot(ctx *NotContext) yammm.Expression {
	return yammm.SExpr{yammm.Op(
		ctx.op.GetText()),
		v.Visit(ctx.right),
	}
}

// VisitGroup implements this visitor method.
func (v *ExprVisitor) VisitGroup(ctx *GroupContext) yammm.Expression {
	return v.Visit(ctx.left)
}

// VisitAt implements this visitor method.
func (v *ExprVisitor) VisitAt(ctx *AtContext) yammm.Expression {
	right := ctx.right
	length := len(right)
	args := make([]yammm.Expression, length)
	for i := range right {
		args[i] = v.Visit(right[i])
	}
	return append(yammm.SExpr{
		yammm.Op("@"),
		v.Visit(ctx.left)},
		args...)
}

// VisitList implements this visitor method.
func (v *ExprVisitor) VisitList(ctx *ListContext) yammm.Expression {
	values := ctx.values
	length := len(values)
	elements := make([]yammm.Expression, length)
	for i := range values {
		elements[i] = v.Visit(values[i])
	}
	return append(yammm.SExpr{yammm.Op("[]")}, elements...)
}

// VisitDatatypeName implements this visitor method.
func (v *ExprVisitor) VisitDatatypeName(ctx *DatatypeNameContext) yammm.Expression {
	return yammm.DatatypeLiteral(ctx.left.GetText())
}

// VisitFcall implements this visitor method.
func (v *ExprVisitor) VisitFcall(ctx *FcallContext) yammm.Expression {
	lhs := v.Visit(ctx.left)      // Operand.
	fname := ctx.name.GetText()   // Function name, must be a literal string.
	args := v.Visit(ctx.args)     // Literal []Expression to be evaluated into arguments.
	params := v.Visit(ctx.params) // Literal []string with parameter names.
	body := v.Visit(ctx.body)     // Lambda Expression to evaluate.
	return yammm.SExpr{yammm.Op(v.MustBeBuiltInFunctionName(fname)), lhs, args, params, body}
}

// VisitArguments implements this visitor method.
func (v *ExprVisitor) VisitArguments(ctx *ArgumentsContext) yammm.Expression {
	a := make([]any, len(ctx.args))
	for i := range ctx.args {
		a[i] = v.Visit(ctx.args[i])
	}
	return yammm.NewLiteral(a)
}

// VisitParameters implements this visitor method.
func (v *ExprVisitor) VisitParameters(ctx *ParametersContext) yammm.Expression {
	p := make([]string, len(ctx.params))
	for i := range ctx.params {
		p[i] = ctx.params[i].GetText()[1:]
	}
	return yammm.NewLiteral(p)
}
func (v *ExprVisitor) binaryExpr(op antlr.Token, left antlr.ParseTree, right antlr.ParseTree) yammm.Expression {
	return yammm.SExpr{yammm.Op(op.GetText()), v.Visit(left), v.Visit(right)}
}

// MustBeBuiltInFunctionName returns the given name if it is a built in function and
// panics otherwise.
func (v *ExprVisitor) MustBeBuiltInFunctionName(n string) string {
	switch n {
	case "map", "filter", "reduce",
		"then", "lest", "with",
		"all", "any", "count", "all_or_none",
		"compact", "unique",
		"compare", "match",
		"len",
		"abs", "min", "max", "floor", "ceil", "round":
		return n
	}
	panic(fmt.Errorf("unreognized function: '%s'", n))
}
