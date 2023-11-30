package binarysearch

/*
	Given a m x n binary matrix mat, find the 0-indexed position of the row that contains the maximum count of ones, and the number of ones in that row. **Each row of matrix is sorted

	In case there are multiple rows that have the maximum count of ones, return the row with the smallest row number.

	TC--->O(m*n)
	SC--->O(1)
	Brute-force
*/
func MaxNumberOfOnes(mat [][]int) int {
	rows := len(mat)
	cols := len(mat[0])
	maxOnes := -1
	pos := -1

	for i := 0; i < rows; i++ {
		counter := 0

		for j := 0; j < cols; j++ {
			counter += mat[i][j]
		}

		if counter > maxOnes {
			maxOnes = counter
			pos = i
		}
	}

	return pos
}

/*
	TC--->O(n*logm)
	SC--->O(1)
	We can optimise the above code by finding the firt occurence of 1 in each row using binary search algo.
*/
func MaxNumberOfOnes2(mat [][]int) int {
	rows := len(mat)
	cols := len(mat[0])
	ans := -1
	max := -1

	for i := 0; i < rows; i++ {
		pos := findStartingOne(mat[i])
		if pos != -1 {
			if cols-pos > max {
				max = cols - pos
				ans = i
			}
		}
	}
	return ans
}

func findStartingOne(row []int) int {
	left, right := 0, len(row)-1
	ans := -1

	for left <= right {
		mid := (left + right) / 2

		if row[mid] == 1 {
			if ans == -1 || ans > mid {
				ans = mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}
