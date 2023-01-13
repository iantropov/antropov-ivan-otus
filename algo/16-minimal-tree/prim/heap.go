package prim

import "math"

type heap struct {
	elements []*node
	keys     []int
	size     int
	count    int
}

type node struct {
	key, priority int
	extracted     bool
}

func NewHeap(size int) *heap {
	h := new(heap)
	h.size = size
	h.count = size
	h.elements = make([]*node, size)
	h.keys = make([]int, size)
	for i := 0; i < size; i++ {
		h.elements[i] = &node{key: i, priority: math.MaxInt}
		h.keys[i] = i
	}
	return h
}

func (h *heap) DecreaseKey(key, newPriority int) {
	i := h.keys[key]
	n := h.elements[i]
	n.priority = newPriority
	h.swim(i)
}

func (h *heap) ExtractMin() (key, priority int) {
	minKey, minPriority := h.elements[0].key, h.elements[0].priority

	h.elements[0].priority = math.MaxInt
	h.elements[0].extracted = true
	h.sink(0)
	h.count--

	return minKey, minPriority
}

func (h *heap) Empty() bool {
	return h.count == 0
}

func (h *heap) Contains(key int) bool {
	return !h.elements[h.keys[key]].extracted
}

func (h *heap) Priority(key int) int {
	return h.elements[h.keys[key]].priority
}

func (h *heap) swim(i int) {
	parent := (i - 1) / 2
	if parent >= 0 && h.elements[i].priority < h.elements[parent].priority {
		h.swap(parent, i)
		h.swim(parent)
	}
}

func (h *heap) sink(i int) {
	root := i
	left := i*2 + 1
	right := i*2 + 2

	node := root
	if left < h.size && h.elements[left].priority < h.elements[node].priority {
		node = left
	}
	if right < h.size && h.elements[right].priority < h.elements[node].priority {
		node = right
	}
	if node == root {
		return
	}
	h.swap(root, node)
	h.sink(node)
}

func (h *heap) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
	h.keys[h.elements[i].key], h.keys[h.elements[j].key] = h.keys[h.elements[j].key], h.keys[h.elements[i].key]
}
