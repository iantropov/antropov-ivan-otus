package fast

// handles only A-Z
var alphabet = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func SearchSubstring(text, pattern string) int {
	newPattern := make([]byte, 0)
	newPattern = append(newPattern, pattern...)
	newPattern = append(newPattern, '@')
	newPattern = append(newPattern, text...)
	pi := buildPi(newPattern)
	for i := 0; i <= len(text); i++ {
		if pi[i+len(pattern)+1] == len(pattern) {
			return i - len(pattern)
		}
	}
	return -1
}

func buildPi(pattern []byte) []int {
	pi := make([]int, len(pattern)+1)
	pi[1] = 0
	for q := 1; q < len(pattern); q++ {
		len := pi[q]
		for len > 0 && pattern[len] != pattern[q] {
			len = pi[len]
		}
		if pattern[len] == pattern[q] {
			len++
		}
		pi[q+1] = len
	}
	return pi
}

func areEqualPrefixAndSuffix(left, right []byte, k int) bool {
	for i := 0; i < k; i++ {
		if left[i] != right[len(right)-k+i] {
			return false
		}
	}
	return true
}
