package splay

import (
	"testing"
)

func TestSplayZigLeftOnSmallTree(t *testing.T) {
	splayTree := NewTree()

	splayTree.InsertWithoutSplay(10)
	splayTree.InsertWithoutSplay(15)
	splayTree.InsertWithoutSplay(5)

	if !splayTree.Search(5) {
		t.Errorf("Element %d should be in the tree", 5)
	}

	if splayTree.root.value != 5 {
		t.Errorf("Element %d should be in the root", 5)
	}
}

func TestSplayZigRightOnSmallTree(t *testing.T) {
	splayTree := NewTree()

	splayTree.InsertWithoutSplay(10)
	splayTree.InsertWithoutSplay(15)
	splayTree.InsertWithoutSplay(5)

	if !splayTree.Search(15) {
		t.Errorf("Element %d should be in the tree", 15)
	}

	if splayTree.root.value != 15 {
		t.Errorf("Element %d should be in the root", 15)
	}
}

func TestSplayZigLeftOnMediumTree(t *testing.T) {
	splayTree := NewTree()

	splayTree.InsertWithoutSplay(20)
	splayTree.InsertWithoutSplay(15)
	splayTree.InsertWithoutSplay(25)
	splayTree.InsertWithoutSplay(10)
	splayTree.InsertWithoutSplay(17)
	splayTree.InsertWithoutSplay(23)
	splayTree.InsertWithoutSplay(30)

	if !splayTree.Search(15) {
		t.Errorf("Element %d should be in the tree", 15)
	}

	if splayTree.root.value != 15 {
		t.Errorf("Element %d should be in the root", 15)
	}
}

func TestSplayDoubleSearchOnMediumTree(t *testing.T) {
	splayTree := NewTree()

	splayTree.InsertWithoutSplay(20)
	splayTree.InsertWithoutSplay(15)
	splayTree.InsertWithoutSplay(25)
	splayTree.InsertWithoutSplay(10)
	splayTree.InsertWithoutSplay(17)
	splayTree.InsertWithoutSplay(23)
	splayTree.InsertWithoutSplay(30)

	if !splayTree.Search(15) {
		t.Errorf("Element %d should be in the tree", 15)
	}

	if splayTree.root.value != 15 {
		t.Errorf("Element %d should be in the root", 15)
	}

	if !splayTree.Search(23) {
		t.Errorf("Element %d should be in the tree", 23)
	}

	if splayTree.root.value != 23 {
		t.Errorf("Element %d should be in the root", 23)
	}
}

func TestSplayZigRightOnMediumTree(t *testing.T) {
	splayTree := NewTree()

	splayTree.InsertWithoutSplay(20)
	splayTree.InsertWithoutSplay(15)
	splayTree.InsertWithoutSplay(25)
	splayTree.InsertWithoutSplay(10)
	splayTree.InsertWithoutSplay(17)
	splayTree.InsertWithoutSplay(23)
	splayTree.InsertWithoutSplay(30)

	if !splayTree.Search(25) {
		t.Errorf("Element %d should be in the tree", 25)
	}

	if splayTree.root.value != 25 {
		t.Errorf("Element %d should be in the root", 25)
	}
}

func TestSplayZigZigLeftOnMediumTree(t *testing.T) {
	splayTree := NewTree()

	splayTree.InsertWithoutSplay(20)
	splayTree.InsertWithoutSplay(15)
	splayTree.InsertWithoutSplay(25)
	splayTree.InsertWithoutSplay(10)
	splayTree.InsertWithoutSplay(17)
	splayTree.InsertWithoutSplay(23)
	splayTree.InsertWithoutSplay(30)

	if !splayTree.Search(10) {
		t.Errorf("Element %d should be in the tree", 10)
	}

	if splayTree.root.value != 10 {
		t.Errorf("Element %d should be in the root", 10)
	}
}

func TestSplayZigZigRightOnMediumTree(t *testing.T) {
	splayTree := NewTree()

	splayTree.InsertWithoutSplay(20)
	splayTree.InsertWithoutSplay(15)
	splayTree.InsertWithoutSplay(25)
	splayTree.InsertWithoutSplay(10)
	splayTree.InsertWithoutSplay(17)
	splayTree.InsertWithoutSplay(23)
	splayTree.InsertWithoutSplay(30)

	if !splayTree.Search(30) {
		t.Errorf("Element %d should be in the tree", 30)
	}

	if splayTree.root.value != 30 {
		t.Errorf("Element %d should be in the root", 30)
	}
}

