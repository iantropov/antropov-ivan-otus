package radix

import (
	"linear-sort/sorting"
	"testing"
)

func TestRadix(t *testing.T) {
	sorting.TestSorting(Sort, t)
}
