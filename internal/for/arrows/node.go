package arrows

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/validation"
)

// Node represents an Arrows Node.
// Internal to the package it holds a state that is built up by this package in order to validate the
// contents. At present the private members are designed for arrows graphs that represent Wyrth
// meta-meta model. Significant changes are expected if this package should also be able to
// handle the normal intance graphs.
type Node struct {
	ID             string                 `json:"id"`
	Position       Position               `json:"position"`
	Caption        string                 `json:"caption"`
	Style          map[string]interface{} `json:"style"` // string key and mix of string and int value
	Labels         []string               `json:"labels"`
	Properties     map[string]string      `json:"properties"`
	nProperties    []*Property            // normalized properties
	inEdges        []*Relationship
	outEdges       []*Relationship
	superTypes     *utils.Set[string]
	primaryKeys    []*Property
	isAbstract     bool
	isMixin        bool
	isDataType     bool
	allProperties  map[string]*Property          // all inherited and possibly mixed in properties
	possibleMixins *utils.Set[string]            // transitively collected mixins
	relIndex       map[string]*utils.Set[string] // rel type to []node caption
	captionPlural  string                        // helper for where the type name is needed in plural
}

// Position represents the position in x,y coordinates in the arrows graph json data.
// It is needed for unmarshalling, but is otherwise unused.
type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Property represents an Arrows Property - in Nodes and Relationships.
type Property struct {
	Name     string
	Type     string
	primary  bool
	optional bool
}

// normalize normalizes the node such that the caption is set from the fist label if caption is empty.
// If caption is duplicated as a label, this label is dropped from the list of labels.
// Caption and labels are whitespace trimmed.
// Comparison is case independent.
// The caption and the labels are forced to have initial caps.
func (n *Node) normalize(ic validation.IssueCollector, forMeta bool) {
	n.normalizeCaptionAndLabels(ic, forMeta)
	n.normalizeProperties(ic, forMeta)
}

func (n *Node) normalizeCaptionAndLabels(ic validation.IssueCollector, forMeta bool) { //nolint:gocognit
	// Set booleans for meta #Abstract, #Mixin, and  #DataType labels.
	for _, label := range n.Labels {
		switch label {
		case ABSTRACT:
			n.isAbstract = true
		case MIXIN:
			n.isMixin = true // This is kept although missing in YAMM - it is changed to inheritance
		case DATATYPE:
			n.isDataType = true
		default:
			if len(label) > 0 && label[0] == '#' {
				ic.Collectf(validation.Error, "unknown # marked label - '%s' is not allowed")
			}
		}
	}
	// Error if meta labels are used in instane model.
	if !forMeta {
		if n.isAbstract {
			ic.Collectf(validation.Error, "Node '%s' s labeled #Abstract - not allowed in instance model")
		}
		if n.isMixin {
			ic.Collectf(validation.Error, "Node '%s' s labeled #Mixin - not allowed in instance model")
		}
		if n.isDataType {
			ic.Collectf(validation.Error, "Node '%s' s labeled #DataType - not allowed in instance model")
		}
	}

	if forMeta {
		// In instance models, caption may be some other Property - for example the name of a person instead
		// of "Person". Removing spaces is not meaningful for that reason.
		n.Caption = strings.ReplaceAll(n.Caption, " ", "")
	}

	if forMeta && !n.isDataType {
		// Skip label processing for #DataType since labels are cue constraints
		for i, label := range n.Labels {
			n.Labels[i] = strings.ReplaceAll(label, " ", "")
		}
		if len(n.Caption) < 1 && len(n.Labels) > 0 {
			// use label 0 as caption and drop label 0
			n.Caption = n.Labels[0]
			if forMeta {
				// This changes the semantics in an instance model since the caption is purely presentation.
				// The typoe must be present among the labels.
				n.Labels = n.Labels[1:]
			}
		}
	}
	// Capitalize caption (for instance, does not matter)
	if forMeta {
		if n.isDataType && strings.HasPrefix(n.Caption, "#") {
			ic.Collectf(validation.Error, "node with ID '%s' has caption '%s' - cannot start with '#'", n.ID, n.Caption)
			if n.Caption == "#DataType" {
				ic.Collectf(validation.Info, "hint: #DataType should be a label, and caption the name of the data type")
			}
		}
		if n.isDataType {
			// Must have first char be a letter in lower case
			var first rune
			for _, c := range n.Caption {
				first = c
				break
			}
			if !(unicode.IsLower(first) && unicode.IsLetter(first)) {
				ic.Collectf(validation.Error, "node with ID '%s' for DataType '%s' - must start with lower case letter", n.ID, n.Caption)
			}
			switch n.Caption {
			case "null",
				"string",
				"number",
				"int",
				"float",
				"bool",
				"struct",
				"list",
				"bytes":
				{
					ic.Collectf(validation.Error, "node with ID '%s' for DataType '%s' - duplicates build in data type", n.ID, n.Caption)
				}
			case
				"uint",
				"uint8",
				"int8",
				"uint16",
				"int16",
				"rune",
				"uint32",
				"int32",
				"uint64",
				"int64",
				"int128",
				"uint128",
				"positive",
				"byte",
				"word":
				{
					ic.Collectf(validation.Error, "node with ID '%s' for DataType '%s' - duplicates build in numeric bounds data type", n.ID, n.Caption) //nolint:lll
				}
			}
		}
		if !n.isDataType {
			// for data types the caption is verbatim, usuall a lower case
			n.Caption = utils.Capitalize(n.Caption)
		}
	}

	// Make labels unique (case independent)
	n.Labels = utils.UniqueFold(n.Labels)

	// Drop all empty labels
	n.Labels = utils.DeleteFold(n.Labels, "")

	// If caption exists in labels (case insensitive) then label is dropped.
	// In instance model this would change semantics, so skip.
	if forMeta {
		n.Labels = utils.DeleteFold(n.Labels, n.Caption)
	}

	// All remaining labels are capitalized unless it is a data type where labels are cue constraints
	if !n.isDataType {
		n.Labels = utils.CapitalizeAll(n.Labels)
	}
	n.captionPlural = thePluralizer.Plural(n.Caption)
}

