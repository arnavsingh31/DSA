package twopointer

import util "github.com/arnavsingh31/DSA/Util"

func MaxArea(arr []int) int {
	start := 0
	end := len(arr) - 1
	area := 0

	for start < end {
		if arr[start] < arr[end] {
			area = util.Max(area, (end-start)*arr[start])
			start++
		} else if arr[start] > arr[end] {
			area = util.Max(area, (end-start)*arr[end])
			end--
		} else {
			area = util.Max(area, (end-start)*arr[start])
			start++
		}
	}
	return area
}
