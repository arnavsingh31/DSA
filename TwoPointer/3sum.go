package main

import "sort"

// fix first number then use two pointer approach to get next 2 numbers of a triplet.
func threeSum(arr []int) [][]int {
	sort.Ints(arr)
	res := [][]int{}
	for i := 0; i < len(arr)-2; i++ {

		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		start := i + 1
		end := len(arr) - 1
		for start < end {
			sum := arr[i] + arr[start] + arr[end]
			if sum == 0 {
				res = append(res, []int{arr[i], arr[start], arr[end]})
				start++

				// skipping duplicate values
				for arr[start] == arr[start-1] && start < end {
					start++
				}

			} else if sum < 0 {
				start++
			} else if sum > 0 {
				end--
			}
		}

	}
	return res
}
