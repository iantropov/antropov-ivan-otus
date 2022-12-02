package merge

func Sort(a []int) []int {
	sortRecursive(a, 0, len(a)-1)
	return a
}

func sortRecursive(a []int, l, r int) {
	if l == r {
		return
	}

	m := (l + r) / 2
	sortRecursive(a, l, m)
	sortRecursive(a, m+1, r)
	merge(a, l, m, r)
}

func merge(a []int, l, m, r int) {
	tmp := make([]int, r-l+1)

	i := l
	j := m + 1
	k := 0

	for ; i <= m && j <= r; k++ {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			i++
		} else {
			tmp[k] = a[j]
			j++
		}
	}

	for ; i <= m; k++ {
		tmp[k] = a[i]
		i++
	}

	for ; j <= r; k++ {
		tmp[k] = a[j]
		j++
	}

	for ti := 0; ti < len(tmp); ti++ {
		a[l+ti] = tmp[ti]
	}
}
