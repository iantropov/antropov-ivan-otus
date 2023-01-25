package naive_test

import (
	"boyer-moore/naive"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchSubstring(t *testing.T) {
	res := naive.SearchSubstring("STRONGSTRING", "STRING")
	require.Equal(t, 6, res)
}
