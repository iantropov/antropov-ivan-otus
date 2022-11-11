package fibonacci

func Recursive(num int) int {
	if num == 1 || num == 2 {
		return 1
	}

	return Recursive(num-1) + Recursive(num-2)
}
