package main

import (
	"fmt"
	"math/rand"
	"time"

	"other-trees/btree"
	"other-trees/random"
	"other-trees/rbt"
	"other-trees/splay"
	"other-trees/treap"
	"other-trees/tree"
)

const N = 51 // 20 - invalid case

func main() {
	fmt.Println("Hello from eleventh homework!")

	numbersSequential := make([]int, N)
	for i := range numbersSequential {
		numbersSequential[i] = i
	}

	numbersRandom := make([]int, N)
	copy(numbersRandom, numbersSequential)
	// rand.Seed(time.Hour.Nanoseconds())
	rand.Shuffle(N, func(i, j int) { numbersRandom[i], numbersRandom[j] = numbersRandom[j], numbersRandom[i] })

	fmt.Println("Prepared a shuffled array of length ", N)
	fmt.Println("=========================================")

	// measureBtree("#sequential", numbersSequential)
	// measureBtree("#random", numbersRandom)

	// measureRandom("#sequential", numbersSequential)
	// measureRandom("#random", numbersRandom)

	measureRbt("#sequential", numbersSequential)
	measureRbt("#random", numbersRandom)

	// measureSplay("#sequential", numbersSequential)
	// measureSplay("#random", numbersRandom)

	// measureTreap("#sequential", numbersSequential)
	// measureTreap("#random", numbersRandom)

}

func measureBtree(name string, numbers []int) {
	btree := btree.NewTree()
	measureTree(btree, name+"-btree: ", numbers)
}

func measureRandom(name string, numbers []int) {
	random := random.NewTree()
	measureTree(random, name+"-random: ", numbers)
}

func measureRbt(name string, numbers []int) {
	rbt := rbt.NewTree()
	measureTree(rbt, name+"-rbt: ", numbers)
}

func measureSplay(name string, numbers []int) {
	splay := splay.NewTree()
	measureTree(splay, name+"-splay: ", numbers)
}

func measureTreap(name string, numbers []int) {
	treap := treap.NewTree()
	measureTree(treap, name+"-treap: ", numbers)
}

func measureTree(tree tree.Tree, name string, numbers []int) {
	start := time.Now()
	startTotal := start
	for _, n := range numbers {
		// fmt.Printf("============= BEFORE INSERT (%d)  =============\n", n)
		// tree.DumpValuesInDetails()
		tree.Insert(n)
		// fmt.Printf("============= AFTER INSERT (%d)  =============\n", n)
	}
	elapsed := time.Since(start)
	fmt.Println("Insertion Time for "+name, elapsed)

	start = time.Now()
	for i := 0; i < N/10; i++ {
		tree.Search(rand.Intn(N))
	}
	elapsed = time.Since(start)
	fmt.Println("Search Time for "+name, elapsed)

	deleteNumbers := make([]int, N)
	copy(deleteNumbers, numbers)
	rand.Shuffle(N, func(i, j int) { deleteNumbers[i], deleteNumbers[j] = deleteNumbers[j], deleteNumbers[i] })

	tree.DumpValuesInDetails()
	start = time.Now()
	for i := 0; i < N; i++ {
		// num := rand.Intn(N)
		num := deleteNumbers[i]
		fmt.Printf("============= WILL REMOVE (%d)  =============\n", num)
		tree.Remove(num)
		fmt.Printf("============= AFTER REMOVAL (%d)  =============\n", num)
		tree.DumpValuesInDetails()
	}
	elapsed = time.Since(start)
	fmt.Println("Remove Time for "+name, elapsed)

	elapsed = time.Since(startTotal)
	fmt.Println("Total processing Time for "+name, elapsed)
	tree.DumpValuesInDetails()
	fmt.Println("=========================================")
}
