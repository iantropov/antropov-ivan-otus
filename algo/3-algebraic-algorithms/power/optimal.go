package power

func Optimal(num float64, pow int) float64 {
	if pow == 1 {
		return num
	}

	res := 1.
	d := num
	for n := pow; n > 0; {
		if n%2 == 1 {
			res *= d
		}
		d *= d
		n /= 2
	}

	return res
}
