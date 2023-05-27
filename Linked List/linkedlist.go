package main

import (
	"fmt"
)

type Node struct {
	next *Node
	data interface{}
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) AddNode(d interface{}) {
	newNode := &Node{data: d, next: nil}

	if list.head == nil {
		list.head = newNode
	} else {
		n := list.head

		for n.next != nil {
			n = n.next
		}
		n.next = newNode
	}
}

func (list *LinkedList) Traverse() {
	n := list.head

	for n.next != nil {
		// fmt.Printf("( %v, currentNodeAddres:[%p] )-> ", *n, &*n)
		fmt.Printf("%v -> ", *n)
		n = n.next
	}
	fmt.Printf("%v", *n)
}

func (list *LinkedList) MoveLastToFront() {
	n := list.head
	lastNode := &Node{}

	for n.next != nil {
		n = n.next
		if n.next.next == nil {
			lastNode = n.next
			n.next = nil
		}
	}

	lastNode.next = list.head
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

func (list *LinkedList) ReverseLinkedList() {
	tempNode := &Node{}
	currNode := list.head
	prevNode := list.head

	for currNode.next != nil {
		tempNode.next = currNode.next
		if currNode == list.head {
			currNode.next = nil
			currNode = tempNode.next
		} else {
			tempNode.next, currNode.next = currNode.next, prevNode
			prevNode = currNode
			currNode = tempNode.next
		}
	}
	currNode.next = prevNode
	list.head = currNode
}

func (list *LinkedList) DeleteGivenKey(key int) {
	currNode := list.head
	prevNode := list.head
	// tempNode := &Node{}

	for currNode.data != key {
		prevNode = currNode
		currNode = currNode.next
	}

	prevNode.next = currNode.next
	currNode.next = nil
}

func (list *LinkedList) DeleteKeyAtGivenLocation(pos int) {
	currPos := 0
	currNode := list.head
	prevNode := &Node{}

	for currPos != pos-1 {
		prevNode = currNode
		currNode = currNode.next
		currPos++
	}

	prevNode.next = currNode.next
	currNode.next = nil
}

func (list *LinkedList) InsertionKeyAtGivenPosition(key interface{}, pos int) {
	newNode := &Node{data: key, next: nil}
	if pos == 1 {
		newNode.next = list.head
		list.head = newNode
		return
	}
	prevNode := &Node{}
	currNode := list.head
	currPos := 0
	for currPos != pos-1 {
		prevNode = currNode
		currNode = currNode.next
		currPos++
	}

	newNode.next = prevNode.next
	prevNode.next = newNode
}

func (list *LinkedList) InsertionAfterGivenNode(key interface{}, after int) {
	newNode := &Node{data: key, next: nil}
	currNode := list.head

	for currNode.data != after {
		currNode = currNode.next
	}
	newNode.next = currNode.next
	currNode.next = newNode

}

func (list *LinkedList) OddEvenLinkedList() {
	oddNode := &Node{}
	evenNode := &Node{}
	startEvenNode := &Node{}
	pos := 1
	currNode := list.head
	for currNode.next != nil {
		if pos%2 == 0 {
			if startEvenNode.next == nil {
				startEvenNode.next = currNode
			}
			if evenNode.next != nil {
				evenNode.next.next = currNode
			}
			evenNode.next = currNode

		} else {
			if oddNode.next != nil {
				oddNode.next.next = currNode
			}
			oddNode.next = currNode
		}
		pos++
		currNode = currNode.next
	}

	// checking whether the last node is odd or even and then handling the node respectively
	if pos%2 == 0 {
		if startEvenNode.next != nil {
			evenNode.next.next = currNode
			evenNode.next = currNode
		} else {
			startEvenNode.next = currNode
			evenNode.next = currNode
		}
		// now joinning the odd and even list
		oddNode.next.next = startEvenNode.next
	} else {
		if oddNode.next != nil {
			oddNode.next.next = currNode
			oddNode.next = currNode
		} else {
			oddNode.next = currNode
		}

		if startEvenNode.next != nil {
			oddNode.next.next = startEvenNode.next
			evenNode.next.next = nil
		}
	}
}

func main() {
	newList := initList()
	newList.PopulateList(5)
	newList.Traverse()
	// newList.MoveLastToFront()
	// newList.ReverseLinkedList()
	// newList.DeleteGivenKey(3)
	// newList.DeleteKeyAtGivenLocation(4)
	// newList.InsertionKeyAtGivenPosition(0, 1)
	// newList.InsertionAfterGivenNode(4.5, 4)
	newList.OddEvenLinkedList()
	fmt.Println()
	newList.Traverse()

}

// checks whether the linked list has a cycle or not
func hasCycle(head *Node) bool {
	if head == nil {
		return false
	}
	currNode := head
	visited := make(map[*Node]bool)

	for currNode.next != nil {
		if visited[currNode] {
			return true
		}
		visited[currNode] = true
		currNode = currNode.next
	}

	return false
}
