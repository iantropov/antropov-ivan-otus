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
	fmt.Printf("5 ^ 5 (optimal) = %d\n", power.Optimal(5, 5))
	fmt.Printf("Fibonnaci(5) (recursive) = %d\n", fibonacci.Recursive(5))
	fmt.Printf("Fibonnaci(5) (iterative) = %d\n", fibonacci.Iterative(5))
	fmt.Printf("Fibonnaci(5) (golden-ratio) = %d\n", fibonacci.GoldenRatio(5))
	fmt.Printf("Fibonnaci(5) (matrix) = %d\n", fibonacci.Matrix(5))
	fmt.Printf("Number of primers less than 100 (brute-force) = %v\n", primes.BruteForce(100))
	fmt.Printf("Number of primers less than 100 (optimal) = %v\n", primes.Optimal(100))
	fmt.Printf("Number of primers less than 100 (eratosthenes) = %v\n", primes.Eratosthenes(100))
	fmt.Printf("Number of primers less than 100 (eratosthenes-with-bits) = %v\n", primes.EratosthenesWithBits(100))
	fmt.Printf("Number of primers less than 100 (linear-eratosthenes) = %v\n", primes.LinearEratosthenes(100))
}
