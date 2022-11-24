package heap

import (
	"heapsort/sorting"
	"testing"
)

func TestHeap(t *testing.T) {
	sorting.TestSorting(Sort, t)
}
