package linkedlist

// LC #328
// Approach I came up with.
func OddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyNode := &ListNode{Next: head}
	prevNode := dummyNode
	currNode := prevNode.Next

	count := 1
	for currNode.Next != nil {
		count++
		prevNode.Next = currNode.Next
		prevNode = currNode
		currNode = currNode.Next

		if count%2 == 0 {
			prevNode.Next = dummyNode.Next
		} else {
			prevNode.Next = currNode.Next
			currNode.Next = dummyNode.Next
		}
	}

	return head
}

// Simplified version of above. T.C--> same as above O(n) and S.C--> O(1)
func OddEvenList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd, even := head, head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead

	return head
}
