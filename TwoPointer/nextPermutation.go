package main

import (
	"sort"
)

// Read:- https://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order
/*
	The following algorithm generates the next permutation lexicographically after a given permutation.
	It changes the given permutation in-place.

	1. Find the largest index k such that a[k] < a[k + 1]. If no such index exists, the permutation is the last permutation.
	2. Find the largest index l greater than k such that a[k] < a[l].
	3. Swap the value of a[k] with that of a[l].
	4. Reverse the sequence from a[k + 1] up to and including the final element a[n].
*/
func nextPermutation(arr []int) {
	k := -1
	l := -1
	// find k index
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			if k == -1 {
				k = i
			} else if k < i {
				k = i
			}
		}
	}
	// if k exist, find index l
	if k != -1 {
		// find index l
		for j := 0; j < len(arr); j++ {
			if arr[j] > arr[k] {
				if l == -1 {
					l = j
				} else if l < j {
					l = j
				}
			}
		}
		// swap value at k and l
		arr[k], arr[l] = arr[l], arr[k]

		// reverse array from k+1 to end of array.
		start := k + 1
		end := len(arr) - 1
		for start <= end {
			arr[start], arr[end] = arr[end], arr[start]
			start++
			end--
		}

	} else {
		// k doesn't exist i.e given permutation is the last permutation. So we simply sort array.
		sort.Ints(arr)
	}
}
