package insertion

import (
	"simple-sortings/sorting"
	"testing"
)

func TestInsertionSorting(t *testing.T) {
	sorting.TestSorting(Sort, t)
}

func TestInsertionWithShiftsSorting(t *testing.T) {
	sorting.TestSorting(SortWithShifts, t)
}
