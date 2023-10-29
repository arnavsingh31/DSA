package slidingwindow

import util "github.com/arnavsingh31/DSA/Util"

/*
	LC #424
	TC--->O(n)
	SC--->O(n)
*/
func CharacterReplacement(s string, k int) int {
	charCountMap := make(map[byte]int)
	left := 0
	res := 0

	for right := 0; right < len(s); right++ {
		charCountMap[s[right]] += 1

		for right-left+1-getMaxCharCount(charCountMap) > k {
			charCountMap[s[left]] -= 1
			left++
		}

		res = util.Max(res, right-left+1)
	}

	return res
}

func getMaxCharCount(windowMap map[byte]int) int {
	max := 0
	for _, val := range windowMap {
		max = util.Max(max, val)
	}
	return max
}
