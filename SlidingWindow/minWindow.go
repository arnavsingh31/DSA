package main

import "log"

func minWindow(s string, t string) string {

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
	log.Printf("mapT is ---->%v", mapT)

	for right := 0; right < sLen; right++ {
		log.Printf("s.right is ----->%v", string(s[right]))
		if val, exist := mapT[s[right]]; exist {
			windowMap[s[right]] += 1
			if windowMap[s[right]] == val {
				charFound++
				log.Printf("charFound=======>%v", charFound)
			}
		}
		log.Printf("window map is ---->%v", windowMap)

		for charFound == len(mapT) && left <= right {
			log.Print("here1")
			log.Printf("left----> %v, right---->%v", left, right)
			if posArr[0] == -1 || (right-left+1) < posArr[0] {
				log.Print("here2")
				posArr[0] = right - left + 1
				posArr[1] = left
				posArr[2] = right
				log.Printf("posArray-------> %v", posArr)
			}
			log.Printf("s.left---%v", string(s[left]))
			if _, exist := windowMap[s[left]]; exist {
				windowMap[s[left]] -= 1
				log.Printf("window map is ---->%v", windowMap)
			}

			if count, exist := mapT[s[left]]; exist && windowMap[s[left]] < count {
				charFound--
				log.Printf("charFound2=======>%v", charFound)
			}

			left++
		}

	}

	if posArr[0] == -1 {
		return ""
	}

	return s[posArr[1] : posArr[2]+1]
}

func main() {
	log.Printf("min window substring is :- %v", minWindow("aa", "aa"))
}

/*
dry run:



*/
