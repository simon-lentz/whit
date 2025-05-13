/*
Package cypher contains logic to generate Cypher from a yammm modeled instance (either a map[string]any, or
a generate go implementation instance that represents the top level document).

To generate Cypher to create data in a Neo4j call the function `Generate(graph any, meta yammm.Meta) []Statement`
with the loaded Graph (from a generated package) and its coresponding yammm.Meta interface. The produced
`Statement` slices contains a sequence of Cypher expressions with parameters in a map. These can be
iterated over and calls can be made to Neo4j in a transaction with the source and parameters.

	// Example go code for a transaction for a single statement
	greeting, err := session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		result, err := transaction.Run(ctx, statement.Source, statement.Parameters)
	})

The generated source for types looks like below. It merges on the primary key(s), and then either sets
or merges the given property values.

	MERGE (p:Person {uid: $primaryKeys.uid})
	ON CREATE SET p = $props
	ON MATCH  SET p += $props

If an instance has more than one label the generated source will also set all of the labels.

	MERGE (p:Person:Entity {uid: $primaryKeys.uid})
	ON CREATE
	  SET p = $props,
	ON MATCH
	  SET p += props

For compositions the part gets a separate node, and a link is generated between composition/part using their primary keys, with
the link type derived from the property name and type. These links do not have properties.

For associations the link is generated between the from/to types using their primary keys. The association may have
properties.

TODO: Generated cypher for associations and compositions are probably not idempotent and has to be changed to use some kind of
MERGE to avoid creating duplicate links.
*/
package cypher

// Statement describes a source string in Cypher syntax that when executed should be
// executed with the given map of parameters.
type Statement struct {
	// Source is the statement with references to $variables.
	Source string `json:"source,omitempty"`
	// Parameters is a map of variable name to value.
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	// NeedsSepTx is set for statements that needs to be executed in a separate transaction.
	NeedsSepTx bool `json:"sepTx,omitempty"`
}

const (
	// UID is the name of the id property in Neo4j since it has a built in numeric id.
	UID = "uid"
)
