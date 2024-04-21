package collections

type ImmutableList[T any] struct {
	elements []T
	size     int
}

func NewImmuList[T any](elements ...T) *ImmutableList[T] {
	newElements := make([]T, len(elements))
	copy(newElements, elements)
	return &ImmutableList[T]{elements: newElements, size: len(newElements)}
}

func (l *ImmutableList[T]) Elements() []T {
	copied := make([]T, len(l.elements))
	copy(copied, l.elements)
	return copied
}

func (l *ImmutableList[T]) ValueAt(index int) (*T, bool) {
	if index < 0 || index >= l.size {
		return nil, false
	}
	return &l.elements[index], true
}

func (l *ImmutableList[T]) Size() int {
	return l.size
}
