package spaceArray

import (
	factorArray "base-data-structures/dynamic-array/factor-array"
	"errors"
)

type SpaceArray[T any] struct {
	items                           factorArray.FactorArray[*factorArray.FactorArray[T]]
	length, rowLimit, rowThreshhold int
	zeroValue                       T
}

func NewSpaceArray[T any](rowThreshhold int) *SpaceArray[T] {
	newSpaceArray := new(SpaceArray[T])
	newSpaceArray.items.Initialize()
	newSpaceArray.rowLimit = rowThreshhold * 2
	newSpaceArray.rowThreshhold = rowThreshhold
	return newSpaceArray
}

func (sa *SpaceArray[T]) Length() int {
	return sa.length
}

func (sa *SpaceArray[T]) Capacity() int {
	return sa.items.Capacity() * sa.rowLimit
}

func (sa *SpaceArray[T]) Push(item T) error {
	if sa.items.Length() == 0 {
		sa.appendRow()
	}

	lastRow, error := sa.items.Get(sa.items.Length())
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
			item.Set(value, slidingIndex)
			return nil
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

		if slidingIndex >= item.Length() {
			slidingIndex -= item.Length()
			continue
		}

		itemLength := item.Length()
		if itemLength < sa.rowLimit {
			item.Add(value, slidingIndex)
			sa.length++
			return nil
		}

		sa.Push(sa.zeroValue)
		error = sa.moveElementsToTheRightFrom(i, slidingIndex)
		if error != nil {
			return error
		}
		item.Set(value, slidingIndex)

		return nil
	}

	return errors.New("array with invalid structure")
}

func (sa *SpaceArray[T]) appendRow() *factorArray.FactorArray[T] {
	newRow := new(factorArray.FactorArray[T])
	newRow.Initialize()
	sa.items.Push(newRow)
	return newRow
}

func (sa *SpaceArray[T]) moveElementsToTheRightFrom(itemIndex, valueInItemIndex int) error {
	item, error := sa.items.Get(itemIndex)
	if error != nil {
		return error
	}

	item.Add(sa.zeroValue, valueInItemIndex)

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
