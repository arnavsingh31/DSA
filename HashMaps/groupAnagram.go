package hashmaps

import (
	"sort"
	"strings"
)

/*
In this approach we first sort the string while looping through the array and after sorting
we use this sorted string as the key for our hashmap and if the key already exist then we simply
add the corresponding string. if it doesn't exist then we initialise empty slice and then add string to it.

#Algorithm:------>

Maintain a map ans : {String -> List} where each key K is a sorted string, and each value is the
list of strings from the initial input that when sorted, are equal to K.

#Complexity Analysis----->

Time Complexity: O(NKlog⁡K), where N is the length of strs, and K is the maximum length of a string in strs.
The outer loop has complexity O(N) as we iterate through each string. Then, we sort each string in O(Klog⁡K) time.

Space Complexity: O(NK), the total information content stored in hashMap.
*/
func GroupAnagram(arr []string) [][]string {
	hashMap := make(map[string][]string)

	for _, str := range arr {
		charArray := strings.Split(str, "")
		sort.SliceStable(charArray, func(i, j int) bool { return charArray[i] < charArray[j] })
		sortedStr := strings.Join(charArray, "")
		if _, exist := hashMap[sortedStr]; !exist {
			hashMap[sortedStr] = []string{}
		}
		hashMap[sortedStr] = append(hashMap[sortedStr], str)
	}

	groupAnagram := [][]string{}
	for _, values := range hashMap {
		groupAnagram = append(groupAnagram, values)
	}
	return groupAnagram
}

/*	Approach 2: Categorize by Count

	Here instead of sorting each string and using it as basis for our hashmap we create an array of size
	26 its index denotes each lowercase charater and the corresponding index value denote the count of a
	character in a string. Then we use this array as key/basis for our hashmap.


	#Algorithm:---->

	We can transform each string s\text{s}s into a character count, count\text{count}count, consisting of
	26 non-negative integers representing the number of a's, b's, c's, etc. We use these counts as the
	basis for our hash map.

	#Complexity Analysis:---->

	Time Complexity: O(NK), where N is the length of strs, and K is the maximum length of a string in strs.
	Counting each string is linear in the size of the string, and we count every string.

	Space Complexity: O(NK), the total information content stored in hashMap.
*/

func GroupAnagram2(arr []string) [][]string {
	hashMap := make(map[[26]int][]string)

	for _, str := range arr {
		countArr := [26]int{}
		for _, char := range str {
			countArr[char-'a']++
		}
		if _, exist := hashMap[countArr]; !exist {
			hashMap[countArr] = []string{}
		}
		hashMap[countArr] = append(hashMap[countArr], str)
	}

	groupAnagram := [][]string{}
	for _, values := range hashMap {
		groupAnagram = append(groupAnagram, values)
	}
	return groupAnagram
}
