package main

import (
	"fmt"
	"math/rand"
	"prefix-tree/triemap"
	"time"
)

const ITERATIONS = 10_000

func randomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

func main() {
	fmt.Println("Hello from fiveteenth homework!")

	fmt.Println("Iterations count: ", ITERATIONS)

	rand.Seed(255)

	start := time.Now()
	m := make(map[string]int)
	for i := 0; i < ITERATIONS; i++ {
		key := randomString(10)
		m[key] = i
	}
	elapsed := time.Since(start)
	fmt.Println("Insertion Time for map: ", elapsed)

	start = time.Now()
	tm := triemap.Constructor[int]()
	for i := 0; i < ITERATIONS; i++ {
		key := randomString(10)
		tm.Put(key, i)
	}
	elapsed = time.Since(start)
	fmt.Println("Insertion Time for triemap", elapsed)

	rand.Seed(255)

	start = time.Now()
	for i := 0; i < ITERATIONS; i++ {
		key := randomString(10)
		if m[key] != i {
			panic("INVALID MAP")
		}
	}
	elapsed = time.Since(start)
	fmt.Println("Search Time for map: ", elapsed)

	start = time.Now()
	for i := 0; i < ITERATIONS; i++ {
		key := randomString(10)
		if tm.Get(key) != i {
			panic("INVALID TRIEMAP")
		}
	}
	elapsed = time.Since(start)
	fmt.Println("Search Time for triemap", elapsed)

	rand.Seed(255)

	start = time.Now()
	for i := 0; i < ITERATIONS; i++ {
		key := randomString(10)
		delete(m, key)
	}
	elapsed = time.Since(start)
	fmt.Println("Remove Time for map: ", elapsed)

	start = time.Now()
	for i := 0; i < ITERATIONS; i++ {
		key := randomString(10)
		tm.Remove(key)
	}
	elapsed = time.Since(start)
	fmt.Println("Remove Time for triemap", elapsed)
}
