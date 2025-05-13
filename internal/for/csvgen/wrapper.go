package csvgen

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"reflect"
	"sort"

	"github.com/wyrth-io/whit/internal/xray"
	"github.com/wyrth-io/whit/internal/yammm"
)

// NewGraph creates a new graph implementing the xray.Wrapper interface making possible to
// treat a csv as an instance graph.
func NewGraph(ctx yammm.Context, gm *Genmodel, reader *csv.Reader) (graph any, err error) {
	t := ctx.LookupType(gm.Typename)
	if t == nil {
		return nil, fmt.Errorf("type in genmodel '%s' not found in schema", gm.Typename)
	}
	pluralName := t.PluralName
	instances, err := NewTopWrapper(ctx, t, gm, reader)
	if err != nil {
		return nil, err
	}
	return map[string]any{pluralName: instances}, nil
}

// NewTopWrapper creates a wrapper for the array of instances in an instance document.
func NewTopWrapper(ctx yammm.Context, targetType *yammm.Type, gm *Genmodel, reader *csv.Reader) (xray.Wrapper, error) {
	// TODO: Get all feature names in a slice from gm.
	header, err := reader.Read()
	if err != nil {
		return nil, err
	}
	headerMap := make(map[string]int, len(header))
	for i := range header {
		headerMap[header[i]] = i
	}
	numFeatures := len(gm.PropertyMap) + len(gm.AssociationMap)
	featureNames := make([]string, 0, numFeatures) // names of properties and associations
	colMap := make(map[string]int, len(gm.PropertyMap))
	for k, v := range gm.PropertyMap {
		colMap[k] = headerMap[v]
		featureNames = append(featureNames, k)
	}
	sort.Strings(featureNames)

	for k := range gm.AssociationMap {
		featureNames = append(featureNames, k)
	}
	allProps := targetType.AllProperties()
	propMap := make(map[string]*yammm.Property, len(allProps))
	for _, p := range allProps {
		propMap[p.Name] = p
	}
	allAssocs := targetType.AllAssociations()
	allAssocsMap := make(map[string]*yammm.Association)
	for i := range allAssocs {
		a := allAssocs[i]
		allAssocsMap[a.PropertyName(ctx)] = a
	}
	// // Make map of those associations that are mapped
	// assocMap := make(map[string]*yammm.Association, len(gm.AssociationMap))
	// for k, v := range gm.AssociationMap {
	// 	assocMap[k] = allAssocsMap[]
	// }
	theTop := &top{
		ctx:          ctx,
		gm:           gm,
		reader:       reader,
		headerMap:    headerMap,
		colMap:       colMap,
		propMap:      propMap,
		assocMap:     allAssocsMap,
		featureNames: featureNames,
	}
	rows := []*row{}
	for {
		r, err := reader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}
		rows = append(rows, newRow(theTop, r))
	}
	theTop.rows = rows
	theTop.numRows = len(rows)
	theTop.numFeatures = numFeatures
	return theTop, nil
}

type top struct {
	ctx          yammm.Context                 // to get meta information
	gm           *Genmodel                     // mappings from model to header
	reader       *csv.Reader                   // for reading the csv
	rows         []*row                        // row wrappers
	numRows      int                           // number of rows
	numFeatures  int                           // number of features
	featureNames []string                      // names of features (properties and relations)
	colMap       map[string]int                // feature name to column number (for properties)
	headerMap    map[string]int                // maps header column names to column index
	propMap      map[string]*yammm.Property    // to get meta information
	assocMap     map[string]*yammm.Association // to get meta information
}

func (w *top) Feature(_ string) xray.Wrapper {
	return nil
}
func (w *top) IsSlice() bool {
	return true
}
func (w *top) IsObject() bool {
	return false
}
func (w *top) FeatureAtIndex(i int) xray.Wrapper {
	return w.rows[i]
}
func (w *top) FeatureNames() []string {
	return []string{}
}
func (w *top) HasCapitalizedFeatureNames() bool {
	return false
}
func (w *top) Len() int {
	return len(w.rows)
}
func (w *top) Value(_ string) any {
	return nil
}
func (w *top) FeatureName(_ string) xray.Wrapper {
	return nil
}
func (w *top) Kind() reflect.Kind {
	// TODO: This is lying, it is a wrapper struct
	return reflect.Slice
}

type row struct {
	top     *top
	data    []string
	lineNbr int
}

func newRow(top *top, data []string) *row {
	return &row{top: top, data: data}
}

func (w *row) Feature(name string) xray.Wrapper {
	return xray.NewWrapper(w.Value(name))
}
func (w *row) IsSlice() bool {
	return false
}
func (w *row) IsObject() bool {
	return true
}
func (w *row) FeatureAtIndex(_ int) xray.Wrapper {
	return nil
}
func (w *row) FeatureNames() []string {
	return w.top.featureNames
}
func (w *row) HasCapitalizedFeatureNames() bool {
	return false
}
func (w *row) Len() int {
	return w.top.numFeatures
}
func (w *row) Value(name string) any {
	// TODO: a map of assoc feature to properties????? Not needed, since this is
	// found in genmodel assoc, lookup the feature name to get where/props and
	// get the header number from header map.

	col, ok := w.top.colMap[name]
	if !ok {
		// If not a property, it can be an association.
		assocModel, ok := w.top.gm.AssociationMap[name]
		if !ok {
			return nil // Does not have this association since it isn't in the genmodel
		}
		a := w.top.assocMap[name]
		aProps := a.PropMap()
		_, aPkProps := a.TargetMap(w.top.ctx)
		// Construct whereMap with the target type's primary keys filled in.
		whereMap := make(map[string]any, len(assocModel.Where))
		for pN, hN := range assocModel.Where {
			col := w.top.headerMap[hN]
			val, err := mapDataType(w.lineNbr, w.data[col], pN, aPkProps)
			if err != nil {
				panic(err) // should not happen if instance is validated
			}
			whereMap[pN] = val
		}
		// Construct the resulting map of association properties and Where-map.
		resultMap := make(map[string]any, len(aProps)+1)
		resultMap["Where"] = whereMap
		for pN, hN := range assocModel.Properties {
			col := w.top.headerMap[hN]
			val, err := mapDataType(w.lineNbr, w.data[col], pN, aProps)
			if err != nil {
				panic(err) // should not happen if instance is validated
			}
			resultMap[pN] = val
		}
		if a.Many {
			// a slice of maps is expected
			return []any{resultMap}
		}
		return resultMap
	}
	// Get property value
	valStr := w.data[col]
	val, err := mapDataType(w.lineNbr, valStr, name, w.top.propMap)
	if err != nil {
		panic(err) // cannot happen if this "graph" was validated.
	}
	return val
}
func (w *row) FeatureName(name string) xray.Wrapper {
	ft := w.Feature(name)
	if ft != nil {
		return xray.NewWrapper(name)
	}
	return nil
}
func (w *row) Kind() reflect.Kind {
	// TODO: This is lying, it is a wrapper struct
	return reflect.Slice
}
