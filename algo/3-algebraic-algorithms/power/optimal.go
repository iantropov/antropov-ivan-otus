package power

func Optimal(num, pow int) int {
	if pow <= 1 {
		return num
	}

	res := 0
	d := 2
	for n := num; n >= 1; {
		d *= d
		if n%2 == 1 {
			res += d
		}
		n /= 2
	}

	return res
}
