package prim

type Edge struct {
	U, V, W int
}

func Prim(g [][]int) []Edge {
	edges := make([][]Edge, len(g))
	for i := 0; i < len(g); i++ {
		edges[i] = make([]Edge, 0)
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] > 0 {
				edges[i] = append(edges[i], Edge{
					U: i,
					V: j,
					W: g[i][j],
				})
			}
		}
	}

	output := make([]Edge, 0)
	heap := NewHeap(len(g))

	heap.DecreaseKey(0, 0)
	u, _ := heap.ExtractMin()
	for !heap.Empty() {
		minEdge := Edge{U: u}
		for _, edge := range edges[u] {
			if heap.Contains(edge.V) && edge.W < heap.Priority(edge.V) {
				minEdge.V = edge.V
				minEdge.W = edge.W
				heap.DecreaseKey(edge.V, edge.W)
			}
		}
		output = append(output, minEdge)
		u, _ = heap.ExtractMin()
	}

	return output
}
