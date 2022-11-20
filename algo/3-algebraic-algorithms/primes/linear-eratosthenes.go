package primes

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
	lp := make([]int, num+1)

	for i := 2; i <= num; i++ {
		if lp[i] == 0 {
			lp[i] = i
			pr = append(pr, i)
		}
		for pi := 0; pi < len(pr); pi++ {
			p := pr[pi]
			if p > lp[i] || p*i > num {
				break
			}
			lp[p*i] = p
		}
	}

	return len(pr)
}

func LinearEratosthenes2(n int) int {
	primes := make([]int, 0)
	mind := make([]int, n+1)

	for i := 2; i <= n; i++ {
		if mind[i] == 0 {
			mind[i] = i
			primes = append(primes, i)
		}

		for j := 0; j < len(primes) && i*primes[j] <= n; j++ {
			if primes[j] <= mind[i] {
				mind[i*primes[j]] = primes[j]
			}
		}
	}

	return len(primes)
}