func TestSplayZigZagLeftOnMediumTree(t *testing.T) {
	splayTree := NewTree()

	splayTree.InsertWithoutSplay(20)
	splayTree.InsertWithoutSplay(15)
	splayTree.InsertWithoutSplay(25)
	splayTree.InsertWithoutSplay(10)
	splayTree.InsertWithoutSplay(17)
	splayTree.InsertWithoutSplay(23)
	splayTree.InsertWithoutSplay(30)

	if !splayTree.Search(17) {
		t.Errorf("Element %d should be in the tree", 17)
	}

	if splayTree.root.value != 17 {
		t.Errorf("Element %d should be in the root", 17)
	}
}

func TestSplayZigZagRightOnMediumTree(t *testing.T) {
	splayTree := NewTree()

	splayTree.InsertWithoutSplay(20)
	splayTree.InsertWithoutSplay(15)
	splayTree.InsertWithoutSplay(25)
	splayTree.InsertWithoutSplay(10)
	splayTree.InsertWithoutSplay(17)
	splayTree.InsertWithoutSplay(23)
	splayTree.InsertWithoutSplay(30)

	if !splayTree.Search(23) {
		t.Errorf("Element %d should be in the tree", 23)
	}

	if splayTree.root.value != 23 {
		t.Errorf("Element %d should be in the root", 23)
	}
}

func TestSplayZigZigLeftOLargeTree(t *testing.T) {
	splayTree := NewTree()

	splayTree.InsertWithoutSplay(30)
	splayTree.InsertWithoutSplay(15)
	splayTree.InsertWithoutSplay(45)
	splayTree.InsertWithoutSplay(10)
	splayTree.InsertWithoutSplay(20)
	splayTree.InsertWithoutSplay(40)
	splayTree.InsertWithoutSplay(50)
	splayTree.InsertWithoutSplay(5)
	splayTree.InsertWithoutSplay(12)
	splayTree.InsertWithoutSplay(17)
	splayTree.InsertWithoutSplay(25)
	splayTree.InsertWithoutSplay(35)
	splayTree.InsertWithoutSplay(42)
	splayTree.InsertWithoutSplay(47)
	splayTree.InsertWithoutSplay(55)

	if !splayTree.Search(5) {
		t.Errorf("Element %d should be in the tree", 5)
	}

	if splayTree.root.value != 5 {
		t.Errorf("Element %d should be in the root", 5)
	}
}

func TestSplayDoubleSearchOLargeTree(t *testing.T) {
	splayTree := NewTree()

	splayTree.InsertWithoutSplay(30)
	splayTree.InsertWithoutSplay(15)
	splayTree.InsertWithoutSplay(45)
	splayTree.InsertWithoutSplay(10)
	splayTree.InsertWithoutSplay(20)
	splayTree.InsertWithoutSplay(40)
	splayTree.InsertWithoutSplay(50)
	splayTree.InsertWithoutSplay(5)
	splayTree.InsertWithoutSplay(12)
	splayTree.InsertWithoutSplay(17)
	splayTree.InsertWithoutSplay(25)
	splayTree.InsertWithoutSplay(35)
	splayTree.InsertWithoutSplay(42)
	splayTree.InsertWithoutSplay(47)
	splayTree.InsertWithoutSplay(55)

	if !splayTree.Search(5) {
		t.Errorf("Element %d should be in the tree", 5)
	}

	if splayTree.root.value != 5 {
		t.Errorf("Element %d should be in the root", 5)
	}

	if !splayTree.Search(20) {
		t.Errorf("Element %d should be in the tree", 20)
	}

	if splayTree.root.value != 20 {
		t.Errorf("Element %d should be in the root", 20)
	}
}

func TestRemove(t *testing.T) {
	splayTree := NewTree()

	splayTree.InsertWithoutSplay(20)
	splayTree.InsertWithoutSplay(15)
	splayTree.InsertWithoutSplay(25)
	splayTree.InsertWithoutSplay(10)
	splayTree.InsertWithoutSplay(17)
	splayTree.InsertWithoutSplay(23)
	splayTree.InsertWithoutSplay(30)

	splayTree.Remove(15)

	if splayTree.Search(15) {
		t.Errorf("Element %d shouldn't be in the tree", 15)
	}
}

func TestInsert(t *testing.T) {
	splayTree := NewTree()

	values := []int{20, 15, 25, 10, 17, 23}
	for _, value := range values {
		splayTree.Insert(value)
		if splayTree.root.value != value {
			t.Errorf("Element %d should be in the root", value)
		}
	}
}
