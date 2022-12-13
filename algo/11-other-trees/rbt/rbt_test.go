package rbt

import (
	"testing"
)

func TestRbtInsert(t *testing.T) {
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

func TestRbtSimpleRemoval1(t *testing.T) {
	tree := buildTree([]int{1})

	tree.Remove(1)
	assertAbsence(t, tree, 1)
}

func TestRbtSimpleRemoval2(t *testing.T) {
	tree := buildTree([]int{1, 2})

	tree.Remove(1)
	assertAbsence(t, tree, 1)
}

func TestRbtSimpleRemoval3(t *testing.T) {
	tree := buildTree([]int{1, 2, 3})

	tree.Remove(1)
	assertAbsence(t, tree, 1)
}

func TestRbtSimpleRemoval4(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4})

	tree.Remove(1)
	assertAbsence(t, tree, 1)
}

func TestRbtSimpleRemoval5(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5})

	tree.Remove(1)
	assertAbsence(t, tree, 1)
}

func TestRbtSimpleRemoval6(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})

	tree.Remove(1)
	assertAbsence(t, tree, 1)
}

func TestRbtSimpleRemoval7(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})

	tree.Remove(1)
	assertAbsence(t, tree, 1)
}

func TestRbtSimpleRemoval8(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})

	tree.Remove(1)
	assertAbsence(t, tree, 2)
}

func TestRbtSimpleRemoval9(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})

	tree.Remove(1)
	assertAbsence(t, tree, 3)
}

func TestRbtSimpleRemoval10(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})

	tree.Remove(1)
	assertAbsence(t, tree, 4)
}

func TestRbtSimpleRemoval11(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})

	tree.Remove(1)
	assertAbsence(t, tree, 5)
}

func TestRbtSimpleRemoval12(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})

	tree.Remove(1)
	assertAbsence(t, tree, 6)
}

func TestRbtComplexRemoval(t *testing.T) {
	tree := buildTree([]int{20, 10, 25, 4, 16, 23, 30, 2, 5, 14, 17, 3, 12, 15, 19, 11})

	assertNodesAreRed(t, tree, []int{10, 14, 3, 19, 11})

	tree.Remove(30)
	if tree.searchNode(30) != nil {
		t.Errorf("The node with %d should be absent", 30)
	}

	assertNodesAreRed(t, tree, []int{10, 19, 23, 3, 11})
}

func buildTree(values []int) *rbtTree {
	tree := NewTree()
	for _, value := range values {
		tree.Insert(value)
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

func assertAbsence(t *testing.T, tree *rbtTree, value int) {
	if tree.searchNode(value) != nil {
		t.Errorf("The node with %d should be absent", value)
	}
}
