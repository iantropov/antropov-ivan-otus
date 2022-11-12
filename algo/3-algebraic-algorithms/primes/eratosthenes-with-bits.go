package primes

func EratosthenesWithBits(num int) int {
	lengthInInt32 := num / 32
	if num%32 > 0 {
		lengthInInt32++
	}
	notPrimers := make([]int32, lengthInInt32)

	for i := 2; i < num; i++ {
		if isNotPrimer(i, notPrimers) == false {
			markNotPrimersWithBits(i, num, notPrimers)
		}
	}

	count := 0
	for i := 2; i < num; i++ {
		if isNotPrimer(i, notPrimers) == false {
			count++
		}
	}
	return count
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
	i := n / 32
	j := n % 32
	mask := int32(01 << j)
	return (notPrimers[i] & mask) == mask
}

func markAsNotPrimer(n int, notPrimers []int32) {
	i := n / 32
	j := n % 32
	notPrimers[i] |= (01 << j)
}
