package rbt

// https://habr.com/ru/company/otus/blog/472040/

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

func (tree *rbtTree) Insert(val int) {
	newNode := tree.root.insert(val, tree.root)
	newNode.balance()
}

func (n *node) insert(val int, parent *node) *node {
	if n == nil {
		return &node{
			val:    val,
			red:    true,
			parent: parent,
		}
	} else if val < n.val {
		return n.left.insert(val, n)
	} else if val > n.val {
		return n.right.insert(val, n)
	} else {
		return nil
	}
}

func (n *node) balance() {
	if n == nil {
		return
	}

	if n.parent == nil {
		if n.red {
			n.red = false
		}
		return
	}

	if n.parent.parent == nil {
		return
	}

	p := n.parent
	g := p.parent
	u := g.getSibling()

	// case 0 - parent is black
	if !p.red {
		return
	}

	// case 1 - uncle is red
	if !g.red && p.red && u.isRed() {
		p.red = false
		u.red = false
		g.red = true
		g.balance()
		return
	}

	// case 2 - uncle is black & (grandpa & parent are on different sides)
	if !g.red && p.red && !u.isRed() && (g.left == p && p.right == n || g.right == p && p.left == n) {
		if g.left == p && p.right == n {
			p.rotateLeft()
		} else {
			p.rotateLeft()
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
	}
}

func (n *node) rotateRight() {
	oldParent := n.parent
	oldLeftRight := n.left.right

	newNode := n.left
	newNode.right = n
	newNode.parent = n.parent

	n.parent = newNode
	n.left = oldLeftRight

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

	oldParent.replaceChild(n, newNode)
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
