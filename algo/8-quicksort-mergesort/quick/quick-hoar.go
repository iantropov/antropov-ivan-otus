package quick

import "quicksort-mergesort/sorting"

func SortHoar(a []int) []int {
	quickHoar(a, 0, len(a)-1)
	return a
}

func quickHoar(a []int, l, r int) {
	if l >= r {
		return
	}

	m := split(a, l, r)
	quick(a, l, m-1)
	quick(a, m+1, r)
}

func splitHoar(a []int, l, r int) int {
	pivot := a[(l+r)/2]

	i := l
	j := r
	for {
		for ; a[i] < pivot; i++ {
		}
		for ; a[j] > pivot; j-- {
		}
		if i >= j {
			return j
		}
		sorting.Swap(a, i, j)
		i++
		j--
	}
}
