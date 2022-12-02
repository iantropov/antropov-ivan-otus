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

func TestHeapifyWithRandomFiles(t *testing.T) {
	sorting.TestSortingWithRandomFiles(SortHeapify, t)
}

func TestHeapifyWithDigitsFiles(t *testing.T) {
	sorting.TestSortingWithDigitsFiles(SortHeapify, t)
}

func TestHeapifyWithSortedFiles(t *testing.T) {
	sorting.TestSortingWithSortedFiles(SortHeapify, t)
}

func TestHeapifyWithReversFiles(t *testing.T) {
	sorting.TestSortingWithReversFiles(SortHeapify, t)
}
