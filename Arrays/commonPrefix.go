package main

import (
	"log"
)

// my solution
func longestCommonPrefix(strs []string) string {
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
	log.Printf("common prefix is :%s", res)
	return res
}

// better solution. simple logic T.C --> O(n.m), where n is the length of strs and m is the length of common prefix.
// S.C--> O(1)
func longestCommonPrefix2(strs []string) string {
	res := ""
	for i := 0; i < len(strs[0]); i++ {
		charByte := strs[0][i]
		flag := true
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) < i+1 || strs[j][i] != charByte {
				flag = false
				break
			}
		}
		if !flag {
			return res
		}
		res += string(charByte)

	}
	return res
}
