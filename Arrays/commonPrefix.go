package arrays

import (
	"strings"
)

// my solution
func LongestCommonPrefix(strs []string) string {
	IndexToCharMap := make(map[int]byte)
	res := ""
	first := strs[0]

	for i := 0; i < len(first); i++ {
		IndexToCharMap[i] = first[i]
	}

	for i := 1; i < len(strs); i++ {
		word := strs[i]
		charByte := IndexToCharMap[0]
		if len(word) == 0 || charByte != word[0] {
			return ""
		}
		matchStr := ""
		for i := 0; i < len(word); i++ {
			if word[i] == IndexToCharMap[i] {
				matchStr += string(word[i])
			} else {
				break
			}
		}
		if len(res) == 0 || len(matchStr) < len(res) {

			res = matchStr
		}
	}

	return res
}

/*
Vertical Scanning.
TC--->O(S) where S is the sum of all characters in all strings.
In the worst case there will be n equal strings with length m and the algorithm performs S = m*n character comparisons.
SC--->O(1)
*/
func LongestCommonPrefix2(strs []string) string {
	prefix := strs[0]
	for i := 0; i < len(prefix); i++ {
		charByte := prefix[i]

		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != charByte {
				return prefix[0:i]
			}
		}
	}
	return prefix
}

/*
Horizontal Scanning
TC--->O(S) where S is the sum of all characters in all strings.
In the worst case all nnn strings are the same. The algorithm compares the string S1 with the other strings [S2..Sn]. There are S character comparisons, where S is the sum of all characters in the input array.
SC--->O(1)
*/
func LongestCommonPrefix3(strs []string) string {
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
		}

		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}
