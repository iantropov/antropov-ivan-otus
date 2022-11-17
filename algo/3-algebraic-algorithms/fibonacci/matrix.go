package fibonacci

import "math/big"

func Matrix(n int) *big.Int {
	if n == 1 || n == 2 {
		return big.NewInt(1)
	}

	return big.NewInt(int64(powerMatrix([][]int{{1, 1}, {1, 0}}, n-1)[0][0]))
}

func powerMatrix(a [][]int, pow int) [][]int {
	if pow <= 1 {
		return a
	}

	res := buildOneMatrix(len(a), len(a[0]))
	d := a
	for n := pow; n > 0; {
		if n%2 == 1 {
			res = multiplyMatrixes(res, d)
		}
		d = multiplyMatrixes(d, d)
		n /= 2
	}

	return res
}

func buildOneMatrix(n, m int) [][]int {
	c := make([][]int, n)

	for i := range c {
		c[i] = make([]int, m)
		c[i][i] = 1
	}

	return c
}

func multiplyMatrixes(a, b [][]int) [][]int {
	c := make([][]int, len(a))

	for i := range c {
		c[i] = make([]int, len(b[0]))
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b[0]); j++ {
			for k := 0; k < len(a[0]); k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return c
}
