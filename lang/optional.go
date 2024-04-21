package lang

import "errors"

var ErrEmptyOptional = errors.New("empty optional")

type Optional[T any] struct {
	value T
	err   error
}

func (o Optional[T]) IsPresent() bool {
	return o.err == nil
}

func (o Optional[T]) OrElse(other T) T {
	if o.err != nil {
		return other
	}
	return o.value
}

func (o Optional[T]) Value() T {
	return o.value
}

func (o Optional[T]) Error() error {
	return o.err
}

func OptionalOf[T any](value T) Optional[T] {
	return Optional[T]{value: value}
}

func Empty[T any]() Optional[T] {
	return Optional[T]{err: ErrEmptyOptional}
}
