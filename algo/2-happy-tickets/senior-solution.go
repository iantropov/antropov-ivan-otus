package main

var memo [][]int

var countsOfSums []int

func findCountOfSum(n, sum, count int) int {
	if n == 0 {
		if sum == 0 {
			return count + 1
		} else {
			return count
		}
	}

	if sum < 0 {
		return count
	}

	if sum > 9*n {
		return count
	}

	if memo[n-1][sum] != -1 {
		return memo[n-1][sum]
	}

	for i := 0; i < 10; i++ {
		count += findCountOfSum(n-1, sum-i, 0)
	}

	memo[n-1][sum] = count

	return count
}

func FindHappyTicketsAsSenior(n int) int {
	numberOfSums := n * 9
	countsOfSums = make([]int, numberOfSums+1)

	memo = make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, (i+1)*9+1)
		for j := 0; j <= (i+1)*9; j++ {
			memo[i][j] = -1
		}
	}

	for sum := 0; sum <= numberOfSums; sum++ {
		countsOfSums[sum] = findCountOfSum(n, sum, 0)
	}

	result := 0
	for i := 0; i <= numberOfSums; i++ {
		result += countsOfSums[i] * countsOfSums[i]
	}

	return result
}
