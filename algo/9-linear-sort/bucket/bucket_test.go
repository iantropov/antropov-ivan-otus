package bucket

import (
	"linear-sort/sorting"
	"testing"
)

func TestSort(t *testing.T) {
	sorting.TestSorting(Sort, t)
}
