package treap

import (
	"fmt"
	"math/rand"
)

type node struct {
	val         int
	pr          float64
	left, right *node
}

type treap struct {
	root *node
}

func NewTree() *treap {
	return &treap{}
}

func (tree *treap) Search(val int) bool {
	return tree.root.searchNode(val) != nil
}

func (tree *treap) Insert(val int) {
	newNode := &node{
		val:   val,
		pr:    rand.Float64(),
		left:  nil,
		right: nil,
	}
	tree.root = tree.root.merge(newNode)
}

func (tree *treap) insertDirect(val int, pr float64) {
	newNode := &node{val, pr, nil, nil}
	tree.root = tree.root.insertDirect(newNode)
}

func (tree *treap) Remove(val int) {
	if tree.root == nil {
		return
	}

	left, right := tree.root.split(val)
	tree.root = left.merge(right)
}

func (tree *treap) DumpValuesInDetails() {
	fmt.Print("ROOT: ")
	tree.root.dumpValuesInDetails()
}

func (tree *treap) CheckForInvariants() bool {
	return tree.root.checkForInvariants()
}

func (n *node) insertDirect(nn *node) *node {
	if n == nil {
		return nn
	}

	if nn.val > n.val {
		n.right = n.right.insertDirect(nn)
	} else {
		n.left = n.left.insertDirect(nn)
	}

	return n
}

func (n *node) split(val int) (*node, *node) {
	var leftTree, rightTree *node

	fmt.Printf("Will split %v with left - %v, right - %v\n", n, n.left, n.right)

	if n.val < val {
		leftTree = &node{n.val, n.pr, nil, nil}
		leftTree.left = n.left
		leftTree.right, rightTree = n.right.split(val)
	} else if n.val > val {
		rightTree = &node{n.val, n.pr, nil, nil}
		rightTree.right = n.right
		leftTree, rightTree.left = n.left.split(val)
	} else {
		leftTree = n.left
		rightTree = n.right
	}
	return leftTree, rightTree
}

func (n *node) merge(other *node) *node {
	if n == nil {
		return other
	} else if other == nil {
		return n
	}

	if n.pr > other.pr {
		if n.val > other.val {
			n.left = n.left.merge(other)
		} else if n.val < other.val {
			n.right = n.right.merge(other)
		} else {
			panic("duplicated element for merge-1")
		}
		return n
	} else {
		if other.val > n.val {
			other.left = other.left.merge(n)
		} else if other.val < n.val {
			other.right = other.right.merge(n)
		} else {
			panic("duplicated element for merge-2")
		}
		return other
	}
}

func (n *node) searchNode(val int) *node {
	if n == nil {
		return nil
	}

	if n.val < val {
		return n.right.searchNode(val)
	} else if n.val > val {
		return n.left.searchNode(val)
	} else {
		return n
	}
}

func (n *node) dup() *node {
	if n == nil {
		return nil
	}

	newNode := &node{n.val, n.pr, nil, nil}
	newNode.left = n.left.dup()
	newNode.right = n.right.dup()
	return newNode
}

func (n *node) dumpValuesInDetails() {
	if n == nil {
		return
	}

	fmt.Printf("node - value: %d, priority: %v, left - %v, right - %v\n", n.val, n.pr, n.left, n.right)
	n.left.dumpValuesInDetails()
	n.right.dumpValuesInDetails()
}

func (n *node) checkForInvariants() bool {
	if n == nil {
		return true
	}

	result := true
	if n.left != nil && n.left.val > n.val {
		fmt.Printf("Node %v has an invalid value from the left - %d\n", n, n.left.val)
		result = false
	}

	if n.right != nil && n.right.val < n.val {
		fmt.Printf("Node %v has an invalid value from the right - %d\n", n, n.right.val)
		result = false
	}

	if n.left != nil && n.left.pr > n.pr {
		fmt.Printf("Node %v has an invalid priority from the left - %v\n", n, n.left.pr)
		result = false
	}

	if n.right != nil && n.right.pr > n.pr {
		fmt.Printf("Node %v has an invalid priority from the right - %v\n", n, n.left.pr)
		result = false
	}

	result = result && n.left.checkForInvariants()
	result = result && n.right.checkForInvariants()

	return result
}
