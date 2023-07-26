package main

import (
	"math"
	"sort"
)

/*
LC #581
TC--> O(n^2)
SC--> O(1)
*/
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findUnsortedSubarray1(arr []int) int {
	left := len(arr)
	right := 0

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				left = min(left, i)
				right = max(right, j)
			}
		}
	}

	if (right - left) < 0 {
		return 0
	}
	return right - left + 1
}

/*
TC--> O(nlogn + n) ~ O(nlogn)
SC--> O(n) since we are making copy of original array
*/
func findUnsortedSubarray2(arr []int) int {
	arrSorted := make([]int, len(arr))
	copy(arrSorted, arr)
	sort.Ints(arrSorted)

	left := len(arr)
	right := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != arrSorted[i] {
			left = min(left, i)
			right = max(right, i)
		}
	}

	if right-left < 0 {
		return 0
	}
	return right - left + 1
}

/*
	TC---> O(n)
	SC---> O(n) using stack
*/

func findUnsortedSubarray3(arr []int) int {
	stack := []int{}
	left := len(arr)
	right := 0
	for i, num := range arr {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > num {
			left = min(left, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	for len(stack) > 0 {
		stack = stack[:len(stack)-1]
	}

	for i := len(arr) - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] < arr[i] {
			right = max(right, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	if right-left < 0 {
		return 0
	}
	return right - left + 1
}

/*
TC---> O(n)
SC---> O(1)
Watch this: https://leetcode.com/problems/shortest-unsorted-continuous-subarray/editorial/comments/562497
for reference.
*/
func findUnsortedSubarray4(arr []int) int {
	var left, right int
	min := math.MaxInt
	max := math.MinInt

	// first left where current element is bigger than next element
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			left = i
			break
		}
	}

	// first index from back where current element is smaller than next element
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] < arr[i-1] {
			right = i
			break
		}
	}

	// finding max and min element with the range [left, right]
	for i := left; i <= right; i++ {
		if max < arr[i] {
			max = arr[i]
		}
		if min > arr[i] {
			min = arr[i]
		}
	}

	// checking whether all elements before left are smaller than min or not, if not then we update left
	//	and break
	for i := 0; i < left; i++ {
		if min < arr[i] {
			left = i
			break
		}
	}

	// checking whether all element after right are greater than max or not. if not then we update right and
	// break (iterating from behind).
	for i := len(arr) - 1; i > right; i-- {
		if max > arr[i] {
			right = i
			break
		}
	}

	if right-left <= 0 {
		return 0
	}
	return right - left + 1
}

/*
	[1,3,4,7,6,2,12,10,9,11]
	left = 3--->1 right = 8--->9
	max = 12  min = 2
*/
