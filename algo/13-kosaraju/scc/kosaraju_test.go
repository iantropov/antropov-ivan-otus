package kosaraju_test

import (
	kosaraju "kosaraju/scc"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKosaraju(t *testing.T) {
	g := [][]int{
		0: {1},
		1: {2},
		2: {0},
		3: {1, 2, 4},
		4: {3, 5},
		5: {2, 6},
		6: {5},
		7: {4, 6, 7},
	}

	result := kosaraju.Kosaraju(g)

	require.Equal(t, result, []int{0, 0, 0, 3, 3, 5, 5, 7})
}
