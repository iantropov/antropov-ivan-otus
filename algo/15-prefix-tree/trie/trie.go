package trie

type Trie struct {
	root *node
}

type node struct {
	char *nodeChar
	end  bool
}

type nodeChar struct {
	c        rune
	nextChar *nodeChar
	nextNode *node
}

func Constructor() Trie {
	root := new(node)
	return Trie{root}
}

func (this *Trie) Insert(word string) {
	n := this.root
	for _, c := range word {
		nc := n.char
		if nc == nil {
			n.char = new(nodeChar)
			n.char.c = c
			nc = n.char
		} else {
			prev := (*nodeChar)(nil)
			for ; nc != nil && nc.c < c; nc = nc.nextChar {
				prev = nc
			}
			if nc == nil {
				newNc := new(nodeChar)
				newNc.c = c
				prev.nextChar = newNc
				nc = newNc
			} else if nc.c != c {
				newNc := new(nodeChar)
				newNc.c = c
				newNc.nextChar = nc
				if prev != nil {
					prev.nextChar = newNc
				} else {
					newNc.nextChar = n.char
					n.char = newNc
				}
				nc = newNc
			}
		}
		if nc.nextNode == nil {
			nc.nextNode = new(node)
		}
		n = nc.nextNode
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
		nodeChar := n.char
		if n.char == nil {
			return nil
		}

		for ; nodeChar.nextChar != nil; nodeChar = nodeChar.nextChar {
			if nodeChar.c == c {
				break
			}
		}
		if nodeChar.c != c {
			return nil
		}
		n = nodeChar.nextNode
		if n == nil {
			return nil
		}
	}
	return n
}
