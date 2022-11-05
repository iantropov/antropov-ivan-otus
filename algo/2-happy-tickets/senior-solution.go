package main

// var memo [][]int

var countsOfSums []int

func findCountOfSum(n, sum, count int) int {
	if n == 0 {
		if sum == 0 {
			return count + 1
		} else {
			return count
		}
	}

	for i := 0; i < 10; i++ {
		count = findCountOfSum(n-1, sum-i, count)
	}

	return count
}

func FindHappyTicketsAsSenior(n int) int {
	numberOfSums := n * 9
	countsOfSums = make([]int, numberOfSums+1)

	for sum := 0; sum <= numberOfSums; sum++ {
		countsOfSums[sum] = findCountOfSum(n, sum, 0)
	}

	result := 0
	for i := 0; i <= numberOfSums; i++ {
		result += countsOfSums[i] * countsOfSums[i]
	}

	return result
}
