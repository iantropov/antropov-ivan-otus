package random

import (
	"fmt"
	"math/rand"
)

type randomTree struct {
	root *node
	size int
}

type node struct {
	value       int
	left, right *node
}

func NewTree() *randomTree {
	return &randomTree{}
}

func (tree *randomTree) InsertWithoutSplay(val int) {
	tree.root = tree.root.insert(val)
}

func (tree *randomTree) Insert(val int) {
	tree.root = tree.root.insert(val)
	tree.size++
	if rand.Intn(tree.size+1) == tree.size {
		tree.root = tree.root.splay(val)
	}
}

func (tree *randomTree) Search(val int) bool {
	node := tree.root.searchNode(val)
	if node == nil {
		return false
	}
	if rand.Intn(tree.size+1) == tree.size {
		tree.root = tree.root.splay(val)
	}
	return true
}

func (tree *randomTree) Remove(val int) {
	node := tree.root.searchNode(val)
	if node == nil {
		return
	}
	tree.size--
	tree.root = tree.root.splay(val)
	tree.root = tree.root.left.merge(tree.root.right)
}

func (t *randomTree) DumpValuesInDetails() {
	t.root.dumpValuesInDetails()
}

func (n *node) insert(val int) *node {
	if n == nil {
		return &node{value: val}
	}

	if val < n.value {
		n.left = n.left.insert(val)
	} else if val > n.value {
		n.right = n.right.insert(val)
	}

	return n
}

func (node *node) searchNode(val int) *node {
	if node == nil {
		return nil
	}

	if val < node.value {
		return node.left.searchNode(val)
	} else if val > node.value {
		return node.right.searchNode(val)
	} else {
		return node
	}
}

func (node *node) splay(val int) *node {
	if node == nil {
		return nil
	} else if node.value == val {
		return node
	}

	if val < node.value {
		if node.left.value == val {
			return node.zigLeft()
		} else if node.left.left != nil && node.left.left.value == val {
			return node.zigZigLeft()
		} else if node.left.right != nil && node.left.right.value == val {
			return node.zigZagLeft()
		} else {
			node.left = node.left.splay(val)
		}
	} else {
		if node.right.value == val {
			return node.zigRight()
		} else if node.right.left != nil && node.right.left.value == val {
			return node.zigZagRight()
		} else if node.right.right != nil && node.right.right.value == val {
			return node.zigZigRight()
		} else {
			node.right = node.right.splay(val)
		}
	}

	return node.splay(val)
}

func (node *node) zigLeft() *node {
	newNode := node.left
	node.left = node.left.right
	newNode.right = node
	return newNode
}

func (node *node) zigRight() *node {
	newNode := node.right
	node.right = node.right.left
	newNode.left = node
	return newNode
}

func (node *node) zigZigLeft() *node {
	newNode := node.left.left

	oldNodeLeft := node.left
	oldNodeLeftRight := node.left.right

	node.left.left = node.left.left.right
	node.left.right = node
	node.left = oldNodeLeftRight

	newNode.right = oldNodeLeft

	return newNode
}

func (node *node) zigZigRight() *node {
	newNode := node.right.right

	oldNodeRight := node.right
	oldNodeRightLeft := node.right.left

	node.right.right = node.right.right.left
	node.right.left = node
	node.right = oldNodeRightLeft

	newNode.left = oldNodeRight

	return newNode
}

func (node *node) zigZagLeft() *node {
	newNode := node.left.right

	oldNodeLeftRightRight := node.left.right.right
	oldNodeLeft := node.left

	node.left.right = node.left.right.left
	node.left = oldNodeLeftRightRight

	newNode.left = oldNodeLeft
	newNode.right = node

	return newNode
}

func (node *node) zigZagRight() *node {
	newNode := node.right.left

	oldNodeRightLeftLeft := node.right.left.left
	oldNodeRight := node.right

	node.right.left = node.right.left.right
	node.right = oldNodeRightLeftLeft

	newNode.left = node
	newNode.right = oldNodeRight

	return newNode
}

func (node *node) split(val int) (*node, *node) {
	splitNode := node.splay(val)
	return splitNode.left, splitNode.right
}

func (n *node) merge(other *node) *node {
	if n == nil {
		return other
	}
	maxNode := n.findMax()
	newRoot := n.splay(maxNode.value)
	newRoot.right = other
	return newRoot
}

func (n *node) findMax() *node {
	max := n
	for ; max.right != nil; max = max.right {

	}
	return max
}

func (n *node) dumpValuesInDetails() {
	if n == nil {
		return
	}

	n.left.dumpValuesInDetails()
	fmt.Printf("Node: %d, left - %v, right = %v\n", n.value, n.left, n.right)
	n.right.dumpValuesInDetails()
}
