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
	// tree.root = mergeTwo(tree.root, newNode)
}

func (tree *treap) InsertWithPriority(val int, pr float64) {
	newNode := &node{
		val:   val,
		pr:    pr,
		left:  nil,
		right: nil,
	}
	tree.root = tree.root.merge(newNode)
	// tree.root = mergeTwo(tree.root, newNode)
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
		return other.dup()
	} else if other == nil {
		return n.dup()
	}

	if n.pr > other.pr {
		if n.val > other.val {
			n.right = n.right.merge(other.right)
			otherRight := other.right
			other.right = nil
			n.left = n.left.merge(other)
			other.right = otherRight
		} else {
			n.left = n.left.merge(other.left)
			otherLeft := other.left
			other.left = nil
			n.right = n.right.merge(other)
			other.left = otherLeft
		}
		return n
	} else {
		if other.val > n.val {
			other.right = other.right.merge(n.right)
			nRight := n.right
			n.right = nil
			other.left = other.left.merge(n)
			n.right = nRight
		} else {
			other.left = other.left.merge(n.left)
			nLeft := n.left
			n.left = nil
			other.right = other.right.merge(n)
			n.left = nLeft
		}
		return other
	}
}

// func mergeTwo(n, other *node) *node {
// 	if n == nil {
// 		return other.dup()
// 	} else if other == nil {
// 		return n.dup()
// 	}

// 	if n.pr > other.pr {
// 		if n.val > other.val {
// 			n.right = n.right.merge(other.right)
// 			otherRight := other.right
// 			other.right = nil
// 			n.left = n.left.merge(other)
// 			other.right = otherRight
// 		} else {
// 			n.left = n.left.merge(other.left)
// 			otherLeft := other.left
// 			other.left = nil
// 			n.right = n.right.merge(other)
// 			other.left = otherLeft
// 		}
// 		return n
// 	} else {
// 		if other.val > n.val {
// 			other.right = other.right.merge(n.right)
// 			nRight := n.right
// 			n.right = nil
// 			other.left = other.left.merge(n)
// 			n.right = nRight
// 		} else {
// 			other.left = other.left.merge(n.left)
// 			nLeft := n.left
// 			n.left = nil
// 			other.right = other.right.merge(n)
// 			n.left = nLeft
// 		}
// 		return other
// 	}
// }

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
	if !n.left.checkForMaxValue(n.val) {
		fmt.Printf("Node %v has an invalid value in the left subtree\n", n)
		result = false
	}

	if !n.right.checkForMinValue(n.val) {
		fmt.Printf("Node %v has an invalid value in the right subtree\n", n)
		result = false
	}

	if !n.left.checkForMaxPriority(n.pr) {
		fmt.Printf("Node %v has an invalid priority in the left subtree\n", n)
		result = false
	}

	if !n.right.checkForMaxPriority(n.pr) {
		fmt.Printf("Node %v has an invalid priority in the right subtree\n", n)
		result = false
	}

	result = result && n.left.checkForInvariants()
	result = result && n.right.checkForInvariants()

	return result
}

func (n *node) checkForMaxValue(max int) bool {
	if n == nil {
		return true
	} else {
		return n.val < max && n.left.checkForMaxValue(max) && n.right.checkForMaxValue(max)
	}
}

func (n *node) checkForMinValue(min int) bool {
	if n == nil {
		return true
	} else {
		return n.val > min && n.left.checkForMinValue(min) && n.right.checkForMinValue(min)
	}
}

func (n *node) checkForMaxPriority(max float64) bool {
	if n == nil {
		return true
	} else {
		return n.pr < max && n.left.checkForMaxPriority(max) && n.right.checkForMaxPriority(max)
	}
}
