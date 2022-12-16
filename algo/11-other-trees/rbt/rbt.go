package rbt

import "fmt"

// https://habr.com/ru/company/otus/blog/472040/
// https://habr.com/ru/company/otus/blog/521034/

type node struct {
	val                 int
	left, right, parent *node
	red                 bool
}

type rbtTree struct {
	root *node
}

func NewTree() *rbtTree {
	return &rbtTree{}
}

func (tree *rbtTree) Search(val int) bool {
	return tree.searchNode(val) != nil
}

func (tree *rbtTree) Insert(val int) {
	if tree.root == nil {
		tree.root = buildNode(val, nil)
		tree.root.red = false
		return
	}

	newNode := tree.root.insert(val, tree.root)
	rotatedNode := newNode.balanceAfterInsertion()
	// tree.root = newNode.findRoot()

	if rotatedNode != nil && rotatedNode.parent == nil {
		tree.root = rotatedNode
	}
}

func (tree *rbtTree) Remove(val int) {
	node := tree.root.searchNode(val)
	if node == nil {
		return
	}

	nodeForBalance := node.remove()

	if tree.root == node && node.parent == nil {
		tree.root = node.left
		if tree.root == nil {
			tree.root = node.right
		}
		return
	}

	tree.root = node.findRoot()
	if nodeForBalance == nil || nodeForBalance.parent == nil {
		return
	}

	nodeForBalance.parent.balanceAfterRemoval()
	tree.root = node.findRoot()
}

func (tree *rbtTree) DumpValuesInDetails() {
	tree.root.dumpValuesInDetails()
}

func (tree *rbtTree) searchNode(val int) *node {
	return tree.root.searchNode(val)
}

func (tree *rbtTree) insertDirectly(val int, red bool) {
	if tree.root == nil {
		tree.root = buildNode(val, nil)
		tree.root.red = red
		return
	}

	newNode := tree.root.insert(val, tree.root)
	newNode.red = red
}

func (tree *rbtTree) checkForInvariants() bool {
	if tree.root == nil {
		return true
	}

	if tree.root.red {
		fmt.Printf("Root should be black: %v\n", tree.root)
		return false
	}

	return tree.root.checkForInvariants()
}

func (n *node) checkForInvariants() bool {
	if n == nil {
		return true
	}

	if n.red && (n.left.isRed() || n.right.isRed()) {
		fmt.Printf("Red node can't have red children: %v\n", n)
		return false
	}

	leftBlackHeight := n.left.findBlackHeight()
	rightBlackHight := n.right.findBlackHeight()
	if leftBlackHeight != rightBlackHight {
		fmt.Printf("Left and right black heights should be equal: %v\n", n)
		return false
	}

	return true
}

func (n *node) findBlackHeight() int {
	if n == nil {
		return 0
	}

	leftBlackHeight := n.left.findBlackHeight()
	if n.red {
		return leftBlackHeight
	} else {
		return leftBlackHeight + 1
	}
}

func (n *node) searchNode(val int) *node {
	if n == nil {
		return nil
	}

	if val < n.val {
		return n.left.searchNode(val)
	} else if val > n.val {
		return n.right.searchNode(val)
	} else {
		return n
	}
}

func (n *node) insert(val int, parent *node) *node {
	if val < n.val {
		if n.left == nil {
			n.left = buildNode(val, n)
			return n.left
		} else {
			return n.left.insert(val, n)
		}
	} else if val > n.val {
		if n.right == nil {
			n.right = buildNode(val, n)
			return n.right
		} else {
			return n.right.insert(val, n)
		}
	} else {
		return nil
	}
}

func buildNode(val int, parent *node) *node {
	return &node{
		val:    val,
		red:    true,
		parent: parent,
	}
}

func (n *node) balanceAfterInsertion() *node {
	if n == nil {
		return nil
	}

	if n.parent == nil {
		if n.red {
			n.red = false
		}
		return nil
	}

	if n.parent.parent == nil {
		return nil
	}

	p := n.parent
	g := p.parent
	u := p.getSibling()

	// case 0 - parent is black
	if !p.red {
		return nil
	}

	// case 1 - uncle is red
	if !g.red && p.red && u.isRed() {
		p.red = false
		u.red = false
		g.red = true
		return g.balanceAfterInsertion()
	}

	// case 2 - uncle is black & (grandpa & parent are on different sides)
	if !g.red && p.red && !u.isRed() && (g.left == p && p.right == n || g.right == p && p.left == n) {
		if g.left == p && p.right == n {
			p.rotateLeft()
		} else {
			p.rotateRight()
		}
		n, p = p, n
	}

	//case 3 - uncle is black & (grandpa & parent are on the same side)
	if !g.red && p.red && !u.isRed() && (g.left == p && p.left == n || g.right == p && p.right == n) {
		if g.left == p && p.left == n {
			g.rotateRight()
		} else {
			g.rotateLeft()
		}
		p.red = false
		g.red = true

		return p
	}

	return nil
}

