package strutils

import "testing"

func TestIsEmpty(t *testing.T) {
	if !IsEmpty("") {
		t.Fail()
	}
	if IsEmpty("a") {
		t.Fail()
	}
}

func TestIsNotEmpty(t *testing.T) {
	if IsNotEmpty("") {
		t.Fail()
	}
	if !IsNotEmpty("a") {
		t.Fail()
	}
}
