package selection

import (
	"heapsort/sorting"
	"testing"
)

func TestSelection(t *testing.T) {
	sorting.TestSorting(Sort, t)
}

func TestSelectionWithRandomFiles(t *testing.T) {
	sorting.TestSortingWithRandomFiles(Sort, t)
}

func TestSelectionWithDigitsFiles(t *testing.T) {
	sorting.TestSortingWithDigitsFiles(Sort, t)
}

func TestSelectionWithSortedFiles(t *testing.T) {
	sorting.TestSortingWithSortedFiles(Sort, t)
}

func TestSelectionWithReversFiles(t *testing.T) {
	sorting.TestSortingWithReversFiles(Sort, t)
}
