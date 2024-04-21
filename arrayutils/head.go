package arrayutils

import (
	"github.com/pythonik/golang-commons/lang"
)

func Head[T any](s []T) lang.Optional[T] {
	if IsEmpty(s) {
		return lang.Empty[T]()
	}
	return lang.OptionalOf(&s[0])
}

func Tail[T any](s []T) lang.Optional[T] {
	if IsEmpty(s) {
		return lang.Empty[T]()
	}
	return lang.OptionalOf(&s[len(s)-1])
}
