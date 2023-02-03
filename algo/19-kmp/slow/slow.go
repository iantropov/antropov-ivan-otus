package slow

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
	for i := 0; i <= len(pattern); i++ {
		line := make([]byte, 0)
		line = append(line, pattern[:i]...)
		for j := 0; j < i; j++ {
			if areEqualPrefixAndSuffix(line, line, j) {
				pi[i] = j
			}
		}
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
