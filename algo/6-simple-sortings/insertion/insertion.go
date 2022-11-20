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
