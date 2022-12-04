package avl

type Node struct {
	val, height int
	left, right *Node
}

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return new(Tree)
}

func (tree *Tree) Search(val int) bool {
	node := tree.root.search(val)
	return node != nil
}

func (tree *Tree) Insert(val int) {
	tree.root = tree.root.insert(val)
}

func (tree *Tree) Remove(val int) {
	tree.root = tree.root.remove(val)
}

func (tree *Tree) DumpValues() []int {
	values := make([]int, 0)
	return tree.root.dumpValues(values)
}

func (node *Node) search(val int) *Node {
	if node == nil {
		return nil
	}

	if val > node.val {
		return node.right.search(val)
	} else if val < node.val {
		return node.left.search(val)
	} else {
		return node
	}
}

func (node *Node) insert(val int) *Node {
	if node == nil {
		return newNode(val)
	}

	if val > node.val {
		node.right = node.right.insert(val)
	} else if val < node.val {
		node.left = node.left.insert(val)
	}

	return node
}

func newNode(val int) *Node {
	newNode := new(Node)
	newNode.val = val
	return newNode
}

func (node *Node) remove(val int) *Node {
	if node == nil {
		return nil
	}

	if val > node.val {
		node.right = node.right.remove(val)
		return node
	} else if val < node.val {
		node.left = node.left.remove(val)
		return node
	} else if node.left != nil && node.right != nil {
		x := node.right.findMin()
		x.right = node.right.removeMin()
		x.left = node.left
		return x
	} else if node.left != nil {
		return node.left
	} else {
		return node.right
	}
}

func (node *Node) findMin() *Node {
	min := node
	for ; min.left != nil; min = min.left {
	}
	return min
}

func (node *Node) removeMin() *Node {
	if node.left != nil {
		node.left = node.left.removeMin()
		return node
	} else {
		return node.right
	}
}

func (n *Node) dumpValues(values []int) []int {
	if n == nil {
		return values
	}

	values = n.left.dumpValues(values)
	values = append(values, n.val)
	return n.right.dumpValues(values)
}
