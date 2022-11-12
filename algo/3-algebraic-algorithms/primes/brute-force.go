package primes

func BruteForce(num int) []int {
	primes := make([]int, 0)
	for i := 2; len(primes) < num; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes
}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
