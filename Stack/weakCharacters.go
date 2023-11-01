package stack

import (
	"sort"
)

/*
LC #1996
TC---> O(nlogn)
SC---> O(1)
*/
func NumberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] != properties[j][0] {
			// sorting on basis of decreasing attack attribute if they are not equal
			return properties[i][0] > properties[j][0]
		} else {
			// if attack attribute are equal then sort on basis of increasing defence. why? increasing defence

			return properties[i][1] > properties[j][1]
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

/*
	why?
	sorting in increasing order of defence becoz if we sort in dec order of defence the  there arise a case
	where the defence of i character is greater than i+1 character but both have same attack attribute,
	due to which we end up counting it as a weak character even though it isn't.

	for example :
	[[10 1] [7 10] [7 2] [6 9] [5 9] [5 1] [4 1] [1 10]]
	at i=1; our max defence becomes equal to 10
	at i=2; maxDefence > character[1] (i.e 2) due to which it count character i=2 as weak but it isn't weak
	as both have same attack attribute(=7)
*/

// func main() {
// 	numberOfWeakCharacters([][]int{{10, 1}, {5, 1}, {7, 10}, {4, 1}, {5, 9}, {6, 9}, {7, 2}, {1, 10}})
// }
