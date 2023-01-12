package uf

type Uf struct {
	components      []int
	heights         []int
	componentsCount int
}

// weighted quick-union solution with path compression

func NewUf(count int) *Uf {
	uf := new(Uf)
	uf.componentsCount = count
	uf.components = make([]int, count)
	uf.heights = make([]int, count)
	for i := 0; i < count; i++ {
		uf.components[i] = i
		uf.heights[i] = 1
	}
	return uf
}

func (uf *Uf) Union(u, v int) {
	uRoot := uf.Find(u)
	vRoot := uf.Find(v)

	if uRoot == vRoot {
		return
	}

	if uf.heights[uRoot] > uf.heights[vRoot] {
		uf.components[vRoot] = uRoot
		uf.heights[uRoot] += uf.heights[vRoot]
	} else {
		uf.components[uRoot] = vRoot
		uf.heights[vRoot] += uf.heights[uRoot]
	}

	uf.componentsCount--
}

func (uf *Uf) Find(u int) int {
	node := u
	for ; uf.components[node] != node; node = uf.components[node] {
	}

	uRoot := node
	for p := u; uf.components[p] != uRoot; p = uf.components[p] {
		uf.components[p] = uRoot
	}

	return uRoot
}

func (uf *Uf) ComponentsCount() int {
	return uf.componentsCount
}
