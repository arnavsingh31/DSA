package hashmaps

/*
	LC #1488
	TC--->O(n^2) {if we use binary search to find whether there is a dryDay available, then TC will be O(nlogn)}
	SC--->O(n) {map + dryDay_array}
*/
func AvoidFlood(rains []int) []int {
	lakeIndexMap := make(map[int]int)
	dryDay := make([]int, 0)

	for i := 0; i < len(rains); i++ {
		if rains[i] == 0 {
			dryDay = append(dryDay, i)
		} else {
			if lastIndex, exist := lakeIndexMap[rains[i]]; exist {
				usedZeroIdx := -1
				for j, val := range dryDay {
					if val > lastIndex {
						rains[val] = rains[i]
						lakeIndexMap[rains[i]] = i
						rains[i] = -1
						usedZeroIdx = j
						break
					}
				}

				if usedZeroIdx == -1 {
					return []int{}
				}

				dryDay = remove(dryDay, usedZeroIdx)

			} else {
				lakeIndexMap[rains[i]] = i
				rains[i] = -1
			}
		}
	}

	if len(dryDay) > 0 {
		for _, val := range dryDay {
			rains[val] = 1
		}
	}
	return rains
}

func remove(arr []int, index int) []int {
	a1 := arr[:index]
	a2 := arr[index+1:]
	return append(a1, a2...)
}
