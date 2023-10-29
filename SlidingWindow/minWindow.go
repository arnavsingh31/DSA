package slidingwindow

// T.C----->O(m+n) m = length of s and n = length of t. S.C ---->O(m+n).
// please refere the link for better understanding:
// https://leetcode.com/problems/minimum-window-substring/editorial/?envType=study-plan-v2&envId=top-interview-150
func MinWindow(s string, t string) string {

	sLen := len(s)
	tLen := len(t)

	if sLen == 0 || tLen == 0 {
		return ""
	}

	left := 0
	posArr := []int{-1, 0, 0}
	mapT := make(map[byte]int)
	windowMap := make(map[byte]int)
	charFound := 0

	for i := 0; i < tLen; i++ {
		mapT[t[i]] += 1
	}

	for right := 0; right < sLen; right++ {

		if val, exist := mapT[s[right]]; exist {
			windowMap[s[right]] += 1
			if windowMap[s[right]] == val {
				charFound++
			}
		}

		for charFound == len(mapT) && left <= right {

			if posArr[0] == -1 || (right-left+1) < posArr[0] {
				posArr[0] = right - left + 1
				posArr[1] = left
				posArr[2] = right
			}

			if _, exist := windowMap[s[left]]; exist {
				windowMap[s[left]] -= 1
			}

			if count, exist := mapT[s[left]]; exist && windowMap[s[left]] < count {
				charFound--
			}

			left++
		}

	}

	if posArr[0] == -1 {
		return ""
	}

	return s[posArr[1] : posArr[2]+1]
}

// Approach 2
type Pair struct {
	index    int
	charByte byte
}

func MinWindow2(s, t string) string {
	sLen := len(s)
	tLen := len(t)

	if sLen == 0 || tLen == 0 {
		return ""
	}

	mapT := make(map[byte]int)

	for i := 0; i < tLen; i++ {
		mapT[t[i]] += 1
	}

	filteredS := []Pair{}

	for i := 0; i < sLen; i++ {
		if _, exist := mapT[s[i]]; exist {
			indexCharPair := Pair{index: i, charByte: s[i]}
			filteredS = append(filteredS, indexCharPair)
		}
	}

	left := 0
	charFound := 0
	windowMap := make(map[byte]int)
	posArr := []int{-1, 0, 0}

	for right := 0; right < len(filteredS); right++ {
		if count, exist := mapT[filteredS[right].charByte]; exist {
			windowMap[filteredS[right].charByte] += 1

			if windowMap[filteredS[right].charByte] == count {
				charFound++
			}
		}

		for charFound == len(mapT) && left <= right {
			rightIndex := filteredS[right].index
			leftIndex := filteredS[left].index
			if posArr[0] == -1 || rightIndex-leftIndex+1 < posArr[0] {
				posArr[0] = rightIndex - leftIndex + 1
				posArr[1] = leftIndex
				posArr[2] = rightIndex
			}

			if _, exist := windowMap[filteredS[left].charByte]; exist {
				windowMap[filteredS[left].charByte] -= 1
			}

			if count, exist := mapT[filteredS[left].charByte]; exist && windowMap[filteredS[left].charByte] < count {
				charFound--
			}
			left++
		}
	}

	if posArr[0] == -1 {
		return ""
	}
	return s[posArr[1] : posArr[2]+1]
}
