package main

// Naive approach. Here T.C--> O(min(a,b))
func gcd(a, b int) int {
	min := small(a, b)

	for min > 0 {
		if a%min == 0 && b%min == 0 {
			break
		}
		min--
	}
	return min

}

func small(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// using Euclidean Algorithm, says that:-
// let b be the smaller than a then gcd(a,b) = gcd(a-b, b)
// This is basic version of algo. instead of doing repeated subtraction we can do modulo operation to optimise
// the algo.
func gcd2(a, b int) int {

	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}

// optimised implimentatioin of Euclidean Algorithm
func gcd3(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd3(b, a%b)
}
