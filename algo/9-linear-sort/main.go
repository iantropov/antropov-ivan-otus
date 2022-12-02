package main

import (
	"fmt"
	"math/rand"
	"time"

	"linear-sort/bucket"
	"linear-sort/counting"
	"linear-sort/radix"
	"linear-sort/sorting"
)

func main() {
	fmt.Println("Hello from nineth homework!")

	N := 1_000_000
	numbers := make([]int, N)
	for i := range numbers {
		numbers[i] = rand.Intn(1000)
	}

	fmt.Println("Sorting of a shuffled array of length ", N)

	measureSorting("#counting", counting.Sort, numbers)
	measureSorting("#radix", radix.Sort, numbers)
	measureSorting("#bucket", bucket.Sort, numbers)
}

func measureSorting(name string, sort sorting.Sort, numbers []int) {
	numbersCopy := make([]int, len(numbers))
	copy(numbersCopy, numbers)

	start := time.Now()
	sorted := sort(numbersCopy)
	elapsed := time.Since(start)

	fmt.Println("Sorting Time for "+name, elapsed)

	for i := 1; i < len(numbers); i++ {
		if sorted[i] < sorted[i-1] {
			panic("FAIL to sort!")
		}
	}
}
