package chessPiece

import (
	countBits "bit-arithmetic/count-bits"
	"bit-arithmetic/tester"
	"strconv"
	"testing"
)

type PlaceChessPiece func(int) uint64

func TestChessPiece(testDir, testSuiteName string, place PlaceChessPiece, countBits countBits.CountBits, t *testing.T) {
	tester.TestWithFiles(testDir, testSuiteName, t, func(inputs []string) []string {
		num, err := strconv.Atoi(inputs[0])
		if err != nil {
			panic(err)
		}
		moves := place(num)
		count := countBits(moves)
		return []string{strconv.FormatInt(int64(count), 10), strconv.FormatUint(moves, 10)}
	}, func(expected, actual string) bool {
		return expected == actual
	})
}
