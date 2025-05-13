package yammm

// Genmodel is a struct for describing additional "genmodel" properties for a particular technology.
type Genmodel struct {
	// Generator is the name of the generator this genmodel is for.
	Generator string `json:"generator"`

	// Genmodels is a mapping from a yammm entity "path" to a genmodel for that entity.
	// For example the key `"Car.regNbr"`would refer to the property regNbr in the type Car.
	Genmodels map[string]any `json:"genmodels"`

	// Definitions define named definitions, such that if a Genmodels entry maps to a string, the
	// definition of that name would be used for that element.
	Definitions map[string]any `json:"definitions"`
}

// NewGenmodel returns a new empty genmodel.
func NewGenmodel() *Genmodel {
	return &Genmodel{Generator: "", Genmodels: make(map[string]any), Definitions: make(map[string]any)}
}
