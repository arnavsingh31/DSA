package linkedlist

// LC #1669
func MergeInBetween(list1, list2 *ListNode, a, b int) *ListNode {
	dummyNode := &ListNode{Next: list1}
	prevNode := dummyNode
	curr1 := prevNode.Next
	curr2 := list2
	count := 0
	for curr1 != nil && count <= b {
		if count == a {
			prevNode.Next = list2
		}

		count++
		prevNode = curr1
		curr1 = curr1.Next
	}

	for curr2.Next != nil {
		curr2 = curr2.Next
	}

	curr2.Next = curr1

	return dummyNode.Next
}
