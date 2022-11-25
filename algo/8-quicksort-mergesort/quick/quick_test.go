package quick

import (
	"quicksort-mergesort/sorting"
	"testing"
)

func TestQuick(t *testing.T) {
	sorting.TestSorting(Sort, t)
}
