package power

func Iterative(num, pow int) int {
	res := 1

	for i := 1; i <= pow; i++ {
		res *= num
	}

	return res
}
