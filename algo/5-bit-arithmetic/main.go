package main

import (
	"bit-arithmetic/rook"
	"fmt"
)

func main() {
	fmt.Println("Hello, the fifth homework!")
	moves := rook.PlaceRook2(35)
	fmt.Println("ROOK(35) - ", moves)
}
