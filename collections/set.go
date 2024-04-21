package collections

type Set[T comparable] struct {
	elements map[T]struct{} // using an empty struct which does not consume memory
}

// NewSet initializes and returns a new instance of a Set.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{elements: make(map[T]struct{})}
}

// Add adds an element to the set.
func (s *Set[T]) Add(element T) {
	s.elements[element] = struct{}{}
}

// Remove removes an element from the set.
func (s *Set[T]) Remove(element T) {
	delete(s.elements, element)
}

// Contains checks if an element is in the set.
func (s *Set[T]) Contains(element T) bool {
	_, exists := s.elements[element]
	return exists
}

// Size returns the number of elements in the set.
func (s *Set[T]) Size() int {
	return len(s.elements)
}

// Elements returns all the elements in the set as a slice.
func (s *Set[T]) Elements() []T {
	keys := make([]T, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}
