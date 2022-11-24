package main

import (
	"fmt"
	"heapsort/heap"
	"heapsort/selection"
	"heapsort/sorting"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello from seventh homework!")

	N := 1_000_000
	numbers := make([]int, N)
	for i := range numbers {
		numbers[i] = i
	}

	rand.Shuffle(N, func(i, j int) { numbers[i], numbers[j] = numbers[j], numbers[i] })

	fmt.Println("Sorting of a shuffled array of length ", N)

	measureSorting("#selection", selection.Sort, numbers)
	measureSorting("#heap", heap.SortHeapify, numbers)
}

func measureSorting(name string, sort sorting.Sort, numbers []int) {
	start := time.Now()

	sorted := sort(numbers)

	elapsed := time.Since(start)
	fmt.Println("Sorting Time for "+name, elapsed)

	for i := 0; i < len(numbers); i++ {
		if sorted[i] != i {
			fmt.Println("FAIL")
		}
	}
}
