package bitwiseoperation

// using Brian Kerningam's algorithm.
// T.C--> O(n*k) where k is the no of set bits in binary representation of i
func CountBits(n int) []int {
	res := []int{}
	for i := 0; i <= n; i++ {
		count := 0
		x := i
		for x > 0 {
			x = x & (x - 1)
			count++
		}
		res = append(res, count)

	}
	return res
}

// T.C----> O(n)
/*
	Consider some numberx, e.g. x = 1382.
	It's binary form is 10101100111.
	Shifting the number x by 1 will remove only the lowest bit:
	x      = 10101100111 = 1383
	x >> 1 = 1010110011  = 691
	How many 1s are in x? Well, it's the number of 1s in ones[x>>1] PLUS an extra 1 IF the number is odd.
	To add the missing 1, we can use the binary and operator x&1.
*/
func CountBits2(n int) []int {
	res := make([]int, n+1)

	for i := 0; i <= n; i++ {
		res[i] = res[i>>1] + i&1
	}
	return res
}
