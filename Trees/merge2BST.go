package trees

/*
	it can also be done by :
	1. converting both bst to sorted array(inorder traversal)
	2. merge 2 sorted array
	3. create a bst from combined sorted array

	but it will use more space. so we implemented them using linked list method.
*/
func Merge2BST(root1, root2 *Node) *Node {
	// step 1 convert both bst to sorted linked list
	head1 := bstToLinkedList(root1)
	head2 := bstToLinkedList(root2)

	// step 2 merge 2 sorted linked list
	head := merge2List(head1, head2)

	// step 3 convert the merged linked list to a balanced BST
	size := linkedListSize(head)
	return linkedListToBST(head, size)
}

func bstToLinkedList(root *Node) *Node {
	var prev *Node

	var dfs func(*Node)

	dfs = func(root *Node) {
		if root == nil {
			return
		}

		dfs(root.Right)
		dfs(root.Left)
		root.Right = prev
		root.Left = nil
		prev = root
	}

	return prev
}

func merge2List(head1, head2 *Node) *Node {
	curr1 := head1
	curr2 := head2

	dummy := &Node{Right: nil}
	prev := dummy

	for curr1 != nil && curr2 != nil {
		if curr1.Val < curr2.Val {
			prev.Right = curr1
			prev = curr1
			curr1 = curr1.Right
		} else {
			prev.Right = curr2
			prev = curr2
			curr2 = curr2.Right
		}
	}

	for curr1 != nil {
		prev.Right = curr1
		prev = curr1
		curr1 = curr1.Right
	}

	for curr2 != nil {
		prev.Right = curr2
		prev = curr2
		curr2 = curr2.Right
	}

	return dummy.Right
}

func linkedListToBST(head *Node, size int) *Node {
	if head == nil || size <= 0 {
		return nil
	}

	// make left sub tree from starting n/2 nodes.
	leftSubTree := linkedListToBST(head, size/2)

	// now create a root
	root := head
	root.Left = leftSubTree

	head = head.Right

	rightSubTree := linkedListToBST(head, size-size/2-1)
	root.Right = rightSubTree

	return root
}

func linkedListSize(head *Node) int {
	if head == nil {
		return 0
	}

	count := 0
	curr := head

	for curr != nil {
		count++
		curr = curr.Right
	}

	return count
}
