package linkedlist

/*
	LC #148
	TC--->O(nlogn)
	SC--->O(1)
*/
func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	head2 := slow.Next
	slow.Next = nil

	left := SortList(head)
	right := SortList(head2)

	return merge2List(left, right)
}

func merge2List(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}

	if head2 == nil {
		return head1
	}

	var head, tail *ListNode

	curr1, curr2 := head1, head2

	for curr1 != nil && curr2 != nil {
		if curr1.Val <= curr2.Val {
			if head == nil {
				head = curr1
			} else {
				tail.Next = curr1
			}
			tail = curr1
			curr1 = curr1.Next
		} else if curr2.Val < curr1.Val {
			if head == nil {
				head = curr2
			} else {
				tail.Next = curr2
			}
			tail = curr2
			curr2 = curr2.Next
		}
	}

	for curr1 != nil {
		tail.Next = curr1
		tail = curr1
		curr1 = curr1.Next
	}

	for curr2 != nil {
		tail.Next = curr2
		tail = curr2
		curr2 = curr2.Next
	}

	return head
}
