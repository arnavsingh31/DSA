package main

import (
	"sort"
)

/*
LC #1996
TC---> O(nlogn)
SC---> O(1)
*/
func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] != properties[j][0] {
			return properties[i][0] > properties[j][0]
		} else {
			return properties[i][1] < properties[j][1]
		}
	})

	maxDefence := -1
	count := 0
	for _, character := range properties {
		if character[1] < maxDefence {
			count++
		} else {
			maxDefence = character[1]
		}

	}

	return count
}

// func main() {
// 	numberOfWeakCharacters([][]int{{10, 1}, {5, 1}, {7, 10}, {4, 1}, {5, 9}, {6, 9}, {7, 2}, {1, 10}})
// }
