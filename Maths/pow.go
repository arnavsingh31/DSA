package maths

// T.C--> O(n) S.C--> O(1)
func Pow(x, n int) int {
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
	Here were are calculating x^n.
	if power i.e n is even, then we can write x^n = x^(n/2) * x^(n/2)
	if power i.e n is odd, then we can write x^n =  x*(x^(n/2) * x^(n/2))
*/
func Pow2(x, n int) int {
	if n == 0 || x == 1 {
		return 1
	}
	if n == 1 {
		return x
	}

	res := Pow2(x, n/2)
	res = res * res

	if n&1 == 0 {
		return res
	} else {
		return res * x
	}
}

// Iterative Power (Binary Exponentiation)
/*
	3^10 = 3^8 * 3^2
	10 in binary is 1010. We traverse starting from LSB to MSB and for each bit we traverse we consider it
	as a corresponding power of x (3). i.e.
	MSB -----> LSB
	1     0    1    0
	\|/  \|/  \|/  \|/
	3^8  3^4  3^2  3^1
	we consider value corresponding to 1 and multiply them correspondingly. here we get:
	answer = 3^8 * 3^2
	----------------------------------------------------------------------------------------
	3^19 = 3^16 * 3^2 * 3^1
	Every number can be written as sum of powers of 2(set bits in binary representation)
	We can traverse through all bits of a number in O(logn) time
*/
func Pow3(x, n int) int {
	res := 1
	for n > 0 {
		if n%2 != 0 {
			res = res * x
		}
		n = n / 2
		x = x * x
	}

	return res
}

// it calculates x raised to n where n can be positive or negative integer and x is float64 type which
// is also either positive or negative. (Leetcode #50)
func Pow4(x float64, n int) float64 {
	res := 1.0
	isNegative := false

	if n < 0 {
		n = -n
		isNegative = true
	}

	for n > 0 {
		if n%2 != 0 {
			res = res * x
		}
		n = n / 2
		x = x * x
	}

	if isNegative {
		return 1 / res
	}
	return res
}
