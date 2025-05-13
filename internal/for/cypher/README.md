# The Cypher package

The `cypher` package produces `Statement` objects. It can produce initialization statements from a Yammm model to set up
indexes and constraints in a Neo4j store (or store compatible with Cypher as of Neo4j 5.11).
It can also produce `Statement` object for create/merge of data from data represented by Go code generated from a model
in Yammm where the data is unmarshaled from json or yaml or constructed programatically.

The cypher package's API is simple.

```go
// to get statements for an instance graph:
statements := cypher.NewMergeGenerator().Process(ctx, graph)
// to generate statements to initialize a Neo4j DB
cypher.GenerateInit(ctx yammm.Context) []Statement
```

The returned statements look like this:
```
type Statement struct {
	Source     string                 `json:"source,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}
```
And fits well with the Neo4j API for go where the two elements are given as arguments.

The `GenerateInit` only requires a Context with a loaded meta model. Such a model
is best read from a DSL Yammm file, but can also be read from a json file (an earlier Marshaled Yammm model), or loaded
from a Yammm generated Go package (as it includes the Marshaled Json as a blob).

```
	ctx := yammm.NewContext()
	if err := ctx.SetModelFromJSON(strings.NewReader(generated.SerializedModel)); err != nil {
        // handle could not load error
    }
	ic := validation.NewIssueCollector()
	if !ctx.Complete(ic) {
        // handle errors reported in the ic
    }
	statements := cypher.GenerateInit(ctx)
```

The `MergeGenerator` requires an unmarshaled (or programatically
created) instance of the model. In a generated package there is always a top type named `Graph` and the API
expects such a graph compatible with the yammmm model loaded into the yammm.Meta Context.

```
	ctx := yammm.NewContext()
	if err := ctx.SetModelFromJSON(strings.NewReader(generated.SerializedModel)); err != nil {
        // handle could not load error
    }
	ic := validation.NewIssueCollector()
	if !ctx.Complete(ic) {
        // handle errors reported in the ic
    }
	statements := cypher.NewMergeGenerator().Process(ctx, g)
```

# General Notes Neo4j
Enterprise version is required to place restrictions on node properties (required, etc.)
Properties do not have a datatype, but values are of different type. Typecasting is required !! (Awful!!)
May get an error when trying to create an index ??
There is no schema in neo4j!! Labels come into existence when they are set in nodes oy oner when constraints are defined
for them.
Neo4j 5.11 has multiple databases served by one server. It also supports data type constraints on properties.

