package trees

import (
	"fmt"
)

/*
LC #257
TC--->O(n)
SC--->O(n)
*/
func BinaryTreePaths(root *Node) []string {
	res := []string{}

	pathHelper(root, "", &res)
	return res
}

func pathHelper(root *Node, path string, res *[]string) {
	if root == nil {
		return
	}

	seperator := ""

	if path != "" {
		seperator = "->"
	}

	path += fmt.Sprintf("%s%d", seperator, root.Val)

	if root.Left == nil && root.Right == nil {
		*res = append(*res, path)
		return
	}

	pathHelper(root.Left, path, res)
	pathHelper(root.Right, path, res)
}
