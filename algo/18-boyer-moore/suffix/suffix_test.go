package suffix_test

import (
	"boyer-moore/suffix"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchSubstring(t *testing.T) {
	res := suffix.SearchSubstring("TRINGSTRONGSTRING", "STRING")
	require.Equal(t, 11, res)

	res = suffix.SearchSubstring("STROTRINGASTRINGSTRONGSTRING", "STRONGSTRING")
	require.Equal(t, 16, res)
}
