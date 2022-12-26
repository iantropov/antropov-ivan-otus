package ucs

import (
	priorityQueue "kosaraju/priority-queue"
)

func Ucs(g, weights [][]int) []int {
	pq := priorityQueue.NewPriorityQueue[int]()
	visited := make([]bool, len(g))
	order := make([]int, 0)

	for u := range g {
		if !visited[u] {
			pq.Enqueue(u, 0)
			order = dfs(g, weights, visited, pq, order)
		}
	}

	return order
}

func dfs(g, weights [][]int, visited []bool, pq *priorityQueue.PriorityQueue[int], order []int) []int {
	for pq.Length() > 0 {
		u, currentWeight, _ := pq.Dequeue()
		for _, v := range g[u] {
			if !visited[v] {
				visited[v] = true
				nextWeight := weights[u][v]
				pq.Enqueue(v, currentWeight+nextWeight)
			}
		}

		order = append(order, u)
	}

	return order
}
