package linkedlist

// Initial solution I came up with.
// T.C---> O(n^2)[outer loop will run (n-1)times for n number of nodes in the list and reverseList
// itself has T.C-->O(n) where n is number of nodes in the list.], S.C--->O(1).
// here we are simply reversing the nodes after the current node till we reach end of list. after reaching
// end our list will be modified with nodes in the desired positons.
func ReorderList(head *ListNode) {
	mainCurr := head
	for mainCurr != nil {
		mainCurr.Next = ReverseLinkedList(mainCurr.Next)
		mainCurr = mainCurr.Next
	}
}

/*
	First we intend to divide list into 2 halves and then we reverse the second half, after
	reversing we iterate over first and second half while changing the order of the list.
	Using fast and slow algorithm we find the splitting node of the list which is the node pointed by slow
	when fast becomes nil or fast.next is nil.
	After finding splitting node, we reverse the second half of list.
	Now we iterate over 2 halves while moving nodes to modify our original list.
	** Note: List 2 will have less nodes than list 1 that's y we iterate till secondHalf is not nil.
	If node from firstHalf is not equal to slow then we simply move to next node in firstHalf list. When node
	becomes equal to slow we know we have reached end of firstHalf list, so we add nil to our last node in the
	firstHalf list. **OR we can simply point slow.Next = nil as we know it is the last node in the firstHalf.
	T.C---> O(n), S.C---->O(1)
*/
func ReorderList2(head *ListNode) {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse second half
	secondHalf := ReverseLinkedList(slow.Next)
	slow.Next = secondHalf
	firstHalf := head
	for secondHalf != nil {
		temp1, temp2 := firstHalf.Next, secondHalf.Next
		firstHalf.Next = secondHalf
		secondHalf.Next = temp1
		firstHalf = temp1
		secondHalf = temp2
	}

	// for firstHalf != slow {
	// 	firstHalf = firstHalf.Next
	// }
	// firstHalf.Next = nil
	slow.Next = nil
}
