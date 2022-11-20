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
