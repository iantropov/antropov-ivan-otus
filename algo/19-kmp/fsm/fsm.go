package fsm

// handles only A-Z
const ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func SearchSubstring(text, pattern string) int {
	return -1
}

func buildDelta(pattern []byte) [][]int {
	delta := make([][]int, len(pattern))
	alphabet := []byte(ALPHABET)
	for i := 0; i < len(pattern); i++ {
		delta[i] = make([]int, len(alphabet))
		for j := 0; j < len(alphabet); j++ {
			line := append(pattern[:i], alphabet[j])
			k := i + 1
			for !areEqualPrefixAndSuffix(pattern, line, k) {
				k--
			}
			delta[i][j] = k
		}
	}
	return delta
}

func areEqualPrefixAndSuffix(left, right []byte, k int) bool {
	return false
}
