package btree

import (
	"fmt"
	"math"
)

type node struct {
	values   []int
	pointers []*node
	size     int
	leaf     bool
	tree     *Btree
}

type Btree struct {
	root *node
	t    int // minimal degree, t >= 2
}

type directNode struct {
	values []int
	leaf   bool
	root   bool
	parent int
	child  int
}

func NewTree(t int) *Btree {
	if t < 2 {
		panic("Minimal degree of B tree should be >= 2")
	}
	return &Btree{nil, t}
}

func (tree *Btree) Search(val int) bool {
	node := tree.root.searchNode(val)
	return node != nil
}

func (tree *Btree) Insert(val int) {
	if tree.root == nil {
		tree.root = tree.buildNode()
		tree.root.insertToRight(val, nil)
		tree.root.leaf = true
		return
	}

	if tree.root.isFull() {
		newRoot := tree.buildNode()
		newRoot.pointers[0] = tree.root

		newVal, newRightChild := tree.root.split()
		newRoot.insertToRight(newVal, newRightChild)
		tree.root = newRoot
		tree.root.leaf = false
	}

	tree.root.insert(val)
}

func (tree *Btree) Remove(val int) {
	node := tree.root.searchNode(val)
	if node == nil {
		return
	}

	if tree.root.size == 1 {
		leftChild := tree.root.pointers[0]
		rightChild := tree.root.pointers[1]
		if leftChild.size == tree.t-1 && rightChild.size == tree.t-1 {
			leftChild.merge(tree.root.values[0], rightChild)
			tree.root = leftChild
		}
	}

	tree.root.remove(val)
}

func (tree *Btree) DumpValuesInDetails() {
	fmt.Printf("ROOT: ")
	tree.root.dump()
}

func (tree *Btree) BuildDirectly(directNodes []directNode) {
	nodes := make([]*node, len(directNodes))
	for i, dn := range directNodes {
		newNode := tree.buildNode()
		copy(newNode.values, dn.values)
		newNode.size = len(dn.values)
		newNode.leaf = dn.leaf
		nodes[i] = newNode
		if !dn.root {
			nodes[dn.parent].pointers[dn.child] = newNode
		}
	}
	tree.root = nodes[0]
}

func (tree *Btree) CheckForInvariants() bool {
	return tree.root.checkForInvariants(true, math.MinInt, math.MaxInt)
}

func (n *node) checkForInvariants(isRoot bool, min, max int) bool {
	if n == nil {
		return true
	}

	result := true
	nonZeroValues := 0
	nonNilPointers := 0
	for i := 0; i < 2*n.tree.t-1; i++ {
		if n.values[i] != 0 {
			nonZeroValues++
		}

		if n.pointers[i] != nil {
			nonNilPointers++
		}
	}
	if n.pointers[2*n.tree.t-1] != nil {
		nonNilPointers++
	}

	if nonZeroValues != n.size {
		fmt.Println("Invalid number of values")
		result = false
	}

	if n.leaf {
		if nonNilPointers != 0 {
			fmt.Println("Pointers should be empty for leaves")
			result = false
		}
	} else {
		if nonNilPointers != n.size+1 {
			fmt.Println("Invalid number of pointers for non-leaf")
			result = false
		}
	}

	if !isRoot && n.size < n.tree.t-1 {
		fmt.Println("Node is too small")
		result = false
	}

	if isRoot && n.size == 0 {
		fmt.Println("Empty root")
		result = false
	}

	result = result && n.checkValues(min, max)

	result = result && n.pointers[0].checkForInvariants(false, min, n.values[0])
	for i := 1; i < n.size; i++ {
		result = result && n.pointers[i].checkForInvariants(false, n.values[i-1], n.values[i])
	}
	result = result && n.pointers[n.size].checkForInvariants(false, n.values[n.size-1], max)

	return result
}

func (n *node) checkValues(min, max int) bool {
	if n == nil {
		return true
	}

	if n.values[0] <= min || n.values[0] >= max {
		fmt.Printf("Value %d from node %p - invalid!\n", n.values[0], n)
		return false
	}
	for i := 1; i < n.size; i++ {
		if n.values[i] <= min || n.values[i] >= max {
			fmt.Printf("Value %d from node %p - is outside the interval!\n", n.values[i], n)
			return false
		}
		if n.values[i] <= n.values[i-1] {
			fmt.Printf("Value %d from node %p - is unsorted!\n", n.values[i], n)
			return false
		}
	}
	return true
}

func (n *node) searchNode(val int) *node {
	if n == nil {
		return nil
	}

	pos := n.findPos(val)
	if pos < n.size && n.values[pos] == val {
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
		n.insertToRight(val, nil)
		return
	}

	pos := n.findPos(val)
	child := n.pointers[pos]
	if child.isFull() {
		newVal, newRightChild := child.split()
		n.insertToRight(newVal, newRightChild)
		if val < newVal {
			child.insert(val)
		} else {
			newRightChild.insert(val)
		}
	} else {
		child.insert(val)
	}
}

