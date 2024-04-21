package collections

import (
	"github.com/pythonik/golang-commons/testutil"
	"testing"
)

func TestBiMap_GetByKey_GetByValue(t *testing.T) {
	biMap := NewBiMap[int, string]()
	biMap.Put(1, "one")
	biMap.Put(2, "two")
	biMap.Put(3, "three")
	value1Str, _ := biMap.GetByKey(1)
	testutil.Assert(value1Str == "one", "value should be one")
	value1, _ := biMap.GetByValue("one")
	testutil.Assert(value1 == 1, "value should be 1")
	value2Str, _ := biMap.GetByKey(2)
	testutil.Assert(value2Str == "two", "value should be two")
	value2, _ := biMap.GetByValue("two")
	testutil.Assert(value2 == 2, "value should be 2")
	value3Str, _ := biMap.GetByKey(3)
	testutil.Assert(value3Str == "three", "value should be three")
	value3, _ := biMap.GetByValue("three")
	testutil.Assert(value3 == 3, "value should be three")
}

func TestImmutableList_Elements(t *testing.T) {
	list := NewImmuList[int](1, 2, 3)
	elements := list.Elements()
	testutil.Assert(len(elements) == 3, "size should be 3")
	testutil.Assert(elements[0] == 1, "element should be 1")
	testutil.Assert(elements[1] == 2, "element should be 2")
	testutil.Assert(elements[2] == 3, "element should be 3")
}

func TestImmutableList_is_Immutable(t *testing.T) {
	list := NewImmuList[int](1, 2, 3)
	elements := list.Elements()
	originalValue := elements[0]
	elements[0] = 4
	v, _ := list.ValueAt(0)
	testutil.Assert(*v == originalValue, "element should be 1")
}
