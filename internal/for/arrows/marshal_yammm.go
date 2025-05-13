package arrows

import (
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/validation"
	"github.com/wyrth-io/whit/internal/yammm"
)

type yammmCodec struct {
	ctx yammm.Context
}

// Requires:
// * Node has processed labels, removed the stereotype labels and set flags (isDataType, etc.)
// * Graph has an index of node id to Node
// * Properties of a type are obtained from Node.nProperties (where primary key and optionality has been resolved).

// buildYammmModel builds a Yammm model of the model expressed in an Arrows graph.
// A completed yammm.Context is returned if there were no Fatal or Error issues reported when
// preparing the arrows graph. If errors occur duing the completion of the Yammm context they
// will be reported in the IssueCollector and the caller should check for errors. In this case
// the Context will be returned, but may not be correctly completed.
// The name parameter is the name of the generated model.
func (g *Graph) buildYammmModel(name string, ic validation.IssueCollector) yammm.Context {
	// Prepare for output (make indexes, link nodes, validate, etc)
	g.prepareMeta(ic)
	if ic.HasFatal() || ic.HasErrors() {
		return nil
	}
	ctx := yammm.NewContext()
	yc := &yammmCodec{ctx: ctx}
	if err := ctx.SetMainModel(&yammm.Model{Name: name}); err != nil {
		panic("could not define a new yammm model") // should not happen
	}
	for _, n := range g.Nodes {
		if n.isDataType {
			if err := yc.processDataType(n); err != nil {
				ic.Collectf(validation.Error, err.Error())
			}
			continue
		}
		// All other nodes are Types
		if err := yc.processType(n); err != nil {
			ic.Collectf(validation.Error, err.Error())
		}
	}
	// Second pass processes all relations from "from nodes"
	for _, n := range g.Nodes {
		if n.isDataType {
			continue
		}
		if err := yc.processRelations(g, n); err != nil {
			ic.Collectf(validation.Error, err.Error())
		}
	}
	ctx.Complete(ic)
	if ic.HasFatal() || ic.HasErrors() {
		return nil
	}
	return ctx
}
func (y *yammmCodec) processDataType(n *Node) error {
	// A data type has a #DATATYPE label that needs to be filtered out.
	// All other labels are constraints to be joined with & (and).
	// For example the labels: "#DataType", "Integer", "0", "1"
	constraint := utils.Filter(n.Labels, func(s string) bool {
		return s != DATATYPE
	})
	// returned dt ignored since there are no line numbers of relevance in arrows data
	_, err := y.ctx.AddDataType(n.Caption, constraint)
	return err
}

func (y *yammmCodec) processType(n *Node) (err error) {
	// A data type has a #DATATYPE label that needs to be filtered out.
	// All other labels are constraints to be joined with & (and).
	typeName := n.Caption
	isAbstract := n.isAbstract
	isMixin := n.isMixin

	properties := []*yammm.Property{}
	for _, p := range n.nProperties {
		properties = append(properties,
			&yammm.Property{Name: p.Name,
				DataType:     []string{p.Type}, // TODO: arrows is just a string, how to specify the constraints?
				Optional:     p.optional,
				IsPrimaryKey: p.primary},
		)
	}
	switch {
	case isAbstract, isMixin:
		_, err = y.ctx.AddAbstractType(typeName, properties)
	default:
		_, err = y.ctx.AddType(typeName, properties)
	}
	if err != nil {
		return err
	}
	// "mixins" (all labels not being stereotypes are inherits)
	for _, m := range utils.Filter(n.Labels, func(s string) bool {
		return !(s == ABSTRACT)
	}) {
		if err = y.ctx.AddInherits(typeName, m); err != nil {
			return err
		}
	}
	return nil
}

func (y *yammmCodec) processRelations(g *Graph, n *Node) (err error) {
	typeName := n.Caption

	// Organize relationships that are from this node.
	inherits := []*Relationship{}
	compositions := []*Relationship{}
	associations := []*Relationship{}

	// TODO: This is not ideal as it iterates over all relationships once
	// per node. Should instead build an index and lookuo per node.
	for _, r := range utils.Filter(g.Relationships, func(r *Relationship) bool {
		return r.FromID == n.ID
	}) {
		switch {
		case r.isInherits:
			inherits = append(inherits, r)
		case r.isComposition:
			compositions = append(compositions, r)
		case r.isAssociation:
			associations = append(associations, r)
		}
	}

	for _, r := range inherits {
		toNodeName := g.nodeByID[r.ToID].Caption
		if err = y.ctx.AddInherits(typeName, toNodeName); err != nil {
			return err
		}
	}
	for _, r := range associations {
		toNodeName := g.nodeByID[r.ToID].Caption
		optional := (r.toMin.String() == "0")
		many := (r.toMax.String() == "M")
		rProperties := []*yammm.Property{}
		for _, prop := range r.nProperties {
			rProperties = append(rProperties,
				&yammm.Property{
					Name:         prop.Name,
					DataType:     []string{prop.Type}, // TODO: How to set constraints the right way
					Optional:     prop.optional,
					IsPrimaryKey: prop.primary,
				})
		}
		if _, err = y.ctx.AddAssociation(typeName, r.Type, toNodeName, optional, many, rProperties, ""); err != nil {
			return err
		}
	}
	for _, r := range compositions {
		toNodeName := g.nodeByID[r.ToID].Caption
		optional := (r.toMin.String() == "0")
		many := (r.toMax.String() == "M")
		if _, err = y.ctx.AddComposition(typeName, r.Type, toNodeName, optional, many, ""); err != nil {
			return err
		}
	}
	return nil
}
