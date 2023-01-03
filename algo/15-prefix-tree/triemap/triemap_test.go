package triemap_test

import (
	"prefix-tree/triemap"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTriemap(t *testing.T) {
	triemap := triemap.Constructor[int]()
	triemap.Put("hello", 1)
	triemap.Put("asd", 2)

	require.Equal(t, 1, triemap.Get("hello"))
	require.Equal(t, 2, triemap.Get("asd"))
	require.Equal(t, 0, triemap.Get("asdasd"))
}
