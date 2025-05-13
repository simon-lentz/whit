# Yammm Concrete Representation of Meta Model

The Yammm concrete representation of a meta-model (schema) is a DSL for expressing a Yammm model.

## Comments
Comments are written using `//` and the comment is the rest of the line.

Documentation comments are written enclosed in `/*` and `*/` and
this comment must be placed immediately before the definition it is documenting. (Although white space and regular comments are ignored it is a bad idea to place a comment between the doc comment and its definition).

## White Space
The language is not white space sensitive - white space is simply ignored when parsing.

These two examples are equivalent in meaning:
```
type X { a Integer}
type X {
    a
    Integer
}
```

## Schema name
A Yammm file must define the name of the model - like this:
```
schema "myFirstSchema"
```
This must come before any other definition.

## Definitions
The Yammm language allows the definition of:
* Types - a struct/class having properties, association and composition relationships to other types and supports inheritance of other types. Types also support invariants; a constraint expression that operates across
properties and relationships.
* Properties of a type are defined with a data type.
* Associations between types can have properties.
* Part types can be parts in a composition of types.
* Data type aliases can be defined.

## Type
A type is defined like this:
```
type MyType { }
```
The type can be abstract (meaning that there are no instances of this type, but there can be instances of derived non-abstract types).

```
abstract type MyAbstractType { }
```
A type can extend other types (abstract or non-abstract):
```
type MyType extends MyAbstractType {}
```

A type can extend multiple other types:
```
type MyType extends ThisType, ThatType {}
```
When a type extends another it will receive all the properties and relationships defined for the type it is extending.
There is no override of a type's traits (properties, relationships) - each trait must have a unique name. If the
same name is used in a chain of extensions it must also have the exact same data type or it will be flagged as an error.


The name of a type must start with an uppercase letter followed by any ASCII letter, digit, or underscore character.

## Properties
Types (and Associations) have properties. A property is defined by its name, if it is a primary key, and if it is required. By default instances of the type are not required to have a value for a property.
A property has a data type.
A property name must start with a lowercase letter followed by any ASCII letters or digits or the underscore character.

```
type MyType {
    name String
}
```
A property can be required:
```
type MyType {
    name String required
}
```
A property can be a primary key - either on its own or in combination with other properties. If a property is a primary key it is also by definition required (which is then not specified).
A primary key (combination) must be unique among all instances of that type.
```
type MyType {
    name String primary
}
```
```
type Car {
    regNbr String primary
    country String primary
}
```
With multiple primary keys - the combination must be unique among all instances of the type.

### ID and primary keys
All types have a built-in property `id UUID primary`. When there are additional primary keys the `id` will
automatically be set to a deterministic v5 UUID based on the package, type name, and the instance values of the primary keys of the type. 
For types that have no other primary keys besides `id`, the UUID must be set in instances. For types with a deterministically set UUID the `id` property is automatically set and cannot be modified.

A UUID can be set in instances using a full v4 UUID in string form, or a short (base57 encoded) UUID, or one of the special forms supported by Yammm. A UUID of `$$<localid>` (where `<localid>` is any string) will be replaced when an instance is processed. This can be done in batch for an entire data file or done when a data file is processed to be stored in a DB. When processed in batch, the `$$localid` references are replaced with an id on the format `$$:<shortuuid>:<localid>` where the shortuuid is a base57 encoded random UUID. This scheme allows a batch update of a file to be followed by additions to the same data file where only the localid is used in new/additional references. Subsequent updates (either in batch or when processing the data for storing in a DB) will resolve the local ids against any already supplied shortuuids. 

Once the data has been stored in the DB the `id` will not retain the `<localid>` part.

Since a DB technology may already have `id` as reserved name this property name may be mapped to something else for that
technology. For example `id` is mapped to `uid` in the schema generated for `Neo4j` and operations will map `id` property in
instance data to `uid` when inserting or updating data. In manual queries against the DB `uid` needs to be used.

## Associations
Associations from one instance of a type to another are entered in the body of the "from type". At present it is not defined if
the reverse direction is available or not as that depends on the technology used when realizing the model.
In Neo4J for example, all relationships are traversable in both directions using the same name for the association.

