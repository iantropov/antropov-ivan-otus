package main

import "fmt"

func main() {
	fmt.Println("Hello from the twentieth homework!")
	var input string
	fmt.Scan(&input)

	output := rle([]byte(input))
	fmt.Println(output)
}

func rle(input []byte) []byte {
	res := make([]byte, 0)
	if len(input) == 0 {
		return res
	}

	count := 1
	for i := 1; i < len(input); i++ {
		if count == 255 || input[i] != input[i-1] {
			res = append(res, byte(count), input[i-1])
			count = 1
		} else {
			count++
		}
	}
	res = append(res, byte(count), input[len(input)-1])
	return res
}
