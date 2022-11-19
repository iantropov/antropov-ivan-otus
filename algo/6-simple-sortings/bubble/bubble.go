package bubble

import "simple-sortings/sorting"

func Sort(a []int) []int {
	for i := len(a) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				sorting.Swap(a, j, j+1)
			}
		}
	}
	return a
}
