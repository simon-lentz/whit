package xray

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/wyrth-io/whit/internal/utils"
)

// Wrapper is a helper to deal with reflection of a wrapped instance document.
type Wrapper interface {
	// Feature returns a Wrapper representing an "object" or a slice of "object" given a name.
	// If feature is not present nil is returned.
	Feature(name string) Wrapper

	// FeatureAtIndex return a Wrapper representing an "object" or a slice of "object" given an index.
	// Must be used with a Wrapper of `reflect.Slice`` type or nil will be returned.
	FeatureAtIndex(i int) Wrapper

	// FeatureNames returns a slice of all names in a Wrapper representing an Object. The returned
	// names are sorted alphabetically.
	FeatureNames() []string

	// HasCapitalizedFeatureNames returns true for a representation where feature names are
	// all capitalized (initial letter is upper case).
	HasCapitalizedFeatureNames() bool

	// Feature name returns a wrapper representing the feature name.
	FeatureName(name string) Wrapper

	// Kind returns the reflected Kind of value held in the wrapperument.
	Kind() reflect.Kind

	// Len returns the length of a slice if this Wrapper represents a Slice, else the number of
	// values or features in an Object.
	Len() int

	// Value returns a "property" value from a Wrapper representing an Object. A "value" is always
	// an interface. This can also be called for "features" to get the interface to the feature.
	Value(name string) any

	// IsObject returns true if this wrapper represents an Object.
	IsObject() bool

	// IsSlice returns true if this wrapper represents a Slice
	IsSlice() bool
}

// NewWrapper returns a Wrapper reflection handler around the given graph which should be one of
// map[string]any, struct, []struct, or []map[string]any. If the given graph is already
// a Wrapper implementation then this value is returned. If any other type of value for graph is
// given nil will be returned.
func NewWrapper(graph any) Wrapper {
	// If graph is already a Wrapper implementation, simply return it.
	if w, ok := graph.(Wrapper); ok {
		return w
	}
	if graph == nil {
		return nil
	}
	//
	val := reflect.ValueOf(graph)
	if val.IsZero() {
		return nil
	}
	// t := reflect.TypeOf(graph)
	t := reflect.Indirect(val).Type()
	val = reflect.Indirect(val) // in case it is a pointer
	// tt := reflect.TypeOf(val)
	kind := val.Kind()
	switch kind {
	case reflect.Map:
	case reflect.Struct:
	case reflect.Slice:
	default:
		return nil
	}
	return &wrapper{node: graph, kind: kind, val: val, typ: t}
}

type wrapper struct {
	val  reflect.Value
	node any
	kind reflect.Kind
	typ  reflect.Type
}

func (w *wrapper) Feature(name string) Wrapper {
	switch w.kind {
	case reflect.Struct:
		// Go object
		f := w.val.FieldByName(name)
		if !f.IsValid() {
			return nil
		}
		if !f.CanInterface() {
			return nil
		}
		return NewWrapper(f.Interface())
	case reflect.Map:
		if obj, ok := w.node.(map[string]any); ok {
			return NewWrapper(obj[name])
		}
		return nil
	case reflect.Slice:
		// cannot take feature of slice
		return nil
	}
	return nil
}

func (w *wrapper) HasCapitalizedFeatureNames() bool {
	return w.kind == reflect.Struct
}

func (w *wrapper) FeatureName(name string) Wrapper {
	f := w.Feature(name)
	if f != nil {
		return NewWrapper(name)
	}
	return nil
}

func (w *wrapper) FeatureAtIndex(i int) Wrapper {
	switch w.kind {
	case reflect.Map:
		return nil
	case reflect.Struct:
		return nil
	case reflect.Slice:
		return NewWrapper(w.val.Index(i).Interface())
	}
	return nil
}

func (w *wrapper) Value(name string) any {
	switch w.kind {
	case reflect.Struct:
		// Go object, all names are initial UC
		name = utils.Capitalize(name)
		f := w.val.FieldByName(name)
		if !f.IsValid() {
			return nil
		}
		if !f.CanInterface() {
			return nil
		}
		// optional fields are pointers
		if f.Kind() == reflect.Ptr {
			f = f.Elem()
		}
		return f.Interface()
	case reflect.Map:
		if obj, ok := w.node.(map[string]any); ok {
			return obj[name]
		}
		return nil
	case reflect.Slice:
		// cannot take feature of slice
		return nil
	}
	return nil
}
func (w *wrapper) Kind() reflect.Kind {
	return w.kind
}

func (w *wrapper) Len() int {
	switch w.kind {
	case reflect.Slice:
		return w.val.Len()
	case reflect.Map:
		if obj, ok := w.node.(map[string]any); ok {
			return len(obj)
		}
	case reflect.Struct:
		return w.val.NumField()
	}
	return -1
}

func (w *wrapper) FeatureNames() (result []string) {
	w.val.Type()
	switch w.kind {
	case reflect.Struct:
		result = w.deepFieldNames(w.typ)
		// Go names have initial upper case - must change to initial LC for
		// // Go object
		// for i := 0; i < w.typ.NumField(); i++ {
		// 	result = append(result, w.typ.Field(i).Name)
		// }
	case reflect.Map:
		if obj, ok := w.node.(map[string]any); ok {
			for x := range obj {
				result = append(result, x)
			}
		}
	case reflect.Slice:
		// cannot get feature names of slice, return empty slice

	case reflect.Pointer:
		fmt.Printf("Oh, it is a pointer\n")
	}
	sort.Strings(result)
	return
}

func (w *wrapper) IsObject() bool {
	return w.kind == reflect.Struct || w.kind == reflect.Map
}

func (w *wrapper) IsSlice() bool {
	return w.kind == reflect.Slice
}

func (w *wrapper) deepFieldNames(t reflect.Type) (result []string) {
	m := make(map[string]struct{})
	w.collectFieldNames(t, m)
	for name := range m {
		result = append(result, name)
	}
	return result
}

func (w *wrapper) collectFieldNames(t reflect.Type, m map[string]struct{}) {
	// Return if not struct or pointer to struct.
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return
	}

	// Iterate through fields collecting names in map.
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		if !sf.Anonymous {
			m[sf.Name] = struct{}{}
		} else {
			// Recurse into anonymous fields. Do not collect the name.
			w.collectFieldNames(sf.Type, m)
		}
	}
}
