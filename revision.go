package main

import (
	"container/heap"
	"math"

	linkedlist "github.com/arnavsingh31/DSA/LinkedList"
	trees "github.com/arnavsingh31/DSA/Trees"
	util "github.com/arnavsingh31/DSA/Util"
)

// heap revision
func bstToHeap(root *trees.Node) *trees.Node {
	arr := inorderTraversal(root)

	solveRec(root, arr, 0)
	return root
}

func solveRec(root *trees.Node, arr []int, index int) {
	if root == nil {
		return
	}

	root.Val = arr[index]
	solveRec(root.Left, arr, index+1)
	solveRec(root.Right, arr, index+1)
}

func inorderTraversal(root *trees.Node) []int {
	if root == nil {
		return []int{}
	}

	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)

	return append(append(append([]int{}, left...), root.Val), right...)
}

// bst ---> in-order traversal kare to we get sorted array
// now traverse existing tree and fill the node val from sorted array in a pre-order .

type HeapNode struct {
	Sum int
	i   int
	j   int
}

type Heap []*HeapNode

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].Sum < h[j].Sum
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*HeapNode))
}

func (h *Heap) Pop() interface{} {
	lastIndex := len(*h) - 1
	deletedNode := (*h)[lastIndex]
	(*h)[lastIndex] = nil
	*h = (*h)[:lastIndex]
	return deletedNode
}

type Pair struct {
	i, j int
}

func kPairs(arr1, arr2 []int, k int) [][]int {
	pq := make(Heap, 0)
	ans := [][]int{}
	visited := make(map[Pair]bool)

	pq = append(pq, &HeapNode{
		Sum: arr1[0] + arr2[0],
		i:   0,
		j:   0,
	})

	heap.Init(&pq)
	visited[Pair{i: 0, j: 0}] = true

	for len(pq) > 0 && k > 0 {
		top := heap.Pop(&pq).(*HeapNode)
		ans = append(ans, []int{arr1[top.i], arr2[top.j]})

		if top.i+1 < len(arr1) && !visited[Pair{i: top.i + 1, j: top.j}] {
			next := &HeapNode{
				Sum: arr1[top.i+1] + arr2[top.j],
				i:   top.i + 1,
				j:   top.j,
			}
			heap.Push(&pq, next)
			visited[Pair{i: top.i + 1, j: top.j}] = true
		}

		if top.j+1 < len(arr2) && !visited[Pair{i: top.i, j: top.j + 1}] {
			next := &HeapNode{
				Sum: arr1[top.i] + arr2[top.j+1],
				i:   top.i,
				j:   top.j + 1,
			}
			heap.Push(&pq, next)
			visited[Pair{i: top.i, j: top.j + 1}] = true
		}
		k--
	}

	return ans
}

// method 1 do bfs at every level check whether we a parent has two child or only left child .
func isCompleteBinaryTree(root *trees.Node) bool {
	queue := []*trees.Node{root}
	flag := false

	for len(queue) > 0 {
		levelLength := len(queue)

		for i := 0; i < levelLength; i++ {

			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				if flag {
					return false
				}
				queue = append(queue, node.Left)
			} else {
				flag = true
			}

			if node.Right != nil {
				if flag {
					return false
				} else {
					queue = append(queue, node.Right)
				}
			} else {
				flag = true
			}
		}
	}
	return true
}

func isCompleteBinaryTree2(root *trees.Node) bool {
	nodeCount := totalNodes(root)
	return solveCBT(root, 0, nodeCount)
}

func totalNodes(root *trees.Node) int {
	if root == nil {
		return 0
	}

	left := totalNodes(root.Left)
	right := totalNodes(root.Right)

	return 1 + left + right
}

func solveCBT(root *trees.Node, index, count int) bool {
	if root == nil {
		return true
	}

	if index >= count {
		return false
	}

	left := solveCBT(root.Left, 2*index+1, count)
	right := solveCBT(root.Right, 2*index+2, count)

	return left && right

}

// type MinHeap []*int

func KLargest(arr []int, k int) int {
	pq := make(MinHeap, 0)
	heap.Init(&pq)

	for i := 0; i < len(arr); i++ {
		heap.Push(&pq, &arr[i])
		if pq.Len() > k {
			heap.Pop(&pq)
		}
	}
	top := *heap.Pop(&pq).(*int)
	return top
}

func KLargestSum(arr []int, k int) int {
	minHeap := make(MinHeap, 0)
	heap.Init(&minHeap)

	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			heap.Push(&minHeap, sum)
			if len(minHeap) > k {
				heap.Pop(&minHeap)
			}
		}
	}

	return minHeap[0]
}

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	lastIndex := len(*h) - 1
	lastElement := (*h)[lastIndex]
	*h = (*h)[:lastIndex]
	return lastElement
}

type CustomMinHeap []*Node
type Node struct {
	row  int
	col  int
	data int
}

func (h CustomMinHeap) Len() int {
	return len(h)
}

func (h CustomMinHeap) Less(i, j int) bool {
	return h[i].data < h[j].data
}

