package primes

import "fmt"

/*
Вход: натуральное число n

Пусть pr - целочисленный массив, поначалу пустой;

	     lp - целочисленный массив, индексируемый от 2 до n, заполненный нулями

	для i := 2, 3, 4, ..., до n:
	  если lp[i] = 0:
	      lp[i] := i
	      pr[] += {i}
	  для p из pr пока p ≤ lp[i] и p*i ≤ n:
	      lp[p*i] := p

Выход: все числа в массиве pr.

https://habr.com/ru/post/452388/
*/

func LinearEratosthenes(num int) int {
	pr := make([]int, 0)
	lp := make([]int, num)
	counts := make([]int, num)

	for i := 2; i < num; i++ {
		if lp[i] == 0 {
			lp[i] = i
			counts[i]++
			pr = append(pr, i)
		}
		for pi := 0; pi < len(pr); pi++ {
			p := pr[pi]
			if p > lp[i] || p*i >= num {
				break
			}
			lp[p*i] = p
			counts[p*i]++
		}
	}

	fmt.Println(counts)
	return len(pr)
}
