package main

import (
	"algebraic-algorithms/primes"
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	ta := 5
	n := primes.EratosthenesWithBits(100)
	fmt.Printf("Number of primers less than 100 (eratosthenes-with-bits) = %v %v\n", n, ta)

}
