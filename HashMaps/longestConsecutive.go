package hashmaps

// brute force T.C ---> O(n^3) S.C--->O(1)
func LongestConsecutive(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	maxLen := 0

	for _, num := range arr {
		currentNum := num
		currentLen := 1

		for arrContains(arr, currentNum+1) {
			currentNum += 1
			currentLen += 1
		}

		if currentLen > maxLen {
			maxLen = currentLen
		}

	}
	return maxLen

}

func arrContains(arr []int, val int) bool {
	for _, num := range arr {
		if num == val {
			return true
		}
	}
	return false
}

// [100, 4, 200, 1, 3, 2]

// Logic is same as brute force sol. but we reduce the lookups by using a hashmap.
func LongestConsecutive2(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	hashMap := make(map[int]int)
	maxLen := 0

	for index, num := range arr {
		hashMap[num] = index
	}

	/*
		Now map looks like this->

		Key -> Value
		100 -> 0
		4 -> 1
		1 -> 2
		3 -> 3
		2 -> 4
	*/

	for _, num := range arr {
		if !mapContains(hashMap, num-1) {
			currentNum := num
			currLen := 1
			for mapContains(hashMap, currentNum+1) {
				currentNum += 1
				currLen += 1
			}
			if currLen > maxLen {
				maxLen = currLen
			}
		}
	}
	return maxLen
}

func mapContains(arrMap map[int]int, val int) bool {
	if _, exist := arrMap[val]; exist {
		return true
	}
	return false
}