Yammm language allows for the specification of the reverse relationship, but this is currently not implemented throughout.

```
type Person {
    --> OWNS Car
}
```
This example defines an association named OWNS specifying that a Person optionally OWNS one car. 
If we instead wanted a person to be able to own more than one car we use the multiplicity `many`.

```
type Person {
    --> OWNS (many) Car // owns 0 or more cars
}
```
Here is a table showing the allowed multiplicities and their meaning in terms of min and max occurrences of the "to type" in the
relationship.

| mutiplicity | min | max |
| ----------- | --- | --- |
|  | 0 | 1 |
| (one) | 0 | 1 |
| (_:one) | 0 | 1 |
| (one:one) | 1 | 1 |
| (many)    | 0 | many |
| (_:many)    | 0 | many |
| (one:many)| 1 | many |

### Reverse name and multiplicity
This is not yet implemented in the model but can be specified.

```
type Person {
    --> OWNS (many) Car / OWNED_BY (many)
}
```
This defines that a person can own 0 or more Car instances and that a Car may be owned by 0 or more people.

An association adds a property to the type with the name of the association concatenated with the name of the "to type". For example `OWNS_Car`, or `OWNS_Cars` (if the relationship is of "many" multiplicity).
### Association properties
An Association can have properties - they are defined the same way as properties for a type, but do not support
the notion of a "primary key". (Note: this may need to change if there is a need to define an index
for association properties in Neo4J).

```
type Person {
    --> OWNS (many) Car { since Date }
}
```
This defines the property `since` having a `Date` data type for the association.

## Composition
A Composition is like an Association but it also defines that the life cycle of a part in the composition is defined
by the composing type. In other words, when the composition is deleted so are all its parts.

```
type Car {
    *-> HAS (one:one) Chassi 
    *-> HAS Engine
    *-> HAS Transmission
    *-> HAS (many) Wheel
}
```
This defines that a `Car` is composed of a `Chassis` (must have one), optionally one `Engine` and optionally one `Transmission`. The names for properties in the type (and in a DB) will be the name of the association followed by an underscore and the name of the type - for example `HAS_Chassi`. When the multiplicity is a max of `many` the
type name will be pluralized - for example, `HAS_Wheels`.

The same type can be involved in several compositions:
Here is an example.
```
type Car {
    *-> LEFT_FRONT (one:one) Wheel
    *-> RIGHT_FRONT (one:one) Wheel
}
```
## Datatypes
Yamm supports the following data types: `String`, `Integer`, `Float`, `Boolean`, `Timestamp`, `Date`, `Pattern`,  `Enum`, `UUID`, and `Spacevector`.

Some of these types can be further detailed with arguments given within brackets - for example, `Integer[0,100]` specifies the range of the integer value to be >= 0 and <= 100.

User-defined data types come in the form of an "alias". In the DSL an alias is formed by a statement on the form:
```
type <Alias> = <built in data type>
```
For example to define the data type `Color`:
```
type Color = Enum["red", "green", "blue"]
```
would allow the use of the data type `Color` instead of repeating the aliased `Enum["red", "blue", "green"]` wherever that specific enum type is needed.

The data type alias cannot refer to another alias. (And can thus not cause infinite reqursion).

### Integer
Represents an integer value in the full range of possible integer values of the underlying technology (64 bits signed).

An `Integer` can be given a min/max range (inclusive) by specifying the min and max values `Integer[min, max]`. The range values are integers or the special value `_` that denotes the smallest possible min value or largest possible max value. For example, `Integer[0,_]` means any non negative integer.

### Float
Represents a floating point value, optionally with a given range just like for `Integer`, but the min/max values can be given as floating point or integer values.

### String
Represents a text value. By default of any length (including the empty string). A String can specify the number of min and max characters allowed in the String. The `_` can be used to represent the smallest possible or largest possible value.
For example, `String[1,_]` defines that the string must have a minimum length of 1. `String[_,80]` that the string can be between 0 and 80 characters long.

### Boolean
Represent a true/false value. It does not accept a range.

