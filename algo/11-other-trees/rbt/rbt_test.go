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

func assertNodesAreRed(t *testing.T, tree *rbtTree, values []int) {
	for _, val := range values {
		if !tree.searchNode(val).red {
			t.Errorf("The node with %d should be red", val)
		}
	}
}
