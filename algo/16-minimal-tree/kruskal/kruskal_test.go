package kruskal_test

import (
	"minimal-tree/kruskal"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKruskal(t *testing.T) {
	edges := kruskal.Kruskal([][]int{
		0: {0, 9, 0, 10, 0, 0, 0, 0, 3},
		1: {9, 0, 4, 0, 8, 0, 0, 0, 16},
		2: {0, 4, 0, 0, 14, 1, 0, 0, 0},
		3: {10, 0, 0, 0, 7, 0, 13, 5, 11},
		4: {0, 8, 14, 7, 0, 12, 15, 0, 0},
		5: {0, 0, 1, 0, 12, 0, 2, 0, 0},
		6: {0, 0, 0, 13, 15, 2, 0, 6, 0},
		7: {0, 0, 0, 5, 0, 0, 6, 0, 0},
		8: {3, 16, 0, 11, 0, 0, 0, 0, 0},
	})

	require.Equal(t, []kruskal.Edge{
		0: {2, 5, 1},
		1: {5, 6, 2},
		2: {0, 8, 3},
		3: {1, 2, 4},
		4: {3, 7, 5},
		5: {6, 7, 6},
		6: {3, 4, 7},
		7: {0, 1, 9},
	}, edges)
}
