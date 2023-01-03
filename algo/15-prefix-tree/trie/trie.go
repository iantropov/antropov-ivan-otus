package trie

type Trie struct {
	root *node
}

type node struct {
	chars [123]*node
	end   bool
}

func Constructor() Trie {
	root := new(node)
	return Trie{root}
}

func (this *Trie) Insert(word string) {
	n := this.root
	for _, c := range word {
		if n.chars[c] == nil {
			n.chars[c] = new(node)
		}
		n = n.chars[c]
	}
	n.end = true
}

func (this *Trie) Search(word string) bool {
	n := this.follow(word)
	return n != nil && n.end
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.follow(prefix) != nil
}

func (this *Trie) follow(s string) *node {
	n := this.root
	for _, c := range s {
		n = n.chars[c]
		if n == nil {
			return nil
		}
	}
	return n
}
