package main

/*
	Idea is to find a number which whose kth bit is set and then we simply do bitwise "and" operation with n.
	if result is 0 then we can say kth bit of n is not set or else it is set.
	T.C--> theta(k)
*/
func iskthBitSet(n, k int) bool {
	x := 1
	for i := 0; i < k-1; i++ {
		x = x * 2
	}

	return x&n != 0
}

/*
	Here instead of finding the number with help of loop we can simply left-shift 1 by k-1, which will give
	us a number whose kth bit is set. and then we do "and" operation with it as above.
	T.C--> theta(1)
*/
func iskthBitSet2(n, k int) bool {
	x := 1 << (k - 1)

	return x&n != 0
}

// same as above but using right-shift operator
func iskthBitSet3(n, k int) bool {
	n = n >> (k - 1)

	return n&1 != 0
}
