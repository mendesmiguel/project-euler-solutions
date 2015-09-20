package main

import (
	"fmt"
	"math"
)

func sieve(n int) []int {
	var p = make([]bool, n)
	var primes = make([]int, 0, n)

	p[0] = true
	p[1] = true

	for i := 2; i < n; i++ {
		for j := i; i*j < n; j++ {
			p[i*j] = true
		}
		if !p[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func largestPrimeFactor(primes []int, n uint64) int {
	var p int
	t := 1

	for _, e := range primes {
		if n%uint64(e) == 0 {
			t *= e
			p = e
		}
		if uint64(t) == n {
			break
		}
	}

	return p
}

func main() {
	var n uint64 = 600851475143
	primes := sieve(int(math.Sqrt(float64(n))))
	fmt.Printf("Largest prime factor of %d is %d\n", n, largestPrimeFactor(primes, n))
}
