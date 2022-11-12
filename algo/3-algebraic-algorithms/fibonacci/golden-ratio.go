package fibonacci

import "math"

func GoldenRatio(n int) int {
	sqrt5 := math.Sqrt(5)
	fi := (1 + sqrt5) / 2

	return int(math.Floor(math.Pow(fi, float64(n))/sqrt5 + 0.5))
}
