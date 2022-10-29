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
	outputSquareWithValue(25, 25, func(x, y int) string {
		if x >= 13-y && x < y+12 && y < 13+x {
			return "#"
		} else {
			return "."
		}
	})
}

// 1 - x > y

// 2 - x == y
// 3 - x == 24 - y
// 4 - x < 30 - y
// 5 - x / 2 == y
// 6 - x < 10 || y < 10
// 7 - x > 15 && y > 15

// 8 - x == 0 || y == 0
// 9 - x > y + 10 || x < y - 10
// 10 - x > y && x / 2 <= y
// 11 - x == 1 || x == 23 || y == 1 || y == 23
// 12 - ?? - circle ?
// 13 - y > 19 - x && y < 29 - x
// 14 - ??
// 15 - x > y + 10 && x < y + 21 || x < y - 10 && x > y - 21
// 16 -
// 17 -
