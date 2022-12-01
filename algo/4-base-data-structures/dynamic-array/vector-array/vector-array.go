package vectorArray

import "errors"

type VectorArray[T any] struct {
	items                    []T
	vector, length, capacity int
	zeroValue                T
}

func NewVectorArray[T any](vector int) *VectorArray[T] {
	va := new(VectorArray[T])
	va.Initialize(vector)
	return va
}

func (va *VectorArray[T]) Initialize(vector int) {
	va.items = make([]T, vector)
	va.vector = vector
	va.capacity = vector
}

func (va *VectorArray[T]) Length() int {
	return va.length
}

func (va *VectorArray[T]) Capacity() int {
	return va.capacity
}

func (va *VectorArray[T]) Get(index int) (T, error) {
	if index < 0 || index >= va.length {
		return va.zeroValue, errors.New("invalid index")
	}

	return va.items[index], nil
}

func (va *VectorArray[T]) Set(value T, index int) error {
	if index < 0 || index >= va.length {
		return errors.New("invalid index")
	}

	va.items[index] = value
	return nil
}

func (va *VectorArray[T]) Push(item T) error {
	if va.length == va.capacity {
		va.resize()
	}

	va.items[va.length] = item
	va.length++

	return nil
}

func (va *VectorArray[T]) Add(item T, index int) error {
	if index < 0 || index >= va.length {
		return errors.New("invalid index")
	}

	if va.length == va.capacity {
		va.resize()
	}

	for i := va.length; i > index; i-- {
		va.items[i] = va.items[i-1]
	}

	va.items[index] = item
	va.length++

	return nil
}

func (va *VectorArray[T]) Pop() (T, error) {
	if va.length == 0 {
		return va.zeroValue, errors.New("invalid index")
	}

	res := va.items[va.length-1]
	va.length--

	return res, nil
}

func (va *VectorArray[T]) Remove(index int) (T, error) {
	if index < 0 || index >= va.length {
		return va.zeroValue, errors.New("invalid index")
	}

	res := va.items[index]
	for i := index; i < va.length-1; i++ {
		va.items[i] = va.items[i+1]
	}
	va.length--
	va.items[va.length] = va.zeroValue

	return res, nil
}

func (va *VectorArray[T]) resize() {
	newItems := make([]T, va.capacity+va.vector)
	for i, item := range va.items {
		newItems[i] = item
	}
	va.items = newItems
	va.capacity += va.vector
}
