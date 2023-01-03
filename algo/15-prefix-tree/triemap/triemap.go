package triemap

type Trie[V any] struct {
	root       *node[V]
	emptyValue V
}

type node[V any] struct {
	chars [123]*node[V]
	end   bool
	value V
}

func Constructor[V any]() Trie[V] {
	root := new(node[V])
	return Trie[V]{root: root}
}

func (this *Trie[V]) Put(key string, value V) {
	n := this.root
	for _, c := range key {
		if n.chars[c] == nil {
			n.chars[c] = new(node[V])
		}
		n = n.chars[c]
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
