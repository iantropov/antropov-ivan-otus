package quick

import (
	"quicksort-mergesort/sorting"
	"testing"
)

func TestQuick(t *testing.T) {
	sorting.TestSorting(Sort, t)
}

func TestQuicHoark(t *testing.T) {
	sorting.TestSorting(SortHoar, t)
}

func TestQuickWithRandomFiles(t *testing.T) {
	sorting.TestSortingWithRandomFiles(Sort, t)
}

func TestQuickHoarWithRandomFiles(t *testing.T) {
	sorting.TestSortingWithRandomFiles(SortHoar, t)
}

func TestQuickWithDigitsFiles(t *testing.T) {
	sorting.TestSortingWithDigitsFiles(Sort, t)
}

func TestQuickHoarWithDigitsFiles(t *testing.T) {
	sorting.TestSortingWithDigitsFiles(SortHoar, t)
}

func TestQuickWithSortedFiles(t *testing.T) {
	sorting.TestSortingWithSortedFiles(SortHoar, t)
}

func TestQuickWithReversFiles(t *testing.T) {
	sorting.TestSortingWithReversFiles(Sort, t)
}

func TestQuicHoarkWithReversFiles(t *testing.T) {
	sorting.TestSortingWithReversFiles(SortHoar, t)
}
