package stack

/*
	LC #2487
	TC---> O(n)
	SC---> O(n)
	using monotonic stack
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNodes(head *ListNode) *ListNode {
	dec_stack := make([]*ListNode, 0)
	currNode := head

	for currNode != nil {
		for len(dec_stack) > 0 && dec_stack[len(dec_stack)-1].Val < currNode.Val {
			dec_stack = dec_stack[:len(dec_stack)-1]
		}
		dec_stack = append(dec_stack, currNode)
		currNode = currNode.Next
	}

	head = dec_stack[0]
	for i := 0; i < len(dec_stack)-1; i++ {
		dec_stack[i].Next = dec_stack[i+1]
	}
	dec_stack[len(dec_stack)-1].Next = nil
	return head
}
