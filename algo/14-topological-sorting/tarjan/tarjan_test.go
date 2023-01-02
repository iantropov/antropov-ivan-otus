package tarjan_test

import (
	"testing"
	"topological-sorting/tarjan"

	"github.com/stretchr/testify/require"
)

func TestSort(t *testing.T) {
	g := [][]int{
		0: {1},
		1: {4},
		2: {3},
		3: {0, 1, 4, 5},
		4: {6},
		5: {4, 7},
		6: {7},
		7: {},
	}

	result := tarjan.Sort(g)
	require.Equal(t, []int{7, 6, 4, 1, 0, 5, 3, 2}, result)
}
