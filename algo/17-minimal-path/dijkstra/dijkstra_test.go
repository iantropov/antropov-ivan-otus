package dijkstra_test

import (
	"minimal-path/dijkstra"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDijkstra(t *testing.T) {
	path := dijkstra.Dijkstra([][]int{
		0: {0, 1, 0, 3, 0},
		1: {0, 0, 3, 1, 2},
		2: {0, 0, 0, 0, 0},
		3: {0, 0, 3, 0, 1},
		4: {0, 0, 2, 0, 0},
	}, 0, 4)

	require.Equal(t, []int{0, 1, 4}, path)
}
