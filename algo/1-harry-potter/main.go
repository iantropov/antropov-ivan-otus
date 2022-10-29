package main

import (
	tm "github.com/buger/goterm"
)

func outputSquareWithValue(width int, height int, getValue func(int, int) string) {
	tm.Clear()

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			tm.MoveCursor(i+1, j+1)
			tm.Print(getValue(i, j))
		}
		tm.Flush()
		// time.Sleep(time.Second)
	}
}

func main() {
	outputSquareWithValue(20, 20, func(i, j int) string {
		if i >= j {
			return "*"
		} else {
			return "."
		}
	})
}
