package trees

import (
	"math"
	"sort"

	util "github.com/arnavsingh31/DSA/Util"
)

/*
	LC #987
	TC--->O(n)

	GFG problem (good question)

*/

func VerticalTraversal(root *Node) [][]int {
	// ans := make([]int, 0)
	ans2 := make([][]int, 0)
	minHd, maxHd := math.MaxInt, math.MinInt
	horizontalDistanceToNodesMap := make(map[int]map[int][]int)
	queue := []*NodeInfo{{root, 0, 0}}
	totalLevel := 0

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			temp := queue[0]
			queue = queue[1:]

			node := temp.node
			horizontalDistance := temp.hd
			nodeLevel := temp.level

			minHd = util.Min(minHd, horizontalDistance)
			maxHd = util.Max(maxHd, horizontalDistance)

			if _, exist := horizontalDistanceToNodesMap[horizontalDistance]; !exist {
				horizontalDistanceToNodesMap[horizontalDistance] = make(map[int][]int)
			}

			horizontalDistanceToNodesMap[horizontalDistance][nodeLevel] = append(horizontalDistanceToNodesMap[horizontalDistance][nodeLevel], node.Val)

			if node.Left != nil {
				queue = append(queue, &NodeInfo{node.Left, horizontalDistance - 1, nodeLevel + 1})
			}

			if node.Right != nil {
				queue = append(queue, &NodeInfo{node.Right, horizontalDistance + 1, nodeLevel + 1})
			}
		}
		totalLevel++

	}

	// sorting on basis of value in case there are multiple nodes on same hd and level.
	for i := minHd; i <= maxHd; i++ {
		for j := 0; j <= totalLevel; j++ {
			sort.Ints(horizontalDistanceToNodesMap[i][j])
		}
	}

	for i := minHd; i <= maxHd; i++ {
		arr := make([]int, 0)

		for j := 0; j <= totalLevel; j++ {
			// for k := 0; k < len(horizontalDistanceToNodesMap[i][j]); k++ {
			// 	ans = append(ans, horizontalDistanceToNodesMap[i][j][k])
			// }
			if _, exist := horizontalDistanceToNodesMap[i][j]; exist {
				arr = append(arr, horizontalDistanceToNodesMap[i][j]...)
			}
		}
		ans2 = append(ans2, arr)
	}

	return ans2
}
