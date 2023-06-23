package main

// Iterative approach. It gives S.C-->O(1) and T.C-->O(n)
func factorial(n int) int {
	res := 1

	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

// recursive approach. It give S.C -> O(n) and T.C--> O(n)
func fact(n int) (res int) {
	if n == 1 {
		return 1
	}

	return n * fact(n-1)
}
