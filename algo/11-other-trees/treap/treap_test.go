package treap

import "testing"

func TestTreapInsertion(t *testing.T) {
	tree := NewTree()
	for i := 1; i <= 10; i++ {
		tree.Insert(i)
		checkPresence(t, tree, i)
		checkInvariants(t, tree)
	}
}

func TestTreapRemove(t *testing.T) {
	tree := NewTree()
	for i := 1; i <= 10; i++ {
		tree.Insert(i)
	}

	tree.Remove(3)
	checkAbsence(t, tree, 3)
	checkInvariants(t, tree)

	tree.Remove(7)
	checkAbsence(t, tree, 7)
	checkInvariants(t, tree)

	tree.Remove(9)
	checkAbsence(t, tree, 9)
	checkInvariants(t, tree)
}

func TestTreapRemove2(t *testing.T) {
	tree := NewTree()
	tree.insertDirect(26, 0.9)
	tree.insertDirect(24, 0.8)
	tree.insertDirect(29, 0.8)
	tree.insertDirect(22, 0.7)
	tree.insertDirect(25, 0.7)
	tree.insertDirect(20, 0.6)
	tree.insertDirect(23, 0.6)
	tree.insertDirect(1, 0.5)
	tree.insertDirect(21, 0.5)
	tree.insertDirect(10, 0.4)
	tree.insertDirect(5, 0.3)
	tree.insertDirect(14, 0.3)

	checkInvariants(t, tree)

	tree.Remove(21)
	checkAbsence(t, tree, 21)
	checkInvariants(t, tree)
}

func checkPresence(t *testing.T, tree *treap, val int) {
	if !tree.Search(val) {
		t.Errorf("Value %d should be presented in the tree", val)
	}
}

func checkAbsence(t *testing.T, tree *treap, val int) {
	if tree.Search(val) {
		t.Errorf("Value %d should be presented in the tree", val)
	}
}

func checkInvariants(t *testing.T, tree *treap) {
	if !tree.CheckForInvariants() {
		tree.DumpValuesInDetails()
		t.Error("Tree is invalid!")
	}
}
