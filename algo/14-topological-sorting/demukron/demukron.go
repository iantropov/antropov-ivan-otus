package demukron

func Sort(g [][]int) []int {
	lenG := len(g)

	a := make([][]int, lenG)
	for i := range a {
		a[i] = make([]int, lenG)
	}

	v := make([]int, lenG)

	for i := range g {
		for _, j := range g[i] {
			a[i][j] = 1
			v[j]++
		}
	}

	res := make([]int, lenG)

	for i := 0; i < lenG; i++ {
		zeroIdx := 0
		for ; v[zeroIdx] != 0; zeroIdx++ {
		}

		res[i] = zeroIdx

		for j := 0; j < lenG; j++ {
			v[j] -= a[zeroIdx][j]
		}
		v[zeroIdx] = -1
	}

	return res
}
