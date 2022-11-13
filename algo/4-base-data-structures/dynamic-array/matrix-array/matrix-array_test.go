package matrixArray

import (
	"testing"
)

func TestNewDynamicArray(t *testing.T) {
	array := NewMatrixArray[int](2)

	length := array.Length()
	if length != 0 {
		t.Errorf("Incorrect length - %d", length)
	}

	capacity := array.Capacity()
	if capacity != 0 {
		t.Errorf("Incorrect capacuty - %v", capacity)
	}
}

func TestPushToDynamicArray(t *testing.T) {
	array := NewMatrixArray[int](2)

	error := array.Push(1)
	if error != nil {
		t.Errorf("Failed to push 1 - %v", error)
	}

	length := array.Length()
	if length != 1 {
		t.Errorf("Incorrect length - %d", length)
	}

	capacity := array.Capacity()
	if capacity != 2 {
		t.Errorf("Incorrect capacuty - %v", capacity)
	}

	error = array.Push(2)
	if error != nil {
		t.Errorf("Failed to push 2 - %v", error)
	}

	length = array.Length()
	if length != 2 {
		t.Errorf("Incorrect length - %d", length)
	}

	capacity = array.Capacity()
	if capacity != 2 {
		t.Errorf("Incorrect capacuty - %v", capacity)
	}

	error = array.Push(3)
	if error != nil {
		t.Errorf("Failed to push 2 - %v", error)
	}

	length = array.Length()
	if length != 3 {
		t.Errorf("Incorrect length - %d", length)
	}

	capacity = array.Capacity()
	if capacity != 4 {
		t.Errorf("Incorrect capacuty - %v", capacity)
	}
}

func TestPopFromDynamicArray(t *testing.T) {
	array := NewMatrixArray[int](2)

	array.Push(1)
	array.Push(2)

	item, error := array.Pop()
	if item != 2 {
		t.Errorf("Invalid value - %v", item)
	}
	if error != nil {
		t.Errorf("Failed to pop 2 - %v", error)
	}

	length := array.Length()
	if length != 1 {
		t.Errorf("Incorrect length - %d", length)
	}

	capacity := array.Capacity()
	if capacity != 2 {
		t.Errorf("Incorrect capacuty - %v", capacity)
	}

	item, error = array.Pop()
	if item != 1 {
		t.Errorf("Invalid value - %v", item)
	}
	if error != nil {
		t.Errorf("Failed to push 1 - %v", error)
	}

	length = array.Length()
	if length != 0 {
		t.Errorf("Incorrect length - %d", length)
	}

	capacity = array.Capacity()
	if capacity != 2 {
		t.Errorf("Incorrect capacuty - %v", capacity)
	}
}

func TestPushAfterPopForDynamicArray(t *testing.T) {
	array := NewMatrixArray[int](2)

	array.Push(1)
	array.Push(2)
	array.Pop()

	error := array.Push(3)
	if error != nil {
		t.Errorf("Failed to push 3 - %v", error)
	}

	length := array.Length()
	if length != 2 {
		t.Errorf("Incorrect length - %d", length)
	}

	capacity := array.Capacity()
	if capacity != 2 {
		t.Errorf("Incorrect capacuty - %v", capacity)
	}

	array.Pop()
	array.Pop()

	error = array.Push(4)
	if error != nil {
		t.Errorf("Failed to push 4 - %v", error)
	}

	length = array.Length()
	if length != 1 {
		t.Errorf("Incorrect length - %d", length)
	}

	capacity = array.Capacity()
	if capacity != 2 {
		t.Errorf("Incorrect capacuty - %v", capacity)
	}
}

func TestAddToDynamicArray(t *testing.T) {
	array := NewMatrixArray[int](2)

	array.Push(1)
	array.Push(2)

	error := array.Add(3, 1)
	if error != nil {
		t.Errorf("Failed to add 3 - %v", error)
	}

	length := array.Length()
	if length != 3 {
		t.Errorf("Incorrect length - %d", length)
	}

	capacity := array.Capacity()
	if capacity != 4 {
		t.Errorf("Incorrect capacuty - %v", capacity)
	}

	item, error := array.Get(0)
	if item != 1 {
		t.Errorf("Incorrect value - %d", item)
	}
	if error != nil {
		t.Errorf("Failed to get 0 - %v", error)
	}

	item, error = array.Get(1)
	if item != 3 {
		t.Errorf("Incorrect value - %d", item)
	}
	if error != nil {
		t.Errorf("Failed to get 1 - %v", error)
	}

	item, error = array.Get(2)
	if item != 2 {
		t.Errorf("Incorrect value - %d", item)
	}
	if error != nil {
		t.Errorf("Failed to get 2 - %v", error)
	}
}

func TestRemoveFromDynamicArray(t *testing.T) {
	array := NewMatrixArray[int](2)

	array.Push(1)
	array.Push(2)
	array.Push(3)

	item, error := array.Remove(1)
	if item != 2 {
		t.Errorf("Incorrect value - %d", item)
	}
	if error != nil {
		t.Errorf("Failed to remove 1 - %v", error)
	}

	length := array.Length()
	if length != 2 {
		t.Errorf("Incorrect length - %d", length)
	}

	capacity := array.Capacity()
	if capacity != 4 {
		t.Errorf("Incorrect capacuty - %v", capacity)
	}

	item, error = array.Get(0)
	if item != 1 {
		t.Errorf("Incorrect value - %d", item)
	}
	if error != nil {
		t.Errorf("Failed to get 0 - %v", error)
	}

	item, error = array.Get(1)
	if item != 3 {
		t.Errorf("Incorrect value - %d", item)
	}
	if error != nil {
		t.Errorf("Failed to get 1 - %v", error)
	}
}
