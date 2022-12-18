package btree

import "fmt"

type node struct {
	values   []int
	pointers []*node
	size     int
	leaf     bool
	t        int // minimal degree, t >= 2
}

type btree struct {
	root *node
	t    int // minimal degree, t >= 2
}

func NewTree(t int) *btree {
	return &btree{nil, t}
}

func (tree *btree) Search(val int) bool {
	node := tree.root.searchNode(val)
	return node != nil
}

func (tree *btree) Insert(val int) {
	if tree.root == nil {
		tree.root = buildNode(tree.t)
		tree.root.insertIntoNode(val, nil)
		tree.root.leaf = true
		return
	}

	if tree.root.isFull() {
		newRoot := buildNode(tree.t)
		newRoot.pointers[0] = tree.root

		newVal, newRightChild := tree.root.split()
		newRoot.insertIntoNode(newVal, newRightChild)
		tree.root = newRoot
		tree.root.leaf = false
	}

	tree.root.insert(val)
}

func (tree *btree) dump() {
	fmt.Printf("ROOT: ")
	tree.root.dump()
}

func (n *node) searchNode(val int) *node {
	if n == nil {
		return nil
	}

	pos := 0
	for ; pos < n.size && n.values[pos] < val; pos++ {
	}
	if pos == n.size {
		return n.pointers[pos]
	} else if n.values[pos] == val {
		return n
	} else {
		return n.pointers[pos].searchNode(val)
	}
}

func (n *node) insert(val int) {
	if n.size == len(n.values) {
		panic("the node is full")
	}

	if n.leaf {
		n.insertIntoNode(val, nil)
		return
	}

	child := n.findChild(val)
	if child.isFull() {
		newVal, newRightChild := child.split()
		n.insertIntoNode(newVal, newRightChild)
		if val < newVal {
			child.insert(val)
		} else {
			newRightChild.insert(val)
		}
	} else {
		child.insert(val)
	}
}

func (n *node) insertIntoNode(val int, rightPointer *node) {
	if n.size == len(n.values) {
		panic("the leaf is full")
	}

	pos := 0
	for ; pos < n.size && n.values[pos] < val; pos++ {
	}

	for i := n.size; i > pos; i-- {
		n.values[i] = n.values[i-1]
		n.pointers[i+1] = n.pointers[i]
	}

	n.values[pos] = val
	n.pointers[pos+1] = rightPointer
	n.size++
}

func (n *node) findChild(val int) *node {
	pos := 0
	for ; pos < n.size && n.values[pos] < val; pos++ {
	}
	return n.pointers[pos]
}

func (n *node) split() (int, *node) {
	if n.size < n.t-1 {
		panic("the leaf is too sparse")
	}

	pos := n.size / 2

	newVal := n.values[pos]
	n.values[pos] = 0

	newRightChild := buildNode(n.t)
	newRightChild.pointers[0] = n.pointers[pos+1]
	n.pointers[pos+1] = nil
	newRightChild.leaf = n.leaf

	for i := pos + 1; i < n.size; i++ {
		newRightChild.values[i-pos-1] = n.values[i]
		newRightChild.pointers[i-pos] = n.pointers[i+1]
		n.values[i] = 0
		n.pointers[i+1] = nil
	}

	newRightChild.size = n.size - pos - 1
	n.size = pos

	return newVal, newRightChild
}

func (n *node) isFull() bool {
	return n.size == len(n.values)
}

func (n *node) dump() {
	if n == nil {
		return
	}

	fmt.Printf("node: %p, leaf - %v, size - %v, values - %v, pointers - %v\n", n, n.leaf, n.size, n.values, n.pointers)
	for _, pointer := range n.pointers {
		pointer.dump()
	}
}

func buildNode(t int) *node {
	newNode := &node{}
	newNode.values = make([]int, 2*t-1)
	newNode.pointers = make([]*node, 2*t)
	newNode.t = t
	return newNode
}
