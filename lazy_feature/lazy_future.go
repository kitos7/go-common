package lazy_feature

import "sync"

// LazyFuture provides a mechanism for lazily computing a value only when it is needed.
// The computation happens only once on the first Get() call.
type LazyFuture[T any] struct {
	compute func() (T, error) // Function that will be executed to compute the value
	value   T                 // Cached value after computation
	err     error             // Cached error after computation
	once    sync.Once         // Ensures the compute function is only called once
}

// NewLazyFuture creates a new LazyFuture that will use the provided compute function
// to calculate its value when Get() is first called.
func NewLazyFuture[T any](compute func() (T, error)) *LazyFuture[T] {
	return &LazyFuture[T]{
		compute: compute,
	}
}

// Get returns the computed value and any error that occurred during computation.
// The computation happens only once on the first call, with subsequent calls
// returning the cached result.
func (f *LazyFuture[T]) Get() (T, error) {
	f.once.Do(func() {
		f.value, f.err = f.compute()
	})
	return f.value, f.err
}
