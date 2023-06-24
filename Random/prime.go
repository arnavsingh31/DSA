package main

import (
	"log"
	"math"
)

// T.C--> O(n) if number is even then T.C will be constant.
func isPrime(a int) bool {
	if a == 1 {
		return false
	}

	for i := 2; i < a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

// T.C-->O(n^1/2)..square root of n.
func isPrime2(a int) bool {
	sqrt_a := int(math.Sqrt(float64(a)))

	for i := 2; i < sqrt_a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

// all prime numbers can be expressed in the form: 6k +1 or 6k -1. hence we can reduce more iteration of loop.
func isPrime3(a int) bool {
	if a == 1 {
		return false
	}

	if a == 2 || a == 3 {
		return true
	}

	if a%2 == 0 || a%3 == 0 {
		return false
	}

	for i := 5; i*i <= a; i += 6 {
		if a%i == 0 || a%(i+2) == 0 {
			return false
		}
	}

	return true
}

func main() {
	log.Print(isPrime(97))
}
