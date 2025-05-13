package utils

import (
	"github.com/mitchellh/hashstructure/v2"
	"github.com/tada/catch"
)

// HashMap is a map supporting any type of key and any type of value.
type HashMap interface {
	// Get returns the mapped value for the given key or nil if value is not present.
	Get(key any) any

	// Get returns the mapped value for the given key, returns nil, false if value is not present.
	GetOk(key any) (any, bool)

	// Put stores the given value at the given key, overwriting any other value bound for that key.
	Put(key, value any)

	// PutUnique stores the given value at the given key unless the key is already bound to a value.
	// If key is already present false is returned and the value is not stored
	PutUnique(key, value any) (unique bool)

	// Remove removes the key from the hashmap. If key was not present this is a no-op.
	Remove(key any)
}

type hashMap struct {
	data map[uint64]any
}

// NewHashMap returns a new HashMap. A HashMap accepts any type of key and any type of value.
func NewHashMap() HashMap {
	return &hashMap{data: make(map[uint64]any)}
}
func (m *hashMap) Get(key any) any {
	if val, ok := m.GetOk(key); ok {
		return val
	}
	return nil
}
func (m *hashMap) GetOk(key any) (any, bool) {
	if val, ok := m.data[m.mustHash(key)]; ok {
		return val, ok
	}
	return nil, false
}
func (m *hashMap) Put(key, value any) {
	m.data[m.mustHash(key)] = value
}
func (m *hashMap) PutUnique(key, value any) bool {
	h := m.mustHash(key)
	if _, ok := m.data[h]; ok {
		return false
	}
	m.data[h] = value
	return true
}
func (m *hashMap) Remove(key any) {
	delete(m.data, m.mustHash(key))
}
func (m *hashMap) mustHash(key any) uint64 {
	h, err := hashstructure.Hash(key, hashstructure.FormatV2, nil)
	if err != nil {
		panic(catch.Error("cannot hash key: %s", err))
	}
	return h
}
