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

func TestStack_Pop(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	value, _ := stack.Pop()
	testutil.Assert(value == 3, "value should be 3")
	value, _ = stack.Pop()
	testutil.Assert(value == 2, "value should be 2")
	value, _ = stack.Pop()
	testutil.Assert(value == 1, "value should be 1")
}

func TestStack_Size(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	size := stack.Size()
	testutil.Assert(size == 3, "size should be 3")
}

func TestStack_IsEmpty(t *testing.T) {
	stack := NewStack[int]()
	testutil.Assert(stack.IsEmpty(), "stack should be empty")
	stack.Push(1)
	testutil.Assert(!stack.IsEmpty(), "stack should not be empty")
}

func TestStack_Peek(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	value, _ := stack.Peek()
	testutil.Assert(value == 3, "value should be 3")
	value, _ = stack.Peek()
	testutil.Assert(value == 3, "value should be 3")
}

func TestSet_Add(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	testutil.Assert(set.Size() == 3, "size should be 3")
}

func TestSet_Contains(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	testutil.Assert(set.Contains(1), "should contain 1")
	testutil.Assert(set.Contains(2), "should contain 2")
	testutil.Assert(set.Contains(3), "should contain 3")
}

func TestSet_Remove(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Remove(1)
	testutil.Assert(set.Size() == 2, "size should be 2")
}

func TestSet_Elements(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	elements := set.Elements()
	testutil.Assert(len(elements) == 3, "size should be 3")
	testutil.Assert(elements[0] == 1, "element should be 1")
	testutil.Assert(elements[1] == 2, "element should be 2")
	testutil.Assert(elements[2] == 3, "element should be 3")
}
