package fibonacci

import "math/big"

func Recursive(num int) *big.Int {
	return recursive(big.NewInt(int64(num)))
}

func recursive(num *big.Int) *big.Int {
	if num.Cmp(big.NewInt(2)) <= 0 {
		return big.NewInt(1)
	}

	num_1 := new(big.Int)
	num_1.Set(num)
	num_1 = num_1.Sub(num_1, big.NewInt(1))

	num_2 := new(big.Int)
	num_2.Set(num)
	num_2 = num_2.Sub(num_2, big.NewInt(2))

	num_1_sum := recursive(num_1)
	return num_1_sum.Add(num_1_sum, recursive(num_2))
}
