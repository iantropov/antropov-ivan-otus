package kruskal

import (
	"minimal-tree/uf"
	"sort"
)

type Edge struct {
	U, V, W int
}

func Kruskal(wg [][]int) []Edge {
	mt := make([]Edge, 0)
	uf := uf.NewUf(len(wg))

	edges := make([]Edge, 0)
	for i := 0; i < len(wg); i++ {
		for j := i + 1; j < len(wg[i]); j++ {
			if wg[i][j] > 0 {
				edges = append(edges, Edge{U: i, V: j, W: wg[i][j]})
			}
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].W < edges[j].W
	})

	for _, e := range edges {
		if !uf.Connected(e.U, e.V) {
			uf.Union(e.U, e.V)
			mt = append(mt, e)
		}
	}

	return mt
}
