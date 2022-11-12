package primes

func Eratosthenes(num int) int {
	notPrimers := make([]bool, num)

	for i := 2; i < num; i++ {
		if notPrimers[i] != true {
			markNotPrimers(i, notPrimers)
		}
	}

	count := 0
	for i := 2; i < num; i++ {
		if notPrimers[i] == false {
			count++
		}
	}
	return count
}

func markNotPrimers(n int, notPrimers []bool) {
	addition := n
	if n > 2 {
		addition += n
	}
	for i := n * n; i < len(notPrimers); i += addition {
		notPrimers[i] = true
	}
}
