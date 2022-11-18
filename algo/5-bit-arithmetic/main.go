package main

import (
	"bit-arithmetic/queen"
	"fmt"
)

func main() {
	fmt.Println("Hello, the fifth homework!")
	moves := queen.PlaceQueen(0)
	fmt.Println("QUEEN(0) - ", moves)
}
