package main

import "math"

// Implemented using singly-linked list
type MinStack struct {
	head *Node
}

type Node struct {
	Val  int
	Prev *Node
	Min  int
}

func MinStackConstructor() MinStack {
	return MinStack{
		head: &Node{
			Min: math.MaxInt32,
		},
	}
}

func (ms *MinStack) Push(val int) {
	min := Min(ms.head.Min, val)
	ms.head = &Node{
		Val:  val,
		Prev: ms.head,
		Min:  min,
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func (ms *MinStack) Pop() {
	tempNode := ms.head.Prev
	ms.head.Prev = nil
	ms.head = tempNode
}

func (ms *MinStack) GetMin() int {
	return ms.head.Min
}

func (ms *MinStack) Top() int {
	return ms.head.Val
}
