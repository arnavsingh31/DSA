package main

import (
	"math"
)

func primeFactors(a int) []int {
	if a <= 1 {
		return []int{}
	}
	sqrt_a := int(math.Sqrt(float64(a)))
	ans := []int{}
	for i := 2; i <= sqrt_a; i++ {
		for a%i == 0 {
			ans = append(ans, i)
			a = a / i
		}
	}
	if a > 1 {
		ans = append(ans, a)
	}
	return ans
}

// optimised using similar approach like in prime numbers.
// T.C--> in worst case number will be prime number so, neither of the two outer loops(loop1&2) will run,
// only loop 3 will run that too no inner loops (loops inside loop 3) will run. so ony outer loop 3 will
// run upto square root of 'a' times. Hence, T.C--> O(n^1/2)
func primeFactors2(a int) []int {
	ans := []int{}
	if a <= 1 {
		return ans
	}
	// loop 1
	for a%2 == 0 {
		ans = append(ans, 2)
		a = a / 2
	}
	// loop2
	for a%3 == 0 {
		ans = append(ans, 3)
		a = a / 3
	}
	sqrt_a := int(math.Sqrt(float64(a)))
	// loop 3
	for i := 5; i <= sqrt_a; i = 6 {
		for a%i == 0 {
			ans = append(ans, i)
			a = a / i
		}

		for a%(i+2) == 0 {
			ans = append(ans, i+2)
			a = a / (i + 2)
		}
	}

	if a > 3 {
		ans = append(ans, a)
	}

	return ans
}
