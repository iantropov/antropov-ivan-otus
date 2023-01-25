package prefix_test

import (
	"boyer-moore/prefix"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchSubstring(t *testing.T) {
	res := prefix.SearchSubstring("STRONGSTRING", "STRING")
	require.Equal(t, 6, res)
}
