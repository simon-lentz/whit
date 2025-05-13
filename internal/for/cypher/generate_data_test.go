// Contains generated code. For those parts: DO NOT EDIT.
package cypher_test

// blob is test data in json instance form.
// TODO: include data type other than basic
var blob = `
{
"Cars": [
	{
		"regNbr": "ABC123",
		"registrationDate": "2017-05-03",
		"model": "Skoda"
	},
	{
		"regNbr": "XYZ123",
		"registrationDate": "2017-05-03",
		"model": "Lada",
		"color": "rust brown"
	}
],
"People": [
	{
		"name": "Anna-Greta",
		"birthday": "1927-12-03",
		"HAS_Head": {
			"id": "$$abc12345",
			"hasHair": true,
			"color": "grey"
		}
	},
	{
		"name": "Henrik",
		"birthday": "1959-03-08",
		"MOTHER_Person": {
			"Where": {
				"name": "Anna-Greta"
			}
		},
		"accidentFreeSince": "1982-01-01",
		"OWNS_VEHICLE_RegisteredVehicle": {
			"fromDate": "2020-10-01",
			"Where": {
				"regNbr": "ABC123"
			}
		},
		"HAS_Head": {
			"id": "$$abc12345",
			"hasHair": true,
			"color": "grey"
		},
		"HAS_Limbs":[
			{
				"id": "$$arm1",
				"kind": "arm"
			},
			{
				"id": "$$leg1",
				"kind": "leg"
			}
		]
	}
]
}`

type Car struct {
	RegisteredVehicle
}
type Person struct {
	Entity
	Birthday string  `json:"birthday"`
	Name     string  `json:"name"`
	HAS_Head     *Head   `json:"HAS_Head,omitempty"`
	HAS_Limbs    []*Limb `json:"HAS_Limbs,omitempty"`
	MOTHER_Person *EDGE_MOTHER_Person `json:"MOTHER_Person,omitempty"`
}
type Head struct {
	Id string `json:"id"`
	HasHair bool   `json:"hasHair"`
	Color   string `json:"color,omitempty"`
}
type Limb struct {
	Id string `json:"id"`
	Kind string `json:"kind"`
}
type Registered struct {
	RegNbr           string `json:"regNbr"`
	RegistrationDate string `json:"registrationDate"`
}
type Vehicle struct {
	Color string `json:"color,omitempty"`
	Model string `json:"model,omitempty"`
}
type RegisteredVehicle struct {
	Vehicle
	Registered
}
type MotorVehicleOwner struct {
	AccidentFreeSince string `json:"accidentFreeSince,omitempty"`
	OWNS_VEHICLE_RegisteredVehicle *EDGE_OWNS_VEHICLE_RegisteredVehicle `json:"OWNS_VEHICLE_RegisteredVehicle,omitempty"`
}
type Entity struct {
	MotorVehicleOwner
}
type Graph struct {
	Cars   []*Car    `json:"Cars,omitempty"`
	People []*Person `json:"People,omitempty"`
}

type EDGE_OWNS_VEHICLE_RegisteredVehicle struct {
	FromDate string `json:"fromDate"`
	ToDate   string `json:"toDate?"`
	Where    struct {
		RegNbr string `json:"regNbr"`
	} `json:"Where"`
}
type EDGE_MOTHER_Person struct {
	Where struct {
		Name string `json:"name"`
	} `json:"Where"`
}

const SerializedModel = `
{
	"name": "cypher_test",
	"types": [
		{
			"name": "Head",
			"plural_name": "Heads",
			"is_part": true,
			"properties": [
				{
					"name": "id",
					"datatype": ["UUID"],
					"primary": true
				},
				{
					"name": "hasHair",
					"datatype": ["Boolean"]
				},
				{
					"name": "color",
					"datatype": ["String"],
					"optional": true
				}
			]
		},
		{
			"name": "Limb",
			"plural_name": "Limbs",
			"is_part": true,
			"properties": [
				{
					"name": "kind",
					"datatype": ["String"]
				},
				{
					"name": "id",
					"datatype": ["UUID"],
					"primary": true
				}
			]
		},
		{
			"name": "Car",
			"plural_name": "Cars",
			"inherits": [
				"RegisteredVehicle"
			]
		},
		{
			"name": "Person",
			"plural_name": "People",
			"properties": [
				{
					"name": "birthday",
					"datatype": ["Date"]
				},
				{
					"name": "name",
					"datatype": ["String"],
					"primary": true
				}
			],
			"compositions": [
				{
					"to": "Head",
					"optional": false,
					"many": false,
					"name": "HAS"
				},
				{
					"to": "Limb",
					"optional": true,
					"many": true,
					"name": "HAS"
				}
			],
			"associations": [
				{
					"name": "MOTHER",
					"to": "Person",
					"optional": true,
					"many": false
				}
			],
			"inherits": [
				"Entity"
			]
		},
		{
			"name": "Registered",
			"plural_name": "Registereds",
			"properties": [
				{
					"name": "regNbr",
					"datatype": ["String"],
					"primary": true
				},
				{
					"name": "registrationDate",
					"datatype": ["String"]
				}
			],
			"is_abstract": true
		},
		{
			"name": "Vehicle",
			"plural_name": "Vehicles",
			"properties": [
				{
					"name": "color",
					"datatype": ["String"],
					"optional": true
				},
				{
					"name": "model",
					"datatype": ["String"],
					"optional": true
				}
			],
			"is_abstract": true
		},
		{
			"name": "RegisteredVehicle",
			"plural_name": "RegisteredVehicles",
			"inherits": [
				"Vehicle",
				"Registered"
			],
			"is_abstract": true
		},
		{
			"name": "MotorVehicleOwner",
			"plural_name": "MotorVehicleOwners",
			"properties": [
				{
					"name": "accidentFreeSince",
					"datatype": ["String"],
					"optional": true
				}
			],
			"associations": [
				{
					"to": "RegisteredVehicle",
					"optional": true,
					"many": false,
					"name": "OWNS_VEHICLE",
					"properties": [
						{
							"name": "fromDate",
							"datatype": ["String"]
						},
						{
							"name": "toDate",
							"datatype": ["String"],
							"optional": true
						}
					]
				}
			],
			"is_abstract": true
		},
		{
			"name": "Entity",
			"plural_name": "Entities",
			"inherits": [
				"MotorVehicleOwner"
			],
			"is_abstract": true
		},
		{
			"name": "Company",
			"plural_name": "Companies",
			"inherits": [
				"Entity"
			]
		}
	],
	"data_types": [
		{
			"name": "xtimeStamp",
			"constraint": ["String"]
		}
	]
}
`