func (n *node) rotateRight() {
	oldParent := n.parent
	oldLeftRight := n.left.right

	newNode := n.left
	newNode.right = n
	newNode.parent = n.parent

	n.parent = newNode
	n.left = oldLeftRight
	if oldLeftRight != nil {
		oldLeftRight.parent = n
	}

	oldParent.replaceChild(n, newNode)
}

func (n *node) rotateLeft() {
	oldParent := n.parent
	oldRightLeft := n.right.left

	newNode := n.right
	newNode.left = n
	newNode.parent = n.parent

	n.parent = newNode
	n.right = oldRightLeft
	if oldRightLeft != nil {
		oldRightLeft.parent = n
	}

	oldParent.replaceChild(n, newNode)
}

func (n *node) remove() *node {
	// case 0 - there are no children
	if n.left == nil && n.right == nil {
		n.parent.replaceChild(n, nil)
		if n.red {
			return nil
		} else {
			return n
		}
	}

	// case 1 - there is the only child
	var onlyChild *node
	if n.left == nil && n.right != nil {
		onlyChild = n.right
	} else if n.left != nil && n.right == nil {
		onlyChild = n.left
	}
	if onlyChild != nil {
		n.parent.replaceChild(n, onlyChild)
		onlyChild.parent = n.parent
		onlyChild.red = false
		return nil
	}

	// case 2 - there are two children
	nextNode := n.findNextNodeFromRight()
	n.red, nextNode.red = nextNode.red, n.red
	n.switchWithRight(nextNode)
	return n.remove()
}

func (n *node) findNextNodeFromRight() *node {
	minNode := n.right
	for ; minNode.left != nil; minNode = minNode.left {
	}
	return minNode
}

func (n *node) switchWithRight(other *node) {
	if n.right == other {
		n.parent.replaceChild(n, other)
		other.parent = n.parent
		other.left, n.left = n.left, other.left
		other.right, n.right = n, other.right
		if other.left != nil {
			other.left.parent = other
		}
		n.parent = other
	} else {
		n.parent.replaceChild(n, other)
		other.parent.replaceChild(other, n)

		n.parent, other.parent = other.parent, n.parent
		n.replaceParentForChildren(other)
		other.replaceParentForChildren(n)
		n.left, other.left = other.left, n.left
		n.right, other.right = other.right, n.right
	}
}

func (n *node) balanceAfterRemoval() *node {
	if n == nil {
		return nil
	}

	if n.left == nil {
		return n.balanceAfterLeftRemoval()
	} else {
		return n.balanceAfterRightRemoval()
	}
}

func (n *node) balanceAfterRightRemoval() *node {
	// case 1 - a red parent with a black left child with black grandchildren
	if n.red && !n.left.red && !n.left.left.isRed() && !n.left.right.isRed() {
		n.red = false
		n.left.red = true
		return nil
	}

	// case 2 - a red parent with a black left child with a left red grandchild
	if n.red && !n.left.red && n.left.left.isRed() {
		n.rotateRight()
		n.red = false
		n.parent.red = true
		n.parent.left.red = false
		return n.parent
	}

	// case 2.5 - a red parent with a black left child with a right red grandchild
	if n.red && !n.left.red && n.left.right.isRed() {
		n.left.right.red = false
		n.left.red = true
		n.left.rotateLeft()
		n.rotateRight()
		return n.parent
	}

	// case 3 - a black parent with a red left child with black grandchildren with right black grand-grandchildren
	if !n.red && n.left.isRed() && !n.left.left.isRed() && !n.left.right.isRed() && !n.left.right.left.isRed() && !n.left.right.right.isRed() {
		n.rotateRight()
		n.parent.red = false
		n.left.red = true
		return n.parent
	}

	// case 4 - a black parent with a red left child with black grandchildren with a left red grand-grandchild
	if !n.red && n.left.isRed() && !n.left.left.isRed() && n.left.right != nil && n.left.right.left.isRed() {
		newBlackNode := n.left.right.left
		n.left.rotateLeft()
		newBlackNode.red = false
		n.rotateRight()
		return n.parent
	}

	// case 5 - a black parent with a black left child with a red right grandchild
	if !n.red && !n.left.red && n.left.right.isRed() {
		n.left.rotateLeft()
		n.left.red = false
		n.rotateRight()
		return n.parent
	}

	// case 5.5 - a black parent with a black left child with a red left grandchild
	if !n.red && !n.left.isRed() && n.left.left.isRed() {
		n.rotateRight()
		if n.right != nil {
			n.right.red = false
		}
		n.parent.left.red = false
		return n.parent
	}

	// case 6
	if !n.red && !n.left.red && !n.left.left.isRed() && !n.left.right.isRed() {
		n.left.red = true

		if n.parent == nil {
			return nil
		}

		if n.parent.left == n {
			return n.parent.balanceAfterLeftRemoval()
		} else {
			return n.parent.balanceAfterRightRemoval()
		}
	}

	// root := n.findRoot()
	// root.dumpValuesInDetails()
	// panic(n.val)
	panic("invalid case")
}

