package kosaraju

func Kosaraju(g [][]int) []int {
	reversedG := make([][]int, len(g))
	visitedForOrder := make([]bool, len(g))
	order := make([]int, 0)

	for u := range g {
		order = dfsForOrder(u, g, reversedG, visitedForOrder, order)
	}

	cc := make([]int, len(g))
	visitedForCc := make([]bool, len(g))
	for i := len(order) - 1; i >= 0; i-- {
		dfsForCc(order[i], order[i], reversedG, visitedForCc, cc)
	}

	return cc
}

func dfsForOrder(u int, g, reversedG [][]int, visited []bool, order []int) []int {
	if visited[u] {
		return order
	}

	visited[u] = true
	for _, v := range g[u] {
		order = dfsForOrder(v, g, reversedG, visited, order)
		reversedG[v] = append(reversedG[v], u)
	}
	order = append(order, u)
	return order
}

func dfsForCc(u, root int, g [][]int, visited []bool, cc []int) {
	if visited[u] {
		return
	}

	visited[u] = true
	cc[u] = root
	for _, v := range g[u] {
		dfsForCc(v, root, g, visited, cc)
	}
}
