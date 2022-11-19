package insertion

import "simple-sortings/sorting"

func Sort(a []int) []int {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				sorting.Swap(a, j, j-1)
			}
		}
	}

	return a
}
