package main

func containsNearbyDuplicate(arr []int, k int) bool {
	hashMap := make(map[int]int)

	for curr_index, num := range arr {
		if mapped_index, exist := hashMap[num]; exist {
			indexDiff := getAbsDiff(curr_index, mapped_index)

			if indexDiff <= k {
				return true
			}
		}
		hashMap[num] = curr_index
	}

	return false
}

func getAbsDiff(i, j int) (diff int) {
	diff = i - j

	if diff < 0 {
		diff *= -1
		return
	}
	return
}
