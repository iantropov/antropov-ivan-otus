package bellmanford_test

import (
	bellmanford "minimal-tree/bellman-ford"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBellmanFord(t *testing.T) {
	path := bellmanford.BellmanFord([][]int{
		0: {0, 1, 0, 3, 0},
		1: {0, 0, 3, 1, 2},
		2: {0, 0, 0, 0, 0},
		3: {0, 0, 3, 0, 1},
		4: {0, 0, 2, 0, 0},
	}, 0, 4)

	require.Equal(t, []int{0, 1, 4}, path)
}
