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

	assertTreeInvariants(t, tree)
}

func TestRbtMediumInsert(t *testing.T) {
	tree := buildTree([]int{1, 3, 0, 7, 5})
	assertNodesAreRed(t, tree, []int{3, 7})
	assertNodesAreBlack(t, tree, []int{1, 0, 5})
	assertTreeInvariants(t, tree)
}

func TestRbtComplexInsert(t *testing.T) {
	tree := buildTree([]int{20, 10, 25, 4, 16, 23, 30, 2, 5, 14, 17, 3, 12, 15, 19, 11})

	assertNodesAreRed(t, tree, []int{14, 20, 3, 11, 19, 23, 30})
	assertNodesAreBlack(t, tree, []int{10, 4, 16, 2, 5, 12, 15})
	assertTreeInvariants(t, tree)
}

func TestRbtRootRemoval1(t *testing.T) {
	tree := buildTree([]int{1})

	tree.Remove(1)
	assertAbsence(t, tree, 1)
	assertTreeInvariants(t, tree)
}

func TestRbtRootRemoval2(t *testing.T) {
	tree := buildTree([]int{1, 2})

	tree.Remove(1)
	assertAbsence(t, tree, 1)
	assertTreeInvariants(t, tree)
}

func TestRbtRootRemoval3(t *testing.T) {
	tree := buildTree([]int{1, 2, 3})

	tree.Remove(2)
	assertAbsence(t, tree, 2)
	assertTreeInvariants(t, tree)
}

func TestRbtRootRemoval4(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4})

	tree.Remove(2)
	assertAbsence(t, tree, 2)
	assertTreeInvariants(t, tree)
}

func TestRbtRootRemoval5(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5})

	tree.Remove(2)
	assertAbsence(t, tree, 2)
	assertTreeInvariants(t, tree)
}

func TestRbtRootRemoval6(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})

	tree.Remove(2)
	assertAbsence(t, tree, 2)
	assertTreeInvariants(t, tree)
}

func TestRbtSimpleRemoval1(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})

	tree.Remove(1)
	assertAbsence(t, tree, 1)
	assertNodesAreRed(t, tree, []int{3, 6})
	assertNodesAreBlack(t, tree, []int{4, 2, 5})
	assertTreeInvariants(t, tree)
}

func TestRbtSimpleRemoval2(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})

	tree.Remove(2)
	assertAbsence(t, tree, 2)
	assertNodesAreRed(t, tree, []int{5})
	assertNodesAreBlack(t, tree, []int{1, 3, 4, 6})
	assertTreeInvariants(t, tree)
}

func TestRbtSimpleRemoval3(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})

	tree.Remove(3)
	assertAbsence(t, tree, 3)
	assertNodesAreRed(t, tree, []int{5})
	assertNodesAreBlack(t, tree, []int{1, 2, 4, 6})
	assertTreeInvariants(t, tree)
}

func TestRbtSimpleRemoval4(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})

	tree.Remove(4)
	assertAbsence(t, tree, 4)
	assertNodesAreRed(t, tree, []int{5})
	assertNodesAreBlack(t, tree, []int{1, 2, 3, 6})
	assertTreeInvariants(t, tree)
}

func TestRbtSimpleRemoval5(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})

	tree.Remove(5)
	assertAbsence(t, tree, 5)
	assertNodesAreRed(t, tree, []int{4})
	assertNodesAreBlack(t, tree, []int{1, 2, 3, 6})
	assertTreeInvariants(t, tree)
}

func TestRbtSimpleRemoval6(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})

	tree.Remove(6)
	assertAbsence(t, tree, 6)
	assertNodesAreRed(t, tree, []int{4})
	assertTreeInvariants(t, tree)
}

func TestRbtSimpleRemoval7(t *testing.T) {
	tree := buildTreeDirectly([]directNode{
		{3, false},
		{0, false},
		{7, true},
		{5, false},
		{8, false},
		{6, true},
	})
	assertNodesAreRed(t, tree, []int{7, 6})
	assertNodesAreBlack(t, tree, []int{3, 0, 5, 8})

	tree.Remove(8)
	assertAbsence(t, tree, 8)
	assertNodesAreRed(t, tree, []int{5, 7})
	assertNodesAreBlack(t, tree, []int{3, 0, 6})
	assertTreeInvariants(t, tree)
}

