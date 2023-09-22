package linkedlist

import "fmt"

func (list *LinkedList) AddNode(num int) {
	newNode := &ListNode{Val: num}
	if list.head == nil {
		list.head = newNode
	} else {
		currNode := list.head
		for currNode.Next != nil {
			currNode = currNode.Next
		}
		currNode.Next = newNode
	}
}

func (list *LinkedList) Traverse() {
	currNode := list.head

	for currNode != nil {
		if currNode.Next == nil {
			fmt.Printf(" %v ", currNode.Val)
			fmt.Println()
		} else {
			fmt.Printf(" %v -->", currNode.Val)
		}
		currNode = currNode.Next
	}
}

func ReverseLinkedList(head *ListNode) *ListNode {
	var prevNode *ListNode
	currNode := head

	for currNode != nil {
		tempNode := currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = tempNode
	}
	head = currNode

	return head
}

func LengthOfList(head *ListNode) int {
	if head == nil {
		return 0
	}

	len := 0
	currNode := head
	for currNode != nil {
		currNode = currNode.Next
		len++
	}
	return len
}

func (list *LinkedList) MoveLastToFront() {
	n := list.head
	lastNode := &ListNode{}

	for n.Next != nil {
		n = n.Next
		if n.Next.Next == nil {
			lastNode = n.Next
			n.Next = nil
		}
	}

	lastNode.Next = list.head
	list.head = lastNode
}

func initList() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) PopulateList(k int) {
	for i := 1; i <= k; i++ {
		list.AddNode(i)
	}
}

func (list *LinkedList) DeleteGivenKey(key int) {
	currNode := list.head
	prevNode := list.head

	for currNode.Val != key {
		prevNode = currNode
		currNode = currNode.Next
	}

	prevNode.Next = currNode.Next
	currNode.Next = nil
}

func (list *LinkedList) DeleteKeyAtGivenLocation(pos int) {
	currPos := 0
	currNode := list.head
	prevNode := &ListNode{}

	for currPos != pos-1 {
		prevNode = currNode
		currNode = currNode.Next
		currPos++
	}

	prevNode.Next = currNode.Next
	currNode.Next = nil
}

func (list *LinkedList) InsertionKeyAtGivenPosition(key int, pos int) {
	newNode := &ListNode{Val: key, Next: nil}
	if pos == 1 {
		newNode.Next = list.head
		list.head = newNode
		return
	}
	prevNode := &ListNode{}
	currNode := list.head
	currPos := 0
	for currPos != pos-1 {
		prevNode = currNode
		currNode = currNode.Next
		currPos++
	}

	newNode.Next = prevNode.Next
	prevNode.Next = newNode
}

func (list *LinkedList) InsertionAfterGivenNode(key int, after int) {
	newNode := &ListNode{Val: key, Next: nil}
	currNode := list.head

	for currNode.Val != after {
		currNode = currNode.Next
	}
	newNode.Next = currNode.Next
	currNode.Next = newNode

}

// checks whether the linked list has a cycle or not
func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	currNode := head
	visited := make(map[*ListNode]bool)

	for currNode.Next != nil {
		if visited[currNode] {
			return true
		}
		visited[currNode] = true
		currNode = currNode.Next
	}

	return false
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := &ListNode{}
	prevNode := &ListNode{}

	// traversing through lists while both are not nil.
	for l1 != nil && l2 != nil {
		newNode := new(ListNode)
		newNode.Val, carry = addDigits(l1.Val, l2.Val, carry)

		if head.Next == nil {
			head.Next = newNode
		} else {
			prevNode.Next = newNode
		}

		prevNode = newNode

		l1 = l1.Next
		l2 = l2.Next
	}

	// Now traversing through remaining nodes of either linked lists
	for l1 != nil {
		l1.Val, carry = addDigits(l1.Val, 0, carry)
		prevNode.Next = l1
		prevNode = l1
		l1 = l1.Next
	}

	for l2 != nil {
		l2.Val, carry = addDigits(l2.Val, 0, carry)
		prevNode.Next = l2
		prevNode = l2
		l2 = l2.Next
	}

	// handling case where sum of last node and previous carry is > 0, in such case we have to create new node and add it to the list.
	if carry > 0 {
		n := &ListNode{Val: carry, Next: nil}
		prevNode.Next = n
	}

	return head.Next
}

func addDigits(d1, d2, prevCarry int) (sum, carry int) {
	tempSum := (d1 + d2 + prevCarry)
	sum = tempSum % 10
	carry = tempSum / 10
	return
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prevNode := &ListNode{}
	head := &ListNode{}

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			if head.Next == nil {
				head.Next = l1
			}
			prevNode.Next = l1
			prevNode = l1
			l1 = l1.Next
		} else if l2.Val < l1.Val {
			if head.Next == nil {
				head.Next = l2
			}
			prevNode.Next = l2
			prevNode = l2
			l2 = l2.Next
		} else if l1.Val == l2.Val {
			if head.Next == nil {
				head.Next = l1
			}
			prevNode.Next = l1
			prevNode = l1
			l1 = l1.Next
			prevNode.Next = l2
			prevNode = l2
			l2 = l2.Next
		}
	}

	if l1 != nil {
		prevNode.Next = l1
	}

	if l2 != nil {
		prevNode.Next = l2
	}

	return head.Next
}
