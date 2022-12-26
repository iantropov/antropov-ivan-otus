package dfs

type dfs struct {
	levelNodes []int
	g          [][]int
	visited    []bool
	finished   bool
}

func NewDfs(g [][]int) *dfs {
	newDfs := &dfs{}

	newDfs.visited = make([]bool, len(g))
	newDfs.levelNodes = make([]int, len(g))
	for i := range g {
		newDfs.levelNodes[i] = i
	}
	newDfs.g = g

	return newDfs
}

func (d *dfs) VisitLevel() []int {
	if d.finished {
		return d.levelNodes
	}

	visitingNodes := d.levelNodes
	d.levelNodes = make([]int, 0)
	planned := make([]bool, len(d.g))
	visited := make([]int, 0)

	for _, u := range visitingNodes {
		if planned[u] {
			continue
		}

		d.visited[u] = true
		visited = append(visited, u)
		for _, v := range d.g[u] {
			if d.visited[v] || planned[v] {
				continue
			}
			d.levelNodes = append(d.levelNodes, v)
			planned[v] = true
		}
	}

	if len(d.levelNodes) == 0 {
		d.finished = true
	}

	return visited
}

func (d *dfs) IsFinished() bool {
	return d.finished
}
