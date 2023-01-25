package tarjan

const (
	Black int = iota
	Red
	Green
)

func Sort(g [][]int) []int {
	vertexes := make([]int, len(g))
	result := make([]int, 0)

	for i := range g {
		if vertexes[i] == Black {
			result = dfs(i, g, vertexes, result)
		}
	}

	return result
}

func dfs(u int, g [][]int, vertexes []int, result []int) []int {
	for _, v := range g[u] {
		if vertexes[v] == Black {
			vertexes[v] = Red
			result = dfs(v, g, vertexes, result)
		}
	}
	vertexes[u] = Green
	result = append(result, u)
	return result
}
