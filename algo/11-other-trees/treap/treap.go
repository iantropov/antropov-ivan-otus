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
	tree.InsertWithPriority(val, rand.Float64())
}

func (tree *treap) InsertWithPriority(val int, pr float64) {
	newNode := &node{
		val:   val,
		pr:    pr,
		left:  nil,
		right: nil,
	}
	l, r := tree.root.split(val)
	l = l.merge(newNode)
	tree.root = l.merge(r)
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

	if n == nil {
		return nil, nil
	}

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

func (l *node) merge(r *node) *node {
	if l == nil {
		return r
	} else if r == nil {
		return l
	}

	newNode := &node{}
	if l.pr > r.pr {
		newNode.val = l.val
		newNode.pr = l.pr
		newNode.left = l.left
		newNode.right = l.right.merge(r)
	} else {
		newNode.val = r.val
		newNode.pr = r.pr
		newNode.right = r.right
		newNode.left = l.merge(r.left)
	}
	return newNode
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
