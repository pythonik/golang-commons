package arrayutils

func IsEmpty[T any](s []T) bool {
	if s == nil {
		return true
	}
	return len(s) == 0
}

func IsNotEmpty[T any](s []T) bool {
	return !IsEmpty(s)
}
