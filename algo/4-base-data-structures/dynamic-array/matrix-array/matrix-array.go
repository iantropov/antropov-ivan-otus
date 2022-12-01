package matrixArray

import "errors"

type MatrixArray[T any] struct {
	items                      [][]T
	rowSize, rowsCount, length int
	zeroValue                  T
}

func NewMatrixArray[T any](rowSize int) *MatrixArray[T] {
	ma := new(MatrixArray[T])
	ma.items = make([][]T, 0)
	ma.rowSize = rowSize
	return ma
}

func (ma *MatrixArray[T]) Length() int {
	return ma.length
}

func (ma *MatrixArray[T]) Capacity() int {
	return ma.rowsCount * ma.rowSize
}

func (ma *MatrixArray[T]) Get(index int) (T, error) {
	if index < 0 || index >= ma.length {
		return ma.zeroValue, errors.New("invalid index")
	}

	return ma.get(index), nil
}

func (ma *MatrixArray[T]) Push(item T) error {
	if ma.length == ma.Capacity() {
		ma.resize()
	}

	ma.set(item, ma.length)
	ma.length++

	return nil
}

func (ma *MatrixArray[T]) Add(item T, index int) error {
	if index < 0 || index >= ma.length {
		return errors.New("invalid index")
	}

	if ma.length == ma.Capacity() {
		ma.resize()
	}

	for i := ma.length; i > index; i-- {
		ma.set(ma.get(i-1), i)
	}

	ma.set(item, index)
	ma.length++

	return nil
}

func (ma *MatrixArray[T]) Pop() (T, error) {
	if ma.length == 0 {
		return ma.zeroValue, errors.New("invalid index")
	}

	res := ma.get(ma.length - 1)
	ma.length--

	return res, nil
}

func (ma *MatrixArray[T]) Remove(index int) (T, error) {
	if index < 0 || index >= ma.length {
		return ma.zeroValue, errors.New("invalid index")
	}

	res := ma.get(index)
	for i := index; i < ma.length-1; i++ {
		ma.set(ma.get(i+1), i)
	}
	ma.length--
	ma.set(ma.zeroValue, ma.length)

	return res, nil
}

func (ma *MatrixArray[T]) get(index int) T {
	return ma.items[index/ma.rowSize][index%ma.rowSize]
}

func (ma *MatrixArray[T]) set(item T, index int) {
	ma.items[index/ma.rowSize][index%ma.rowSize] = item
}

func (ma *MatrixArray[T]) resize() {
	newItems := make([][]T, ma.rowsCount+1)
	for i, item := range ma.items {
		newItems[i] = item
	}
	ma.items = newItems
	ma.items[ma.rowsCount] = make([]T, ma.rowSize)
	ma.rowsCount++
}
