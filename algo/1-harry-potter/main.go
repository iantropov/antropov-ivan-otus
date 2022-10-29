package main

import (
	"time"

	tm "github.com/buger/goterm"
)

func outputSquareWithValue(width, height int, value string) {
	tm.Clear()

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			tm.MoveCursor((i+1)*(len(value)), j+1)
			tm.Print(value)
			tm.Flush()
		}
		time.Sleep(time.Second)
	}
}

func main() {
	outputSquareWithValue(10, 10, "*")
}
