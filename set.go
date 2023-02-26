package gs

// Set is set of unique comparable values.
type Set[T comparable] struct {
	values map[T]struct{}
}

// NewSet returns initialized Set structure.
func NewSet[T comparable]() Set[T] {
	return Set[T]{
		values: map[T]struct{}{},
	}
}

// Add appends element to set.
func (set Set[T]) Add(value T) Set[T] {
	set.values[value] = struct{}{}

	return set
}

// Remove deletes element from set.
func (set Set[T]) Remove(value T) Set[T] {
	delete(set.values, value)

	return set
}

// Contains checks element is exists in set.
func (set Set[T]) Contains(value T) bool {
	_, ok := set.values[value]

	return ok
}

// Len returns count of set elements.
func (set Set[T]) Len() int {
	return len(set.values)
}

// Values returns slice of set elements.
func (set Set[T]) Values() []T {
	result := make([]T, 0, len(set.values))
	for k := range set.values {
		result = append(result, k)
	}

	return result
}
