# Package test
```mermaid
classDiagram
    directionTB
class SimonDemo{
    value int ?
}
SimonDemo *-- "0..*" Part : PARTS
SimonDemo *-- "0..1" Engine : ENGINE
SimonDemo *-- "1..1" Transmission : HAS_TRANSMISSION_SYSTEM
SimonDemo --> "0..1" OtherOne : OTHER_ONE ...
SimonDemo --> "0..*" OtherMany : OTHER_MANY ...
SimonDemo --> "1..1" OtherOneOne : OTHER_ONE_ONE ...
SimonDemo --> "1..*" OtherOneMany : OTHER_ONE_MANY ...
class OtherOne
class OtherMany
class OtherOneOne
class OtherOneMany
class Part
class Engine
class Transmission
```

## Types
### Engine
### OtherMany
### OtherOne
### OtherOneMany
### OtherOneOne
### Part
### SimonDemo
Docs for Simon Demo. Blah blah blah Lore ipsum blah.

| property | type | kind | description |
| -------- | ---- | ---- | ----------- |
| value | int |  | <p></p> |

#### Associations

* **OTHER_ONE** --> (one) [OtherOne](#otherone)<br>
If there is another one it is linked with this link (showing (one))
    | property | type | kind | description |
    | -------- | ---- | ---- | ----------- |
    | someProperty | int |  | <p>Some documentation</p> |

* **OTHER_MANY** --> (many) [OtherMany](#othermany)<br>
If there is another one it is linked with this link (showing (many))
    | property | type | kind | description |
    | -------- | ---- | ---- | ----------- |
    | name | string & strings.MinRunes(1) & strings.MaxRunes(80) |  | <p>Some documentation</p> |

* **OTHER_ONE_ONE** --> (one:one) [OtherOneOne](#otheroneone)<br>
If there is another one it is linked with this link (showing (one:one)
    | property | type | kind | description |
    | -------- | ---- | ---- | ----------- |
    | anInt | int |  | <p>Some documentation</p> |

* **OTHER_ONE_MANY** --> (one:many) [OtherOneMany](#otheronemany)<br>
If there is another one it is linked with this link (showing (one:many))
    | property | type | kind | description |
    | -------- | ---- | ---- | ----------- |
    | anInt | int |  | <p>Some documentation</p> |
#### Compositions
* **PARTS** &#9670;-> (many) [Part](#part)
* **ENGINE** &#9670;-> (one) [Engine](#engine)
* **HAS_TRANSMISSION_SYSTEM** &#9670;-> (one:one) [Transmission](#transmission)
### Transmission
