package hashmaps

import (
	"strings"
)

// using 2 hashmaps
func RansomNote(ransomNote, magazine string) bool {
	ransomMap := make(map[string]int)
	magazineMap := make(map[string]int)

	for _, char := range ransomNote {
		count, exist := ransomMap[string(char)]

		if exist {
			ransomMap[string(char)] = count + 1
		} else {
			ransomMap[string(char)] = 1
		}
	}

	for _, char := range magazine {
		count, exist := magazineMap[string(char)]

		if exist {
			magazineMap[string(char)] = count + 1
		} else {
			magazineMap[string(char)] = 1
		}
	}

	for key, val := range ransomMap {
		if val > magazineMap[key] {
			return false
		}
	}
	return true
}

// alternative approach--> using single hashmap.

func RansomNote2(ransomNote, magazine string) bool {
	ransomMap := make(map[string]int)

	for _, char := range ransomNote {
		count, exist := ransomMap[string(char)]

		if exist {
			ransomMap[string(char)] = count + 1
		} else {
			ransomMap[string(char)] = 1
		}
	}

	for _, char := range magazine {
		count, exist := ransomMap[string(char)]

		if exist {
			ransomMap[string(char)] = count - 1
		}
	}

	for _, val := range ransomMap {
		if val > 0 {
			return false
		}
	}

	return true
}

// APPROACH #3 using Count() function from strings Package to count each letter occurances.
func RansomNote3(ransomNote, magazine string) bool {
	for _, char := range ransomNote {
		if strings.Count(ransomNote, string(char)) > strings.Count(magazine, string(char)) {
			return false
		}
	}
	return true
}
