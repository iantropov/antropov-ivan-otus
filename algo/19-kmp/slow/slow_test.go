package slow_test

import (
	"kmp/slow"
	"testing"
)

func TestFsm(t *testing.T) {
	res := slow.SearchSubstring("AABAABAABAAABA", "AABAABAAABA")
	if res != 3 {
		t.Error("Invalid result", res)
	}
}
