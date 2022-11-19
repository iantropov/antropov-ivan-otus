package sorting

import "testing"

type Sort func([]int) []int

func Swap(a []int, i, j int) {
	t := a[i]
	a[i] = a[j]
	a[j] = t
}

func TestSorting(sort Sort, t *testing.T) {
	array := []int{1, 23, 4, 7, 2, 9, 0, 2, 89, 41}
	expectedSortedArray := []int{0, 1, 2, 2, 4, 7, 9, 23, 41, 89}

	actualSortedArray := sort(array)

	for i := 0; i < len(array); i++ {
		if actualSortedArray[i] != expectedSortedArray[i] {
			t.Errorf("Incorrect actual value - %d. Expected - %d\n", actualSortedArray[i], expectedSortedArray[i])
		}
	}
}
