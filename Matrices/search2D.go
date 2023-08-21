package main

/*
	Using Binary Search
	LC #74
	TC--->O(log(m*n)), since binary seacrh has logN TC and here N = m*n.
	SC--->O(1)
*/
func search2D(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	start := 0
	end := rows*cols - 1

	mid := (start + end) / 2

	for start <= end {
		element := matrix[mid/cols][mid%cols]

		if element == target {
			return true
		}

		if element < target {
			start++
		}

		if element > target {
			end--
		}

		mid = (start + end) / 2
	}

	return false
}
