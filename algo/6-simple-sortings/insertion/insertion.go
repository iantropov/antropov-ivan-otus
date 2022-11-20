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

func SortWithShifts(a []int) []int {
	for i := 1; i < len(a); i++ {
		aI := a[i]
		j := i
		for ; j > 0 && a[j-1] > aI; j-- {
			a[j] = a[j-1]
		}
		a[j] = aI
	}

	return a
}

func SortWithBinarySearch(a []int) []int {
	for i := 1; i < len(a); i++ {
		lo := 0
		hi := i - 1
		for hi-lo > 1 {
			m := (lo + hi) / 2
			if a[m] >= a[i] {
				hi = m
			} else {
				lo = m
			}
		}

		dst := 0
		if a[hi] < a[i] {
			dst = hi + 1
		} else if a[lo] < a[i] {
			dst = lo + 1
		} else {
			dst = lo
		}

		aI := a[i]
		for k := i; k > dst; k-- {
			a[k] = a[k-1]
		}
		a[dst] = aI
	}

	return a
}
