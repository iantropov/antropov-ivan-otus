package fsm_test

import (
	"kmp/fsm"
	"testing"
)

func TestFsm(t *testing.T) {
	res := fsm.SearchSubstring("AABAABAABAAABA", "AABAABAAABA")
	if res != 3 {
		t.Error("Invalid result", res)
	}
}
