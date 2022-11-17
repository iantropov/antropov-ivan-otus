package countBits

type CountBits func(uint64) int

func CountBitsBySubstraction(value uint64) int {
	count := 0
	for value > 0 {
		value &= value - 1
		count++
	}
	return count
}

func CountBitsStraightforward(value uint64) int {
	count := 0
	for value > 0 {
		if value&1 == 1 {
			count++
		}
		value >>= 1
	}
	return count
}
