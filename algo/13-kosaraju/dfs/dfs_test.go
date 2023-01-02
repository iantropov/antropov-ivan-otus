package dfs_test

import (
	"kosaraju/dfs"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDfsWithLevels1(t *testing.T) {
	g := [][]int{
		0: {1, 3},
		1: {},
		2: {3},
		3: {},
	}

	dfs := dfs.NewDfs(g)
	require.Equal(t, false, dfs.IsFinished())

	require.Equal(t, []int{0, 2}, dfs.VisitLevel())
	require.Equal(t, false, dfs.IsFinished())

	require.Equal(t, []int{1, 3}, dfs.VisitLevel())
	require.Equal(t, true, dfs.IsFinished())

	require.Equal(t, []int{}, dfs.VisitLevel())
	require.Equal(t, true, dfs.IsFinished())
}

func TestDfsWithLevels2(t *testing.T) {
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

	dfs := dfs.NewDfs(g)
	require.Equal(t, false, dfs.IsFinished())

	require.Equal(t, []int{0, 2, 3, 5, 7}, dfs.VisitLevel())
	require.Equal(t, false, dfs.IsFinished())

	require.Equal(t, []int{1, 4, 6}, dfs.VisitLevel())
	require.Equal(t, true, dfs.IsFinished())

	require.Equal(t, []int{}, dfs.VisitLevel())
	require.Equal(t, true, dfs.IsFinished())
}
