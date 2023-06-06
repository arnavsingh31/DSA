package main

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS := make(map[string]int)
	mapT := make(map[string]int)

	for i := 0; i < len(s); i++ {
		if val, exist := mapS[string(s[i])]; exist {
			mapS[string(s[i])] = val + 1
		} else {
			mapS[string(s[i])] = 1
		}

		if val, exist := mapT[string(t[i])]; exist {
			mapT[string(t[i])] = val + 1
		} else {
			mapT[string(t[i])] = 1
		}
	}

	for key, value := range mapS {
		if value != mapT[key] {
			return false
		}
	}
	return true
}

// APPROACH #2
func isAnagram2(s, t string) bool {
	charsToCountMap := make(map[rune]int)

	if len(s) != len(t) {
		return false
	}

	for _, char := range s {
		charsToCountMap[char] += 1
	}

	for _, char := range t {
		if charCount, exist := charsToCountMap[char]; !exist || charCount == 0 {
			return false
		}
		charsToCountMap[char]--
	}
	return true
}
