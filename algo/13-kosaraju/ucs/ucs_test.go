package ucs_test

import (
	"kosaraju/ucs"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUcs(t *testing.T) {
	g := [][]int{
		0: {1, 2, 3},
		1: {4},
		2: {3, 4},
		3: {4},
		4: {},
	}

	weights := [][]int{
		0: {0, 2, 1, 1, 0},
		1: {0, 0, 0, 0, 3},
		2: {0, 0, 0, 5, 1},
		3: {0, 0, 0, 0, 3},
		4: {0, 0, 0, 0, 0},
	}

	order := ucs.Ucs(g, weights)
	require.Equal(t, []int{0, 3, 2, 1, 4}, order)
}
