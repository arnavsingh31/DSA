package main

func maxArea(arr []int) int {
	start := 0
	end := len(arr) - 1
	area := 0

	for start < end {
		if arr[start] < arr[end] {
			area = getMaxArea(area, (end-start)*arr[start])
			start++
		} else if arr[start] > arr[end] {
			area = getMaxArea(area, (end-start)*arr[end])
			end--
		} else {
			area = getMaxArea(area, (end-start)*arr[start])
			start++
		}
	}
	return area
}

func getMaxArea(a, b int) int {
	if a > b {
		return a
	}
	return b
}
