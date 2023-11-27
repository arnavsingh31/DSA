package binarysearch

/*
	LC #4
	TC--->O(n+m)
	SC--->O(n+m)
	Brute force {merge 2 sorted array and then find median}
*/
func MedianOf2SortedArray(nums1, nums2 []int) float64 {
	arr := make([]int, 0)
	median := 0.0

	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			arr = append(arr, nums1[i])
			i++
		} else if nums2[j] < nums1[i] {
			arr = append(arr, nums2[j])
			j++
		} else {
			arr = append(arr, nums1[i], nums2[j])
			i++
			j++
		}
	}

	for i < len(nums1) {
		arr = append(arr, nums1[i])
		i++
	}

	for j < len(nums2) {
		arr = append(arr, nums2[j])
		j++
	}

	k := len(nums1) + len(nums2)

	if k&1 == 0 {
		median = float64(arr[k/2-1]+arr[k/2]) / 2
	} else {
		median = float64(arr[k/2])
	}

	return median
}

/*
	TC--->O(m+n)
	SC--->O(1)
	same as above but without using extra space
*/
func MedianOf2SortedArray2(nums1, nums2 []int) float64 {
	counter := 0
	n, m := len(nums1), len(nums2)
	i, j := 0, 0
	mid, beforeMid := -1, -1
	median := 0.0

	k := (n + m) / 2
	for i < n && j < m {
		if nums1[i] < nums2[j] {
			if counter == k {
				mid = nums1[i]
			}

			if counter == k-1 {
				beforeMid = nums1[i]
			}
			counter++
			i++
		} else {
			if counter == k {
				mid = nums2[j]
			}

			if counter == k-1 {
				beforeMid = nums2[j]
			}
			counter++
			j++
		}
	}

	for i < n {
		if counter == k {
			mid = nums1[i]
		}

		if counter == k-1 {
			beforeMid = nums1[i]
		}
		counter++
		i++
	}

	for j < m {
		if counter == k {
			mid = nums2[j]
		}

		if counter == k-1 {
			beforeMid = nums2[j]
		}
		counter++
		j++
	}

	if (n+m)&1 == 0 {
		median = float64(mid+beforeMid) / 2
	} else {
		median = float64(mid)
	}
	return median
}