func (h CustomMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *CustomMinHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *CustomMinHeap) Pop() interface{} {
	lastIndex := len(*h) - 1
	lastElement := (*h)[lastIndex]
	(*h)[lastIndex] = nil
	*h = (*h)[:lastIndex]
	return lastElement
}

func merge(matrix [][]int) []int {
	ans := make([]int, 0)
	rows := len(matrix)
	cols := len(matrix[0])

	pq := make(CustomMinHeap, 0)
	heap.Init(&pq)

	for i := 0; i < rows; i++ {
		heap.Push(&pq, &Node{
			row:  i,
			col:  0,
			data: matrix[i][0],
		})
	}

	for len(pq) > 0 {
		top := heap.Pop(&pq).(*Node)
		ans = append(ans, top.data)
		if top.col+1 < cols {
			heap.Push(&pq, &Node{
				row:  top.row,
				col:  top.col + 1,
				data: matrix[top.row][top.col+1],
			})
		}
	}
	return ans
}

// merge k sorted array of different sizes
func merge2(matrix [][]int) []int {
	ans := make([]int, 0)
	rows := len(matrix)

	pq := make(CustomMinHeap, 0)
	heap.Init(&pq)

	for i := 0; i < rows; i++ {
		heap.Push(&pq, &Node{
			row:  i,
			col:  0,
			data: matrix[i][0],
		})
	}

	for len(pq) > 0 {
		top := heap.Pop(&pq).(*Node)
		ans = append(ans, top.data)

		if top.col+1 < len(matrix[top.row]) {
			heap.Push(&pq, &Node{
				row:  top.row,
				col:  top.col + 1,
				data: matrix[top.row][top.col+1],
			})
		}
	}
	return ans
}

type LNode struct {
	Next *LNode
	Val  int
}

func mergeLists(arr []*LNode) *LNode {
	pq := make(PriorityQueue, 0)
	var head, tail *LNode

	heap.Init(&pq)

	for i := 0; i < len(arr); i++ {
		if arr[i] != nil {
			heap.Push(&pq, arr[i])
		}
	}

	for len(pq) > 0 {
		top := heap.Pop(&pq).(*LNode)

		//make sorted list
		if head == nil {
			head = top
			tail = top
		} else {
			tail.Next = top
			tail = top
		}

		if top.Next != nil {
			heap.Push(&pq, top.Next)
		}
	}
	return head
}

type PriorityQueue []*LNode

func (h PriorityQueue) Len() int {
	return len(h)
}

func (h PriorityQueue) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h PriorityQueue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PriorityQueue) Push(x interface{}) {
	*h = append(*h, x.(*LNode))
}

func (h *PriorityQueue) Pop() interface{} {
	lastIndex := len(*h) - 1
	lastItem := (*h)[lastIndex]
	(*h)[lastIndex] = nil
	*h = (*h)[:lastIndex]
	return lastItem
}

func joinRopes(arr []int) int {
	pq := make(MinHeap, 0)

	pq = append(pq, arr...)
	heap.Init(&pq)

	cost := 0
	for len(pq) > 1 {
		top1 := heap.Pop(&pq).(int)
		top2 := heap.Pop(&pq).(int)

		sum := top1 + top2
		cost += sum
		heap.Push(&pq, sum)
	}
	return cost
}

func smallestRange(matrix [][]int) []int {
	pq := make(CustomMinHeap, 0)
	rows := len(matrix)
	ans := make([]int, 2)

	heap.Init(&pq)

	max := math.MinInt
	min := math.MaxInt

	for i := 0; i < rows; i++ {
		max = util.Max(max, matrix[i][0])
		min = util.Min(min, matrix[i][0])
		heap.Push(&pq, &Node{
			row:  i,
			col:  0,
			data: matrix[i][0],
		})
	}

	start := min
	end := max

	for len(pq) > 0 {
		top := heap.Pop(&pq).(*Node)

		if end-start > max-top.data {
			start = top.data
			end = max
		}

		if top.col+1 < len(matrix[top.row]) {
			next := &Node{
				row:  top.row,
				col:  top.col + 1,
				data: matrix[top.row][top.col+1],
			}
			heap.Push(&pq, next)
			max = util.Max(max, matrix[top.row][top.col+1])
		} else {
			break
		}
	}

	ans[0] = start
	ans[1] = end

	return ans
}

// -------------------------------------------------------------------------------------------------------//

// Linked list revision

func add2Num(l1 *linkedlist.ListNode, l2 *linkedlist.ListNode) *linkedlist.ListNode {
	sum, carry := 0, 0
	var head, tail *linkedlist.ListNode

	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		sum = sum % 10

		newNode := &linkedlist.ListNode{
			Val:  sum,
			Next: nil,
		}

		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}
		sum = carry
	}

	if carry > 0 {
		newNode := &linkedlist.ListNode{
			Val:  carry,
			Next: nil,
		}
		tail.Next = newNode
		tail = newNode
	}

	return head
}

