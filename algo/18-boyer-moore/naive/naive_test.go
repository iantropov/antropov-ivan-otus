package naive_test

import (
	"boyer-moore/naive"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchSubstring(t *testing.T) {
	res := naive.SearchSubstring("TRINGSTRONGSTRING", "STRING")
	require.Equal(t, 11, res)

	res = naive.SearchSubstring("STROTRINGASTRINGSTRONGSTRING", "STRONGSTRING")
	require.Equal(t, 16, res)
}
