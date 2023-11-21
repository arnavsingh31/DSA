package hashmaps

import (
	"sort"
)

/*
LC #451
TC--->O(nlogn)
SC--->O(n)
Brute force {Accepted}
*/
func SortCharcterByFrequency(s string) string {
	freqMap := make(map[byte]int)
	arr := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		freqMap[s[i]] += 1
		arr = append(arr, s[i])
	}

	sort.Slice(arr, func(i, j int) bool {
		if freqMap[arr[i]] == freqMap[arr[j]] {
			return arr[i] > arr[j]
		}

		return freqMap[arr[i]] > freqMap[arr[j]]
	})

	return string(arr)
}

/*
TC--->O()
SC--->O()
*/
func SortCharcterByFrequency2(s string) string {

}
