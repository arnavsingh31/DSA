package linkedlist

/*
	LC #160
	Here we traverse through both LL simultaneously until nodeA == nodeB. We basically try bring both pointer
	of the lists at same level/distance from intersecting point/node. we know that both LL might have different
	lengths so inorder to make distance from intersecting point/node same we traverse the pointerA and pointerB
	additional distance of L2 and L1 respectively. where L1 is length of first LL and L2 is length of second
	LL. Once the either pointer finish traversing their current LL and become nil we immediately point them to
	their start of other LL. This way they both pointer will meet at a node which will be our intersecting node.
*/
func IntersectingNode(headA, headB *ListNode) *ListNode {
	currNodeA := headA
	currNodeB := headB

	for currNodeA != currNodeB {
		if currNodeA != nil {
			currNodeA = currNodeA.Next
		} else {
			currNodeA = headB
		}

		if currNodeB != nil {
			currNodeB = currNodeB.Next
		} else {
			currNodeB = headA
		}
	}

	return currNodeA
}

/*
	Here we find the length of both LL and find the max and diff between both lengths. Based on the diff
	we iterate through LL with larger length till diff. Now both pointers A and B will have to travel
	same distance till intersecting node. Hence when they become equal we have found our intersecting node.
*/
func IntersectingNode2(headA, headB *ListNode) *ListNode {
	lenA, lenB := LengthOfList(headA), LengthOfList(headB)
	currNodeA := headA
	currNodeB := headB

	max, diff := maxWithDiff(lenA, lenB)

	if max == lenA {
		for i := 0; i < diff; i++ {
			currNodeA = currNodeA.Next
		}
	} else if max == lenB {
		for i := 0; i < diff; i++ {
			currNodeB = currNodeB.Next
		}
	}

	for currNodeA != currNodeB {
		currNodeA = currNodeA.Next
		currNodeB = currNodeB.Next
	}

	return currNodeA
}

func maxWithDiff(x, y int) (int, int) {
	if x > y {
		return x, x - y
	}
	return y, y - x
}
