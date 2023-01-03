package trie_test

import (
	"prefix-tree/trie"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTrie(t *testing.T) {
	/**
	 * Your Trie object will be instantiated and called as such:
	 * obj := Constructor();
	 * obj.Insert(word);
	 * param_2 := obj.Search(word);
	 * param_3 := obj.StartsWith(prefix);
	 */

	trie := trie.Constructor()
	trie.Insert("hello")

	require.True(t, trie.Search("hello"))
	require.True(t, trie.StartsWith("hel"))

	require.False(t, trie.Search("asd"))
	require.False(t, trie.StartsWith("asd"))
}
