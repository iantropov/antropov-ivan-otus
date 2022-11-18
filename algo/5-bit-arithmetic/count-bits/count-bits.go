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

func CountBitsWithPrecount(value uint64) int {
	bits := initBits()

	count := 0

	for value > 0 {
		count += bits[value&255]
		value >>= 8
	}
	return count

}

func initBits() [256]int {
	bits := [256]int{}
	for i := 0; i < 255; i++ {
		bits[i] = CountBitsBySubstraction(uint64(i))
	}
	return bits
}
