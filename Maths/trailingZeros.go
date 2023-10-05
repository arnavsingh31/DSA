package maths

import util "github.com/arnavsingh31/DSA/Util"

func TrailingZeros(n int) int {
	count_two := 0
	count_five := 0

	for i := 2; i <= n; i++ {
		temp := i

		for temp > 1 {
			if temp%2 == 0 {
				count_two++
				temp = temp / 2

			} else if temp%5 == 0 {
				count_five++
				temp = temp / 5

			} else {
				break
			}
		}
	}

	return util.Min(count_two, count_five)
}

/*
	n! = n x n-1 x n-2 ... x 1
	That means trailing zero is contributed by the number of (5, 2) pair in prime factor decomposition of n!

	With the value growing on n, n can be re-written as n = 5^k x 2^q x (other prime factors ...).

	The key-point is the k, the exponent of 5, because k ≦ q always for every possible n! with trailing zero.

	Observation:

	5! = 5 x 4 x 3 x 2 x 1 = 5^1 x 2^1 x ( 4 x 3 ) = 5^1 x 2^3 x other prime factors= 120

	10! = 10 x 9 x ... 5 x ...x 2 x 1 = 5^2 x 2^8 x other prime factors = 3628800

	15! = 15 x 14 x ... x 2 x1 = 5^3 x 2^11 x other prime factors = 1307674368000

	It's clear that only a 5 and a 2 can make a 0 at the end.
	But the 2 is always more(no less) than 5.
	like between 1, 5. The 2 get 3 times(2 is 1 time, and 4 is 2 times) but 5 just 1 time.(You can prove it by yourself)
	And 5 can give me 5 1 time, 10: 1 time ...
	25 is 5*5, so it's 2 time. 50: 2 time ...
	125: 3 time...

	So I just count how many 5 it contains, plus how many 25, plus how many 125, plus....
*/

func TrailingZeroes2(n int) int {
	zeros := 0

	for i := 5; i <= n; i = i * 5 {
		zeros += n / i
	}

	return zeros
}
