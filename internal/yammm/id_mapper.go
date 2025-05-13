package yammm

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/lithammer/shortuuid/v4"
	"github.com/wyrth-io/whit/internal/tc"
	"github.com/wyrth-io/whit/internal/utils"
)

// IDMapper is a listener for UUID properties that needs to be replaced once all are known.
type IDMapper struct {
	BaseGraphListener
	local2uuid   map[string]uuid.UUID
	replacements map[string]uuid.UUID
	locals       *utils.Set[string]
}

// NewIDMapper returns a new IDMapper for mapping special references to UUIDs into actual uuid.UUID values.
func NewIDMapper() *IDMapper {
	return &IDMapper{
		local2uuid:   make(map[string]uuid.UUID),
		replacements: make(map[string]uuid.UUID),
		locals:       utils.NewSet[string](),
	}
}

// Map creates a mapping of all UUID references in the given graph. The graph must have been validated.
// A replacement map is returned that will enable lookup of $$:global:local and $$local strings to the
// corresponding UUID. This enables easy lookup given property vales on this form.
func (m *IDMapper) Map(ctx Context, graph any) (replacements map[string]uuid.UUID, err error) {
	walker := NewGraphWalker(ctx, m)
	walker.Walk(graph)
	var uid uuid.UUID
	var ok bool
	for _, local := range m.locals.Slices() {
		if uid, ok = m.local2uuid[local]; !ok {
			// Not defined, define it using a random v4 UUID
			uid, err = uuid.NewRandom()
			if err != nil {
				return nil, err
			}
			m.local2uuid[local] = uid
			// Store entries for $$:globalshort:local => uid
			m.replacements[fmt.Sprintf("$$:%s:%s", shortuuid.DefaultEncoder.Encode(uid), local)] = uid
		}
		// store local to uid (also for the already defined)
		m.replacements[fmt.Sprintf("$$%s", local)] = uid
	}
	return m.replacements, nil
}

var uuidFilter = func(p *Property) bool { return p.DataType[0] == tc.UUIDS }
var propNameMapper = func(p *Property) string { return p.Name }

// OnProperties processes UUID properties.
func (m *IDMapper) OnProperties(_ Context, t *Type, propMap map[string]any) {
	// Only operate on properties of UUID type
	ids := utils.Filter(t.AllProperties(), uuidFilter)
	m.processProperties(utils.Map(ids, propNameMapper), propMap)
}
func (m *IDMapper) processProperties(ids []string, propMap map[string]any) {
	if len(ids) < 1 {
		return
	}
	// For the UUID properties, get the value from the propmap and see if it is a
	// $$local or $$:globalShortUUID:local. If so, add locals to a set of locals
	// and remember the already resolved $$:global:local uuids.
	for _, k := range ids {
		if v, ok := propMap[k]; ok {
			switch x := v.(type) {
			case string:
				if len(x) < 3 {
					continue
				}
				if strings.HasPrefix(x, "$$") {
					// is it resolved with shortUUID global part?
					if x[2] == ':' {
						parts := strings.Split(x, ":")
						if len(parts) < 3 || len(parts[2]) < 1 { // bad format, but not an error
							continue
						}
						uid, _ := shortuuid.DefaultEncoder.Decode(parts[1]) // Is validated, ignore error.
						m.local2uuid[parts[2]] = uid
						m.replacements[x] = uid
						m.replacements[fmt.Sprintf("$$%s", parts[2])] = uid
					} else {
						// it is a local reference, add it to the set of locals.
						m.locals.Add(x[2:])
					}
				}
			default:
			}
		}
	}
}

// OnEdge handles UUID properties of associations and primary keys.
func (m *IDMapper) OnEdge(_ Context,
	a *Association, assocPropMap map[string]any,
	_ *Type, _ map[string]any, t2 *Type, toPks map[string]any,
) {
	// Handle uuid references in association properties
	ids := utils.Filter(a.Properties, uuidFilter)
	m.processProperties(utils.Map(ids, propNameMapper), assocPropMap)

	// Hande uuid references in toType primary key values
	ids = utils.Filter(t2.AllPrimaryKeys(), uuidFilter)
	m.processProperties(utils.Map(ids, propNameMapper), toPks)
}
