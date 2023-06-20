package main

func firstOccurence(haystack, needle string) int {
	firstIndex := -1

	for i := 0; i < len(haystack); i++ {
		firstCharByte := needle[0]
		if firstCharByte == haystack[i] {
			tempPtr := i
			flag := false
			for j := 0; j < len(needle); j++ {
				if tempPtr < len(haystack) && haystack[tempPtr] == needle[j] {
					temp := haystack[i : tempPtr+1]
					if temp == needle {
						flag = true
					}
					tempPtr++
				} else {
					break
				}
			}

			if flag {
				firstIndex = i
				break
			}
		}
	}
	return firstIndex
}