func (n *node) balanceAfterLeftRemoval() *node {
	// case 1 - a red parent with black children with black grandchildren
	if n.red && !n.right.red && !n.right.left.isRed() && !n.right.right.isRed() {
		n.red = false
		n.right.red = true
		return nil
	}

	// case 2 - a red parent with a black right child with a right red grandchild
	if n.red && !n.right.red && n.right.right.isRed() {
		n.rotateLeft()
		n.red = false
		n.parent.red = true
		n.parent.right.red = false
		return n.parent
	}

	// case 2.5 - a red parent with a black right child with a left red grandchild
	if n.red && !n.right.red && n.right.left.isRed() {
		n.right.left.red = false
		n.right.red = true
		n.right.rotateRight()
		n.rotateLeft()
		return n.parent
	}

	// case 3 - a black parent with a red right child with black grandchildren with left black grand-grandchildren
	if !n.red && n.right.isRed() && !n.right.right.isRed() && !n.right.left.left.isRed() && !n.right.left.right.isRed() {
		n.rotateLeft()
		n.parent.red = false
		n.right.red = true
		return n.parent
	}

	// case 4 - a black parent with a red right child with black grandchildren with a right red grand-grandchild
	if !n.red && n.right.isRed() && n.right.left != nil && !n.right.left.red && n.right.left.right.isRed() {
		newBlackNode := n.right.left.right
		n.right.rotateRight()
		newBlackNode.red = false
		n.rotateLeft()
		return n.parent
	}

	// case 5 - a black parent with a black right child with a red left grandchild
	if !n.red && !n.right.isRed() && n.right.left.isRed() {
		n.right.rotateRight()
		n.right.red = false
		n.rotateLeft()
		return n.parent
	}

	// case 5.5 - a black parent with a black right child with a red right grandchild
	if !n.red && !n.right.isRed() && n.right.right.isRed() {
		n.rotateLeft()
		if n.left != nil {
			n.left.red = false
		}
		n.parent.right.red = false
		return n.parent
	}

	// case 6
	if !n.red && !n.right.isRed() && !n.right.left.isRed() && !n.right.right.isRed() {
		n.right.red = true

		if n.parent == nil {
			return nil
		}

		if n.parent.right == n {
			return n.parent.balanceAfterRightRemoval()
		} else {
			return n.parent.balanceAfterLeftRemoval()
		}
	}

	panic("invalid case")
}

func (n *node) getParent() *node {
	if n == nil {
		return nil
	} else {
		return n.parent
	}
}

func (n *node) getLeft() *node {
	if n == nil {
		return nil
	} else {
		return n.left
	}
}

func (n *node) getRight() *node {
	if n == nil {
		return nil
	} else {
		return n.right
	}
}

func (n *node) isRed() bool {
	return n != nil && n.red
}

func (n *node) getSibling() *node {
	if n == nil || n.parent == nil {
		return nil
	}

	if n.parent.left == n {
		return n.parent.right
	} else {
		return n.parent.left
	}
}

func (n *node) replaceChild(oldChild, newChild *node) {
	if n == nil {
		return
	}

	if n.left == oldChild {
		n.left = newChild
	} else {
		n.right = newChild
	}
}

func (n *node) replaceParentForChildren(parent *node) {
	if n.left != nil {
		n.left.parent = parent
	}
	if n.right != nil {
		n.right.parent = parent
	}
}

func (n *node) findRoot() *node {
	root := n
	for ; root.parent != nil; root = root.parent {
	}
	return root
}

func (n *node) dumpValuesInDetails() {
	if n == nil {
		return
	}

	n.left.dumpValuesInDetails()
	fmt.Printf("Node: %d, red - %v, left - %v, right = %v, parent = %v\n", n.val, n.red, n.left, n.right, n.parent)
	n.right.dumpValuesInDetails()
}
