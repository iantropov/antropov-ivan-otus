package bubble

import (
	"simple-sortings/sorting"
	"testing"
)

func TestBubbleSorting(t *testing.T) {
	sorting.TestSorting(Sort, t)
}

func TestBubbleOptimizedSorting(t *testing.T) {
	sorting.TestSorting(SortOptimized, t)
}

func TestBubbleWithRandomFiles(t *testing.T) {
	sorting.TestSortingWithRandomFiles(Sort, t)
}

func TestBubbleWithDigitsFiles(t *testing.T) {
	sorting.TestSortingWithDigitsFiles(Sort, t)
}

func TestBubbleWithSortedFiles(t *testing.T) {
	sorting.TestSortingWithSortedFiles(Sort, t)
}

func TestBubbleWithReversFiles(t *testing.T) {
	sorting.TestSortingWithReversFiles(Sort, t)
}
