package main

import (
	"fmt"
	"math/rand"
	"quicksort-mergesort/merge"
	"quicksort-mergesort/quick"
	"quicksort-mergesort/sorting"
	"time"
)

func main() {
	fmt.Println("Hello from eighth homework!")

	N := 10_000_000
	numbers := make([]int, N)
	for i := range numbers {
		numbers[i] = i
	}

	rand.Shuffle(N, func(i, j int) { numbers[i], numbers[j] = numbers[j], numbers[i] })

	fmt.Println("Sorting of a shuffled array of length ", N)

	measureSorting("#quick", quick.Sort, numbers)
	measureSorting("#merge", merge.Sort, numbers)
}

func measureSorting(name string, sort sorting.Sort, numbers []int) {
	numbersCopy := make([]int, len(numbers))
	copy(numbersCopy, numbers)

	start := time.Now()
	sorted := sort(numbersCopy)
	elapsed := time.Since(start)

	fmt.Println("Sorting Time for "+name, elapsed)

	for i := 0; i < len(numbers); i++ {
		if sorted[i] != i {
			fmt.Println("FAIL")
		}
	}
}
