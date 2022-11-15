package spaceArray

import "testing"

func TestSpaceArrayCreation(t *testing.T) {
	sa := NewSpaceArray[int](2)

	length := sa.Length()
	if length != 0 {
		t.Errorf("Invalid length - %d", length)
	}

	capacity := sa.Capacity()
	if capacity != 0 {
		t.Errorf("Invalid capacity - %d", capacity)
	}
}

func TestPushIntoSpaceArray(t *testing.T) {
	sa := NewSpaceArray[int](2)

	error := sa.Push(1)
	if error != nil {
		t.Errorf("Failed to push 1 - %v", error)
	}
	length := sa.Length()
	if length != 1 {
		t.Errorf("Invalid length - %d", length)
	}
	capacity := sa.Capacity()
	if capacity != 5 {
		t.Errorf("Invalid capacity - %d", capacity)
	}

	error = sa.Push(2)
	if error != nil {
		t.Errorf("Failed to push 2 - %v", error)
	}
	length = sa.Length()
	if length != 2 {
		t.Errorf("Invalid length - %d", length)
	}
	capacity = sa.Capacity()
	if capacity != 5 {
		t.Errorf("Invalid capacity - %d", capacity)
	}

	error = sa.Push(3)
	if error != nil {
		t.Errorf("Failed to push 3 - %v", error)
	}
	length = sa.Length()
	if length != 3 {
		t.Errorf("Invalid length - %d", length)
	}
	capacity = sa.Capacity()
	if capacity != 10 {
		t.Errorf("Invalid capacity - %d", capacity)
	}
}

func TestPopFromSpaceArray(t *testing.T) {
	sa := NewSpaceArray[int](2)
	sa.Push(1)
	sa.Push(2)
	sa.Push(3)

	value, error := sa.Pop()
	if value != 3 {
		t.Errorf("Invalid value - %d", value)
	}
	if error != nil {
		t.Errorf("Failed to pop 3 - %v", error)
	}
	length := sa.Length()
	if length != 2 {
		t.Errorf("Invalid length - %d", length)
	}
	capacity := sa.Capacity()
	if capacity != 10 {
		t.Errorf("Invalid capacity - %d", capacity)
	}

	value, error = sa.Pop()
	if value != 2 {
		t.Errorf("Invalid value - %d", value)
	}
	if error != nil {
		t.Errorf("Failed to pop 2 - %v", error)
	}
	length = sa.Length()
	if length != 1 {
		t.Errorf("Invalid length - %d", length)
	}
	capacity = sa.Capacity()
	if capacity != 10 {
		t.Errorf("Invalid capacity - %d", capacity)
	}

	value, error = sa.Pop()
	if value != 1 {
		t.Errorf("Invalid value - %d", value)
	}
	if error != nil {
		t.Errorf("Failed to pop 1 - %v", error)
	}
	length = sa.Length()
	if length != 0 {
		t.Errorf("Invalid length - %d", length)
	}
	capacity = sa.Capacity()
	if capacity != 10 {
		t.Errorf("Invalid capacity - %d", capacity)
	}
}

func TestAddAfterPushInSpaceArray(t *testing.T) {
	sa := NewSpaceArray[int](2)
	sa.Push(1)
	sa.Push(2)
	sa.Push(6)

	error := sa.Add(5, 2)
	if error != nil {
		t.Errorf("Failed to add 5 - %v", error)
	}
	length := sa.Length()
	if length != 4 {
		t.Errorf("Invalid length - %d", length)
	}
	capacity := sa.Capacity()
	if capacity != 10 {
		t.Errorf("Invalid capacity - %d", capacity)
	}

	error = sa.Add(3, 2)
	if error != nil {
		t.Errorf("Failed to add 3 - %v", error)
	}
	length = sa.Length()
	if length != 5 {
		t.Errorf("Invalid length - %d", length)
	}
	capacity = sa.Capacity()
	if capacity != 10 {
		t.Errorf("Invalid capacity - %d", capacity)
	}

	error = sa.Add(4, 3)
	if error != nil {
		t.Errorf("Failed to add 4 - %v", error)
	}
	length = sa.Length()
	if length != 6 {
		t.Errorf("Invalid length - %d", length)
	}
	capacity = sa.Capacity()
	if capacity != 10 {
		t.Errorf("Invalid capacity - %d", capacity)
	}

	value, _ := sa.Get(0)
	if value != 1 {
		t.Errorf("Invalid value at 0 - %d", value)
	}
	value, _ = sa.Get(1)
	if value != 2 {
		t.Errorf("Invalid value at 1 - %d", value)
	}
	value, _ = sa.Get(2)
	if value != 3 {
		t.Errorf("Invalid value at 2 - %d", value)
	}
	value, _ = sa.Get(3)
	if value != 4 {
		t.Errorf("Invalid value at 3 - %d", value)
	}
	value, _ = sa.Get(4)
	if value != 5 {
		t.Errorf("Invalid value at 4 - %d", value)
	}
	value, _ = sa.Get(5)
	if value != 6 {
		t.Errorf("Invalid value at 5 - %d", value)
	}
}

func TestRemoveAfterPushInSpaceArray(t *testing.T) {
	sa := NewSpaceArray[int](2)
	sa.Push(1)
	sa.Push(2)
	sa.Add(2, 1)
	sa.Add(2, 1)
	sa.Push(3)

	value, error := sa.Remove(2)
	if value != 2 {
		t.Errorf("Invalid value - %d", value)
	}
	if error != nil {
		t.Errorf("Failed to remove at 2 - %v", error)
	}
	length := sa.Length()
	if length != 4 {
		t.Errorf("Invalid length - %d", length)
	}
	capacity := sa.Capacity()
	if capacity != 10 {
		t.Errorf("Invalid capacity - %d", capacity)
	}

	value, error = sa.Remove(2)
	if value != 2 {
		t.Errorf("Invalid value - %d", value)
	}
	if error != nil {
		t.Errorf("Failed to remove at 2 - %v", error)
	}
	length = sa.Length()
	if length != 3 {
		t.Errorf("Invalid length - %d", length)
	}
	capacity = sa.Capacity()
	if capacity != 10 {
		t.Errorf("Invalid capacity - %d", capacity)
	}

	value, _ = sa.Get(0)
	if value != 1 {
		t.Errorf("Invalid value at 0 - %d", value)
	}
	value, _ = sa.Get(1)
	if value != 2 {
		t.Errorf("Invalid value at 1 - %d", value)
	}
	value, _ = sa.Get(2)
	if value != 3 {
		t.Errorf("Invalid value at 2 - %d", value)
	}
}
