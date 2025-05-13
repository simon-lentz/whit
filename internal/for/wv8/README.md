# Yammm Weaviate support

Notes about weaviate support.

## Classes
Weaviate does not support inheritance. All inherited traits are therefore flattened.

## Properties
Property names allow for initial UC letter or _ but is generally in LC.
Documentation vaguely says that names are treated the same - which could possibly mean they are
case independant.

In any case we should make properties and relationships case independently unique in Yammm (easiest way to ensure they won't clash for some technology).

The property `id` is reserved and must be a valid UUID in string form.

## DataTypes
Weaviate has a `date` data type that is RFC3339, no other formats are supported. This means we have to
map instances that comply with some other format to RFC3339. Our Date type will need to pad a time (midnight).
Padding will need to take place for other formats not including all of 3339.

Objects cannot be embedded, must be a relation to another class.

## Relationships
Relationships are links and cannot have properties. This means that an intermediate class is needed when
the relationship has properties.

Relationships are one way but can be made two way by adding a back-reference property in the target class.

References are always possibly many.

# Actions to take

* Make all traits be validated to be case independently unique
* Make reverse relations work - validate reverse name uniqueness from the to-type's POV.

# Gendata
The model supports the following under the gendata key `"wv8"`:

For Type/Class:
* `"InvertedIndexConfig"` should be the JSON data for a Weaviate `models.InvertedIndexConfig`
* `"ModuleConfig"` JSON data for modules installed into Weaviate - documented as being `interface{}`.
* `"MultiTenancyConfig"` should be JSON data for Weaviate `models.MultiTenancyConfig`.
* `"ReplicationConfig"` should be JSON data for Weaviate `models.ReplicationConfig`.
* `"ShardingConfig"` should be JSON data - - documented as being `interface{}`.
* `"VectorIndexConfig"` should be JSON data - - documented as being `interface{}`.
* `"VectorIndexType"` should be a string in JSON format.
* `"Vectorizer"` should be a string in JSON format.

