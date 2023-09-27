package trees

import (
	"math"

	util "github.com/arnavsingh31/DSA/Util"
)

/*
GFG problem
*/
func TopView(root *Node) []int {
	if root == nil {
		return []int{}
	}

	ans := make([]int, 0)
	queue := []*NodeHorizontalDistance{{node: root, hd: 0}}
	nodeDistanceMap := make(map[int]int)
	minHd, maxHd := math.MaxInt, math.MinInt

	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]

		node := temp.node
		hd := temp.hd

		minHd = util.Min(minHd, hd)
		maxHd = util.Max(maxHd, hd)

		if _, exist := nodeDistanceMap[hd]; !exist {
			nodeDistanceMap[hd] = node.Val
		}

		if node.Left != nil {
			queue = append(queue, &NodeHorizontalDistance{node.Left, hd - 1})
		}

		if node.Right != nil {
			queue = append(queue, &NodeHorizontalDistance{node.Right, hd + 1})
		}
	}

	for i := minHd; i <= maxHd; i++ {
		if val, exist := nodeDistanceMap[i]; exist {
			ans = append(ans, val)
		}
	}
	return ans
}
