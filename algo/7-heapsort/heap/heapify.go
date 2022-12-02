package heap

import "heapsort/sorting"

func SortHeapify(a []int) []int {
	for i := len(a)/2 - 1; i >= 0; i-- {
		heapify(a, len(a), i)
	}

	for i := len(a) - 1; i > 0; i-- {
		sorting.Swap(a, i, 0)
		heapify(a, i, 0)
	}

	return a
}

func heapify(a []int, n, i int) {
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
	heapify(a, n, node)

}
