package vectorArray

import "errors"

type VectorArray[T any] struct {
	items                    []T
	vector, length, capacity int
	zeroValue                T
}

func NewVectorArray[T any](vector int) *VectorArray[T] {
	sa := new(VectorArray[T])
	sa.items = make([]T, vector)
	sa.vector = vector
	sa.capacity = vector
	return sa
}

func (sa *VectorArray[T]) Length() int {
	return sa.length
}

func (sa *VectorArray[T]) Capacity() int {
	return sa.capacity
}

func (sa *VectorArray[T]) Get(index int) (T, error) {
	if index < 0 || index >= sa.length {
		return sa.zeroValue, errors.New("invalid index")
	}

	return sa.items[index], nil
}

func (sa *VectorArray[T]) Push(item T) error {
	if sa.length == sa.capacity {
		sa.resize()
	}

	sa.items[sa.length] = item
	sa.length++

	return nil
}

func (sa *VectorArray[T]) Add(item T, index int) error {
	if index < 0 || index >= sa.length {
		return errors.New("invalid index")
	}

	if sa.length == sa.capacity {
		sa.resize()
	}

	for i := index; i < sa.length; i++ {
		sa.items[i+1] = sa.items[i]
	}

	sa.items[index] = item
	sa.length++

	return nil
}

func (sa *VectorArray[T]) Pop() (T, error) {
	if sa.length == 0 {
		return sa.zeroValue, errors.New("invalid index")
	}

	res := sa.items[sa.length-1]
	sa.length--

	return res, nil
}

func (sa *VectorArray[T]) Remove(index int) (T, error) {
	if index < 0 || index >= sa.length {
		return sa.zeroValue, errors.New("invalid index")
	}

	res := sa.items[index]
	for i := index; i < sa.length-1; i++ {
		sa.items[i] = sa.items[i+1]
	}
	sa.length--

	return res, nil
}

func (sa *VectorArray[T]) resize() {
	newItems := make([]T, sa.capacity+sa.vector)
	for i, item := range sa.items {
		newItems[i] = item
	}
	sa.items = newItems
	sa.capacity += sa.vector
}
