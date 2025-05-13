This is a document exploring a concrete syntax for YAMMM

// Define the schema's name
schema "YammmExploration"
// Define the type/struct/entity, give its name (and optionally plural form - generated automatically), its "super type" (extend) and
// any mixins.

/* Documentation for Car
*/
type Car, Cars extend Vehicle mixin CommonPartA, CommonPartB {
    /* documentation for regNbr */
    regNbr RegNbr                 // property with custom data type
    color ? string                // optional property
    serialNbr + string            // required property (same as no +)

    --> MADE_BY(01) / MADE(0M) Manufacturer   // association with name of association / reverse name (reverse not yet implemented in whit)
    *-> HAS Engine 11 {               // composition
      installed date              // relationship with property
    }
    o-> HAS 0M Example            // Aggregation (no ownership), not implemented in whit
}
abstract type Vehicle {}
type Manufacturer { }
type Engine { }

mixin CommonPartA {
    id string
}
mixin CommonPartB {
    comment string
}
datatype RegNbr = Pattern["[A-Z]{3}[0-9]{3}"]

// Idea for optional, required, etc.
// Using "puppet file system" data types optionality can be controlled by data type - i.e. Optional[String] means it can be missing/null, but is
// otherwise a string - which may mean it comes back as an empty String. Using Optionak[String[1]] means it can be missing, but if it is a string
// it must have a minimum length of 1.
// Primary key can be denoted by +