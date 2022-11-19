package bubble

import (
	"simple-sortings/sorting"
	"testing"
)

func TestBubbleSorting(t *testing.T) {
	sorting.TestSorting(Sort, t)
}
