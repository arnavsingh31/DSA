package main

func rotateList(head *ListNode, k int) *ListNode {
	if head == nil || (k == lengthOfList(head)) {
		return head
	}

	k %= lengthOfList(head)

	head = recursionRotate(head, k)

	return head
}

func recursionRotate(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}

	dummyNode := &ListNode{Next: head}
	prevNode := dummyNode
	currNode := prevNode.Next

	for currNode != nil {
		if currNode.Next == nil {
			prevNode.Next = nil
			currNode.Next = dummyNode.Next
			dummyNode.Next = currNode
			break
		}
		prevNode = currNode
		currNode = currNode.Next
	}

	return recursionRotate(dummyNode.Next, k-1)
}
