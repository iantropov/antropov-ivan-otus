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

func TestTrie2(t *testing.T) {
	trie := trie.Constructor()
	trie.Insert("apple")
	require.True(t, trie.Search("apple"))

	trie.Insert("app")
	require.True(t, trie.Search("app"))
}

func TestTrie3(t *testing.T) {
	trie := trie.Constructor()
	require.False(t, trie.Search("apple"))
}

func TestTrie4(t *testing.T) {
	trie := trie.Constructor()

	trie.Insert("hello")
	require.False(t, trie.Search("hell"))
	require.False(t, trie.Search("helloa"))
	require.True(t, trie.Search("hello"))

	require.True(t, trie.StartsWith("hell"))
	require.False(t, trie.StartsWith("helloa"))
	require.True(t, trie.StartsWith("hell"))
}

func TestTrie5(t *testing.T) {
	trie := trie.Constructor()

	trie.Insert("app")
	trie.Insert("apple")
	trie.Insert("beer")
	trie.Insert("add")
	trie.Insert("jam")
	trie.Insert("rental")

	require.False(t, trie.Search("apps"))
	require.True(t, trie.Search("app"))
	require.False(t, trie.Search("ad"))
	require.False(t, trie.Search("applepie"))
	require.False(t, trie.Search("rest"))
	require.False(t, trie.Search("jan"))
	require.False(t, trie.Search("rent"))
	require.True(t, trie.Search("beer"))
	require.True(t, trie.Search("jam"))

	require.False(t, trie.StartsWith("apps"))
	require.True(t, trie.StartsWith("app"))
	require.True(t, trie.StartsWith("ad"))
	require.False(t, trie.StartsWith("applepie"))
	require.False(t, trie.StartsWith("rest"))
	require.False(t, trie.StartsWith("jan"))
	require.True(t, trie.StartsWith("rent"))
	require.True(t, trie.StartsWith("beer"))
	require.True(t, trie.StartsWith("jam"))
}
