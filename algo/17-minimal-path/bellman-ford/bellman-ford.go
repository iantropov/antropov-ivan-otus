package bellmanford

import "math"

type edge struct {
	u, v, w int
}

func BellmanFord(g [][]int, u, v int) []int {
	edges := make([]edge, len(g))
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] > 0 {
				edges = append(edges, edge{
					u: i,
					v: j,
					w: g[i][j],
				})
			}
		}
	}

	priors := make([]int, len(g))

	distancesFromU := make([]int, len(g))
	for i := 0; i < len(g); i++ {
		distancesFromU[i] = math.MaxInt
	}
	distancesFromU[u] = 0

	for i := 0; i < len(g)-1; i++ {
		for _, e := range edges {
			if distancesFromU[e.v] > distancesFromU[e.u]+e.w {
				distancesFromU[e.v] = distancesFromU[e.u] + e.w
				priors[e.v] = e.u
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
