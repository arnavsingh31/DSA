package main

import "sort"

/*
	LC #1337
	TC--->O(m*n + nlogn)
	SC--->O(n), n-> number of rows
*/
func kWeakestRows(matrix [][]int, k int) []int {
	rows := len(matrix)
	cols := len(matrix[0])
	soldiers := make([][]int, rows)
	ans := make([]int, k)

	for i := 0; i < rows; i++ {
		soldierCount := 0
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 1 {
				soldierCount++
			}
		}
		soldiers[i] = []int{i, soldierCount}
	}

	sort.Slice(soldiers, func(i int, j int) bool {
		// first we check if soldier count is same or not. If same, then sort by row-index else sort by soldier count.
		return (soldiers[i][1] == soldiers[j][1] && soldiers[i][0] < soldiers[j][0]) || (soldiers[i][1] < soldiers[j][1])
	})

	for i := 0; i < k; i++ {
		ans[i] = soldiers[i][0]
	}

	return ans
}
