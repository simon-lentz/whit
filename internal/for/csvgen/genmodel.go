package csvgen

// Genmodel describes a mapping of csv rows to instances of a yammm model.
type Genmodel struct {
	// Typename is the name of the yammm type to map to. For example "Person".
	// This type must exist in the schema being used.
	Typename string `json:"type"`

	// PropertyMap yamm property name to value of csv column name. For example "name":"NAME".
	PropertyMap map[string]string `json:"property_map"`

	// AssociationMap maps from a Yammm Assoiation to columns making out the primary key and
	// association properties. For example:
	//   "association_map": {
	//     "OWNS_Thing": {
	//	     "where": { "thingID": "THINGNBR", ...},
	//       "properties": {
	// 		  	"since": "OWNED_SINCE"
	// 		  }
	//     }
	//  }
	// This allows mapping several columns to primary keys and association properties.
	AssociationMap map[string]*AssocModel `json:"association_map"`
}

// AssocModel describes the primary keys used to form an association and any properties for
// the association. The keys in the maps are the yammm names.
type AssocModel struct {
	Where      map[string]string `json:"where"`
	Properties map[string]string `json:"property_map,omitempty"`
}
