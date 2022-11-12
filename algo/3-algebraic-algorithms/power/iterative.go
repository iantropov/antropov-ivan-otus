package power

func Iterative(num float64, pow int) float64 {
	res := 1.0

	for i := 1; i <= pow; i++ {
		res *= num
	}

	return res
}
