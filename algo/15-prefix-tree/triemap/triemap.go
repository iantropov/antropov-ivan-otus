package triemap

type Trie[V any] struct {
	root       *node[V]
	emptyValue V
}

type node[V any] struct {
	chars  [123]*node[V]
	parent *node[V]
	end    bool
	value  V
	size   int
}

func Constructor[V any]() Trie[V] {
	root := new(node[V])
	return Trie[V]{root: root}
}

func (this *Trie[V]) Put(key string, value V) {
	n := this.root
	n.size++
	for _, c := range key {
		if n.chars[c] == nil {
			n.chars[c] = new(node[V])
			n.chars[c].parent = n
		}
		n = n.chars[c]
		n.size++
	}
	n.end = true
	n.value = value
}

func (this *Trie[V]) Get(key string) V {
	n := this.follow(key)
	if n != nil && n.end {
		return n.value
	} else {
		return this.emptyValue
	}
}

func (this *Trie[V]) Remove(key string) {
	n := this.follow(key)
	if n == nil || !n.end {
		return
	}

	n.end = false
	n.value = this.emptyValue

	p := n
	for ; p != nil; p = p.parent {
		p.size--
		if p.size == 0 {
			p.parent.removeChild(p)
		}
	}
}

func (this *Trie[V]) follow(s string) *node[V] {
	n := this.root
	for _, c := range s {
		n = n.chars[c]
		if n == nil {
			return nil
		}
	}
	return n
}

func (n *node[V]) removeChild(child *node[V]) {
	if n == nil {
		return
	}

	for i := range n.chars {
		if n.chars[i] == child {
			n.chars[i] = nil
		}
	}
}
