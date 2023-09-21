package trees

/*
	LC #106
	TC---> O(n)
	SC---> O(n)
*/
func BuildTree2(inorder, postorder []int) *Node {
	if len(inorder) == 0 {
		return nil
	}

	root := &Node{Val: postorder[len(postorder)-1]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}

	root.Left = BuildTree2(inorder[:i], postorder[:i])
	root.Right = BuildTree2(inorder[i+1:], postorder[i:len(postorder)-1])

	return root
}

/*
	inorder = [9,3,15,20,7]
	postorder = [9,15,7,20,3]

	root -> 3
	i- > 1
	in[:1] = [9]
	post[:1] = [9]

	root-> 9
	i-> 0
	in[:0] = []
	post[:0] = []

*/
