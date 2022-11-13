package factorArray

import "errors"

type FactorArray[T any] struct {
	items                    []T
	factor, length, capacity int
	zeroValue                T
}

func NewFactorArray[T any]() *FactorArray[T] {
	sa := new(FactorArray[T])
	sa.items = make([]T, 0)
	sa.factor = 2
	return sa
}

func (sa *FactorArray[T]) Length() int {
	return sa.length
}

func (sa *FactorArray[T]) Capacity() int {
	return sa.capacity
}

func (sa *FactorArray[T]) Get(index int) (T, error) {
	if index < 0 || index >= sa.length {
		return sa.zeroValue, errors.New("invalid index")
	}

	return sa.items[index], nil
}

func (sa *FactorArray[T]) Push(item T) error {
	if sa.length == sa.capacity {
		sa.resize()
	}

	sa.items[sa.length] = item
	sa.length++

	return nil
}

func (sa *FactorArray[T]) Add(item T, index int) error {
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

func (sa *FactorArray[T]) Pop() (T, error) {
	if sa.length == 0 {
		return sa.zeroValue, errors.New("invalid index")
	}

	res := sa.items[sa.length-1]
	sa.length--

	return res, nil
}

func (sa *FactorArray[T]) Remove(index int) (T, error) {
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

func (sa *FactorArray[T]) resize() {
	newItems := make([]T, sa.capacity*sa.factor+1)
	for i, item := range sa.items {
		newItems[i] = item
	}
	sa.items = newItems
	sa.capacity = len(newItems)
}
