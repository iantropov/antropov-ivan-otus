package btree

type node struct {
	values   []int
	pointers []*node
	size     int
	leaf     bool
}

type btree struct {
	root   *node
	degree int
}

func NewTree(degree int) *btree {
	return &btree{nil, degree}
}

func (tree *btree) Insert(val int) {
	if tree.root == nil {
		tree.root = buildNode(tree.degree)
		tree.root.insertIntoNode(val, nil)
		tree.root.leaf = true
		return
	}

	if tree.root.size == tree.degree-1 {
		newRoot := buildNode(tree.degree)

		tree.root
	}

	newNode := tree.root.insert(val)
	if newNode != nil {
		tree.root = newNode
	}
}

func (n *node) insert(val int) *node {
	degree := len(n.pointers)

	if n.leaf && n.size < degree-1 {
		n.insertIntoNode(val, nil)
		return nil
	} else if n.leaf {
		newVal, newChild := n.split()
		newNode := buildNode()
		n.insertIntoNode(newVal, newChild)
		if val < newVal {
			n.insert(val)
		} else {
			newChild.insert(val)
		}

	}

	childPos := n.findChildPosFor(val)

	child := n.pointers[childPos]

	if child.size < degree-1 {
		return child.insert(val)
	} else {
		newVal, newChild := child.split()
		n.insertIntoNode(newVal, newChild)
		if val < newVal {
			return child.insert(val)
		} else {
			return newChild.insert(val)
		}
	}
}

func (n *node) insertIntoNode(val int, pointer *node) {
	if n.size == len(n.values) {
		panic("the leaf is full")
	}

	pos := 0
	for ; n.values[pos] < val; pos++ {
	}

	for i := n.size; i > pos; i-- {
		n.values[i] = n.values[i-1]
		n.pointers[i] = n.pointers[i-1]
	}

	n.values[pos] = val
	n.pointers[pos] = pointer
}

func (n *node) findChildPosFor(val int) int {
	pos := 0
	for ; n.values[pos] < val; pos++ {
	}
	return pos - 1
}

func (n *node) split() (int, *node) {
	if n.size < len(n.values)/2+1 {
		panic("the leaf is too sparse")
	}

	pos := n.size / 2
	if n.size%2 == 1 {
		pos++
	}

	newVal := n.values[pos]
	newChild := buildNode(len(n.pointers))
	for i := pos + 1; i < n.size; i++ {
		newChild.values[i-pos+1] = n.values[i]
		newChild.pointers[i-pos+1] = n.pointers[i]
		n.values[i] = 0
		n.pointers[i] = nil
	}

	newChild.size = n.size - pos
	n.size = pos

	return newVal, newChild
}

func buildNode(degree int) *node {
	newNode := &node{}
	newNode.values = make([]int, degree-1)
	newNode.pointers = make([]*node, degree)
	return newNode
}
