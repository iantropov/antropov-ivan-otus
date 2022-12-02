package radix

func Sort(a []int) []int {
	max := a[0]
	for _, el := range a {
		if el > max {
			max = el
		}
	}

	arrays := [][]int{
		a,
		make([]int, len(a)),
	}
	d := make([]int, 10)
	for r := 1; r <= max; r *= 10 {
		for i := range d {
			d[i] = 0
		}

		for _, el := range arrays[0] {
			elRadix := getRadix(el, r)
			d[elRadix]++
		}

		for i := 1; i < len(d); i++ {
			d[i] += d[i-1]
		}

		for i := len(a) - 1; i >= 0; i-- {
			elRadix := getRadix(arrays[0][i], r)
			d[elRadix]--
			pos := d[elRadix]
			arrays[1][pos] = arrays[0][i]
		}

		arrays[0], arrays[1] = arrays[1], arrays[0]
	}

	return arrays[0]
}

func getRadixCount(num int) int {
	res := 0
	for ; num > 0; num /= 10 {
		res++
	}
	return res
}

func getRadix(num, radix int) int {
	num = num % (radix * 10)
	return num / radix
}
