package main

import (
	"bit-arithmetic/rook"
	"fmt"
)

func main() {
	fmt.Println("Hello, the fifth homework!")
	moves := rook.PlaceRook2(35)
	fmt.Println("ROOK(35) - ", moves)

	var a uint64 = 0xFF
	var pos uint64 = 1 << 35
	var b uint64 = 0x101010101010101
	var res uint64 = (a << ((pos >> 3) << 3)) ^ (b << (pos & 7))
	fmt.Println("ANSWER - ", res)

}
