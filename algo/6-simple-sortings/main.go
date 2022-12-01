package main

import (
	"fmt"
	"math/rand"
	"simple-sortings/bubble"
	"simple-sortings/insertion"
	"simple-sortings/shell"
	"simple-sortings/sorting"
	"time"
)

func main() {
	fmt.Println("Hello from sixth homework!")

	N := 1_000_000
	numbers := make([]int, N)
	for i := range numbers {
		numbers[i] = i
	}

	rand.Shuffle(N, func(i, j int) { numbers[i], numbers[j] = numbers[j], numbers[i] })

	fmt.Println("Sorting of a shuffled array of length ", N)

	measureSorting("#bubble", bubble.Sort, numbers)
	measureSorting("#insertion", insertion.Sort, numbers)
	measureSorting("#insertion-with-shifts", insertion.SortWithShifts, numbers)
	measureSorting("#insertion-with-binary-search", insertion.SortWithBinarySearch, numbers)
	measureSorting("#shell-naive", shell.SortNaive, numbers)
	measureSorting("#shell", shell.Sort, numbers)
	measureSorting("#shell-gap3", shell.SortWithGap3, numbers)
	measureSorting("#shell-gap2k", shell.SortWithGap2k, numbers)
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
