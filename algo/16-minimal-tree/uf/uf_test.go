package uf_test

import (
	"minimal-tree/uf"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUf(t *testing.T) {
	uf := uf.NewUf(10)

	require.Equal(t, 10, uf.ComponentsCount())

	uf.Union(4, 3)
	uf.Union(3, 8)
	uf.Union(6, 5)
	uf.Union(9, 4)
	uf.Union(2, 1)
	uf.Union(8, 9)
	uf.Union(5, 0)
	uf.Union(7, 2)
	uf.Union(6, 1)
	uf.Union(1, 0)
	uf.Union(6, 7)

	require.Equal(t, 2, uf.ComponentsCount())
}
