package splay

type splayTree struct {
	root *node
}

type node struct {
	value       int
	left, right *node
}

func NewTree() *splayTree {
	return &splayTree{}
}

func (tree *splayTree) Insert(val int) {
	t1, t2 := tree.split(val)
	tree.root = &node{val, t1, t2}
}

func (tree *splayTree) Search(val int) bool {
	node := tree.root.searchNode(val)
	if node == nil {
		return false
	}
	tree.root = tree.root.splay(val)
	return true
}

func (tree *splayTree) Remove(val int) {

}

func (node *node) searchNode(val int) *node {
	if node == nil {
		return nil
	}

	if val < node.value {
		return node.left.searchNode(val)
	} else if val > node.value {
		return node.right.searchNode(val)
	} else {
		return node
	}
}

func (node *node) splay(val int) *node {
	if node == nil {
		return nil
	} else if node.value == val {
		return node
	}

	if val < node.value {
		if node.left.value == val {
			return node.zigLeft()
		} else if node.left.left.value == val {
			return node.zigZigLeft()
		} else if node.left.right.value == val {
			return node.zigZagLeft()
		} else {
			node.left = node.left.splay(val)
		}
	} else {
		if node.right.value == val {
			return node.zigRight()
		} else if node.right.left.value == val {
			return node.zigZagRight()
		} else if node.right.right.value == val {
			return node.zigZigRight()
		} else {
			node.right = node.right.splay(val)
		}
	}

	return node.splay(val)
}

func (node *node) zigLeft() *node {
	newNode := node.left
	node.left = node.left.right
	newNode.right = node
	return newNode
}

func (node *node) zigRight() *node {
	newNode := node.right
	node.right = node.right.left
	newNode.left = node
	return newNode
}

func (node *node) zigZigLeft() *node {
	newNode := node.left.left

	oldNodeLeft := node.left
	oldNodeLeftRight := node.left.right

	node.left.left = node.left.left.right
	node.left.right = node
	node.left = oldNodeLeftRight

	newNode.right = oldNodeLeft

	return newNode
}

func (node *node) zigZigRight() *node {
	newNode := node.right.right

	oldNodeRight := node.right
	oldNodeRightLeft := node.right.left

	node.right.right = node.right.right.left
	node.right.left = node
	node.right = oldNodeRightLeft

	newNode.left = oldNodeRight

	return newNode
}

func (node *node) zigZagLeft() *node {
	newNode := node.left.right

	oldNodeLeftRightRight := node.left.right.right
	oldNodeLeft := node.left

	node.left.right = node.left.right.left
	node.left = oldNodeLeftRightRight

	newNode.left = oldNodeLeft
	newNode.right = node

	return newNode
}

func (node *node) zigZagRight() *node {
	return nil
}

func (tree *splayTree) split(val int) (*node, *node) {
	return nil, nil
}

func (tree *splayTree) merge(otherTree *splayTree) {

}
