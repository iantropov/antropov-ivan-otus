package radix

func Sort(a []int) []int {
	max := a[0]
	for _, el := range a {
		if el > max {
			max = el
		}
	}

	radixCount := getRadixCount(max)

	arrays := [][]int{
		a,
		make([]int, len(a)),
	}
	d := make([]int, 10)
	for r := 1; r <= radixCount; r++ {
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
	}

	return arrays[1]
}

func getRadixCount(num int) int {
	res := 0
	for ; num > 0; num /= 10 {
		res++
	}
	return res
}

func getRadix(num, radix int) int {
	num = num % (radix + 1)
	return num / radix
}
