package main

import (
	"algebraic-algorithms/fibonacci"
	"algebraic-algorithms/power"
	"fmt"
)

func main() {
	fmt.Println("Hello, the third chapter!")
	fmt.Printf("5 ^ 3 = %d\n", power.Iterative(5, 3))
	fmt.Printf("Fibonnaci(5) (recursive) = %d\n", fibonacci.Recursive(5))
}
