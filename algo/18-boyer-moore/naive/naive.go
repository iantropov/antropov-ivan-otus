package naive

func SearchSubstring(text, mask string) int {
	textRunes := []rune(text)
	maskRunes := []rune(mask)
	var i, j int
	for i = 0; i < len(textRunes)-len(maskRunes)+1; i++ {
		for j = 0; j < len(maskRunes) && textRunes[i+j] == maskRunes[j]; j++ {
		}
		if j == len(maskRunes) {
			return i
		}
	}
	return -1
}
