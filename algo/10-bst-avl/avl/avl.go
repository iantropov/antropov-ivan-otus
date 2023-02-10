package avl

type Node struct {
	val, height int
	left, right *Node
}

type Tree struct {
	root *Node
}

// Really Slow implementation - because of inaccurate balancing after insertion / removal
// Now - rebalnce all tree
// Should - rebalance only on the way from affected node towards root (ancestors of affected node)

func NewTree() *Tree {
	return new(Tree)
}

func (tree *Tree) Search(val int) bool {
	node := tree.root.search(val)
	return node != nil
}

func (tree *Tree) Insert(val int) {
	tree.root = tree.root.insert(val)
	tree.root = tree.root.rebalance()
}

func (tree *Tree) Remove(val int) {
	tree.root = tree.root.remove(val)
	tree.root = tree.root.rebalance()
}

func (tree *Tree) DumpValues() []int {
	values := make([]int, 0)
	return tree.root.dumpValues(values)
}

func (tree *Tree) SearchStrictLeft(val int) *Node {
	node := tree.root.searchStrictLeft(val)
	return node
}

func (n *Node) searchStrictLeft(val int) *Node {
	if n == nil {
		return nil
	}

	if val < n.val {
		if n.left != nil && n.left.val > val {
			return n.left.searchStrictLeft(val)
		} else if n.left != nil && n.left.val <= val {
			leftRes := n.left.searchStrictLeft(val)
			if leftRes == nil || leftRes.val < val {
				return n
			} else {
				return leftRes
			}
		} else if n.left == nil {
			return n
		}
	} else if val >= n.val {
		return n.right.searchStrictLeft(val)
	}

	return nil
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
	newNode.height = 1
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

func (node *Node) getHeight() int {
	if node == nil {
		return 0
	} else {
		return node.height
	}
}

func (node *Node) calculateHeight() int {
	if node == nil {
		return 0
	} else {
		leftHeight := node.left.getHeight()
		rightHeight := node.right.getHeight()
		height := leftHeight
		if rightHeight > leftHeight {
			height = rightHeight
		}
		node.height = height + 1
		return node.height
	}
}

func (node *Node) rebalance() *Node {
	if node == nil {
		return nil
	}

	newNode := node
	newNode.left = node.left.rebalance()
	newNode.right = node.right.rebalance()

	leftHeight := node.left.getHeight()
	rightHeight := node.right.getHeight()

	if leftHeight > rightHeight+1 {
		newNode = node.rotateRight()
	} else if rightHeight > leftHeight+1 {
		newNode = node.rotateLeft()
	}

	newNode.calculateHeight()

	return newNode
}

func (node *Node) rotateLeft() *Node {
	if node.right.left.getHeight() > node.right.right.getHeight() {
		return node.bigRotateLeft()
	} else {
		return node.smallRotateLeft()
	}
}

func (node *Node) rotateRight() *Node {
	if node.left.right.getHeight() > node.left.left.getHeight() {
		return node.bigRotateRight()
	} else {
		return node.smallRotateRight()
	}
}

func (node *Node) smallRotateLeft() *Node {
	newNode := node.right
	node.right = newNode.left
	newNode.left = node
	newNode.left.calculateHeight()
	newNode.right.calculateHeight()
	return newNode
}

func (node *Node) smallRotateRight() *Node {
	newNode := node.left
	node.left = newNode.right
	newNode.right = node
	newNode.left.calculateHeight()
	newNode.right.calculateHeight()
	return newNode
}

func (node *Node) bigRotateLeft() *Node {
	node.right = node.right.smallRotateRight()
	return node.smallRotateLeft()
}

func (node *Node) bigRotateRight() *Node {
	node.left = node.left.smallRotateLeft()
	return node.smallRotateRight()
}
