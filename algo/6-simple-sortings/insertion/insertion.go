package insertion

import "simple-sortings/sorting"

func Sort(a []int) []int {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && a[j-1] > a[j]; j-- {
			sorting.Swap(a, j-1, j)
		}
	}

	return a
}
