package primes

func EratosthenesWithBits(num int) int {
	correctedN := num / 2
	lengthInInt32 := correctedN / 32
	if correctedN%32 > 0 {
		lengthInInt32++
	}
	notPrimers := make([]int32, lengthInInt32)

	for i := 3; i < num; i++ {
		if i%2 == 1 && isNotPrimer(i, notPrimers) == false {
			markNotPrimersWithBits(i, num, notPrimers)
		}
	}

	count := 0
	for i := 3; i < num; i++ {
		if i%2 == 1 && isNotPrimer(i, notPrimers) == false {
			count++
		}
	}
	return count + 1 // add '2'
}

func markNotPrimersWithBits(n, num int, notPrimers []int32) {
	addition := n
	if n > 2 {
		addition += n
	}
	for i := n * n; i < num; i += addition {
		markAsNotPrimer(i, notPrimers)
	}
}

func isNotPrimer(n int, notPrimers []int32) bool {
	correctedN := n/2 - 1
	i := correctedN / 32
	j := correctedN % 32
	mask := int32(01 << j)
	return (notPrimers[i] & mask) == mask
}

func markAsNotPrimer(n int, notPrimers []int32) {
	correctedN := n/2 - 1
	i := correctedN / 32
	j := correctedN % 32
	notPrimers[i] |= (01 << j)
}
