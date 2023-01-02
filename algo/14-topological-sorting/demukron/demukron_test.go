package demukron_test

import (
	"testing"
	"topological-sorting/demukron"

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

	result := demukron.Sort(g)
	require.Equal(t, []int{2, 3, 0, 1, 5, 4, 6, 7}, result)
}
