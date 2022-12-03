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
}

func TestBstBigRemovals(t *testing.T) {
	values := []int{15, 10, 20, 5, 7, 6, 8, 21}
	bst := prepareBST(values)
	bst.Remove(7)
	checkPresenceFor(t, bst, []int{10, 20, 5, 6, 8, 21})
	checkBSTOrder(t, bst)

	bst = prepareBST(values)
	bst.Remove(10)
	checkPresenceFor(t, bst, []int{15, 20, 5, 7, 6, 8})
	checkBSTOrder(t, bst)
}

func TestBstRootRemoval(t *testing.T) {
	bst := prepareBST([]int{15, 10, 20, 5, 7, 6, 8, 21})
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
