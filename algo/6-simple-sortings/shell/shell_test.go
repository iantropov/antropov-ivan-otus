package shell

import (
	"simple-sortings/sorting"
	"testing"
)

func TestShellNaiveSorting(t *testing.T) {
	sorting.TestSorting(SortNaive, t)
}

func TestShellSorting(t *testing.T) {
	sorting.TestSorting(Sort, t)
}

func TestShellWithGap3Sorting(t *testing.T) {
	sorting.TestSorting(SortWithGap3, t)
}

func TestShellWithGap2kSorting(t *testing.T) {
	sorting.TestSorting(SortWithGap2k, t)
}

func TestShellWithRandomFiles(t *testing.T) {
	sorting.TestSortingWithRandomFiles(Sort, t)
}

func TestShellWithDigitsFiles(t *testing.T) {
	sorting.TestSortingWithDigitsFiles(Sort, t)
}

func TestShellWithSortedFiles(t *testing.T) {
	sorting.TestSortingWithSortedFiles(Sort, t)
}

func TestShellWithReversFiles(t *testing.T) {
	sorting.TestSortingWithReversFiles(Sort, t)
}
