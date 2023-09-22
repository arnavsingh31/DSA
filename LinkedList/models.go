package linkedlist

type LinkedList struct {
	head *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type SpecialNode struct {
	Val    int
	Next   *SpecialNode
	Random *SpecialNode
}