### Timestamp
Represents a timestamp in string form (the underlying storage type is a String). The format of the strings is determined by the data type. The format RFC 3339 is used by default. To specify the format, the "go lang time format string" is given. For example, `Timestamp["2006-01-02"]` specifies that the timestamp must be a date in Year, Month, and Date order with hyphens between the three.

Note that while the underlying type is a String, a driver for a particular type of database may translate this
into an actual time data type.
### Date
Represents a date in string form (the underlying storage type is a string). A date is always expressed in the format "2006-01-02". Other date formats can be used by defining a custom Timestamp. For example:
```
type USDate = Timestamp["01-02-2006"]
```

Note that while the underlying type is a String, a driver for a particular type of database may translate this
into an actual time data type.
### Enum
Represents a string value that must match one of the literals specified in the type.
```
Enum["red", "green", "blue"]
```
Specifies that the string must be one of "red", "green" or "blue".

### Pattern
Represents a string value that matches a regular expression value.
The regexp syntax is that of Golang. One or more patterns can be specified, and the instance value must match at least one of them.
```
Pattern["^[A-Z]{3,3}$", "^[0-9]{3,3}$"]
```
In this example, the string must either be three upper case letters A-Z or three digits 0-9. 

### UUID
Represents a string defining a UUID value in string format. The formats accepted are:
* A local reference on the form `$$<localid>` where `<localid>` is any string. For example `$$1`, `$$2`,`$$theCar`, etc.
* A resolved local reference on the form `$$:<shortuuid>:<localid>` where the `<shortuuid>` is a base57 encoded UUID resolution used for all occurences of the associated `<localid<`.
* A shortuuid string (22 base57 characters).
* A standard UUID string (36 hex characters including "-" separator between groups).

Note that the built in `id UUID primary` is automatically assigned a deterministic (SHA1) v5 UUID based on the other primary keys if such exists, and may then not be set by instances. For types without any natural primary keys (where `id` is the only primary key) the `id` must be set.

### Spavevector
Represents a vector in a multi-dimensional space. The number of dimensions must be specified within brackets. The typical use of this is to store vectorized text.

# Invariants
While properties on their own can specify the particular data type of that property, it cannot on its own specify constraints based on other properties. As an example, you may want to restrict that a couple of properties should all be filled out, or none of them, or you may want to restrict that the "ownedSince" data of a car ownership needs to on or after the manufacturing date of the car being owned. Such constraints are what
invariants are for.

The yammm DSL invariant expressions are (somewhat) inspired by UML's OCL (Object Constraint Language).

An invariant consists of two things; (i) a string that describes the constraint, which is used in an error message should validation of an instance not meet the constraints, and (ii) a boolean expression that if it evaluates to false will fail the validation.

An invariant can be placed anywhere in a type definition and it starts with an exclamation mark `!`. For example:
```
type Example {
    a Integer
    b Integer
    c Integer

    ! "a must be greater than the sum of b and c"
      a > b + c
}
```
## Literal Values
Literal values are integers, floating point numbers, strings, booleans, regular expressions,
lists/arrays of literal values, and nil/undef. For example:
```
1              // integer
1.2            // float
"hello"        // string
true           // boolean
false          // boolean
/[a-z]+/       // regular expression
[1,2,"hello"]  // array 
_              // nil/undefined value
```
Arrays can contain expressions. For example:
```
[1+2, 2+3]     // results in [3, 5]
```

### Undefined/Nil values
A "nil" value represents "no/missing value" or "undefined". It is the value of non-required
properties that have no value set. It can be compared against other values and is only equal to itself.
When used in boolean operations, nil acts as  boolean false. 
A literal nil is obtained by a single `_`.

For example:
```
_ == _     // true, nil is nil
_ < false  // true. See Comparisons why that is.
_ == false // false, only nil is equal to nil
!_         // true, since nil is considered boolean false
```

