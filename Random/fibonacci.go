package main

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}
	res := fib(n-1) + fib(n-2)
	return res
}

// iterative sol is better than recursive becoz S.C->> is O(1) while in recursive S.C-> O(n)
func fib2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	prev1, prev2 := 0, 1
	res := 0
	for i := 2; i <= n; i++ {
		res = prev1 + prev2
		prev1 = prev2
		prev2 = res
	}
	return res
}
