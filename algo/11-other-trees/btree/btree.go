package btree

import "fmt"

type node struct {
	values   []int
	pointers []*node
	size     int
	leaf     bool
	tree     *btree
}

type btree struct {
	root *node
	t    int // minimal degree, t >= 2
}

func NewTree(t int) *btree {
	if t < 2 {
		panic("Minimal degree of B tree should be >= 2")
	}
	return &btree{nil, t}
}

func (tree *btree) Search(val int) bool {
	node := tree.root.searchNode(val)
	return node != nil
}

func (tree *btree) Insert(val int) {
	if tree.root == nil {
		tree.root = tree.buildNode()
		tree.root.insertIntoNode(val, nil)
		tree.root.leaf = true
		return
	}

	if tree.root.isFull() {
		newRoot := tree.buildNode()
		newRoot.pointers[0] = tree.root

		newVal, newRightChild := tree.root.split()
		newRoot.insertIntoNode(newVal, newRightChild)
		tree.root = newRoot
		tree.root.leaf = false
	}

	tree.root.insert(val)
}

func (tree *btree) Remove(val int) {
	node := tree.root.searchNode(val)
	if node == nil {
		return
	}

	tree.root.remove(val)
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
	if n.size < n.tree.t-1 {
		panic("the leaf is too sparse")
	}

	pos := n.size / 2

	newVal := n.values[pos]
	n.values[pos] = 0

	newRightChild := n.tree.buildNode()
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

func (n *node) remove(val int) {
	if n.leaf {
		n.removeFromNode(val)
		return
	}

	if n.contains(val) {
		pos := n.findPos(val)
		child := n.pointers[pos]
		nextChild := n.pointers[pos+1]
		if child.size >= n.tree.t {
			n.values[pos] = child.removePredecessor(val)
		} else if nextChild.size >= n.tree.t {
			n.values[pos] = nextChild.removeSuccessor(val)
		} else {
			child.merge(n.values[pos+1], nextChild)
			n.removeFromNode(val)
		}
		return
	}

	pos := n.findPos(val)
	child := n.pointers[pos]
	leftChild := n.pointers[pos-1]
	rightChild := n.pointers[pos+1]
	if child.size < n.tree.t {
		if leftChild.size < n.tree.t {
			if rightChild.size < n.tree.t {
				valueToMoveDown := n.values[pos]
				n.removeFromNode(valueToMoveDown)
				leftChild.merge(valueToMoveDown, child)
				leftChild.remove(val)
				return
			} else {
				child.insertIntoNode(rightChild.values[0], rightChild.pointers[0])
				rightChild.removeFromNode(rightChild.values[0])
				child.remove(val)
				return
			}
		} else {
			child.insertIntoNode(leftChild.values[leftChild.size-1], leftChild.pointers[leftChild.size])
			leftChild.removeFromNode(leftChild.values[leftChild.size-1])
			child.remove(val)
			return
		}
	} else {
		child.remove(val)
	}
}

// TODO
func (n *node) findPos(val int) int {
	pos := 0
	for ; pos < n.size && n.values[pos] < val; pos++ {
	}
	return pos
}

func (n *node) removeFromNode(val int) {
	if n.size < n.tree.t {
		panic("the leaf is too sparse")
	}

	pos := 0
	for ; pos < n.size && n.values[pos] < val; pos++ {
	}

	for i := pos; i < n.size; i++ {
		n.values[i] = n.values[i+1]
		n.pointers[i+1] = n.pointers[i+2]
	}

	n.size--
}

func (n *node) removePredecessor(val int) int {
	return 0
}

func (n *node) removeSuccessor(val int) int {
	return 0
}

func (n *node) merge(val int, sibling *node) {
	if n.size+sibling.size+1 > 2*n.tree.t-1 {
		panic("Too many values for merge")
	}

	n.values[n.size] = val
	n.pointers[n.size+1] = sibling.pointers[0]
	for i, value := range sibling.values {
		n.values[i+n.size+1] = value
		n.pointers[i+n.size+2] = sibling.pointers[i+1]
	}
	n.size += sibling.size + 1
}

func (n *node) isFull() bool {
	return n.size == len(n.values)
}

func (n *node) contains(val int) bool {
	for i := 0; i < n.size; i++ {
		if n.values[i] == val {
			return true
		}
	}
	return false
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

func (tree *btree) buildNode() *node {
	newNode := &node{}
	newNode.values = make([]int, 2*tree.t-1)
	newNode.pointers = make([]*node, 2*tree.t)
	newNode.tree = tree
	return newNode
}
