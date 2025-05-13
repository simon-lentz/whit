package utils

// Filter returns a new slice containing the elements from slices or which the given
// filterFunc returned true.
func Filter[T any](slices []T, filterFunc func(T) bool) (result []T) {
	for _, v := range slices {
		if filterFunc(v) {
			result = append(result, v)
		}
	}
	return
}

// Partition returns two new slices where the first (trueResult) contains all elements for which
// the given function returned true, and the second (falseResult) the rest. The original slice
// is not changed.
func Partition[T any](slices []T, filterFunc func(T) bool) (trueResult []T, falseResult []T) {
	for _, v := range slices {
		if filterFunc(v) {
			trueResult = append(trueResult, v)
		} else {
			falseResult = append(falseResult, v)
		}
	}
	return
}

// All returns true if the given predicateFunc returns true for all elements in slices.
func All[T any](slices []T, predicateFunc func(T) bool) (allTrue bool) {
	if len(slices) == 0 {
		return false
	}
	for _, v := range slices {
		if !predicateFunc(v) {
			return false
		}
	}
	return true
}

// Any returns true if the given predicateFunc returns true for any element in slices.
// It will stop the iteration over slices when the predicateFunc returns true.
func Any[T any](slices []T, predicateFunc func(T) bool) (oneTrue bool) {
	for i := range slices {
		if predicateFunc(slices[i]) {
			return true
		}
	}
	return false
}

// Map returns a new slice of the type the mapFunc returns where there is one element
// in the result for every element in the input.
func Map[T any, N any](slices []T, mapFunc func(T) N) (mapped []N) {
	mapped = make([]N, len(slices))
	for i, v := range slices {
		mapped[i] = mapFunc(v)
	}
	return
}

// Reduce reduces all of the elements in slices of type T into a single value of type R by
// calling the reduce func for each element in slices. For the first element, the reduceFunc
// receives the initial value as its previous value, and for all subsequent calls it will get
// the value it returned in the previous iteration.
func Reduce[T any, R any](slices []T, initial R, reduceFunc func(val T, previous R) R) R {
	previous := initial
	for _, v := range slices {
		previous = reduceFunc(v, previous)
	}
	return previous
}

// Index returns the index of the element for which the predicate function returns true.
// If no call to the predicate function returned true, a value < 0 is returned.
func Index[T any](slices []T, predicateFunc func(T) bool) int {
	for i, v := range slices {
		if predicateFunc(v) {
			return i
		}
	}
	return -1
}

// Find returns a pointer to the element in slices for which the predicate function returns true.
// If no call to the predicate function returned true, nil is returned.
func Find[T any](slices []T, predicateFunc func(T) bool) *T {
	for i, v := range slices {
		if predicateFunc(v) {
			return &slices[i]
		}
	}
	return nil
}

// PtrToLastSlice returns a pointer to the last element in the given slice.
func PtrToLastSlice[T any](slices []T) *T {
	return &slices[len(slices)-1]
}
