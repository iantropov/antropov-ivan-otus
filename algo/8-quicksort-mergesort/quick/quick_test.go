package quick

import (
	"quicksort-mergesort/sorting"
	"testing"
)

func TestQuick(t *testing.T) {
	sorting.TestSorting(Sort, t)
}

func TestQuickWithRandomFiles(t *testing.T) {
	sorting.TestSortingWithRandomFiles(Sort, t)
}

func TestQuickWithDigitsFiles(t *testing.T) {
	sorting.TestSortingWithDigitsFiles(Sort, t)
}

func TestQuickWithSortedFiles(t *testing.T) {
	sorting.TestSortingWithSortedFiles(Sort, t)
}

func TestQuickWithReversFiles(t *testing.T) {
	sorting.TestSortingWithReversFiles(Sort, t)
}
