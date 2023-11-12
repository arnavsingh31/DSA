package binarysearch

/*
	Find max integer which on squaring is <= n
	TC--->O(logn)
	SC--->O(1)
	Here we apply Binary Search on answers which is b/w (1->n).
*/
func SqrtOfInt(n int) int {
	left, right := 1, n

	for left < right {
		mid := (left + right) / 2

		if mid*mid <= n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}
