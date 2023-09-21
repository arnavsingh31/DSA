package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Next  *TreeNode
}

type Node struct {
	Left  *Node
	Right *Node
	Val   int
}
