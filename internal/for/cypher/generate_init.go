package cypher

import (
	"fmt"
	"strings"

	"github.com/tada/catch"
	"github.com/wyrth-io/whit/internal/tc"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/yammm"
)

// GenerateInit generates Neo4j Cypher source from a Yammm modeled Graph. It returns a slice of Statement;
// a combination of a cypher source string and parameters map for the source.
// The generated cypher code is for initialization of a Neo4j store and consists of create statements
// for indexes and constraints. Panics if given context is not completed.
func GenerateInit(ctx yammm.Context) (result []Statement) {
	if !ctx.IsCompleted() {
		panic("context not completed")
	}
	// For all types - get their primary keys, and required properties.
	// PKs are unique IS NODE KEY (once for 'id', and once in combination for the rest)
	// Required are constrained by IS NOT NULL.

	nameOfProp := func(prop *yammm.Property) string { return prop.Name }
	requiredPredicate := func(prop *yammm.Property) bool { return !(prop.Optional || prop.IsPrimaryKey) } // TODO: and not Primary
	ctx.EachType(func(t *yammm.Type) {
		if !(t.IsAbstract) {
			// PK constraints, one for id and one for the rest of the Pks.
			result = append(result, generatePKeysConstraints(t.Name, utils.Map(t.AllPrimaryKeys(), nameOfProp))...)
			// All required, that are not also PKs
			requiredProps := utils.Filter(t.AllProperties(), requiredPredicate)
			result = append(result, generateRequiredConstrainta(t.Name, utils.Map(requiredProps, nameOfProp))...)
			result = append(result, generateDataTypeConstraints(t.Name, t.AllProperties())...)
			result = append(result, generatePropertyIndexes(t.Name, t.AllProperties())...)
		}
	})
	return result
}

// generatePKeysConstraints returns cypher Statements producing a unique constraint for a list of property names.
// The `id` property name is always unique on its own. Other primary keys must be unique in combination.
// The returned slice will have one or two entries depending on if the type has other primary keys than id.
func generatePKeysConstraints(label string, propnames []string) []Statement {
	statements := []Statement{}
	pksNoID := utils.Filter(propnames, func(p string) bool { return p != "id" })
	if len(pksNoID) > 0 {
		// Produce comma separated list of props prefixed with ".n"
		props := strings.Join(utils.Map(pksNoID, func(s string) string {
			return fmt.Sprintf("n.%s", s)
		}), ", ")
		constraintName := fmt.Sprintf("%s_primary_keys", label)
		statements = append(statements, Statement{
			Source: fmt.Sprintf(`CREATE CONSTRAINT %s IF NOT EXISTS FOR (n:%s) REQUIRE (%s) IS NODE KEY`,
				constraintName, label, props),
			Parameters: map[string]any{},
		})
	}
	constraintName := fmt.Sprintf("%s_primary_uid", label)
	statements = append(statements, Statement{
		Source: fmt.Sprintf(`CREATE CONSTRAINT %s IF NOT EXISTS FOR (n:%s) REQUIRE (n.uid) IS NODE KEY`,
			constraintName, label),
		Parameters: map[string]any{},
	})
	return statements
}

// generateRequiredConstrainta returns cypher Statements producing a unique constraint for a list of property names.
// One statement is needed per constraint since IS_NOT_NULL is per property!
func generateRequiredConstrainta(label string, propnames []string) []Statement {
	statements := []Statement{}
	for i := range propnames {
		p := propnames[i]
		if p == "id" {
			p = UID // neo4j already has "id" as a special autonumeric id
		}
		constraintName := fmt.Sprintf("%s_required_property_%s", label, p)
		statements = append(statements, Statement{
			Source: fmt.Sprintf(`CREATE CONSTRAINT %s IF NOT EXISTS FOR (n:%s) REQUIRE (n.%s) IS NOT NULL`,
				constraintName,
				label,
				p),
			Parameters: map[string]any{},
		})
	}
	return statements
}

// generateDataTypeConstraints produces one cypher data type constraint per property.
// This works for integer, float, string, and boolean (which are in upper case in Cypher).
// The date and time types are difficult to map exactly as yammm supports different formats
// where cypher dictates one format. For date/time yammm will map to string.
func generateDataTypeConstraints(tname string, props []*yammm.Property) []Statement {
	// cname, label, propname, datatypename
	template := `CREATE CONSTRAINT %s IF NOT EXISTS FOR (n:%s) REQUIRE n.%s IS :: %s`
	statements := make([]Statement, len(props))
	for i := range props {
		cypherTypeName := BaseTypeName(props[i])
		propName := props[i].Name
		if propName == "id" {
			propName = "uid"
		}
		constraintName := fmt.Sprintf("%s_%s_dt", tname, propName)
		statements[i].Source = fmt.Sprintf(template, constraintName, tname, propName, cypherTypeName)
	}
	return statements
}

// generatePropertyIndexes produces one cypher statement per property that needs
// an index. (Currently only Spacevector properties).
func generatePropertyIndexes(tname string, props []*yammm.Property) []Statement {
	template := `CALL db.index.vector.createNodeIndex($idxName, $nodeLabel, $propertyName, $dim, "cosine")`
	statements := make([]Statement, 0, len(props))
	for i := range props {
		p := props[i]
		bt := p.BaseType()
		if bt.Kind() != tc.SpacevectorKind {
			continue
		}
		propName := props[i].Name
		if propName == "id" {
			propName = "uid"
		}
		idxName := fmt.Sprintf("%s_%s_idx", tname, propName)
		statements = append(statements, Statement{
			Source: template,
			Parameters: map[string]any{
				"idxName":      idxName,
				"nodeLabel":    tname,
				"propertyName": p.Name,
				"dim":          bt.Dim(),
			},
			NeedsSepTx: true,
		})
	}
	return statements
}

// BaseTypeName returns the Cypher data type for a Property's yammm.BaseType.
func BaseTypeName(p *yammm.Property) string {
	x := p.TypeChecker().BaseType()
	switch x.Kind() {
	case tc.IntKind:
		return "INTEGER"
	case tc.FloatKind:
		return "FLOAT"
	case tc.BoolKind:
		return "BOOLEAN"
	case tc.StringKind:
		return "STRING"
	case tc.SpacevectorKind:
		return "LIST<FLOAT NOT NULL>"
	default:
		panic(catch.Error("internal error: cypher schema init does not handle this base type: %s",
			tc.GoName(x)))
	}
}
