package main

import (
	"bit-arithmetic/bishop"
	"fmt"
)

func main() {
	fmt.Println("Hello, the fifth homework!")
	moves := bishop.PlaceBishop(0)
	fmt.Println("BISHOP(0) - ", moves)
}
