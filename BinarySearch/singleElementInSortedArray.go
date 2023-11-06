package binarysearch

/*
	LC #540
	TC--->O(n)
	SC--->O(1)
	Brute-force
*/
func SingleElementInSortedArray(arr []int) int {
	n := len(arr)
	ans := -1

	if n == 1 {
		return arr[0]
	}

	if arr[0] != arr[1] {
		return arr[0]
	}

	if arr[n-1] != arr[n-2] {
		return arr[n-1]
	}

	for i := 1; i < n-1; i++ {
		if arr[i] != arr[i-1] && arr[i] != arr[i+1] {
			ans = arr[i]
			break
		}
	}

	return ans
}

/*
	TC--->O(logn)
	SC--->O(1)
	Binary search
*/
func SingleElementInSortedArray2(arr []int) int {
	n := len(arr)
	ans := -1
	// handling all edge cases before, so that we don't have to handle them inside loop.
	if n == 1 {
		return arr[0]
	}

	if arr[0] != arr[1] {
		return arr[0]
	}

	if arr[n-1] != arr[n-2] {
		return arr[n-1]
	}

	left, right := 1, n-1

	for left <= right {
		mid := (left + right) / 2

		// checking if mid element is not equal to mid+1 and mid-1 , then this is our single element.
		if arr[mid] != arr[mid-1] && arr[mid] != arr[mid+1] {
			ans = arr[mid]
			break
		}

		// since every number except one exist in pairs. Pair have their indexes either (even, odd) or (odd, even)
		// if mid is even index then we check if pair are occuring in (even, odd) or (odd, even) pattern.
		// key thing to note is that pairing start from even index.
		// first we check if our mid is even or odd index.
		// 1. if pairs exist in (even, odd) index pair then all elements till mid index, are paired and hence
		// we can reduce search space by left = mid + 1, but if pairs occuring in (odd, even) indices then
		// all elements from mid to end are paired, so we reduce search space by right = mid - 1.
		// similar process we follow when our mid is odd index.
		if mid&1 == 0 {
			if arr[mid] == arr[mid+1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if arr[mid] == arr[mid+1] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return ans
}
