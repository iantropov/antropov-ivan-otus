package avl

import (
	"fmt"
	"testing"
)

func TestTreeInsert(t *testing.T) {
	tree := prepareTree([]int{15, 10, 20, 5, 7, 6, 8})
	checkTreeOrder(t, tree)
}

func TestTreeSearch(t *testing.T) {
	values := []int{15, 10, 20, 5, 7, 6, 8}
	tree := prepareTree(values)

	checkPresenceFor(t, tree, values)

	checkAbsenceFor(t, tree, []int{17, 27, 0})
}

func TestTreeInvalidRemoval(t *testing.T) {
	values := []int{15, 10, 20, 5, 7, 6, 8, 21}
	tree := prepareTree(values)
	tree.Remove(70)
	checkPresenceFor(t, tree, values)
}

func TestTreeSmallRemovals(t *testing.T) {
	tree := prepareTree([]int{15, 10, 20, 5, 7, 6, 8, 21})
	tree.Remove(21)
	checkPresenceFor(t, tree, []int{15, 10, 20, 5, 7, 6, 8})
	checkTreeOrder(t, tree)

	tree.Remove(5)
	checkPresenceFor(t, tree, []int{15, 10, 20, 7, 6, 8})
	checkTreeOrder(t, tree)

	tree.Remove(6)
	checkPresenceFor(t, tree, []int{15, 10, 20, 7, 8})
	checkTreeOrder(t, tree)

	tree.Remove(8)
	checkPresenceFor(t, tree, []int{15, 10, 20, 7})
	checkTreeOrder(t, tree)

	tree.Remove(7)
	checkPresenceFor(t, tree, []int{15, 10, 20})
	checkTreeOrder(t, tree)

	tree.Remove(10)
	checkPresenceFor(t, tree, []int{15, 20})
	checkTreeOrder(t, tree)

	tree.Remove(20)
	checkPresenceFor(t, tree, []int{15})
	checkTreeOrder(t, tree)

	tree.Remove(15)
	if tree.root != nil {
		t.Error("Root should be empty")
	}
}

func TestTreeBigRemovalsInLeftSubtree(t *testing.T) {
	values := []int{15, 10, 20, 5, 12, 11, 13, 21}
	tree := prepareTree(values)
	tree.Remove(12)
	checkPresenceFor(t, tree, []int{15, 10, 20, 5, 11, 13, 21})
	checkTreeOrder(t, tree)

	tree = prepareTree(values)
	tree.Remove(10)
	checkPresenceFor(t, tree, []int{15, 20, 5, 12, 11, 13, 21})
	checkTreeOrder(t, tree)

	tree = prepareTree([]int{15, 10, 20, 5, 13, 21, 12, 14, 11})
	tree.Remove(13)
	checkPresenceFor(t, tree, []int{15, 10, 20, 5, 21, 12, 14, 11})
	checkTreeOrder(t, tree)

	tree = prepareTree([]int{15, 9, 20, 5, 12, 11, 13, 10})
	tree.Remove(12)
	checkPresenceFor(t, tree, []int{15, 9, 20, 5, 11, 13, 10})
	checkTreeOrder(t, tree)
}

func TestTreeBigRemovalsInRightSubtree(t *testing.T) {
	tree := prepareTree([]int{15, 9, 20, 17, 25, 16, 19})
	tree.Remove(17)
	checkPresenceFor(t, tree, []int{15, 9, 20, 25, 16, 19})
	checkTreeOrder(t, tree)
	tree.Remove(17)

	tree = prepareTree([]int{15, 9, 20, 17, 25, 16, 19})
	tree.Remove(20)
	checkPresenceFor(t, tree, []int{15, 9, 17, 25, 16, 19})
	checkTreeOrder(t, tree)
	tree.Remove(20)
	tree.Remove(17)

	tree.Remove(17)
	checkPresenceFor(t, tree, []int{15, 9, 25, 16, 19})
	checkTreeOrder(t, tree)
	tree.Remove(15)
	checkPresenceFor(t, tree, []int{9, 25, 16, 19})
	checkTreeOrder(t, tree)
	tree.Remove(15)
	tree.Remove(20)
	tree.Remove(17)

	tree = prepareTree([]int{15, 9, 20, 17, 25, 23, 26})
	tree.Remove(25)
	checkPresenceFor(t, tree, []int{15, 9, 20, 17, 23, 26})
	checkTreeOrder(t, tree)

	tree = prepareTree([]int{15, 9, 20, 17, 25, 23, 26})
	tree.Remove(20)
	checkPresenceFor(t, tree, []int{15, 9, 17, 25, 23, 26})
	checkTreeOrder(t, tree)
}

