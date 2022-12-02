package selection

import "heapsort/sorting"

func Sort(a []int) []int {
	for i := len(a) - 1; i > 0; i-- {
		max := i
		for j := i - 1; j >= 0; j-- {
			if a[j] > a[max] {
				max = j
			}
		}
		sorting.Swap(a, i, max)
	}

	return a
}
