package rbt

import (
	"testing"
)

type directNode struct {
	val int
	red bool
}

func TestRbtSimpleInsert(t *testing.T) {
	tree := NewTree()

	tree.Insert(1)
	tree.Insert(2)
	assertNodesAreRed(t, tree, []int{2})

	tree.Insert(3)
	assertNodesAreRed(t, tree, []int{1, 3})

	tree.Insert(4)
	assertNodesAreRed(t, tree, []int{4})

	tree.Insert(5)
	assertNodesAreRed(t, tree, []int{3, 5})

	tree.Insert(6)
	assertNodesAreRed(t, tree, []int{4, 6})
}

func TestRbtComplexInsert(t *testing.T) {
	tree := buildTree([]int{20, 10, 25, 4, 16, 23, 30, 2, 5, 14, 17, 3, 12, 15, 19, 11})

	assertNodesAreRed(t, tree, []int{14, 20, 3, 11, 19, 23, 30})
	assertNodesAreBlack(t, tree, []int{10, 4, 16, 2, 5, 12, 15})
}

func TestRbtRootRemoval1(t *testing.T) {
	tree := buildTree([]int{1})

	tree.Remove(1)
	assertAbsence(t, tree, 1)
}

func TestRbtRootRemoval2(t *testing.T) {
	tree := buildTree([]int{1, 2})

	tree.Remove(1)
	assertAbsence(t, tree, 1)
}

func TestRbtRootRemoval3(t *testing.T) {
	tree := buildTree([]int{1, 2, 3})

	tree.Remove(2)
	assertAbsence(t, tree, 2)
}

func TestRbtRootRemoval4(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4})

	tree.Remove(2)
	assertAbsence(t, tree, 2)
}

func TestRbtRootRemoval5(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5})

	tree.Remove(2)
	assertAbsence(t, tree, 2)
}

func TestRbtRootRemoval6(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})

	tree.Remove(2)
	assertAbsence(t, tree, 2)
}

func TestRbtSimpleRemoval1(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})

	tree.Remove(1)
	assertAbsence(t, tree, 1)
	assertNodesAreRed(t, tree, []int{3, 6})
	assertNodesAreBlack(t, tree, []int{4, 2, 5})
}

func TestRbtSimpleRemoval2(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})

	tree.Remove(2)
	assertAbsence(t, tree, 2)
	assertNodesAreRed(t, tree, []int{5})
	assertNodesAreBlack(t, tree, []int{1, 3, 4, 6})
}

func TestRbtSimpleRemoval3(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})

	tree.Remove(3)
	assertAbsence(t, tree, 3)
	assertNodesAreRed(t, tree, []int{5})
	assertNodesAreBlack(t, tree, []int{1, 2, 4, 6})
}

func TestRbtSimpleRemoval4(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})

	tree.Remove(4)
	assertAbsence(t, tree, 4)
	assertNodesAreRed(t, tree, []int{5})
	assertNodesAreBlack(t, tree, []int{1, 2, 3, 6})
}

func TestRbtSimpleRemoval5(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})

	tree.Remove(5)
	assertAbsence(t, tree, 5)
	assertNodesAreRed(t, tree, []int{4})
	assertNodesAreBlack(t, tree, []int{1, 2, 3, 6})
}

func TestRbtSimpleRemoval6(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})

	tree.Remove(6)
	assertAbsence(t, tree, 6)
	assertNodesAreRed(t, tree, []int{4})
}

func TestRbtComplexRemoval(t *testing.T) {
	tree := buildTreeDirectly([]directNode{
		{20, false},
		{10, true},
		{25, false},
		{4, false},
		{16, false},
		{23, false},
		{30, false},
		{2, false},
		{5, false},
		{14, true},
		{17, false},
		{3, true},
		{12, false},
		{15, false},
		{19, true},
		{11, true},
	})
	assertNodesAreRed(t, tree, []int{10, 14, 3, 19, 11})
	assertNodesAreBlack(t, tree, []int{20, 25, 4, 16, 23, 30, 2, 5, 17, 12, 15})

	tree.Remove(30)
	assertAbsence(t, tree, 30)
	assertNodesAreRed(t, tree, []int{10, 19, 23, 3, 11})
	assertNodesAreBlack(t, tree, []int{16, 20, 4, 14, 17, 25, 2, 5, 12, 15})
}

func buildTree(values []int) *rbtTree {
	tree := NewTree()
	for _, value := range values {
		tree.Insert(value)
	}
	return tree
}

func buildTreeDirectly(values []directNode) *rbtTree {
	tree := NewTree()
	for _, value := range values {
		tree.insertDirectly(value.val, value.red)
	}
	return tree
}

func assertNodesAreRed(t *testing.T, tree *rbtTree, values []int) {
	for _, val := range values {
		if !tree.searchNode(val).red {
			t.Errorf("The node with %d should be red", val)
		}
	}
}

func assertNodesAreBlack(t *testing.T, tree *rbtTree, values []int) {
	for _, val := range values {
		if tree.searchNode(val).red {
			t.Errorf("The node with %d should be black", val)
		}
	}
}

func assertAbsence(t *testing.T, tree *rbtTree, value int) {
	if tree.searchNode(value) != nil {
		t.Errorf("The node with %d should be absent", value)
	}
}
