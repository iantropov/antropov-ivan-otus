package fibonacci

import (
	"math"
	"math/big"
)

func GoldenRatio(n int) *big.Int {
	sqrt5 := math.Sqrt(5)
	fi := (1 + sqrt5) / 2

	return big.NewInt(int64(math.Floor(math.Pow(fi, float64(n))/sqrt5 + 0.5)))
}
