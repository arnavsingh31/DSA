package binarysearch

/*
	Problem Statement: Given two numbers N and M, find the Nth root of M. The nth root of a number M is defined as a number X when raised to the power N equals M. If the ‘nth root is not an integer, return -1.
*/
func NthRoot(n, m int) int {
	left, right := 1, m

	for left <= right {
		mid := (left + right) / 2

		nthVal := getNthVal(mid, n, m)

		if nthVal == 0 {
			return mid
		} else if nthVal == 1 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func getNthVal(mid, n, m int) int {
	ans := 1
	for i := 1; i <= n; i++ {
		ans *= mid
		if ans > m {
			return 2
		}
	}

	if ans == m {
		return 0
	}

	return 1
}