func (n *Node) normalizeProperties(ic validation.IssueCollector, forMeta bool) {
	if n.isDataType && len(n.Properties) != 0 {
		ic.Collectf(validation.Error, "node: '%s' is a DataType and cannot have properties", n.Caption)
	}
	np := make([]*Property, 0, len(n.Properties))
	for key, typeString := range n.Properties {
		nname, primary, optional := normalizePropertyName(key)
		if nname == "" {
			ic.Collectf(validation.Error, "node '%s' property '%s' - illegal property name\n", n.Caption, key)
		}
		if forMeta && primary && optional {
			ic.Collectf(validation.Error, "node '%s' property '%s' cannot be both primary (+) and optional (?)\n", n.Caption, key)
		}
		if !forMeta && (primary || optional) {
			ic.Collectf(validation.Error, "node '%s' property '%s' has illegal ending of '+ or '?'\n", n.Caption, key)
		}

		np = append(np, &Property{
			Name:     nname,
			Type:     typeString,
			primary:  primary,
			optional: optional,
		})
		n.nProperties = np
	}
}

// normalizePropertyName validates the name against the pattern ^([a-z][a-zA-Z0-9_]+)([?+]?)$ and
// returns an empty name if it does not match. The trailing '+' or '?' are return as primary and optional
// flags.
func normalizePropertyName(name string) (nname string, primary bool, optional bool) {
	namePattern := regexp.MustCompile(`^([a-z][a-zA-Z0-9_]+)([?+]?)$`)
	nname = strings.TrimSpace(name)
	nname = utils.DeCapitalize(nname)
	matches := namePattern.FindStringSubmatch(nname)
	if len(matches) > 0 {
		nname = matches[1]
		if matches[0] == "" {
			nname = ""
		}
		if matches[2] == "+" {
			primary = true
		}
		if matches[2] == "?" {
			optional = true
		}
	} else {
		nname = ""
	}
	return
}

// HasIncomingAssociations returns true if there are incoming relationships.
func (n *Node) HasIncomingAssociations() bool {
	for _, r := range n.inEdges {
		if r.isAssociation {
			return true
		}
	}
	return false
}
