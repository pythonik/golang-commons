package testutil

func Assert(condition bool, message string) {
	if !condition {
		panic("Assertion failed: " + message)
	}
}
