package main

func iterateOverDigits(currentN, originalN, count, sum1, sum2 int) int {
	if currentN == 0 {
		if sum1 == sum2 {
			return count + 1
		} else {
			return count
		}
	}

	for i := 0; i < 10; i++ {
		if currentN > originalN {
			count = iterateOverDigits(currentN-1, originalN, count, sum1+i, sum2)
		} else {
			count = iterateOverDigits(currentN-1, originalN, count, sum1, sum2+i)
		}
	}

	return count
}

func FindHappyTicketsAsMiddle(n int) int {
	return iterateOverDigits(2*n, n, 0, 0, 0)
}
