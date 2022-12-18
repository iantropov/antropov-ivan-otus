package btree

import "testing"

func TestInsertion1(t *testing.T) {
	tree := NewTree(2)
	for i := 1; i <= 10; i++ {
		tree.Insert(i)
		checkPresence(t, tree, i)
		tree.dump()
	}
}

func TestInsertion2(t *testing.T) {
	tree := NewTree(2)
	for i := 1; i <= 5; i++ {
		tree.Insert(i * 10)
		checkPresence(t, tree, i*10)
		tree.dump()
	}

	for i := 1; i <= 5; i++ {
		tree.Insert(5 + i*10)
		checkPresence(t, tree, 5+i*10)
		tree.dump()
	}
}

func TestRemoval1(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})
	tree.Remove(6)
	checkAbsence(t, tree, 6)
	tree.dump()
}

func TestRemoval2(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})
	tree.Remove(5)
	checkAbsence(t, tree, 5)
	tree.dump()
}

func TestRemoval3(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})
	tree.Remove(3)
	checkAbsence(t, tree, 3)
	tree.dump()
}

func TestRemoval4(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})
	tree.Remove(1)
	checkAbsence(t, tree, 1)
	tree.dump()
}

func TestRemoval5(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})
	tree.Remove(4)
	checkAbsence(t, tree, 4)
	tree.dump()
}

func TestRemoval6(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 4, 5, 6})
	tree.Remove(2)
	checkAbsence(t, tree, 2)
	tree.dump()
}

func buildTree(values []int) *btree {
	tree := NewTree(2)
	for _, val := range values {
		tree.Insert(val)
	}
	return tree
}

func checkPresence(t *testing.T, tree *btree, val int) {
	if !tree.Search(val) {
		t.Errorf("Tree should contain value %d", val)
	}
}

func checkAbsence(t *testing.T, tree *btree, val int) {
	if tree.Search(val) {
		t.Errorf("Tree shouldn't contain value %d", val)
	}
}
