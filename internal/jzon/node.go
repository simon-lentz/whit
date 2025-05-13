package jzon

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"github.com/tada/catch"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/xray"
)

// Context describes one JSON source context. Use NewContext() to create a context.
type Context interface {
	// Line returns the source line for an offset into the JSON source this represents.
	Line(offset int) int

	// Col returns the source column for an offset into the JSON source this represents.
	Col(offset int) int

	// UnmarshalNode unmarshals the JSON source into a Node.
	UnmarshalNode() (Node, error)
}

// context implements Context.
type context struct {
	lineoffsets []int
	sourceName  string
	source      string
	decoder     *json.Decoder
	sLen        int
	acceptNull  bool
	allTokens   []stackEntry
	tokenIdx    int
}

var nlRe = regexp.MustCompile(`\n`)

// Option is an option to NewContext.
type Option string

const (
	// AcceptNull Option makes the Context accept JSON null values.
	AcceptNull = Option("accept-null")
)

// NewContext returns a new Context for unmarshaling JSON content into a Node(s).
// The option "accept-null" makes the context produce IsNull nodes, without this option
// the UnmarshalNode will return an error.
func NewContext(sourceName, source string, options ...Option) Context {
	ctx := &context{sourceName: sourceName, source: source}
	found := nlRe.FindAllStringIndex(source, -1)
	ctx.lineoffsets = make([]int, len(found))
	for i, loc := range found {
		ctx.lineoffsets[i] = loc[0]
	}
	ctx.decoder = json.NewDecoder(strings.NewReader(source))
	ctx.decoder.UseNumber()
	ctx.sLen = len(source)
	for i := range options {
		if options[i] == AcceptNull {
			ctx.acceptNull = true
		}
	}
	return ctx
}

func (ctx *context) Line(offset int) int {
	return sort.SearchInts(ctx.lineoffsets, offset) + 1
}

func (ctx *context) Col(offset int) int {
	lineIdx := sort.SearchInts(ctx.lineoffsets, offset)
	var col int
	if lineIdx > 0 {
		col = offset - ctx.lineoffsets[lineIdx-1]
	} else {
		col = offset + 1
	}
	return col
}

// Node is an interface for JSON Object, Array or Value elements obtained from Context.UnmarshalNode.
// A Node has reflection methods IsXXX to obtain the type of node, and specific methods XXXValue()
// allows retreival of the underlying value (int64, float64, bool, string, []Node, and [Node]Node).
// The methods Property() and PropertyNames are specific to a Node representing an Object.
// The Line() and Col() methods returns the JSON source text line number at the start of what
// the node represents, and the column on that line.
type Node interface {
	xray.Wrapper
	xray.Position

	// IsObject returns true if this represents a JSON Object.
	IsObject() bool

	// IsArray returns true if this represents a JSON Array.
	IsArray() bool

	// IsInt returns true if this represents an Integer JSON Number.
	IsInt() bool

	// IsFloat returns true if this represents a Float JSON Number.
	IsFloat() bool

	// IsBool returns true if this represents a JSON Boolean.
	IsBool() bool

	// IsString returns true if this represents a JSON String.
	IsString() bool

	// IsNull returns true if this represents JSON null value.
	IsNull() bool

	// ObjectValue returns a map of associations of property name node to property value node.
	// Panics if this does not represent a JSON Object.
	ObjectValue() map[Node]Node

	// ArrayValue returns a slice of the Nodes being the elemnts of the Array.
	// Panics if this does not represent an Array.
	ArrayValue() []Node

	// IntValue returns an int64 value. Panics if this is not representing an Integer value.
	IntValue() int64

	// FloatValue returns a float64 value. Panics if this is not representing a Float value.
	FloatValue() float64

	// BoolValue returns a boolean value. Panics if this is not representing a Boolean value.
	BoolValue() bool

	// StringValue returns a string value. Panics if this is not representing a String value.
	StringValue() string

	// RawValue returns the underlying value represented by the node as an any. (Reflection will
	// be needed to figure out what it is).
	RawValue() any

	// Property returns the property value node for the given name or nil if property is not
	// a property of the object. Panics if this is not representing a JSON Object.
	Property(name string) Node

	// PropertyName returns the Node representing the property name.
	PropertyName(name string) Node

	// PropertyNames returns a slice of all set property values.
	// Panics if this is not representing a JSON Object.
	PropertyNames() []string

	// Label produces a label string with "[sourceName:line:col]" with conditional inclusion of
	// parts depending on if they are set or not.
	Label() string
}

type nodeImpl struct {
	offset int
	ctx    *context
}

