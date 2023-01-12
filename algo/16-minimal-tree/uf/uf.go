package uf

type Uf struct {
	components      []int
	componentsCount int
}

// non-weighted quick-find solution with path compression

func NewUf(count int) *Uf {
	uf := new(Uf)
	uf.componentsCount = count
	uf.components = make([]int, count)
	for i := 0; i < count; i++ {
		uf.components[i] = i
	}
	return uf
}

func (uf *Uf) Union(u, v int) {
	uRoot := uf.Find(u)
	vRoot := uf.Find(v)

	if uRoot == vRoot {
		return
	}

	for i, w := range uf.components {
		if w == uRoot {
			uf.components[i] = vRoot
		}
	}

	uf.componentsCount--
}

func (uf *Uf) Find(u int) int {
	uRoot := uf.components[u]

	for p := u; uf.components[p] != uRoot; p = uf.components[p] {
		uf.components[p] = uRoot
	}

	return uRoot
}

func (uf *Uf) ComponentsCount() int {
	return uf.componentsCount
}

// package weightedQuickUnion

// type weightedQuickUnion struct {
// 	components      []int
// 	heights         []int
// 	componentsCount int
// }

// func NewUnionFind(size int) *weightedQuickUnion {
// 	newComponents := make([]int, size)
// 	newHeights := make([]int, size)
// 	for i := range newComponents {
// 		newComponents[i] = i
// 		newHeights[i] = 1
// 	}
// 	return &weightedQuickUnion{components: newComponents, componentsCount: size, heights: newHeights}
// }

// func (qu *weightedQuickUnion) Union(p, q int) {
// 	pRoot := qu.Find(p)
// 	qRoot := qu.Find(q)

// 	if pRoot == qRoot {
// 		return
// 	}

// 	if qu.heights[pRoot] > qu.heights[qRoot] {
// 		qu.components[qRoot] = pRoot
// 		qu.heights[pRoot] += qu.heights[qRoot]
// 	} else {
// 		qu.components[pRoot] = qRoot
// 		qu.heights[qRoot] += qu.heights[pRoot]
// 	}

// 	qu.componentsCount--
// }

// func (qu *weightedQuickUnion) Find(p int) int {
// 	root := p
// 	for ; qu.components[root] != root; root = qu.components[root] {
// 	}

// 	// path compression

// 	node := p
// 	nextNode := node
// 	for {
// 		if nextNode == root {
// 			break
// 		}
// 		nextNode = qu.components[nextNode]
// 		qu.components[node] = root
// 	}

// 	return root
// }

// func (qu *weightedQuickUnion) Connected(p, q int) bool {
// 	return qu.Find(p) == qu.Find(q)
// }

// func (qu *weightedQuickUnion) Count() int {
// 	return qu.componentsCount
// }
