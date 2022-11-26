package external

import (
	"math/rand"
	"testing"
	"time"
)

func TestExternal(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	GenerateTextFile("input.txt", 10, 100)
	SortBySplit("input.txt", "output.txt", 10, 3)
}
