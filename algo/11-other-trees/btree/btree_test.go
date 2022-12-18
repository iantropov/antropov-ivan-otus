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

func checkPresence(t *testing.T, tree *btree, val int) {
	if !tree.Search(val) {
		t.Errorf("Tree should contain value %d", val)
	}
}