func TestRbtMediumRemoval1(t *testing.T) {
	tree := buildTree([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	tree.Remove(4)
	assertAbsence(t, tree, 4)
	assertNodesAreRed(t, tree, []int{6, 9})
	assertNodesAreBlack(t, tree, []int{3, 1, 7, 0, 2, 5, 8})
	assertTreeInvariants(t, tree)

	tree.DumpValuesInDetails()

	tree.Remove(1)
	assertAbsence(t, tree, 1)
	assertNodesAreRed(t, tree, []int{0, 7, 6, 9})
	assertNodesAreBlack(t, tree, []int{3, 2, 5, 8})
	assertTreeInvariants(t, tree)
}

func TestRbtMediumRemoval2(t *testing.T) {
	tree := buildTreeDirectly([]directNode{
		{15, false},
		{7, false},
		{23, false},
		{3, false},
		{11, false},
		{19, false},
		{31, true},
		{27, false},
		{39, false},
	})
	assertNodesAreRed(t, tree, []int{31})
	assertNodesAreBlack(t, tree, []int{15, 7, 23, 27, 39})

	tree.Remove(3)
	assertAbsence(t, tree, 3)
	assertNodesAreRed(t, tree, []int{11})
	assertNodesAreBlack(t, tree, []int{23, 15, 31, 7, 19, 27, 39})
	assertTreeInvariants(t, tree)
}

func TestRbtMediumRemoval3(t *testing.T) {
	tree := buildTreeDirectly([]directNode{
		{8, false},
		{5, false},
		{11, false},
		{2, false},
		{6, false},
		{9, false},
		{15, false},
		{0, true},
		{13, true},
	})
	assertNodesAreRed(t, tree, []int{0, 13})
	assertNodesAreBlack(t, tree, []int{8, 5, 2, 6, 9, 15})

	tree.Remove(5)
	assertAbsence(t, tree, 5)
	assertNodesAreRed(t, tree, []int{13})
	assertNodesAreBlack(t, tree, []int{8, 2, 0, 6, 11, 9, 15})
	assertTreeInvariants(t, tree)
}

func TestRbtComplexRemoval1(t *testing.T) {
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
	assertTreeInvariants(t, tree)

	tree.Remove(30)
	assertAbsence(t, tree, 30)
	assertNodesAreRed(t, tree, []int{10, 19, 23, 3, 11})
	assertNodesAreBlack(t, tree, []int{16, 20, 4, 14, 17, 25, 2, 5, 12, 15})
	assertTreeInvariants(t, tree)
}

func TestRbtComplexRemoval2(t *testing.T) {
	tree := buildTreeDirectly([]directNode{
		{35, false},
		{26, false},
		{45, false},
		{17, false},
		{33, false},
		{40, false},
		{48, false},
		{7, true},
		{21, false},
		{31, false},
		{34, false},
		{37, false},
		{43, true},
		{47, false},
		{49, false},
		{3, false},
		{11, false},
		{42, false},
		{44, false},
		{50, true},
		{14, true},
	})
	assertNodesAreRed(t, tree, []int{7, 43, 14, 50})
	assertNodesAreBlack(t, tree, []int{35, 26, 45, 17, 33, 40, 48, 21, 31, 34, 37, 47, 49, 3, 11, 42, 44})
	assertTreeInvariants(t, tree)

	tree.Remove(21)
	assertAbsence(t, tree, 21)
	// assertNodesAreRed(t, tree, []int{10, 19, 23, 3, 11})
	// assertNodesAreBlack(t, tree, []int{16, 20, 4, 14, 17, 25, 2, 5, 12, 15})
	assertTreeInvariants(t, tree)
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

func assertTreeInvariants(t *testing.T, tree *rbtTree) {
	if !tree.checkForInvariants() {
		t.Error("The tree violates RBT invariants!")
	}
}
