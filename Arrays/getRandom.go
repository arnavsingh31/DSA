package main

import (
	"math/rand"
)

type RandomizedSet struct {
	RSList []int
	Map    map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		Map:    make(map[int]int),
		RSList: []int{},
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, exist := rs.Map[val]; exist {
		return false
	}

	rs.RSList = append(rs.RSList, val)
	rs.Map[val] = len(rs.RSList) - 1
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	arr_len := len(rs.RSList)

	if index, exist := rs.Map[val]; exist {
		rs.RSList[index], rs.RSList[arr_len-1] = rs.RSList[arr_len-1], rs.RSList[index]
		rs.Map[rs.RSList[index]] = index
		delete(rs.Map, val)
		rs.RSList = rs.RSList[:arr_len-1]
		return true
	} else {
		return false
	}
}

func (rs *RandomizedSet) GetRandom() int {
	random_index := rand.Intn(len(rs.RSList))
	return rs.RSList[random_index]
}

// func main() {
// 	r_set := Constructor()
// 	log.Printf("insert returned %v", r_set.Insert(3))
// 	log.Printf("insert returned %v", r_set.Insert(3))
// 	log.Printf("getRandom returned %v", r_set.GetRandom())
// 	log.Printf("getRandom returned %v", r_set.GetRandom())
// 	log.Printf("insert returned %v", r_set.Insert(1))
// 	log.Printf("remove returned %v", r_set.Remove(3))
// 	log.Printf("getRandom returned %v", r_set.GetRandom())
// 	log.Printf("getRandom returned %v", r_set.GetRandom())
// 	log.Printf("insert returned %v", r_set.Insert(0))
// 	log.Printf("remove returned %v", r_set.Remove(0))
// }