func (*nodeImpl) IsObject() bool { return false }
func (*nodeImpl) IsArray() bool  { return false }
func (*nodeImpl) IsInt() bool    { return false }
func (*nodeImpl) IsFloat() bool  { return false }
func (*nodeImpl) IsBool() bool   { return false }
func (*nodeImpl) IsString() bool { return false }
func (*nodeImpl) IsNull() bool   { return false }
func (n *nodeImpl) Line() int    { return n.ctx.Line(n.offset) }
func (n *nodeImpl) Col() int     { return n.ctx.Col(n.offset) }

func (n *nodeImpl) RawValue() any {
	panic(catch.Error("Internal error: derived type should have implemented"))
}

func (n *nodeImpl) IntValue() int64 {
	panic(catch.Error("not an int value node"))
}
func (n *nodeImpl) FloatValue() float64 {
	panic(catch.Error("not a float value node"))
}
func (n *nodeImpl) BoolValue() bool {
	panic(catch.Error("not a bool value node"))
}
func (n *nodeImpl) StringValue() string {
	panic(catch.Error("not a string value node"))
}
func (n *nodeImpl) ArrayValue() []Node {
	panic(catch.Error("not an array value node"))
}
func (n *nodeImpl) ObjectValue() map[Node]Node {
	panic(catch.Error("not an object value node"))
}
func (n *nodeImpl) Property(_ string) Node {
	panic(catch.Error("not an object value node"))
}
func (n *nodeImpl) PropertyNames() []string {
	panic(catch.Error("not an object value node"))
}
func (n *nodeImpl) PropertyName(string) Node {
	panic(catch.Error("not an object value node"))
}

// Wrapper interface methods.

func (n *nodeImpl) Feature(_ string) xray.Wrapper     { return nil }
func (n *nodeImpl) FeatureName(_ string) xray.Wrapper { return nil }
func (n *nodeImpl) FeatureNames() []string            { return nil }
func (n *nodeImpl) Len() int                          { return -1 }
func (n *nodeImpl) FeatureAtIndex(_ int) xray.Wrapper { return nil }
func (n *nodeImpl) IsSlice() bool                     { return false }
func (n *nodeImpl) Kind() reflect.Kind                { return reflect.Invalid }
func (n *nodeImpl) Value(_ string) any                { return nil }
func (n *nodeImpl) Label() string                     { return xray.Label(n, n.ctx.sourceName) }
func (n *nodeImpl) HasCapitalizedFeatureNames() bool  { return false }

type nullNode struct {
	nodeImpl
}

func (n *nullNode) IsNull() bool  { return true }
func (n *nullNode) RawValue() any { return nil }

// NewNullNode returns a Node representing a JSON Null value.
func (ctx *context) newNullNode(offset int) Node {
	return &nullNode{nodeImpl: nodeImpl{offset: offset, ctx: ctx}}
}

type objNode struct {
	nodeImpl
	value   map[Node]Node
	propMap map[string]Node
	keyMap  map[string]Node
}

// NewObjNode returns a new Node representing an map[Node]Node (JSON Object) value.
func (ctx *context) newObjNode(obj map[Node]Node, offset int) Node {
	n := &objNode{
		value:    obj,
		propMap:  make(map[string]Node, len(obj)),
		keyMap:   make(map[string]Node, len(obj)),
		nodeImpl: nodeImpl{offset: offset, ctx: ctx},
	}
	for k, v := range obj {
		key := k.StringValue()
		n.propMap[key] = v
		n.keyMap[key] = k
	}
	return n
}
func (*objNode) IsObject() bool                      { return true }
func (n *objNode) ObjectValue() map[Node]Node        { return n.value }
func (n *objNode) Property(name string) Node         { return n.propMap[name] }
func (n *objNode) PropertyName(name string) Node     { return n.keyMap[name] }
func (n *objNode) PropertyNames() []string           { return utils.Keys(n.propMap) }
func (n *objNode) Feature(name string) xray.Wrapper  { return n.propMap[name] }
func (n *objNode) FeatureNames() []string            { return utils.Keys(n.propMap) }
func (n *objNode) Len() int                          { return len(n.propMap) }
func (n *objNode) Kind() reflect.Kind                { return reflect.Struct }
func (n *objNode) RawValue() any                     { return n.propMap }
func (n *objNode) FeatureName(s string) xray.Wrapper { return n.keyMap[s] }
func (n *objNode) Value(name string) any {
	if v, ok := n.propMap[name]; ok {
		return v.RawValue()
	}
	return nil
}

type arrayNode struct {
	nodeImpl
	children []Node
}

