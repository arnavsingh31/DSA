package main

import (
	"fmt"
	"math"
)

/*
LC #835(Linear Transformation)
TC--->O(N^4), square matrix--> m = n
SC--->O(N^2)

Linear Transformation:
find cells with ones. to move a 1 from A to a 1 from B, it'll take some transformation  (dx, dy).
for another 1 cell in A to also overlap with a 1 cell in B will take the SAME (dx,dy).
make a map of possible (dx,dy) and count the ones that have the same (dx,dy)

find [[i,j]...] of 1s for A, B
find all combos of each filtered array
calc the dx,dy and set to a map.  find the mapping that has the highest count and return

T: filtering takes O(N^2). combos takes a double for loop of N^2 * N^2 in the worst case.
so N^2 + N^2 * N^2 ---> O(N^4)
S: O(N^2) for arrays space in the worst case of A and B filled with 1s
*/
func largestOverlap(img1, img2 [][]int) int {
	rows := len(img1)
	cols := len(img1)
	arr1, arr2 := make([]Pos, 0), make([]Pos, 0)
	transformationVectorMap := make(map[string]int, 0)
	overlaps := 0

	// extract all 1's position from both image
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if img1[i][j] == 1 {
				arr1 = append(arr1, Pos{i, j})
			}
			if img2[i][j] == 1 {
				arr2 = append(arr2, Pos{i, j})
			}
		}
	}

	// apply linear transformation for all possible pairs in both arr1 and arr2.
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			x1, y1 := arr1[i].row, arr1[i].col
			x2, y2 := arr2[j].row, arr2[j].col

			dx := x2 - x1
			dy := y2 - y1

			key := getKey(dx, dy)
			transformationVectorMap[key] += 1
			overlaps = int(math.Max(float64(overlaps), float64(transformationVectorMap[key])))
		}
	}
	return overlaps
}

func getKey(dx, dy int) string {
	return fmt.Sprintf("%d#%d", dx, dy)
}