## Variables
There are no freely assignable variables, but variables come into play as function parameters.
See "Functions" below.
Variables start with a `$` dollar sign. Built in parameters are named `0`-`n` and user-defined
parameters (to use instead of numbered parameters) cannot start with a digit. For example,
you may see something like:
```
$0 + $1 // Add the two parameters
```
All properties and relationships are available via their respective name (and are thus a kind of variable).
For example:
```
type Example {
    a Integer
    b Integer
    c Integer
    ! "Sum of a, b, and c cannot be larger than 10"
      a + b + c <= 10
}
```
In some cases, for example, when you want to perform more than one operation on the outcome of
another you can use the `with` function as that introduces a new variable. For example, if you have the
same (somewhat lengthy) expression more than once
```
[1,2,3,4]->map {$0 * 2}->reduce {$0 + $1} > 0 && [1,2,3,4]->map {$0*2}->reduce { $0 + $1} < 3
```
Then this can be written:
```
[1,2,3,4]->map {$0 * 2}->reduce {$0 + $1} -> with |$v| {$v > 0 && $v < 3}
```
The functions `then` and `lest` are variations on this theme where `then` acts as "if not nil, do this" and
`lest` as "if nil, do this".

## Expressions
The following expressions are supported:
| Operator(s)       | Meaning |
| ----              | ---- |
| `+ - * / %`       | Arithmetic operators, for example `1+2` with standard precedence. |
| `< <= > >= == !=` | Comparison operators. See more below. |
| `=~ !~`           | Match, not-match. See more below. |
| `in`              | Tests if lhs is equal to any value in an array on the rhs. |
| `!` `&&` `||` `^`    | Boolean operators unary *not*, *and*, *or*, and *xor* |
| *lhs* `[expr, expr...]` | Slice operation, taking one or a range of values from lhs. See more below. |
| *lhs* `?` `{` *then* `:` *else* `}` | if the *lhs* expr is true the result is the evaluation of the *then* expr, else the evaluation of the *else* expression. |
| *lhs* `->` *name* [`(`*args* `)`] [ <code>\|</code> *params* <code>\|</code> `] `{` *expr* `}` | Calls the named function. See below. |
| *lhs* `.` *name*  | Takes the property *name* from the object result from *lhs* |
| *name*            | Shorthand for `$self.`*name*                                |

## Comparison Operators
Comparison operatos compare the left and right side expressions. Comparisons across types place data types in an order:
* nil
* false, true
* numeric, integers and floats
* strings
* arrays

For example:
```
1 < 2           // true
false < 1       // true
true <  1       // true
1 < 1.2         // true (integer converted to floating point before comparison)
1 < "hello"     // true
1 < "1"         // true numbers in string form are not converted
[1,2] < [1,5]   // true
[1,2] < [1,2,0] // true (equal common length), the longer array makes it greater.
```

## Match Operators
The match operators `=~` (match) and `!~` (not match) match a lhs value against a data type or (if lhs is a string) a regular expression. For example:
```
"hello" =~ /el/      // true
10 =~ Integer[1,100] // true
"hello" !~ /el/      // false
10 !~ Integer[1,100] // false
```

## Slice operator
The slice operator can take one value from the lhs, or a range of values. As a special case, if
the lhs is the name of a data type (for example `Integer`), the slicing is applied to all possible
integers (for example `Integer[1,10]` creates a data type instance that matches integers between 1 and 10).

Here are some examples:
```
[1,2,3,4,5][1]   // produces 2
[1,2,3,4,5][1,1] // produces [2]
[1,2,3,4,5][1,3] // produces [2, 3, 4]
Integer[1,10].   // produces an Integer data type matching integers 1-10.
```
## Function Calls
Function calls are performed using the `->` operator. The lhs is the first "argument" to the function given by name. For example:
```
[1,2,3]->len // the length of the array (which is 3)
```
A function call can be made with additional arguments (depending on what the function requires/accepts). These are given within parentheses after the function name. For example:
```
1->compare(2) // produces -1 since 2 is greater than 1
```
It is accepted to have no arguments and still include the parentheses. For example:
```
[1,2,3]->len()
```

Some function accepts an additional expression to (typically) be evaluated with different parameters
for different values produced by the function. This expression is given within brackets `{ }`. For example:
```
[1,2,3]->reduce { $0 + $1 } // produces the sum of 1,2,3 which is 6.
```

The function sets up a context for evaluation of the expression where it can set variables for the expression to operate on. By default these parameters will be `$` followed by the parameter number, for example, `$0`, `$1` etc. The parameters can be named by stating their names enclosed in pipe <code>\|</code> delimiters.
For example:
```
[1,2,3]->reduce |$memo, $val| { $memo + $val } // produces the sum of 1,2,3 which is 6.
```
When parameters have been named, the default numerical variables are not set.

## Functions
The available functions are:

| signature | description |
| ------    | ----
| *lhs*<code>->**abs**()</code>              | Absolute value of integer or float values                |
| *lhs*<code>->**all**() \|$0\| {}</code>    | Is `true` if all elements in *lhs* evaluates to `true`   |
| *lhs*<code>->**all_or_none**() \|$0\| {}</code>  | Is `true` if all elements in *lhs* evaluates to `true` or none of them do  |
| *lhs*<code>->**any**() \|$0\| {}</code>    | Is `true` if any element in lhs evaluates to `true`      |
| *lhs*<code>->**ceil**()</code>             | Math ceil for float values, accepts integers as well     |
| *lhs*<code>->**count**() \|$0\| {}</code>  | Counts the number of times the evaluation returns `true` |
| *lhs*<code>->**compact**()</code>          | Returns new array where all nil/undefined values are removed      |
| *lhs*<code>->**compare**(*arg*)</code>     | Compares the *lhs* with the given arg and returns 1,0,-1 depending on which is greater  |
| *lhs*<code>->**filter**() \|$0\| {}</code> | Returns a new array with all elements for which the expression is `true`  |
| *lhs*<code>->**floor**()</code>            | Math floor for float values, which also accepts integers (no-op)  |
| *lhs*<code>->**len()**</code>              | Returns the length of an array or string                          |
| *lhs*<code>->**lest**() {}</code>          | Returns the value of the expression if *lhs* is nil               |
| *lhs*<code>->**map**() \|$0\| {}</code>    | Returns a new array with each result of evaluating the expression |
| *lhs*<code>->**max**(*arg*?)</code>        | Returns the greater of two values, or the greatest in *lhs* if no argument is given |
| *lhs*<code>->**min**(*arg*?)</code>        | Returns the smaller of two values, or the greatest in *lhs* if no argument is given |
| *lhs*<code>->**match**(*arg*)</code>       | Matches *lhs* with the given argument regular expression and returns array of match and submatches |
| *lhs*<code>->**reduce**(*start*?) \|$0,$1\| {}</code> | Evaluates the expression for each element in *lhs* and feeds in previous result. Also accepts optional start value. |
| *lhs*<code>->**round**()</code>            | Math round-to-even for float values, which also accepts integers (no-op)  |
| *lhs*<code>->**then**() \|$0\| {}</code>   | Returns the value of the expression if *lhs* is not nil |
| *lhs*<code>->**unique**()</code>           | Returns a new array with *lhs* reduced to consist of only unique entries |
| *lhs*<code>->**with**() \|$0\| {}</code>   | Returns the value of the expression |

## Property values
Expressions have access to the instance being validated by referencing the variable `$self`.
The properties of `$self` can however be directly referenced by using just the property name.
These examples are all equivalent:
```
$self.a
$self."a"
a
[$self, "a"].with { $0.$1}
```
# Generator Models
Some generators may support a generator model where additional metadata for the schema, types, properties, and relationships can be specified. These are not part of the model per se, but it is expected that all
generators conform to the same type of generator model.

# Commands
To parse and validate a Yammm source file:
```
whit yammm parse mymodel.yammm
```
If no errors were found nothing will be printed and the command will exit with status 0. Errors from parsing as well as generating
the Yammm model will be printed out and the exit status will be non-zero.

To parse, validate, and produce an output file with the model expressed in JSON format:
```
whit yammm parse --out mymodel.json mymodel.yammm
```
To produce markdown documentation with a diagram of the model and documentation text (from the produced JSON)
```
whit yammm convert --out mymodel.md --format=md mymodel.json
```

Whit-commands that operate on a schema can take the yammm DSL file directly as an argument.
