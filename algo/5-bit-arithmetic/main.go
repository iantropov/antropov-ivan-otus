package main

import (
	"bit-arithmetic/king"
	"fmt"
)

func main() {
	fmt.Println("Hello, the fifth homework!")
	pos, moves := king.PlaceKing(47)
	fmt.Println("KING(35) - ", pos, moves)
}
