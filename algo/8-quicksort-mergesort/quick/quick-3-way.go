package quick

import "quicksort-mergesort/sorting"

// https://www.geeksforgeeks.org/3-way-quicksort-dutch-national-flag/

func Sort3Way(a []int) []int {
	quick3Way(a, 0, len(a)-1)
	return a
}

func quick3Way(a []int, l, r int) {
	if l >= r {
		return
	}

	j, i := split3Way(a, l, r)
	quick3Way(a, l, j-1)
	quick3Way(a, i+1, r)
}

func split3Way(a []int, l, r int) (int, int) {
	pivot := a[r]

	p := l - 1
	q := r
	i := l - 1
	j := r

	for {
		for {
			i++
			if a[i] >= pivot {
				break
			}
		}

		for {
			j--
			if j == l || a[j] <= pivot {
				break
			}
		}

		if i >= j {
			break
		}

		sorting.Swap(a, i, j)

		if a[i] == pivot {
			p++
			sorting.Swap(a, p, i)
		}

		if a[j] == pivot {
			q--
			sorting.Swap(a, j, q)
		}
	}

	sorting.Swap(a, i, r)

	j = i
	for k := l; k < p; k++ {
		j--
		sorting.Swap(a, k, j)
	}

	for k := r - 1; k > q; k-- {
		i++
		sorting.Swap(a, k, i)
	}

	return j, i
}
