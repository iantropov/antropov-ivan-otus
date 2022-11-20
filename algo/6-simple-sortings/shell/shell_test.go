package shell

import (
	"simple-sortings/sorting"
	"testing"
)

func TestShellNaiveSorting(t *testing.T) {
	sorting.TestSorting(SortNaive, t)
}
