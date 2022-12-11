package rbt

type node struct {
	val                 int
	left, right, parent *node
	isRed               bool
}

type rbtTree struct {
	root *node
}

func NewTree() *rbtTree {
	return &rbtTree{}
}

func (tree *rbtTree) Insert(val int) {
	tree.root = tree.root.insert(val)
}

func (node *node) insert(val int) *node {
	return node
}
