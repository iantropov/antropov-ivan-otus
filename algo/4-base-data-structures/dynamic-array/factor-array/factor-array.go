package factorArray

import "errors"

type FactorArray[T any] struct {
	items                    []T
	factor, length, capacity int
	zeroValue                T
}

func NewFactorArray[T any]() *FactorArray[T] {
	fa := new(FactorArray[T])
	fa.Initialize()
	return fa
}

func (fa *FactorArray[T]) Initialize() {
	fa.items = make([]T, 0)
	fa.factor = 2
}

func (fa *FactorArray[T]) Length() int {
	return fa.length
}

func (fa *FactorArray[T]) Capacity() int {
	return fa.capacity
}

func (fa *FactorArray[T]) Get(index int) (T, error) {
	if index < 0 || index >= fa.length {
		return fa.zeroValue, errors.New("invalid index")
	}

	return fa.items[index], nil
}

func (fa *FactorArray[T]) Set(value T, index int) error {
	if index < 0 || index >= fa.length {
		return errors.New("invalid index")
	}

	fa.items[index] = value
	return nil
}

func (fa *FactorArray[T]) Push(item T) error {
	if fa.length == fa.capacity {
		fa.resize()
	}

	fa.items[fa.length] = item
	fa.length++

	return nil
}

func (fa *FactorArray[T]) Add(item T, index int) error {
	if index < 0 || index >= fa.length {
		return errors.New("invalid index")
	}

	if fa.length == fa.capacity {
		fa.resize()
	}

	for i := fa.length; i > index; i-- {
		fa.items[i] = fa.items[i-1]
	}

	fa.items[index] = item
	fa.length++

	return nil
}

func (fa *FactorArray[T]) Pop() (T, error) {
	if fa.length == 0 {
		return fa.zeroValue, errors.New("invalid index")
	}

	res := fa.items[fa.length-1]
	fa.length--

	return res, nil
}

func (fa *FactorArray[T]) Remove(index int) (T, error) {
	if index < 0 || index >= fa.length {
		return fa.zeroValue, errors.New("invalid index")
	}

	res := fa.items[index]
	for i := index; i < fa.length-1; i++ {
		fa.items[i] = fa.items[i+1]
	}
	fa.length--

	return res, nil
}

func (fa *FactorArray[T]) resize() {
	newItems := make([]T, fa.capacity*fa.factor+1)
	for i, item := range fa.items {
		newItems[i] = item
	}
	fa.items = newItems
	fa.capacity = len(newItems)
}