func (n *node) insertToRight(val int, rightPointer *node) {
	if n.size == len(n.values) {
		panic("Too many values for insertion to right")
	}

	pos := n.findPos(val)
	for i := n.size; i > pos; i-- {
		n.values[i] = n.values[i-1]
		n.pointers[i+1] = n.pointers[i]
	}

	n.values[pos] = val
	n.pointers[pos+1] = rightPointer
	n.size++
}

func (n *node) insertToLeft(val int, pointer *node) {
	if n.size == len(n.values) {
		panic("Too many values for insertion to left")
	}

	for i := n.size; i > 0; i-- {
		n.values[i] = n.values[i-1]
		n.pointers[i+1] = n.pointers[i]
	}
	n.pointers[1] = n.pointers[0]
	n.values[0] = val
	n.pointers[0] = pointer
	n.size++
}

func (n *node) findPos(val int) int {
	pos := 0
	for ; pos < n.size && n.values[pos] < val; pos++ {
	}
	return pos
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

func (n *node) merge(val int, sibling *node) {
	if n.size+sibling.size+1 > 2*n.tree.t-1 {
		panic("Too many values for merge")
	}

	n.values[n.size] = val
	n.pointers[n.size+1] = sibling.pointers[0]
	for i := 0; i < sibling.size; i++ {
		n.values[i+n.size+1] = sibling.values[i]
		n.pointers[i+n.size+2] = sibling.pointers[i+1]
	}
	n.size += sibling.size + 1
}

func (n *node) remove(val int) {
	if n.leaf {
		n.removeFromRight(val)
		return
	}

	pos := n.findPos(val)
	child := n.pointers[pos]
	if pos < n.size && n.values[pos] == val {
		rightChild := n.pointers[pos+1]
		if child.size >= n.tree.t {
			predecessor := child.findMax()
			n.values[pos] = predecessor
			child.remove(predecessor)
		} else if rightChild.size >= n.tree.t {
			successor := rightChild.findMin()
			n.values[pos] = successor
			rightChild.remove(successor)
		} else {
			child.merge(val, rightChild)
			n.removeFromRight(val)
			child.remove(val)
		}
		return
	}

	var leftChild *node = nil
	var rightChild *node = nil
	if pos > 0 {
		leftChild = n.pointers[pos-1]
	}
	if pos < n.size {
		rightChild = n.pointers[pos+1]
	}
	if child.size >= n.tree.t {
		child.remove(val)
	} else {
		if leftChild != nil && leftChild.size >= n.tree.t {
			child.insertToLeft(n.values[pos-1], leftChild.pointers[leftChild.size])
			n.values[pos-1] = leftChild.values[leftChild.size-1]
			leftChild.removeFromRight(leftChild.values[leftChild.size-1])
			child.remove(val)
			return
		}

		if rightChild != nil && rightChild.size >= n.tree.t {
			child.insertToRight(n.values[pos], rightChild.pointers[0])
			n.values[pos] = rightChild.values[0]
			rightChild.removeFromLeft(rightChild.values[0])
			child.remove(val)
			return
		}

		if leftChild != nil {
			leftChild.merge(n.values[pos-1], child)
			n.removeFromRight(n.values[pos-1])
			leftChild.remove(val)
			return
		}

		if rightChild != nil {
			child.merge(n.values[pos], rightChild)
			n.removeFromRight(n.values[pos])
			child.remove(val)
			return
		}

		panic("invalid case")
	}
}

func (n *node) findMax() int {
	child := n
	for ; child.pointers[child.size] != nil; child = child.pointers[child.size] {

	}
	return child.values[child.size-1]
}

func (n *node) findMin() int {
	child := n
	for ; child.pointers[0] != nil; child = child.pointers[0] {

	}
	return child.values[0]
}

func (n *node) removeFromRight(val int) {
	if n.size < n.tree.t-1 {
		panic("the leaf is too sparse")
	}

	pos := n.findPos(val)
	for i := pos; i < n.size-1; i++ {
		n.values[i] = n.values[i+1]
		n.pointers[i+1] = n.pointers[i+2]
	}

	n.values[n.size-1] = 0
	n.pointers[n.size] = nil
	n.size--
}

func (n *node) removeFromLeft(val int) {
	if n.size < n.tree.t-1 {
		panic("the leaf is too sparse")
	}

	pos := n.findPos(val)
	for i := pos; i < n.size-1; i++ {
		n.values[i] = n.values[i+1]
		n.pointers[i] = n.pointers[i+1]
	}
	n.pointers[n.size-1] = n.pointers[n.size]

	n.values[n.size-1] = 0
	n.pointers[n.size] = nil
	n.size--
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

func (tree *Btree) buildNode() *node {
	newNode := &node{}
	newNode.values = make([]int, 2*tree.t-1)
	newNode.pointers = make([]*node, 2*tree.t)
	newNode.tree = tree
	return newNode
}
