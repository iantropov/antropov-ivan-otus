package merge

import (
	"quicksort-mergesort/sorting"
	"testing"
)

func TestMerge(t *testing.T) {
	sorting.TestSorting(Sort, t)
}

func TestMergeWithRandomFiles(t *testing.T) {
	sorting.TestSortingWithRandomFiles(Sort, t)
}

func TestMergeWithDigitsFiles(t *testing.T) {
	sorting.TestSortingWithDigitsFiles(Sort, t)
}

func TestMergeWithSortedFiles(t *testing.T) {
	sorting.TestSortingWithSortedFiles(Sort, t)
}

func TestMergeWithReversFiles(t *testing.T) {
	sorting.TestSortingWithReversFiles(Sort, t)
}
