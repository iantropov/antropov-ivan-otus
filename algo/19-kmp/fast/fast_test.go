package fast_test

import (
	"kmp/fast"
	"testing"
)

func TestFast(t *testing.T) {
	res := fast.SearchSubstring("AABAABAABAAABA", "AABAABAAABA")
	if res != 3 {
		t.Error("Invalid result", res)
	}
}