// NewArrayNode returns a new Node representing a slice of Node values.
func (ctx *context) newArrayNode(values []Node, offset int) Node {
	return &arrayNode{children: values, nodeImpl: nodeImpl{offset: offset, ctx: ctx}}
}
func (n *arrayNode) IsArray() bool                     { return true }
func (n *arrayNode) ArrayValue() []Node                { return n.children }
func (n *arrayNode) RawValue() any                     { return n.children }
func (n *arrayNode) IsSlice() bool                     { return true }
func (n *arrayNode) Len() int                          { return len(n.children) }
func (n *arrayNode) FeatureAtIndex(i int) xray.Wrapper { return n.children[i] }
func (n *arrayNode) Kind() reflect.Kind                { return reflect.Slice }

type intNode struct {
	nodeImpl
	value int64
}

func (n *intNode) IntValue() int64 { return n.value }
func (n *intNode) IsInt() bool     { return true }
func (n *intNode) RawValue() any   { return n.value }

// NewIntNode returns a new Node representing an int64 value.
func (ctx *context) newIntNode(value int64, offset int) Node {
	return &intNode{value: value, nodeImpl: nodeImpl{offset: offset, ctx: ctx}}
}

type floatNode struct {
	nodeImpl
	value float64
}

// NewFloatNode returns a new Node representing a float64 value.
func (ctx *context) newFloatNode(n float64, offset int) Node {
	return &floatNode{value: n, nodeImpl: nodeImpl{offset: offset, ctx: ctx}}
}
func (n *floatNode) FloatValue() float64 { return n.value }
func (n *floatNode) IsFloat() bool       { return true }
func (n *floatNode) RawValue() any       { return n.value }

type boolNode struct {
	nodeImpl
	value bool
}

// NewBoolNode returns a new Node representing a bool value.
func (ctx *context) newBoolNode(value bool, offset int) Node {
	return &boolNode{value: value, nodeImpl: nodeImpl{offset: offset, ctx: ctx}}
}
func (n *boolNode) BoolValue() bool { return n.value }
func (n *boolNode) IsBool() bool    { return true }
func (n *boolNode) RawValue() any   { return n.value }

type stringNode struct {
	nodeImpl
	value string
}

// NewStringNode returns a new Node representing a String value.
func (ctx *context) newStringNode(s string, offset int) Node {
	return &stringNode{value: s, nodeImpl: nodeImpl{offset: offset, ctx: ctx}}
}
func (n *stringNode) IsString() bool      { return true }
func (n *stringNode) StringValue() string { return n.value }
func (n *stringNode) RawValue() any       { return n.value }

type stackEntry struct {
	token  json.Token
	offset int
}

func (ctx *context) readAllTokens() error {
	for {
		// Position of next token start (if any). May refer to whitespace or elided delimiters.
		newOffset := int(ctx.decoder.InputOffset())
		// skip to first significant character.
		for i := newOffset; i < ctx.sLen; i++ {
			switch ctx.source[i] {
			// skip runes that are whitespace or elided "tokens" for : and ,
			case ':', ' ', '\t', '\n', ',':
				newOffset++
				continue
			}
			break
		}
		token, eof := ctx.decoder.Token()
		if errors.Is(eof, io.EOF) {
			return nil
		}
		if eof != nil {
			return fmt.Errorf(fmt.Sprintf("[%s:%d:%d] Invalid JSON: %s",
				ctx.sourceName, ctx.Line(newOffset), ctx.Col(newOffset),
				eof.Error()))
		}
		ctx.allTokens = append(ctx.allTokens, stackEntry{token: token, offset: newOffset})
	}
}

func (ctx *context) nextToken() (t json.Token, offset int) {
	if ctx.tokenIdx >= len(ctx.allTokens) {
		return nil, -1
	}
	entry := ctx.allTokens[ctx.tokenIdx]
	t = entry.token
	offset = entry.offset
	ctx.tokenIdx++
	return
}
func (ctx *context) peekToken() (t json.Token, offset int) {
	entry := ctx.allTokens[ctx.tokenIdx]
	return entry.token, entry.offset
}

