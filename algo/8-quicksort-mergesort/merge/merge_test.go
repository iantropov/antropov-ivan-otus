package merge

import (
	"quicksort-mergesort/sorting"
	"testing"
)

func TestMerge(t *testing.T) {
	sorting.TestSorting(Sort, t)
}
