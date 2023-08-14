package main

/*
	LC #88
	TC--->O(m+n)
	SC--->O(1)
*/
func merge(arr1, arr2 []int, m, n int) {
	index := m + n - 1
	m--
	n--

	for m >= 0 && n >= 0 {
		if arr1[m] >= arr2[n] {
			arr1[index] = arr1[m]
			index--
			m--
		} else {
			arr1[index] = arr2[n]
			index--
			n--
		}
	}

	for n >= 0 {
		arr1[index] = arr2[n]
		index--
		n--
	}
}
