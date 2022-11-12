package primes

import "math"

func Optimal(n int) []int {
	primes := make([]int, 0)
	for i := 2; i < n; i++ {
		if isPrimeWithPrimes(i, primes) {
			primes = append(primes, i)
		}
	}
	return primes
}

func isPrimeWithPrimes(n int, primes []int) bool {
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 0; i < len(primes) && primes[i] <= sqrtN; i++ {
		if n%primes[i] == 0 {
			return false
		}
	}
	return true
}
