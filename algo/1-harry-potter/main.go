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
		if y < x+10 && y > x-10 && y > 14-x && y < 34-x {
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
// 16 - y < x+10 && y > x-10 && y > 14-x && y < 34-x
// 17 - ??

// 18 - (x != 0 || y != 0) && (x < 2 || y < 2)
// 19 - x == 0 || y == 0 || x == 24 || y == 24
// 20 - !(x%2 == 1 && y%2 == 0 || x%2 == 0 && y%2 == 1)
// 21 - ??
// 22 - ??
// 23 - ??
// 24 - x == y || x == 24-y
// 25 - x%6 == 0 || y%6 == 0
