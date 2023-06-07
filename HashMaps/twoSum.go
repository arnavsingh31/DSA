package main

// 2-PASS
func twoSum(arr []int, target int) []int {
	hashMap := make(map[int]int)

	for index, value := range arr {
		hashMap[value] = index
	}

	for index, value := range arr {
		compliment := target - value

		if val, exist := hashMap[compliment]; exist && val != index {
			return []int{index, val}
		}
	}
	return nil
}

// 1-PASS
func twoSum2(arr []int, target int) []int {
	hashMap := make(map[int]int)

	for index, value := range arr {
		compliment := target - value
		if val, exist := hashMap[compliment]; exist && val != index {
			return []int{val, index}
		}
		hashMap[value] = index
	}

	return nil
}
