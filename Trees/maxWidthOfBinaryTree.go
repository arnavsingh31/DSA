package trees

import util "github.com/arnavsingh31/DSA/Util"

type NodeIndexPair struct {
	node  *TreeNode
	index int
}

func MaxWidthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*NodeIndexPair{{
		root,
		0,
	}}

	maxWidth := 0

	for len(queue) > 0 {
		levelLength := len(queue)
		minIndex := queue[0].index
		var start, end int

		for i := 0; i < levelLength; i++ {
			currPair := queue[0]
			currIndex := currPair.index - minIndex
			currNode := currPair.node
			queue = queue[1:]

			if i == 0 {
				start = currIndex
			}

			if i == levelLength-1 {
				end = currIndex
			}

			if currPair.node.Left != nil {
				queue = append(queue, &NodeIndexPair{
					node:  currNode.Left,
					index: 2*currIndex + 1,
				})
			}

			if currPair.node.Right != nil {
				queue = append(queue, &NodeIndexPair{
					node:  currNode.Right,
					index: 2*currIndex + 2,
				})
			}

		}

		maxWidth = util.Max(maxWidth, end-start+1)
	}

	return maxWidth
}
