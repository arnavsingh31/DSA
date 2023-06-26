package main

import "log"

// prints all divisors of a number in ascending order*
func allDivisors(n int) {
	var i int
	for i = 1; i*i < n; i++ {
		if n%i == 0 {
			log.Print(i)
		}
	}

	for j := i; j >= 1; j-- {
		if n%j == 0 {
			log.Print(n / j)
		}
	}
}

func main() {
	allDivisors(12)
}
