package main

// T.C--> O(n) S.C--> O(1)
func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	res := x
	for i := 2; i <= n; i++ {
		res = res * x
	}

	return res
}

/*
	T.C--> O(logn) S.C--> O(logn)
*/
func pow2(x, n int) int {
	if n == 0 || x == 1 {
		return 1
	}
	if n == 1 {
		return x
	}

	res := pow2(x, n/2)
	res = res * res

	if n%2 == 0 {
		return res
	} else {
		return res * x
	}
	// if n%2 == 0 {
	// 	res = pow2(x, n/2) * pow2(x, n/2)
	// } else {
	// 	res = pow2(x, n-1) * x
	// }
}
