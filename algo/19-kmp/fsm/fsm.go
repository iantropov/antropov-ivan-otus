package fsm

// handles only A-Z
var alphabet = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func SearchSubstring(text, pattern string) int {
	delta := buildDelta([]byte(pattern))
	patternIdx := 0
	for i, t := range []byte(text) {
		patternIdx = delta[patternIdx][t-alphabet[0]]
		if patternIdx == len(pattern) {
			return i - len(pattern) + 1
		}
	}
	return -1
}

func buildDelta(pattern []byte) [][]int {
	delta := make([][]int, len(pattern))
	for i := 0; i < len(pattern); i++ {
		delta[i] = make([]int, len(alphabet))
		for j := 0; j < len(alphabet); j++ {
			line := append([]byte(nil), pattern[:i]...)
			line = append(line, alphabet[j])
			k := i + 1
			for k > 0 && !areEqualPrefixAndSuffix(pattern, line, k) {
				k--
			}
			delta[i][j] = k
		}
	}
	return delta
}

func areEqualPrefixAndSuffix(left, right []byte, k int) bool {
	for i := 0; i < k; i++ {
		if left[i] != right[len(right)-k+i] {
			return false
		}
	}
	return true
}
