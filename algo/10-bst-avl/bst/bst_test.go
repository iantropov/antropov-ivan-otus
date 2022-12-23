package bst

import (
	"fmt"
	"testing"
)

func TestBstInsert(t *testing.T) {
	bst := prepareBST([]int{15, 10, 20, 5, 7, 6, 8})
	checkBSTOrder(t, bst)
}

func TestBstSearch(t *testing.T) {
	values := []int{15, 10, 20, 5, 7, 6, 8}
	bst := prepareBST(values)

	checkPresenceFor(t, bst, values)

	checkAbsenceFor(t, bst, []int{17, 27, 0})
}

func TestBstInvalidRemoval(t *testing.T) {
	values := []int{15, 10, 20, 5, 7, 6, 8, 21}
	bst := prepareBST(values)
	bst.Remove(70)
	checkPresenceFor(t, bst, values)
}

func TestBstSmallRemovals(t *testing.T) {
	bst := prepareBST([]int{15, 10, 20, 5, 7, 6, 8, 21})
	bst.Remove(21)
	checkPresenceFor(t, bst, []int{15, 10, 20, 5, 7, 6, 8})
	checkBSTOrder(t, bst)

	bst.Remove(5)
	checkPresenceFor(t, bst, []int{15, 10, 20, 7, 6, 8})
	checkBSTOrder(t, bst)

	bst.Remove(6)
	checkPresenceFor(t, bst, []int{15, 10, 20, 7, 8})
	checkBSTOrder(t, bst)

	bst.Remove(8)
	checkPresenceFor(t, bst, []int{15, 10, 20, 7})
	checkBSTOrder(t, bst)

	bst.Remove(7)
	checkPresenceFor(t, bst, []int{15, 10, 20})
	checkBSTOrder(t, bst)

	bst.Remove(10)
	checkPresenceFor(t, bst, []int{15, 20})
	checkBSTOrder(t, bst)

	bst.Remove(20)
	checkPresenceFor(t, bst, []int{15})
	checkBSTOrder(t, bst)

	bst.Remove(15)
	if bst.root != nil {
		t.Error("Root should be empty")
	}
}

func TestBstBigRemovalsInLeftSubtree(t *testing.T) {
	values := []int{15, 10, 20, 5, 12, 11, 13, 21}
	bst := prepareBST(values)
	bst.Remove(12)
	checkPresenceFor(t, bst, []int{15, 10, 20, 5, 11, 13, 21})
	checkBSTOrder(t, bst)

	bst = prepareBST(values)
	bst.Remove(10)
	checkPresenceFor(t, bst, []int{15, 20, 5, 12, 11, 13, 21})
	checkBSTOrder(t, bst)

	bst = prepareBST([]int{15, 10, 20, 5, 13, 21, 12, 14, 11})
	bst.Remove(13)
	checkPresenceFor(t, bst, []int{15, 10, 20, 5, 21, 12, 14, 11})
	checkBSTOrder(t, bst)

	bst = prepareBST([]int{15, 9, 20, 5, 12, 11, 13, 10})
	bst.Remove(12)
	checkPresenceFor(t, bst, []int{15, 9, 20, 5, 11, 13, 10})
	checkBSTOrder(t, bst)
}

func TestBstBigRemovalsInRightSubtree(t *testing.T) {
	bst := prepareBST([]int{15, 9, 20, 17, 25, 16, 19})
	bst.Remove(17)
	checkPresenceFor(t, bst, []int{15, 9, 20, 25, 16, 19})
	checkBSTOrder(t, bst)
	bst.Remove(17)

	bst = prepareBST([]int{15, 9, 20, 17, 25, 16, 19})
	bst.Remove(20)
	checkPresenceFor(t, bst, []int{15, 9, 17, 25, 16, 19})
	checkBSTOrder(t, bst)
	bst.Remove(20)
	bst.Remove(17)

	bst.Remove(17)
	checkPresenceFor(t, bst, []int{15, 9, 25, 16, 19})
	checkBSTOrder(t, bst)
	bst.Remove(15)
	checkPresenceFor(t, bst, []int{9, 25, 16, 19})
	checkBSTOrder(t, bst)
	bst.Remove(15)
	bst.Remove(20)
	bst.Remove(17)

	bst = prepareBST([]int{15, 9, 20, 17, 25, 23, 26})
	bst.Remove(25)
	checkPresenceFor(t, bst, []int{15, 9, 20, 17, 23, 26})
	checkBSTOrder(t, bst)

	bst = prepareBST([]int{15, 9, 20, 17, 25, 23, 26})
	bst.Remove(20)
	checkPresenceFor(t, bst, []int{15, 9, 17, 25, 23, 26})
	checkBSTOrder(t, bst)
}

func TestBstRootRemoval(t *testing.T) {
	bst := prepareBST([]int{15, 10, 20})
	bst.Remove(15)
	checkPresenceFor(t, bst, []int{10, 20})
	checkBSTOrder(t, bst)

	bst = prepareBST([]int{15, 10, 8, 21, 7, 22})
	bst.Remove(15)
	checkPresenceFor(t, bst, []int{10, 8, 21, 7, 22})
	checkBSTOrder(t, bst)

	bst = prepareBST([]int{15, 10, 20, 5, 7, 6, 8, 21})
	bst.Remove(15)
	checkPresenceFor(t, bst, []int{10, 20, 5, 7, 6, 8, 21})
	checkBSTOrder(t, bst)
}

func prepareBST(values []int) *BST {
	bst := CreateBST()

	for _, val := range values {
		bst.Insert(val)
	}

	return bst
}

func checkBSTOrder(t *testing.T, bst *BST) {
	values := bst.DumpValues()
	fmt.Println(values)

	for i := 1; i < len(values); i++ {
		if values[i] < values[i-1] {
			t.Error("Invalid BST order", values)
		}
	}
}

func checkPresenceFor(t *testing.T, bst *BST, values []int) {
	for _, val := range values {
		if !bst.Search(val) {
			t.Errorf("Value %d should be in the BST", val)
		}
	}
}

func checkAbsenceFor(t *testing.T, bst *BST, values []int) {
	for _, val := range values {
		if bst.Search(val) {
			t.Errorf("Value %d shouldn't be in the BST", val)
		}
	}
}