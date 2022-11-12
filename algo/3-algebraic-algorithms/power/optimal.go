package power

func Optimal(num, pow int) int {
	if pow <= 1 {
		return num
	}

	res := 1
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