func copyListRandom(head *linkedlist.SpecialNode) *linkedlist.SpecialNode {
	if head == nil {
		return nil
	}

	var newHead, newTail *linkedlist.SpecialNode
	oldToNewNodeMap := make(map[*linkedlist.SpecialNode]*linkedlist.SpecialNode)

	currNode := head
	for currNode != nil {
		newNode := &linkedlist.SpecialNode{Val: currNode.Val}

		// populate both maps
		oldToNewNodeMap[currNode] = newNode

		if newHead == nil {
			newHead = newNode
			newTail = newNode
		} else {
			newTail.Next = newNode
			newTail = newNode
		}

		currNode = currNode.Next
	}

	// iterate over both list and populate random pointer
	currNode = head
	copyCurr := newHead
	for currNode != nil {
		copyCurr.Random = oldToNewNodeMap[currNode.Random]
		currNode = currNode.Next
		copyCurr = copyCurr.Next
	}

	return newHead
}

func deleteNfromEnd(head *linkedlist.ListNode, n int) *linkedlist.ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	for n > 0 {
		fast = fast.Next
		n--
	}

	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return head
}

func deletedNode(node *linkedlist.ListNode) {
	prevNode := node

	for node.Next != nil {
		node.Val = node.Next.Val
		prevNode = node
		node = node.Next
	}

	prevNode.Val = node.Val
	prevNode.Next = nil
}

func intersection(headA, headB *linkedlist.ListNode) *linkedlist.ListNode {
	currA := headA
	currB := headB

	for currA != currB {
		if currA != nil {
			currA = currA.Next
		} else {
			currA = headB
		}

		if currB != nil {
			currB = currB.Next
		} else {
			currB = headA
		}
	}

	return currA
}

func cycle2(head *linkedlist.ListNode) *linkedlist.ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}

func mergeBetween(list1, list2 *linkedlist.ListNode, a int, b int) *linkedlist.ListNode {
	curr1 := list1
	curr2 := list2

	for curr2.Next != nil {
		curr2 = curr2.Next
	}

	var start, prev *linkedlist.ListNode
	count := 0

	for count <= b {
		if count == a {
			start = prev
		}

		prev = curr1
		curr1 = curr1.Next
		count++
	}

	start.Next = list2
	curr2.Next = curr1

	return list1
}

func nextGreaterNode(head *linkedlist.ListNode) []int {
	// using monotonic stack we can find next greater node for each node in the list
	stack := make([]int, 0)
	temp := make([]int, 0)
	count := 0
	currNode := head

	for currNode != nil {
		temp = append(temp, currNode.Val)
		count++
		currNode = currNode.Next
	}

	ans := make([]int, count)
	for i := 0; i < len(temp); i++ {
		for len(stack) > 0 && temp[stack[len(stack)-1]] < temp[i] {
			topIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[topIndex] = temp[i]
		}

		stack = append(stack, i)
	}

	return ans
}

func oddEvenList(head *linkedlist.ListNode) *linkedlist.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd := head
	even := head.Next

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

func isPalindrome(head *linkedlist.ListNode) bool {
	var prev *linkedlist.ListNode
	currNode := head
	length := 0
	count := 0

	for currNode != nil {
		length++
		currNode = currNode.Next
	}

	// reverse till half of length
	currNode = head
	for count < length/2 {
		temp := currNode.Next
		currNode.Next = prev
		prev = currNode
		currNode = temp
		count++
	}

	// in case length of ll is odd then move the second half list starting by 1 node (as this node is common)
	if length%2 != 0 {
		currNode = currNode.Next
	}

	currNode2 := prev
	for currNode2 != nil && currNode != nil {
		if currNode2.Val != currNode.Val {
			return false
		}
		currNode2 = currNode2.Next
		currNode = currNode.Next
	}

	return true
}

func partitionList(head *linkedlist.ListNode, x int) *linkedlist.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var smallHead, bigHead, smallCurr, bigCurr *linkedlist.ListNode
	currNode := head

	for currNode != nil {
		if x <= currNode.Val {
			if bigHead == nil {
				bigHead = currNode
				bigCurr = currNode
			} else {
				bigCurr.Next = currNode
				bigCurr = currNode
			}
		} else {
			if smallHead == nil {
				smallHead = currNode
				smallCurr = currNode
			} else {
				smallCurr.Next = currNode
				smallCurr = currNode
			}
		}
		currNode = currNode.Next
	}

	if smallHead == nil {
		return bigHead
	}

	smallCurr.Next = bigHead

	if bigHead != nil {
		bigCurr.Next = nil
	}

	return smallHead
}

// func removeDuplicates(head *linkedlist.ListNode) *linkedlist.ListNode {
// 	if head == nil && head.Next == nil {
// 		return head
// 	}

// 	nodeCountMap := make(map[int]int)
// 	currNode := head
// 	for currNode != nil {
// 		nodeCountMap[currNode.Val] += 1
// 		currNode = currNode.Next
// 	}

// 	currNode = head

// 	for currNode != nil{
// 		if nodeCountMap[currNode.Val] == 1 {

// 		}
// 	}

// }
