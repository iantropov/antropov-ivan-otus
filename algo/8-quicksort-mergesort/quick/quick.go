package quick

import "quicksort-mergesort/sorting"

func Sort(a []int) []int {
	quick(a, 0, len(a)-1)
	return a
}

func split(a []int, l, r int) int {
	pivot := a[r]

	m := l - 1
	for j := l; j <= r; j++ {
		if a[j] <= pivot {
			m++
			sorting.Swap(a, m, j)
		}
	}
	return m
}

func quick(a []int, l, r int) {
	if l >= r {
		return
	}

	m := split(a, l, r)
	quick(a, l, m-1)
	quick(a, m+1, r)
}
