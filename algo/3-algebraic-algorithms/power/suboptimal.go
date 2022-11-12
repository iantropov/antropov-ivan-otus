package power

func Suboptimal(num float64, pow int) float64 {
	if pow == 0 {
		return 1.
	}

	baseCounter := num
	powerCounter := 1
	for {
		if powerCounter*2 > pow {
			break
		}
		powerCounter *= 2
		baseCounter *= baseCounter
	}

	if powerCounter == pow {
		return baseCounter
	} else {
		return baseCounter * Iterative(num, pow-powerCounter)
	}
}
