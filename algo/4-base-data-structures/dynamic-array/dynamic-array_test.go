package dynamicArray

import (
	singleArray "base-data-structures/dynamic-array/single-array"
	"strings"
	"testing"
)

func Test_stacks(t *testing.T) {
	cases := []struct {
		newStack func() DynamicArray[int]
		name     string
	}{
		{func() DynamicArray[int] { return singleArray.NewSingleArray[int]() }, "Single Array"},
	}
	for _, tc := range cases {
		t.Run(strings.Join([]string{tc.name, "testNewDynamicArray"}, " : "), func(t *testing.T) {
			testNewDynamicArray(tc.newStack(), t)
		})
		t.Run(strings.Join([]string{tc.name, "testPushToDynamicArray"}, " : "), func(t *testing.T) {
			testPushToDynamicArray(tc.newStack(), t)
		})
		t.Run(strings.Join([]string{tc.name, "testPopFromDynamicArray"}, " : "), func(t *testing.T) {
			testPopFromDynamicArray(tc.newStack(), t)
		})
		t.Run(strings.Join([]string{tc.name, "testPushAfterPopForDynamicArray"}, " : "), func(t *testing.T) {
			testPushAfterPopForDynamicArray(tc.newStack(), t)
		})
		t.Run(strings.Join([]string{tc.name, "testAddToDynamicArray"}, " : "), func(t *testing.T) {
			testAddToDynamicArray(tc.newStack(), t)
		})
		t.Run(strings.Join([]string{tc.name, "testRemoveFromDynamicArray"}, " : "), func(t *testing.T) {
			testRemoveFromDynamicArray(tc.newStack(), t)
		})
	}
}

func testNewDynamicArray(array DynamicArray[int], t *testing.T) {
	length := array.Length()
	if length != 0 {
		t.Errorf("Incorrect length - %d", length)
	}

	capacity := array.Capacity()
	if capacity != 0 {
		t.Errorf("Incorrect capacuty - %v", capacity)
	}
}

func testPushToDynamicArray(array DynamicArray[int], t *testing.T) {
	error := array.Push(1)
	if error != nil {
		t.Errorf("Failed to push 1 - %v", error)
	}

	length := array.Length()
	if length != 1 {
		t.Errorf("Incorrect length - %d", length)
	}

	capacity := array.Capacity()
	if capacity != 1 {
		t.Errorf("Incorrect capacuty - %v", capacity)
	}

	error = array.Push(2)
	if error != nil {
		t.Errorf("Failed to push 1 - %v", error)
	}

	length = array.Length()
	if length != 2 {
		t.Errorf("Incorrect length - %d", length)
	}

	capacity = array.Capacity()
	if capacity != 2 {
		t.Errorf("Incorrect capacuty - %v", capacity)
	}
}

func testPopFromDynamicArray(array DynamicArray[int], t *testing.T) {
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

func testPushAfterPopForDynamicArray(array DynamicArray[int], t *testing.T) {
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

func testAddToDynamicArray(array DynamicArray[int], t *testing.T) {
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
	if capacity != 3 {
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

func testRemoveFromDynamicArray(array DynamicArray[int], t *testing.T) {
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
	if capacity != 3 {
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
