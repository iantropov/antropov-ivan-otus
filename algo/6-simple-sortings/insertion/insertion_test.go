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

func TestInsertionWithBinarySearchSorting(t *testing.T) {
	sorting.TestSorting(SortWithBinarySearch, t)
}

func TestInsertionWithBinarySearchWithRandomFiles(t *testing.T) {
	sorting.TestSortingWithRandomFiles(SortWithBinarySearch, t)
}

func TestInsertionWithBinarySearchWithDigitsFiles(t *testing.T) {
	sorting.TestSortingWithDigitsFiles(SortWithBinarySearch, t)
}

func TestInsertionWithBinarySearchWithSortedFiles(t *testing.T) {
	sorting.TestSortingWithSortedFiles(SortWithBinarySearch, t)
}

func TestInsertionWithBinarySearchWithReversFiles(t *testing.T) {
	sorting.TestSortingWithReversFiles(SortWithBinarySearch, t)
}
