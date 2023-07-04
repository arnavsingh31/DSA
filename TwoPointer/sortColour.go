package main

// LeetCode #75

func sortColour(arr []int) {
	colourCountMap := make(map[int]int)

	for _, color := range arr {
		colourCountMap[color] += 1
	}

	for i := 0; i < len(arr); i++ {
		if colourCountMap[0] > 0 {
			arr[i] = 0
			colourCountMap[0] -= 1
		} else if colourCountMap[1] > 0 {
			arr[i] = 1
			colourCountMap[1] -= 1
		} else {
			arr[i] = 2
			colourCountMap[2] -= 1
		}
	}
}

// 3 pointer approach
func sortColour2(arr []int) {
	left, i, right := 0, 0, len(arr)-1

	for i <= right {
		if arr[i] == 0 {
			arr[i], arr[left] = arr[left], arr[i]
			left++
			i++
		} else if arr[i] == 2 {
			arr[i], arr[right] = arr[right], arr[i]
			right--
		} else {
			i++
		}
	}
}
