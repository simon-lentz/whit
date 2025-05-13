# Package example
```mermaid
classDiagram
    directionTB
class Vehicle{
    <<Abstract>>
    example Integer ?
}
class Person{
    name String +
}
Person --> "1..*" Car : OWNS ...
class Car{
    regnbr String[6, 8] +
}
Car --|> Vehicle : Extends
Car --|> Something : Extends
Car --|> TotallyDifferent : Extends
Car --|> CommonStuff : Extends
Car *-- "1..*" AutoPart : AUTOPARTS
class AutoPart{
    <<Abstract>>
    serialNumber String +
}
class Engine{
    power Pattern["[0-9]+kW"] ?
}
Engine --|> AutoPart : Extends
class Transmission{
    type Enum["manual", "auto", "cvt", "semi-auto"]
}
Transmission --|> AutoPart : Extends
class CommonStuff{
    <<Abstract>>
    comment String[_, 512] ?
}
class BaseExample{
    propS String ?
    propSOneTen String[1, 10] ?
    propSOneMax String[1, _] ?
    propI Integer ?
    propIFive10 Integer[5, 10] ?
    propIZeroMax Integer[0, _] ?
    propIMinZero Integer ?
    propF Float ?
    propFZeroMax Float[0, _] ?
    propFMinZero Float[_, 0] ?
    propB Boolean ?
    propD Date ?
    propT Timestamp["2006-01-02T15:04:05Z07:00"] ?
    propTCustom Timestamp["20060102"] ?
    type Integer ?
    datatype Integer ?
    mixin Integer ?
    schema Integer ?
    required Integer ?
    primary Integer ?
    extends Integer ?
    includes Integer ?
    abstract Integer ?
    mixin Integer ?
}
class RequiredExample{
    mustBeSet Boolean
}
class Example{
    id String +
}
Example --|> BaseExample : Extends
Example --|> RequiredExample : Extends
Example --|> CommonStuff : Extends
class Powerplant{
    capacity Pattern["[0-9]+MW"] ?
}
class Something
class TotallyDifferent
```

## Types
### AutoPart
* Abstract: true

| property | type | kind | description |
| -------- | ---- | ---- | ----------- |
| serialNumber | String | primary | <p></p> |
### BaseExample
| property | type | kind | description |
| -------- | ---- | ---- | ----------- |
| propS | String |  | <p></p> |
| propSOneTen | String[1, 10] |  | <p></p> |
| propSOneMax | String[1, _] |  | <p></p> |
| propI | Integer |  | <p></p> |
| propIFive10 | Integer[5, 10] |  | <p></p> |
| propIZeroMax | Integer[0, _] |  | <p></p> |
| propIMinZero | Integer |  | <p></p> |
| propF | Float |  | <p></p> |
| propFZeroMax | Float[0, _] |  | <p></p> |
| propFMinZero | Float[_, 0] |  | <p></p> |
| propB | Boolean |  | <p></p> |
| propD | Date |  | <p></p> |
| propT | Timestamp["2006-01-02T15:04:05Z07:00"] |  | <p></p> |
| propTCustom | Timestamp["20060102"] |  | <p></p> |
| type | Integer |  | <p></p> |
| datatype | Integer |  | <p></p> |
| mixin | Integer |  | <p></p> |
| schema | Integer |  | <p></p> |
| required | Integer |  | <p></p> |
| primary | Integer |  | <p></p> |
| extends | Integer |  | <p></p> |
| includes | Integer |  | <p></p> |
| abstract | Integer |  | <p></p> |
| mixin | Integer |  | <p></p> |
### Car
* Extends: [Vehicle](#vehicle), [Something](#something), [TotallyDifferent](#totallydifferent), [CommonStuff](#commonstuff)

Car is a car...

| property | type | kind | description |
| -------- | ---- | ---- | ----------- |
| regnbr | String[6, 8] | primary | <p>regNbr is the registration number of the car. It is a string between 6 and 8 characters long.</p> |
#### Compositions
* **AUTOPARTS** &#9670;-> (one:many) [AutoPart](#autopart)
### CommonStuff
* Abstract: true

| property | type | kind | description |
| -------- | ---- | ---- | ----------- |
| comment | String[_, 512] |  | <p></p> |
### Engine
* Extends: [AutoPart](#autopart)

| property | type | kind | description |
| -------- | ---- | ---- | ----------- |
| power | Pattern["[0-9]+kW"] |  | <p></p> |
### Example
* Extends: [BaseExample](#baseexample), [RequiredExample](#requiredexample), [CommonStuff](#commonstuff)

| property | type | kind | description |
| -------- | ---- | ---- | ----------- |
| id | String | primary | <p></p> |
### Person
| property | type | kind | description |
| -------- | ---- | ---- | ----------- |
| name | String | primary | <p>name is the name of a person. This is a very long lore ipsum text for the purpose of<br>    seeing what the output of multiple lines of description will look like. Does blank lines<br>    survive?<br><br>    After blank line.<br>    And the all lived happiley everafter.<br>    THE END <br></p> |

#### Associations

* **OWNS** --> (one:many) [Car](#car)<br>
    | property | type | kind | description |
    | -------- | ---- | ---- | ----------- |
    | since | Date |  | <p></p> |
### Powerplant
| property | type | kind | description |
| -------- | ---- | ---- | ----------- |
| capacity | Pattern["[0-9]+MW"] |  | <p></p> |
### RequiredExample
| property | type | kind | description |
| -------- | ---- | ---- | ----------- |
| mustBeSet | Boolean | required | <p></p> |
### Something
### TotallyDifferent
### Transmission
* Extends: [AutoPart](#autopart)

| property | type | kind | description |
| -------- | ---- | ---- | ----------- |
| type | Enum["manual", "auto", "cvt", "semi-auto"] | required | <p></p> |
### Vehicle
* Abstract: true

| property | type | kind | description |
| -------- | ---- | ---- | ----------- |
| example | Integer |  | <p></p> |
