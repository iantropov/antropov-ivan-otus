package main

import (
	"fmt"
	"math/rand"
	"time"

	"bst-avl/bst"
)

const N = 150_000

func main() {
	fmt.Println("Hello from tenth homework!")

	numbersSequential := make([]int, N)
	for i := range numbersSequential {
		numbersSequential[i] = i
	}

	numbersRandom := make([]int, N)
	copy(numbersRandom, numbersSequential)
	rand.Seed(time.Hour.Nanoseconds())
	rand.Shuffle(N, func(i, j int) { numbersRandom[i], numbersRandom[j] = numbersRandom[j], numbersRandom[i] })

	fmt.Println("Sorting of a shuffled array of length ", N)

	measureBST("#sequential", numbersSequential)
	measureBST("#random", numbersRandom)
}

func measureBST(name string, numbers []int) {
	start := time.Now()

	bst := bst.CreateBST()

	for _, n := range numbers {
		bst.Insert(n)
	}

	for i := 0; i < N/10; i++ {
		bst.Search(rand.Intn(N))
	}

	for i := 0; i < N/10; i++ {
		num := rand.Intn(N)
		bst.Remove(num)
	}

	elapsed := time.Since(start)

	fmt.Println("Processing Time for "+name, elapsed)
}
