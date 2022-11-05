package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("There are %d happy tickets with %d digits.\n", FindHappyTicketsAsMiddle(i), i)
	}
}
