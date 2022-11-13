package fibonacci

import "math/big"

func Iterative(n int) *big.Int {
	if n == 1 || n == 2 {
		return big.NewInt(1)
	}

	num := big.NewInt(0)
	num_1 := big.NewInt(1)
	num_2 := big.NewInt(1)

	for i := 3; i <= n; i++ {
		num = num.Add(num_1, num_2)

		num_2 = new(big.Int)
		num_2.Set(num_1)

		num_1 = new(big.Int)
		num_1.Set(num)
	}

	return num
}
