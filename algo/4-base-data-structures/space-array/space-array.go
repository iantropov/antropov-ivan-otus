package spaceArray

import (
	singleArray "base-data-structures/dynamic-array/single-array"
	vectorArray "base-data-structures/dynamic-array/vector-array"
	"errors"
)

type SpaceArray[T any] struct {
	items                           singleArray.SingleArray[*vectorArray.VectorArray[T]]
	length, rowLimit, rowThreshhold int
	zeroValue                       T
}

func NewSpaceArray[T any](rowThreshhold int) *SpaceArray[T] {
	newSpaceArray := new(SpaceArray[T])
	newSpaceArray.rowThreshhold = rowThreshhold
	newSpaceArray.rowLimit = rowThreshhold * 2
	return newSpaceArray
}

func (sa *SpaceArray[T]) Length() int {
	return sa.length
}

func (sa *SpaceArray[T]) Capacity() int {
	return sa.items.Length() * (sa.rowLimit + 1)
}

func (sa *SpaceArray[T]) Push(item T) error {
	if sa.items.Length() == 0 {
		sa.appendRow()
	}

	lastRow, error := sa.items.Get(sa.items.Length() - 1)
	if error != nil {
		return error
	}

	if lastRow.Length() >= sa.rowThreshhold {
		lastRow = sa.appendRow()
	}

	lastRow.Push(item)
	sa.length++

	return nil
}

func (sa *SpaceArray[T]) Pop() (T, error) {
	if sa.length == 0 {
		return sa.zeroValue, errors.New("empty array")
	}

	for i := sa.items.Length() - 1; i >= 0; i-- {
		item, error := sa.items.Get(i)
		if error != nil {
			return sa.zeroValue, error
		}

		if item.Length() > 0 {
			res, error := item.Pop()
			if error != nil {
				return sa.zeroValue, error
			}

			sa.length--

			return res, nil
		}
	}

	return sa.zeroValue, errors.New("array with invalid length")
}

func (sa *SpaceArray[T]) Get(index int) (T, error) {
	if sa.length <= index {
		return sa.zeroValue, errors.New("invalid index")
	}

	slidingIndex := index
	for i := 0; i < sa.items.Length(); i++ {
		item, error := sa.items.Get(i)
		if error != nil {
			return sa.zeroValue, error
		}

		if slidingIndex >= item.Length() {
			slidingIndex -= item.Length()
		} else {
			res, error := item.Get(slidingIndex)
			if error != nil {
				return sa.zeroValue, error
			}

			return res, nil
		}
	}

	return sa.zeroValue, errors.New("array with invalid structure")
}

func (sa *SpaceArray[T]) Set(value T, index int) error {
	if sa.length <= index {
		return errors.New("invalid index")
	}

	slidingIndex := index
	for i := 0; i < sa.items.Length(); i++ {
		item, error := sa.items.Get(i)
		if error != nil {
			return error
		}

		if slidingIndex >= item.Length() {
			slidingIndex -= item.Length()
		} else {
			return item.Set(value, slidingIndex)
		}
	}

	return errors.New("array with invalid structure")
}

func (sa *SpaceArray[T]) Add(value T, index int) error {
	if sa.length <= index {
		return errors.New("invalid index")
	}

	slidingIndex := index
	for i := 0; i < sa.items.Length(); i++ {
		item, error := sa.items.Get(i)
		if error != nil {
			return error
		}

		if slidingIndex >= sa.rowLimit {
			slidingIndex -= sa.rowLimit
			continue
		}

		itemLength := item.Length()
		if itemLength < sa.rowLimit {
			if slidingIndex == item.Length() {
				error = item.Push(value)
			} else {
				error = item.Add(value, slidingIndex)
			}
			if error != nil {
				return error
			}

			sa.length++
			return nil
		}

		error = sa.moveElementsToTheRightFrom(slidingIndex, i)
		if error != nil {
			return error
		}

		error = item.Set(value, slidingIndex)
		if error != nil {
			return error
		}

		sa.length++

		return nil
	}

	return errors.New("array with invalid structure")
}

func (sa *SpaceArray[T]) Remove(index int) (T, error) {
	if sa.length <= index {
		return sa.zeroValue, errors.New("invalid index")
	}

	slidingIndex := index
	for i := 0; i < sa.items.Length(); i++ {
		item, error := sa.items.Get(i)
		if error != nil {
			return sa.zeroValue, error
		}

		if slidingIndex >= item.Length() {
			slidingIndex -= item.Length()
			continue
		}

		value, error := item.Remove(slidingIndex)
		if error != nil {
			return sa.zeroValue, error
		}

		sa.length--
		return value, nil
	}

	return sa.zeroValue, errors.New("array with invalid structure")
}

func (sa *SpaceArray[T]) appendRow() *vectorArray.VectorArray[T] {
	newRow := new(vectorArray.VectorArray[T])
	newRow.Initialize(sa.rowLimit + 1) // 1 - a reserve for 'moveElementsToTheRightFrom' operation
	sa.items.Push(newRow)
	return newRow
}

func (sa *SpaceArray[T]) moveElementsToTheRightFrom(valueInItemIndex, itemIndex int) error {
	item, error := sa.items.Get(itemIndex)
	if error != nil {
		return error
	}

	error = item.Add(sa.zeroValue, valueInItemIndex)
	if error != nil {
		return error
	}

	relay := sa.zeroValue
	for i := itemIndex + 1; i < sa.items.Length(); i++ {
		item, error := sa.items.Get(i)
		if error != nil {
			return error
		}

		previousItem, error := sa.items.Get(i - 1)
		if error != nil {
			return error
		}

		relay, error = previousItem.Remove(previousItem.Length() - 1)
		if error != nil {
			return error
		}

		item.Add(relay, 0)
	}

	return nil
}
