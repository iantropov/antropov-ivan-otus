package external

import (
	"math/rand"
	"testing"
	"time"
)

func TestSortBySplit(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	GenerateTextFile("input.txt", 10, 100)
	SortBySplit("input.txt", "output.txt", 10, 3)
}

func TestSortByMerge(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	GenerateTextFile("input.txt", 10, 100)
	SortByMerge("input.txt", "output.txt", 10)
}
