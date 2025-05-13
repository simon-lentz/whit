package utils

import (
	"fmt"
	"strings"
)

// Set is a set of compareable.
type Set[T comparable] struct {
	set map[T]interface{}
}

// NewSet creates a new set. Optionally with given values added.
func NewSet[T comparable](values ...T) *Set[T] {
	s := &Set[T]{set: make(map[T]interface{})}
	for _, v := range values {
		s.Add(v)
	}
	return s
}

// NewSetFrom creates a new set from a slice of V and a transforming function V ->T.
// Returns a Set[T].
func NewSetFrom[T comparable, V any](slice []V, f func(v V) T) *Set[T] {
	set := NewSet[T]()
	for _, v := range slice {
		value := f(v)
		set.Add(value)
	}
	return set
}

// Add adds given value to set and returns true if it was not already in the set.
func (s *Set[T]) Add(val T) bool {
	if _, ok := s.set[val]; ok {
		return false
	}
	var empty interface{}
	s.set[val] = empty
	return true
}

// Remove returns a new set with the given value removed from the set. The original set is
// unchanged. If the value was not contained the original set is returned.
func (s *Set[T]) Remove(val T) *Set[T] {
	if _, ok := s.set[val]; !ok {
		return s
	}
	s2 := NewSet(val)
	return s.Diff(s2)
}

// Size returns the number of entries in the set.
func (s *Set[T]) Size() int {
	return len(s.set)
}

// Each calls the given function for each value in the set.
func (s *Set[T]) Each(f func(val T)) {
	if s == nil {
		return
	}
	for k := range s.set {
		f(k)
	}
}

// Contains returns true if the set has the given value.
func (s *Set[T]) Contains(val T) bool {
	if _, ok := s.set[val]; ok {
		return true
	}
	return false
}

// Union returns the union of this with a number of other sets.
func (s *Set[T]) Union(sets ...*Set[T]) *Set[T] {
	result := NewSet[T]()
	for k := range s.set {
		result.Add(k)
	}
	for _, s2 := range sets {
		s2.Each(func(val T) { result.Add(val) })
	}
	return result
}

// Slices returns the set values as a slice.
func (s *Set[T]) Slices() (result []T) {
	for k := range s.set {
		result = append(result, k)
	}
	return
}

// Intersection returns a set with the values common to this and the given set.
func (s *Set[T]) Intersection(s2 *Set[T]) *Set[T] {
	result := NewSet[T]()
	for x := range s.set {
		if s2.Contains(x) {
			result.Add(x)
		}
	}
	return result
}

// Empty returns true if the set is empty. This is the same as `set.Size() == 0“.
func (s *Set[T]) Empty() bool {
	return len(s.set) == 0
}

// Diff returns a new set with the elements in s that are not in s2.
func (s *Set[T]) Diff(s2 *Set[T]) *Set[T] {
	result := NewSet[T]()
	for x := range s.set {
		if !s2.Contains(x) {
			result.Add(x)
		}
	}
	return result
}

// Equal returns tru if both sets have the same elements.
func (s *Set[T]) Equal(s2 *Set[T]) bool {
	if len(s.set) != len(s2.set) {
		return false
	}
	return s.All(func(v T) bool {
		return s2.Contains(v)
	})
}

// All returns true if the predicate function returns true for all elements in this.
func (s *Set[T]) All(predicate func(v T) bool) bool {
	for x := range s.set {
		if !predicate(x) {
			return false
		}
	}
	return true
}

// String returns an easy to read string on the form Set[elem, elem, ...,].
func (s *Set[T]) String() string {
	var b strings.Builder
	b.WriteString("Set[")
	values := make([]T, 0, len(s.set))
	for x := range s.set {
		values = append(values, x)
	}
	for _, x := range values {
		b.WriteString(fmt.Sprintf("%v,", x))
	}
	b.WriteString("]")
	return b.String()
}