func (ctx *context) parseObject() Node {
	result := map[Node]Node{}
	t, objOffset := ctx.nextToken()
	// Must start with '{'.
	if delim, ok := t.(json.Delim); !(ok && delim == '{') {
		panic(catch.Error("[%s:%d:%d] Expected a '{' ",
			ctx.sourceName, ctx.Line(objOffset), ctx.Col(objOffset)))
	}
	// until NextToken return } expect tuples of string key, value.
	for {
		t, offset := ctx.peekToken()
		if t == nil {
			if offset == -1 {
				panic(catch.Error("[%s:%d:%d] Expected a string key, got EOF",
					ctx.sourceName, ctx.Line(ctx.sLen), ctx.Col(ctx.sLen)))
			}
			panic(catch.Error("[%s:%d:%d] Expected a string key, got '%s'",
				ctx.sourceName, ctx.Line(offset), ctx.Col(offset), string(ctx.source[offset])))
		}
		switch tt := t.(type) {
		case json.Delim:
			_, _ = ctx.nextToken()
			if tt == '}' {
				return ctx.newObjNode(result, objOffset)
			}
		default:
			// ctx.pushBack(tt, offset)
		}
		key := ctx.parseString()
		value := ctx.parseValue()
		result[key] = value
	}
}
func (ctx *context) parseArray() Node {
	values := []Node{}
	t, objOffset := ctx.nextToken()
	// Must start with '['.
	if delim, ok := t.(json.Delim); !(ok && delim == '[') {
		panic(catch.Error("[%s:%d:%d] Expected a '[' ",
			ctx.sourceName, ctx.Line(objOffset), ctx.Col(objOffset)))
	}
	for {
		t, offset := ctx.peekToken()
		if t == nil {
			if offset == -1 {
				panic(catch.Error("[%s:%d:%d] Expected ',' and value or ']', got EOF",
					ctx.sourceName, ctx.Line(ctx.sLen), ctx.Col(ctx.sLen)))
			}
			got := string(ctx.source[offset])
			panic(catch.Error("[%s:%d:%d] Expected ',' and value or ']', got '%s'",
				ctx.sourceName, ctx.Line(offset), ctx.Col(offset), got))
		}
		switch val := t.(type) {
		case json.Delim:
			switch val {
			case ']':
				_, _ = ctx.nextToken()
				return ctx.newArrayNode(values, objOffset)
			case '}':
				panic(catch.Error("[%s:%d:%d] Expected ',' and value or ']', got '%s'",
					ctx.sourceName, ctx.Line(offset), ctx.Col(offset), val))

			default:
				values = append(values, ctx.parseValue())
			}
		default:
			values = append(values, ctx.parseValue())
		}
	}
}
func (ctx *context) parseString() Node {
	t, offset := ctx.nextToken()
	if t == nil || offset == -1 {
		panic(catch.Error("[%s:%d:%d] Expected a string key, got EOF",
			ctx.sourceName, ctx.Line(ctx.sLen), ctx.Col(ctx.sLen)))
	}
	switch s := t.(type) {
	case string:
		return ctx.newStringNode(s, offset)
	default:
		panic(catch.Error("[%s:%d:%d] Expected a string key, got: %s",
			ctx.sourceName, ctx.Line(offset), ctx.Col(offset), t))
	}
}

func (ctx *context) parseValue() Node {
	token, offset := ctx.peekToken()
	if token == nil && offset == -1 {
		panic(catch.Error("[%s:%d:%d] Expected an object, array or value, got EOF",
			ctx.sourceName, ctx.Line(ctx.sLen), ctx.Col(ctx.sLen)))
	}
	if token == nil {
		if ctx.acceptNull {
			return ctx.newNullNode(offset)
		}
		panic(catch.Error("[%s:%d:%d] JSON Null value is not accepted",
			ctx.sourceName, ctx.Line(offset), ctx.Col(offset)))
	}
	switch t := token.(type) {
	case json.Delim:
		switch t.String() {
		case "[":
			return ctx.parseArray()
		case "{":
			return ctx.parseObject()
		}
	case string:
		_, _ = ctx.nextToken() // throw away peeked token
		return ctx.newStringNode(t, offset)

	case bool:
		_, _ = ctx.nextToken() // throw away peeked token
		return ctx.newBoolNode(t, offset)

	case json.Number:
		_, _ = ctx.nextToken() // throw away peeked token
		if n, err := t.Int64(); err == nil {
			return ctx.newIntNode(n, offset)
		} else if n, err := t.Float64(); err == nil {
			return ctx.newFloatNode(n, offset)
		}
	default:
		fmt.Printf("Token %s, %v : %d\n", token, token, offset)
		panic(catch.Error("[%s:%d:%d] Expected an object, array or value, got %s",
			ctx.source, ctx.Line(offset), ctx.Col(offset), t))
	}
	return nil // placeholder
}

func (ctx *context) UnmarshalNode() (n Node, err error) {
	// error from readAllTokens means invalid JSON.
	err = ctx.readAllTokens()
	if err == nil {
		err = catch.Do(func() {
			n = ctx.parseValue()
		})
	}
	return n, err
}
