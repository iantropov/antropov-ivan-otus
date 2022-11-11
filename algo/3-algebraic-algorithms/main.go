package main

import (
	"algebraic-algorithms/fibonacci"
	"algebraic-algorithms/power"
	"algebraic-algorithms/primes"
	"fmt"
)

func main() {
	fmt.Println("Hello, the third chapter!")
	fmt.Printf("5 ^ 5 (iterative) = %d\n", power.Iterative(5, 5))
	fmt.Printf("5 ^ 5 (suboptimal) = %d\n", power.Suboptimal(5, 5))
	fmt.Printf("Fibonnaci(5) (recursive) = %d\n", fibonacci.Recursive(5))
	fmt.Printf("Fibonnaci(5) (iterative) = %d\n", fibonacci.Iterative(5))
	fmt.Printf("5 first primers (brute-force) = %v\n", primes.BruteForce(5))
}
