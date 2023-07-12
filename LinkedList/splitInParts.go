package main

// LC #725
/*
	T.C--->  O(N + k), where N is the number of nodes in the given list.
	If k is large, it could still require creating many new empty lists.
	S.C---> O(k), the additional space used in writing the answer.
*/
func splitList(root *ListNode, k int) []*ListNode {
	ans := []*ListNode{}
	var nilNode *ListNode
	var partSize, extra int

	dummyNode := &ListNode{Next: root}
	prevNode := dummyNode
	currNode := prevNode.Next

	len := 0
	for currNode != nil {
		currNode = currNode.Next
		len++
	}

	if len < k {
		partSize = 1
		extra = 0
	} else {
		partSize = len / k
		extra = len % k
	}

	currNode = root
	for currNode != nil {
		head := currNode
		for i := 0; i < partSize; i++ {
			prevNode = currNode
			currNode = currNode.Next
		}
		if extra > 0 {
			prevNode = currNode
			currNode = currNode.Next
			extra--
		}

		prevNode.Next = nil
		ans = append(ans, head)
	}

	if len < k {
		for i := 0; i < k-len; i++ {
			ans = append(ans, nilNode)
		}
	}
	return ans
}
