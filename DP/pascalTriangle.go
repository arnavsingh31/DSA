package main

/*
	LC #118
	TC--->O(n*n)
	SC--->O(n*n)
*/
func generate(numRows int) [][]int {
	ans := [][]int{{1}}

	for i := 0; i < numRows-1; i++ {
		tempArr := []int{0}
		tempArr = append(append(tempArr, ans[len(ans)-1]...), []int{0}...)
		row := []int{}
		// length of new row will be len(prevRow)+1...here len(ans)+1 will also work.
		for j := 0; j < len(ans[len(ans)-1])+1; j++ {
			row = append(row, tempArr[j]+tempArr[j+1])
		}
		ans = append(ans, row)
	}

	return ans
}
