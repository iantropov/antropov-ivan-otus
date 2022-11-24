package heap

import (
	"heapsort/sorting"
	"testing"
)

func TestHeapSwimSink(t *testing.T) {
	sorting.TestSorting(SortSwinkSink, t)
}

func TestHeapHeapify(t *testing.T) {
	sorting.TestSorting(SortHeapify, t)
}
