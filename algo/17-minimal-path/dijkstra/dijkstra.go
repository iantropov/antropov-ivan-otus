package dijkstra

import (
	"minimal-path/heap"
)

func Dijkstra(g [][]int, u, v int) []int {
	adj := make([][]int, len(g))
	for i := 0; i < len(g); i++ {
		adj[i] = make([]int, 0)
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] > 0 {
				adj[i] = append(adj[i], j)
			}
		}
	}

	priors := make([]int, len(g))

	h := heap.NewHeap(len(g))
	h.DecreaseKey(u, 0)

	for !h.Empty() {
		i, p := h.ExtractMin()
		for _, j := range adj[i] {
			if h.Contains(j) && h.Priority(j) > p+g[i][j] {
				h.DecreaseKey(j, p+g[i][j])
				priors[j] = i
			}
		}
	}

	output := make([]int, 1)
	output[0] = v
	for n := v; priors[n] != u; n = priors[n] {
		output = append(output, priors[n])
	}
	output = append(output, u)

	for i, j := 0, len(output)-1; i < j; {
		output[i], output[j] = output[j], output[i]
		i++
		j--
	}

	return output
}
