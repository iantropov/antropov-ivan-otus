package counting

func Sort(a []int) []int {
	max := a[0]
	for _, el := range a {
		if el > max {
			max = el
		}
	}

	counts := make([]int, max+1)
	for _, el := range a {
		counts[el]++
	}

	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	res := make([]int, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		counts[a[i]]--
		pos := counts[a[i]]
		res[pos] = a[i]
	}

	return res
}
