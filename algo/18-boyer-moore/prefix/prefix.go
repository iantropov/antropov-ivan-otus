package prefix

func SearchSubstring(text, mask string) int {
	textRunes := []rune(text)
	maskRunes := []rune(mask)

	shifts := calculateShifts(maskRunes)
	var i, j int
	for i = 0; i < len(textRunes)-len(maskRunes)+1; {
		for j = len(maskRunes) - 1; j >= 0 && textRunes[i+j] == maskRunes[j]; j-- {
		}

		if j == -1 {
			return i
		}

		if s, ok := shifts[textRunes[i+len(maskRunes)-1]]; ok {
			i += s
		} else {
			i += len(maskRunes)
		}
	}
	return -1
}

func calculateShifts(maskRunes []rune) map[rune]int {
	shifts := make(map[rune]int)
	for i := 0; i < len(maskRunes)-1; i++ {
		shifts[maskRunes[i]] = len(maskRunes) - i - 1
	}
	return shifts
}
