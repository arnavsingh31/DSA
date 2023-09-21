package trees

/*
	LC #105
	TC---> O(n)
	SC---> O(n)
*/
func BuildTree(preorder, inorder []int) *Node {
	if len(inorder) == 0 {
		return nil
	}

	root := &Node{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}

	root.Left = BuildTree(preorder[1:i+1], inorder[:i])
	root.Right = BuildTree(preorder[i+1:], inorder[i+1:])
	return root
}

/*
	[3,9,20,15,7]-preorder
	[9,3,15,20,7]-inorder
	root --> 3
	i = 1
	newPre = [1:2]=[9]
	newIn = [:1] = [9]

	root --> 9
	i = 0
	newPre = [1:1] = []
	newIn = [:0] = []
*/
