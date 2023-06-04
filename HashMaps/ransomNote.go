package main

// using 2 hashmaps
func ransomNote(ransomNote, magazine string) bool {
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
