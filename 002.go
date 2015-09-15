package main

import "fmt"

func main() {
	fmt.Println(fib(4000000))
}

func fib(n int64) int64 {
	var a, b, sum int64
	a, b, sum = 1, 1, 0

	for a < n {
		if a % 2 == 0 {
			sum += a
		}
		a, b = b, b + a
	}
	return sum
}