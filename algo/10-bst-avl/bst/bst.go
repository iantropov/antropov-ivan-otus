package bst

type Node struct {
	val                 int
	parent, left, right *Node
}

type BST struct {
	root *Node
}

func CreateBST() *BST {
	bst := new(BST)
	return bst
}

func (t *BST) Insert(val int) {
	if t.root == nil {
		t.root = new(Node)
		t.root.val = val
	} else {
		t.root.insert(val)
	}
}

func (t *BST) Remove(val int) {
	node := t.root.search(val)
	if node == t.root {
		t.root = nil
	} else {
		node.remove()
	}
}

func (t *BST) Search(val int) bool {
	node := t.root.search(val)
	return node != nil
}

func (t *BST) DumpValues() []int {
	values := make([]int, 0)
	return t.root.dumpValues(values)
}

func (n *Node) insert(val int) {
	if val < n.val {
		if n.left != nil {
			n.left.insert(val)
		} else {
			n.left = new(Node)
			n.left.val = val
			n.left.parent = n
		}
	} else if val > n.val {
		if n.right != nil {
			n.right.insert(val)
		} else {
			n.right = new(Node)
			n.right.val = val
			n.right.parent = n
		}
	}
}

func (n *Node) search(val int) *Node {
	if n == nil {
		return nil
	}

	if val < n.val {
		return n.left.search(val)
	} else if val > n.val {
		return n.right.search(val)
	} else {
		return n
	}
}

func (n *Node) remove() {
	if n == nil {
		return
	}

	if n.left == nil && n.right == nil {
		n.parent.replaceChild(n, nil)
	} else if n.left != nil {
		n.parent.replaceChild(n, n.left)
	} else if n.right != nil {
		n.parent.replaceChild(n, n.right)
	} else {
		maxLeftChild := n.findMaxLeftChild()
		maxLeftChildParent := maxLeftChild.parent
		n.parent.replaceChild(n, maxLeftChild)
		maxLeftChildParent.replaceChild(maxLeftChild, n)
		n.remove()
	}
}

func (n *Node) replaceChild(currentChild, newChild *Node) {
	if n.left == currentChild {
		n.left = newChild
	} else {
		n.right = newChild
	}

	if newChild != nil {
		newChild.parent = n
	}
}

func (n *Node) findMaxLeftChild() *Node {
	maxLeftChild := n.left
	for ; maxLeftChild.right != nil; maxLeftChild = maxLeftChild.right {
	}
	return maxLeftChild
}

func (n *Node) dumpValues(values []int) []int {
	if n == nil {
		return values
	}

	values = n.left.dumpValues(values)
	values = append(values, n.val)
	return n.right.dumpValues(values)
}
