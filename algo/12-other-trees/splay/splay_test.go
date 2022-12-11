package splay

import (
	"testing"
)

func TestSplayZigLeftOnSmallTree(t *testing.T) {
	splayTree := NewTree()

	splayTree.Insert(10)
	splayTree.Insert(15)
	splayTree.Insert(5)

	if !splayTree.Search(5) {
		t.Errorf("Element %d should be in the tree", 5)
	}

	if splayTree.root.value != 5 {
		t.Errorf("Element %d should be in the root", 5)
	}
}

func TestSplayZigRightOnSmallTree(t *testing.T) {
	splayTree := NewTree()

	splayTree.Insert(10)
	splayTree.Insert(15)
	splayTree.Insert(5)

	if !splayTree.Search(15) {
		t.Errorf("Element %d should be in the tree", 15)
	}

	if splayTree.root.value != 15 {
		t.Errorf("Element %d should be in the root", 15)
	}
}

func TestSplayZigLeftOnMediumTree(t *testing.T) {
	splayTree := NewTree()

	splayTree.Insert(20)
	splayTree.Insert(15)
	splayTree.Insert(25)
	splayTree.Insert(10)
	splayTree.Insert(17)
	splayTree.Insert(23)
	splayTree.Insert(30)

	if !splayTree.Search(15) {
		t.Errorf("Element %d should be in the tree", 15)
	}

	if splayTree.root.value != 15 {
		t.Errorf("Element %d should be in the root", 15)
	}
}

func TestSplayDoubleSearchOnMediumTree(t *testing.T) {
	splayTree := NewTree()

	splayTree.Insert(20)
	splayTree.Insert(15)
	splayTree.Insert(25)
	splayTree.Insert(10)
	splayTree.Insert(17)
	splayTree.Insert(23)
	splayTree.Insert(30)

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

	splayTree.Insert(20)
	splayTree.Insert(15)
	splayTree.Insert(25)
	splayTree.Insert(10)
	splayTree.Insert(17)
	splayTree.Insert(23)
	splayTree.Insert(30)

	if !splayTree.Search(25) {
		t.Errorf("Element %d should be in the tree", 25)
	}

	if splayTree.root.value != 25 {
		t.Errorf("Element %d should be in the root", 25)
	}
}

func TestSplayZigZigLeftOnMediumTree(t *testing.T) {
	splayTree := NewTree()

	splayTree.Insert(20)
	splayTree.Insert(15)
	splayTree.Insert(25)
	splayTree.Insert(10)
	splayTree.Insert(17)
	splayTree.Insert(23)
	splayTree.Insert(30)

	if !splayTree.Search(10) {
		t.Errorf("Element %d should be in the tree", 10)
	}

	if splayTree.root.value != 10 {
		t.Errorf("Element %d should be in the root", 10)
	}
}

func TestSplayZigZigRightOnMediumTree(t *testing.T) {
	splayTree := NewTree()

	splayTree.Insert(20)
	splayTree.Insert(15)
	splayTree.Insert(25)
	splayTree.Insert(10)
	splayTree.Insert(17)
	splayTree.Insert(23)
	splayTree.Insert(30)

	if !splayTree.Search(30) {
		t.Errorf("Element %d should be in the tree", 30)
	}

	if splayTree.root.value != 30 {
		t.Errorf("Element %d should be in the root", 30)
	}
}

func TestSplayZigZagLeftOnMediumTree(t *testing.T) {
	splayTree := NewTree()

	splayTree.Insert(20)
	splayTree.Insert(15)
	splayTree.Insert(25)
	splayTree.Insert(10)
	splayTree.Insert(17)
	splayTree.Insert(23)
	splayTree.Insert(30)

	if !splayTree.Search(17) {
		t.Errorf("Element %d should be in the tree", 17)
	}

	if splayTree.root.value != 17 {
		t.Errorf("Element %d should be in the root", 17)
	}
}

func TestSplayZigZagRightOnMediumTree(t *testing.T) {
	splayTree := NewTree()

	splayTree.Insert(20)
	splayTree.Insert(15)
	splayTree.Insert(25)
	splayTree.Insert(10)
	splayTree.Insert(17)
	splayTree.Insert(23)
	splayTree.Insert(30)

	if !splayTree.Search(23) {
		t.Errorf("Element %d should be in the tree", 23)
	}

	if splayTree.root.value != 23 {
		t.Errorf("Element %d should be in the root", 23)
	}
}

func TestSplayZigZigLeftOLargeTree(t *testing.T) {
	splayTree := NewTree()

	splayTree.Insert(30)
	splayTree.Insert(15)
	splayTree.Insert(45)
	splayTree.Insert(10)
	splayTree.Insert(20)
	splayTree.Insert(40)
	splayTree.Insert(50)
	splayTree.Insert(5)
	splayTree.Insert(12)
	splayTree.Insert(17)
	splayTree.Insert(25)
	splayTree.Insert(35)
	splayTree.Insert(42)
	splayTree.Insert(47)
	splayTree.Insert(55)

	if !splayTree.Search(5) {
		t.Errorf("Element %d should be in the tree", 5)
	}

	if splayTree.root.value != 5 {
		t.Errorf("Element %d should be in the root", 5)
	}
}

func TestSplayDoubleSearchOLargeTree(t *testing.T) {
	splayTree := NewTree()

	splayTree.Insert(30)
	splayTree.Insert(15)
	splayTree.Insert(45)
	splayTree.Insert(10)
	splayTree.Insert(20)
	splayTree.Insert(40)
	splayTree.Insert(50)
	splayTree.Insert(5)
	splayTree.Insert(12)
	splayTree.Insert(17)
	splayTree.Insert(25)
	splayTree.Insert(35)
	splayTree.Insert(42)
	splayTree.Insert(47)
	splayTree.Insert(55)

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
