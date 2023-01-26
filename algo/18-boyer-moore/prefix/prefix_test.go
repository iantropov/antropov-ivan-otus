package prefix_test

import (
	"boyer-moore/prefix"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchSubstring(t *testing.T) {
	res := prefix.SearchSubstring("TRINGSTRONGSTRING", "STRING")
	require.Equal(t, 11, res)

	res = prefix.SearchSubstring("STROTRINGASTRINGSTRONGSTRING", "STRONGSTRING")
	require.Equal(t, 16, res)
}
