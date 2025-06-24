package slices

// Set represents a generic set data structure using maps
type Set[T comparable] map[T]struct{}

// NewSet creates and initializes a new empty Set
func NewSet[T comparable]() Set[T] {
	return make(map[T]struct{})
}

// Map applies the mapper function to each element in the slice and returns a new slice
// containing the results
func Map[T any, R any](slice []T, mapper func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = mapper(v)
	}
	return result
}

// Filter returns a new slice containing only the elements that satisfy the predicate function
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// ToSet converts a slice to a Set
func ToSet[T comparable](slice []T) Set[T] {
	set := make(Set[T])
	for _, v := range slice {
		set[v] = struct{}{}
	}
	return set
}

// Has checks if the given value exists in the Set
func (s Set[T]) Has(val T) bool {
	_, ok := s[val]
	return ok
}

// Set adds the given value to the Set
func (s Set[T]) Set(val T) {
	s[val] = struct{}{}
}

// Delete removes the given value from the Set
func (s Set[T]) Delete(val T) {
	delete(s, val)
}

// ToSlice converts the Set to a slice
func (s Set[T]) ToSlice() []T {
	r := make([]T, 0, len(s))
	for v := range s {
		r = append(r, v)
	}
	return r
}

// Contains checks if a slice contains the given value
func Contains[T comparable](slice []T, val T) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

// NotContains checks if a slice does not contain the given value
func NotContains[T comparable](slice []T, val T) bool {
	return !Contains(slice, val)
}

// ToMap converts a slice to a map using the keyMapper function to generate
// key-value pairs for each element in the slice
func ToMap[K comparable, E, V any](slice []E, keyMapper func(E) (K, V)) map[K]V {
	m := make(map[K]V)
	for _, v := range slice {
		k, val := keyMapper(v)
		m[k] = val
	}
	return m
}