func TestTreeRootRemoval(t *testing.T) {
	tree := prepareTree([]int{15, 10, 20})
	tree.Remove(15)
	checkPresenceFor(t, tree, []int{10, 20})
	checkTreeOrder(t, tree)

	tree = prepareTree([]int{15, 10, 8, 21, 7, 22})
	tree.Remove(15)
	checkPresenceFor(t, tree, []int{10, 8, 21, 7, 22})
	checkTreeOrder(t, tree)

	tree = prepareTree([]int{15, 10, 20, 5, 7, 6, 8, 21})
	tree.Remove(15)
	checkPresenceFor(t, tree, []int{10, 20, 5, 7, 6, 8, 21})
	checkTreeOrder(t, tree)
}

func TestTreeSmallRotateLeft(t *testing.T) {
	tree := NewTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(20)

	if tree.root.height != 3 {
		t.Error("Invalid tree height (before rotation)", tree.root.height)
	}
	if tree.root.left.height != 1 {
		t.Error("Invalid left-tree height (before rotation)", tree.root.left.height)
	}
	if tree.root.right.height != 2 {
		t.Error("Invalid right-tree height (before rotation)", tree.root.right.height)
	}

	tree.Insert(25)

	if tree.root.height != 3 {
		t.Error("Invalid tree height (after rotation)", tree.root.height)
	}
	if tree.root.left.height != 1 {
		t.Error("Invalid left-tree height (after rotation)", tree.root.left.height)
	}
	if tree.root.right.height != 2 {
		t.Error("Invalid right-tree height (after rotation)", tree.root.right.height)
	}
}

func TestTreeSmallRotateRight(t *testing.T) {
	tree := NewTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(4)

	if tree.root.height != 3 {
		t.Error("Invalid tree height (before rotation)", tree.root.height)
	}
	if tree.root.left.height != 2 {
		t.Error("Invalid left-tree height (before rotation)", tree.root.left.height)
	}
	if tree.root.right.height != 1 {
		t.Error("Invalid right-tree height (before rotation)", tree.root.right.height)
	}

	tree.Insert(3)

	if tree.root.height != 3 {
		t.Error("Invalid tree height (after rotation)", tree.root.height)
	}
	if tree.root.left.height != 2 {
		t.Error("Invalid left-tree height (after rotation)", tree.root.left.height)
	}
	if tree.root.right.height != 1 {
		t.Error("Invalid right-tree height (after rotation)", tree.root.right.height)
	}
}

func TestTreeBigRotateLeft(t *testing.T) {
	tree := NewTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(20)
	tree.Insert(15)
	tree.Insert(25)

	tree.Insert(14)
	if tree.root.height != 3 {
		t.Error("Invalid tree height (after rotation)", tree.root.height)
	}
	if tree.root.left.height != 2 {
		t.Error("Invalid left-tree height (after rotation)", tree.root.left.height)
	}
	if tree.root.right.height != 2 {
		t.Error("Invalid right-tree height (after rotation)", tree.root.right.height)
	}
}

func TestTreeBigRotateRight(t *testing.T) {
	tree := NewTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(20)
	tree.Insert(2)
	tree.Insert(7)

	tree.Insert(6)
	if tree.root.height != 3 {
		t.Error("Invalid tree height (after rotation)", tree.root.height)
	}
	if tree.root.left.height != 2 {
		t.Error("Invalid left-tree height (after rotation)", tree.root.left.height)
	}
	if tree.root.right.height != 2 {
		t.Error("Invalid right-tree height (after rotation)", tree.root.right.height)
	}
}

func prepareTree(values []int) *Tree {
	tree := NewTree()

	for _, val := range values {
		tree.Insert(val)
	}

	return tree
}

func checkTreeOrder(t *testing.T, tree *Tree) {
	values := tree.DumpValues()
	fmt.Println(values)

	for i := 1; i < len(values); i++ {
		if values[i] < values[i-1] {
			t.Error("Invalid Tree order", values)
		}
	}
}

func checkPresenceFor(t *testing.T, tree *Tree, values []int) {
	for _, val := range values {
		if !tree.Search(val) {
			t.Errorf("Value %d should be in the Tree", val)
		}
	}
}

func checkAbsenceFor(t *testing.T, tree *Tree, values []int) {
	for _, val := range values {
		if tree.Search(val) {
			t.Errorf("Value %d shouldn't be in the Tree", val)
		}
	}
}
