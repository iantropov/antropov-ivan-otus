package shell

import "simple-sortings/sorting"

func SortNaive(a []int) []int {
	for h := len(a) / 2; h > 0; h /= 2 {
		for i := 0; i < h; i++ {
			for j := i + h; j < len(a); j += h {
				for k := j; k >= h && a[k-h] > a[k]; k -= h {
					sorting.Swap(a, k-h, k)
				}
			}
		}
	}
	return a
}

func Sort(a []int) []int {
	for h := len(a) / 2; h > 0; h /= 2 {
		for i := h; i < len(a); i++ {
			for j := i; j >= h && a[j-h] > a[j]; j -= h {
				sorting.Swap(a, j-h, j)
			}
		}
	}

	return a
}
