package main

import (
	"fmt"
	"math/rand"
	"time"

	"bst-avl/avl"
	"bst-avl/bst"
	"bst-avl/tree"
)

const N = 100_000

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

	fmt.Println("Prepared a shuffled array of length ", N)
	fmt.Println("=========================================")

	measureBST("#sequential", numbersSequential)
	measureBST("#random", numbersRandom)

	measureAvl("#sequential", numbersSequential)
	measureAvl("#random", numbersRandom)
}

func measureBST(name string, numbers []int) {
	bst := bst.CreateBST()
	measureTree(bst, name+"-bst: ", numbers)
}

func measureAvl(name string, numbers []int) {
	avl := avl.NewTree()
	measureTree(avl, name+"-avl: ", numbers)
}

func measureTree(tree tree.Tree, name string, numbers []int) {
	start := time.Now()
	startTotal := start
	for _, n := range numbers {
		tree.Insert(n)
	}
	elapsed := time.Since(start)
	fmt.Println("Insertion Time for "+name, elapsed)

	start = time.Now()
	for i := 0; i < N/10; i++ {
		tree.Search(rand.Intn(N))
	}
	elapsed = time.Since(start)
	fmt.Println("Search Time for "+name, elapsed)

	start = time.Now()
	for i := 0; i < N/10; i++ {
		num := rand.Intn(N)
		tree.Remove(num)
	}
	elapsed = time.Since(start)
	fmt.Println("Remove Time for "+name, elapsed)

	elapsed = time.Since(startTotal)
	fmt.Println("Total processing Time for "+name, elapsed)
	fmt.Println("=========================================")
}
