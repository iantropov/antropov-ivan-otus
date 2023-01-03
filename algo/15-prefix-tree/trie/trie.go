package trie

const ASCII_BASE = 97

type Trie struct {
	root *node
}

type node struct {
	chars [26]*node
	end   bool
}

func Constructor() Trie {
	root := new(node)
	return Trie{root}
}

func (this *Trie) Insert(word string) {
	n := this.root
	for _, c := range word {
		if n.chars[c-ASCII_BASE] == nil {
			n.chars[c-ASCII_BASE] = new(node)
		}
		n = n.chars[c-ASCII_BASE]
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
		n = n.chars[c-ASCII_BASE]
		if n == nil {
			return nil
		}
	}
	return n
}
