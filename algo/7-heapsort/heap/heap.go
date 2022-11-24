package heap

import "heapsort/sorting"

func Sort(a []int) []int {
	for i := len(a) - 1; i > 0; i-- {
		swim(a, len(a), i)
	}

	for i := len(a) - 1; i > 0; i-- {
		sorting.Swap(a, i, 0)
		sink(a, i, 0)
	}

	return a
}

func swim(a []int, n, i int) {
	parent := (i - 1) / 2

	if parent >= 0 && a[i] > a[parent] {
		sorting.Swap(a, i, parent)
		swim(a, n, parent)
	}
}

func sink(a []int, n, i int) {
	root := i
	left := i*2 + 1
	right := i*2 + 2

	node := root
	if left < n && a[left] > a[node] {
		node = left
	}
	if right < n && a[right] > a[node] {
		node = right
	}
	if node == i {
		return
	}
	sorting.Swap(a, i, node)
	sink(a, n, node)
}
