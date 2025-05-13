package utils

// HasKey returns true if the given map has the given key, false otherwise.
func HasKey[K comparable, V any](m map[K]V, key K) (ok bool) {
	_, ok = m[key]
	return
}

// Keys returns the keys of a map.
func Keys[K comparable, V any](m map[K]V) (result []K) {
	for k := range m {
		result = append(result, k)
	}
	return
}

// Values returns the values of a map.
func Values[K comparable, V any](m map[K]V) (result []V) {
	for _, v := range m {
		result = append(result, v)
	}
	return
}

// FilterMap returns a new map where the entries in the given map for which the predicate function
// returns false are not present.
func FilterMap[K comparable, V any](m map[K]V, predicate func(K, V) bool) (result map[K]V) {
	result = make(map[K]V, len(m))
	for k, v := range m {
		if predicate(k, v) {
			result[k] = v
		}
	}
	return result
}
