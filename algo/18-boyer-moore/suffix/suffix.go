package suffix

import "fmt"

func SearchSubstring(text, mask string) int {
	textRunes := []rune(text)
	maskRunes := []rune(mask)
	shifts := makeShifts(maskRunes)
	fmt.Println(shifts)
	for i := 0; i < len(text)-len(mask)+1; {
		j := len(mask) - 1
		for ; j >= 0 && textRunes[i+j] == maskRunes[j]; j-- {

		}
		if j < 0 {
			return i
		}
		i += shifts[len(mask)-1-j]
	}
	return -1
}

func makeShifts(mask []rune) []int {
	shifts := make([]int, len(mask))
	shifts[0] = 1
	for i := 1; i < len(mask); i++ {
		for j := len(mask) - i - 1; j >= 0; j-- {
			if areRangesEqual(mask, j, len(mask)-i, i) {
				shifts[i] = len(mask) - 1 - j
				break
			}
		}
		if shifts[i] == 0 {
			shifts[i] = len(mask)
		}
	}
	return shifts
}

func areRangesEqual(mask []rune, startA, startB, length int) bool {
	for i := 0; i < length; i++ {
		if mask[startA+i] != mask[startB+i] {
			return false
		}
	}
	return true
}
