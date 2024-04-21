package lang

import (
	"errors"
	"github.com/pythonik/golang-commons/testutil"
	"testing"
)

func TestEmpty(t *testing.T) {
	empty := Empty[string]()
	if empty.Error() == nil {
		t.Fail()
	}
	testutil.Assert(!empty.IsPresent(), "should not be present")
	testutil.Assert(errors.Is(empty.Error(), ErrEmptyOptional), "expect empty optional error")

}

type ObjectForTest struct {
	name string
}

func TestOptional_OrElse(t *testing.T) {
	var o *ObjectForTest
	var defaultObj = &ObjectForTest{name: "default"}
	maybe := OptionalOf(o)
	value := maybe.OrElse(defaultObj)
	testutil.Assert(value.name == "default", "value should be default")
}
