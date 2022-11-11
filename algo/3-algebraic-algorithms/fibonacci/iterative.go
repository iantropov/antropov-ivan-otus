package fibonacci

func Iterative(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	num := 0
	num_1 := 1
	num_2 := 1

	for i := 3; i <= n; i++ {
		num = num_1 + num_2
		num_2 = num_1
		num_1 = num
	}

	return num
}
