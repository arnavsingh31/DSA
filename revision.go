package main

import (
	"container/heap"
	"math"
	"sort"

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

func removeDuplicates(head *linkedlist.ListNode) *linkedlist.ListNode {
	if head == nil && head.Next == nil {
		return head
	}

	dummy := &linkedlist.ListNode{Next: head}
	prev := dummy
	currNode := prev.Next

	for currNode != nil {
		if currNode.Next != nil && currNode.Val == currNode.Next.Val {
			for currNode.Next != nil && currNode.Val == currNode.Next.Val {
				currNode = currNode.Next
			}
			prev.Next = currNode.Next
		} else {
			prev = currNode
		}
		currNode = currNode.Next
	}

	return dummy.Next
}

func removeElements(head *linkedlist.ListNode, val int) *linkedlist.ListNode {
	if head == nil {
		return head
	}

	dummy := &linkedlist.ListNode{Next: head}
	prev := dummy
	currNode := prev.Next

	for currNode != nil {
		if currNode.Val == val {
			prev.Next = currNode.Next
		} else {
			prev = currNode
		}
		currNode = currNode.Next

	}
	return dummy.Next
}

func reorder(head *linkedlist.ListNode) {
	if head.Next == nil {
		return
	}

	currNode := head
	for currNode != nil {
		newHead := reverse(currNode.Next)
		currNode.Next = newHead
		currNode = currNode.Next
	}
}

func reverse(head *linkedlist.ListNode) *linkedlist.ListNode {
	if head == nil {
		return head
	}
	var prev *linkedlist.ListNode
	currNode := head
	for currNode != nil {
		temp := currNode.Next
		currNode.Next = prev
		prev = currNode
		currNode = temp
	}

	return prev
}

func reorder2(head *linkedlist.ListNode) {
	if head.Next == nil {
		return
	}

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse 2 nd half
	currNodeRev := slow.Next
	var prev *linkedlist.ListNode
	slow.Next = nil
	for currNodeRev != nil {
		temp := currNodeRev.Next
		currNodeRev.Next = prev
		prev = currNodeRev
		currNodeRev = temp
	}

	currNode := head
	currNodeRev = prev
	for currNodeRev != nil {
		temp := currNode.Next
		tempRev := currNodeRev.Next

		currNode.Next = currNodeRev
		currNodeRev.Next = temp

		currNode = temp
		currNodeRev = tempRev
	}
}

func reverseBetween(head *linkedlist.ListNode, left, right int) *linkedlist.ListNode {
	if head.Next == nil {
		return head
	}

	dummy := &linkedlist.ListNode{Next: head}
	prev := dummy
	currNode := prev.Next
	count := 1

	for count < left {
		prev = currNode
		currNode = currNode.Next
		count++
	}
	nodeBeforeLeft := prev

	leftNode := currNode
	//now reverse from left to right
	for count < right {
		temp := currNode.Next
		currNode.Next = prev
		prev = currNode
		currNode = temp
		count++
	}

	nodeBeforeLeft.Next = prev
	leftNode.Next = currNode

	return dummy.Next
}

func reverseInKGroup(head *linkedlist.ListNode, k int) *linkedlist.ListNode {
	currNode := head
	len := 0
	for currNode != nil {
		len++
		currNode = currNode.Next
	}

	return helper1(head, k, len)
}

func helper1(head *linkedlist.ListNode, k int, totalNodes int) *linkedlist.ListNode {
	if head == nil || head.Next == nil || totalNodes < k {
		return head
	}

	// reversed first k nodes
	var prev, temp *linkedlist.ListNode
	currNode := head
	count := 1

	for count <= k {
		temp = currNode.Next
		currNode.Next = prev
		prev = currNode
		currNode = temp
		count++
	}

	head.Next = helper1(temp, k, totalNodes-k)

	return prev
}

func rotateList(head *linkedlist.ListNode, k int) *linkedlist.ListNode {
	len := 0
	currNode := head
	for currNode != nil {
		len++
		currNode = currNode.Next
	}

	k = k % len
	if k == 0 {
		return head
	}

	return helper2(head, k)
}

func helper2(head *linkedlist.ListNode, k int) *linkedlist.ListNode {

	var prev *linkedlist.ListNode
	currNode := head

	for currNode.Next != nil {
		prev = currNode
		currNode = currNode.Next
	}

	prev.Next = nil
	currNode.Next = head
	head = currNode

	return helper2(head, k-1)
}

func rotateList2(head *linkedlist.ListNode, k int) *linkedlist.ListNode {
	len := 0
	originalHead := head
	currNode := head
	for currNode != nil {
		len++
		currNode = currNode.Next
	}

	if head == nil || k == 0 || k == len || len == 1 || k%len == 0 {
		return head
	}

	k = k % len
	n := len - k - 1
	currNode = head
	for currNode.Next != nil && n > 0 {
		currNode = currNode.Next
		n--
	}

	head = currNode.Next
	currNode.Next = nil

	currNode = head
	for currNode.Next != nil {
		currNode = currNode.Next
	}
	currNode.Next = originalHead

	return head
}

func split(head *linkedlist.ListNode, k int) []*linkedlist.ListNode {
	ans := make([]*linkedlist.ListNode, 0)
	length := 0
	currNode := head
	for currNode != nil {
		length++
		currNode = currNode.Next
	}

	minGroupSize := length / k
	extraNode := length % k

	currNode = head
	for currNode != nil {
		// handle extra nodes after equal division of nodes
		count := minGroupSize
		if extraNode >= 1 {
			count += 1
			extraNode--
		}

		// add head of each group

		ans = append(ans, head)

		for count > 1 {
			currNode = currNode.Next
			count--
		}

		head = currNode.Next
		currNode.Next = nil
		currNode = head
	}

	for len(ans) < k {
		ans = append(ans, nil)
	}

	return ans
}

func swapNode(head *linkedlist.ListNode, k int) *linkedlist.ListNode {
	var ktnNodeFromBeg, kthNodeFromEnd *linkedlist.ListNode
	slow, fast := head, head
	n := k
	for n > 1 {
		fast = fast.Next
		n--
	}
	ktnNodeFromBeg = fast

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	kthNodeFromEnd = slow

	ktnNodeFromBeg.Val, kthNodeFromEnd.Val = kthNodeFromEnd.Val, ktnNodeFromBeg.Val

	return head
}

func swapPairs(head *linkedlist.ListNode) *linkedlist.ListNode {
	return helper3(head, 2)
}

func helper3(head *linkedlist.ListNode, k int) *linkedlist.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *linkedlist.ListNode
	currNode := head
	for k > 0 {
		k--
		temp := currNode.Next
		currNode.Next = prev
		prev = currNode
		currNode = temp
	}
	head.Next = helper3(currNode, 2)

	return prev
}

//-----------------------------------------Trees-----------------------------------------------------------//

func averageOfLevel(root *trees.Node) []float64 {
	ans := make([]float64, 0)

	if root == nil {
		return ans
	}

	queue := []*trees.Node{root}

	for len(queue) > 0 {
		levelLength := len(queue)
		sum := 0.0
		for i := 0; i < levelLength; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += float64(node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		ans = append(ans, sum/float64(levelLength))
	}

	return ans
}

func buildTree(preorder []int, inorder []int) *trees.Node {
	root := &trees.Node{Val: preorder[0]}
	i := 0

	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}

	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}

func flatten(root *trees.Node) {
	helper4(root)
}

func helper4(root *trees.Node) *trees.Node {
	if root == nil {
		return root
	}

	leftTail := helper4(root.Left)
	rightTail := helper4(root.Right)

	if root.Left != nil {
		leftTail.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}

	if rightTail != nil {
		return rightTail
	}

	if leftTail != nil {
		return leftTail
	}

	return root
}

func flatten2(root *trees.Node) {
	var prev *trees.Node

	var dfs func(*trees.Node)
	dfs = func(root *trees.Node) {
		if root == nil {
			return
		}

		dfs(root.Right)
		dfs(root.Left)
		root.Right = prev
		root.Left = nil
		prev = root
	}

	dfs(root)
}

func lowestCommonAncestor(root, p, q *trees.Node) *trees.Node {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		return right
	} else if right == nil {
		return left
	}

	return root
}

func zigZagTraversal(root *trees.Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*trees.Node{root}
	ans := make([][]int, 0)

	leftToRight := true
	for len(queue) > 0 {
		levelSize := len(queue)
		arr := make([]int, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if leftToRight {
				arr[i] = node.Val
			} else {
				arr[levelSize-1-i] = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		ans = append(ans, arr)
		leftToRight = !leftToRight
	}
	return ans
}

func buildTree2(inorder, postorder []int) *trees.Node {
	n := len(inorder) // both have same length
	nodeToIndex := make(map[int]int)

	for i := 0; i < len(inorder); i++ {
		nodeToIndex[inorder[i]] = i
	}
	rootIndex := n - 1
	return makeTree(inorder, postorder, nodeToIndex, &rootIndex, 0, n-1)
}

func makeTree(inorder, postorder []int, nodeToIndex map[int]int, index *int, inorderStart, inorderEnd int) *trees.Node {
	if *index < 0 || inorderStart > inorderEnd {
		return nil
	}

	element := postorder[*index]
	node := &trees.Node{Val: element}
	pos := nodeToIndex[element]
	*index--

	node.Right = makeTree(inorder, postorder, nodeToIndex, index, pos+1, inorderEnd)
	node.Left = makeTree(inorder, postorder, nodeToIndex, index, inorderStart, pos-1)

	return node
}

func minPathSum(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i > 0 && j > 0 {
				grid[i][j] = grid[i][j] + util.Min(grid[i-1][j], grid[i][j-1])
			} else if i > 0 {
				grid[i][j] = grid[i][j] + grid[i-1][j]
			} else if j > 0 {
				grid[i][j] = grid[i][j] + grid[i][j-1]
			}

		}
	}

	return grid[rows-1][cols-1]
}

func numIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])
	visited := make(map[Cell]bool)
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	isCellValid := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < rows && j < cols
	}

	bfs := func(i, j int) {
		queue := []Cell{{i, j}}

		for len(queue) > 0 {
			cell := queue[0]
			queue = queue[1:]

			x, y := cell.row, cell.col
			visited[cell] = true

			for _, dir := range directions {
				newX, newY := x+dir[0], y+dir[1]

				if isCellValid(newX, newY) && grid[newX][newY] == '1' && !visited[Cell{newX, newY}] {
					queue = append(queue, Cell{newX, newY})
					visited[Cell{newX, newY}] = true
				}
			}
		}
	}

	islandCount := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' && !visited[Cell{i, j}] {
				bfs(i, j)
				islandCount++
			}
		}
	}

	return islandCount
}

type Cell struct {
	row, col int
}

func gameOfLife(board [][]int) {
	rows := len(board)
	cols := len(board[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {-1, 1}, {-1, -1}, {1, -1}, {1, 1}}

	isCellValid := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < rows && j < cols
	}

	liveNeighbourCount := func(i, j int) int {
		count := 0
		for _, dir := range directions {
			x, y := i+dir[0], j+dir[1]

			if isCellValid(x, y) && (board[x][y] == 1 || board[x][y] == 3) {
				count++
			}
		}
		return count
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			aliveNeighbour := liveNeighbourCount(i, j)
			if board[i][j] == 1 && (aliveNeighbour > 3 || aliveNeighbour < 2) {
				board[i][j] = 3
			} else if board[i][j] == 0 && aliveNeighbour == 3 {
				board[i][j] = 2
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 2 {
				board[i][j] = 1
			} else if board[i][j] == 3 {
				board[i][j] = 0
			}
		}
	}
}

func getAllSubsequence(arr []int, target int) [][]int {
	ans := make([][]int, 0)
	solveRecSub(arr, target, 0, &ans, []int{}, 0)
	return ans
}

func solveRecSub(arr []int, target, sum int, ans *[][]int, temp []int, index int) {
	if index == len(arr) {
		if sum == target {
			new := append([]int{}, temp...)
			*ans = append(*ans, new)
		}
		return
	}

	// take
	sum += arr[index]
	temp = append(temp, arr[index])
	solveRecSub(arr, target, sum, ans, temp, index+1)

	temp = temp[:len(temp)-1]
	sum -= arr[index]

	// not take
	solveRecSub(arr, target, sum, ans, temp, index+1)
}

func solveNQueens(n int) [][]string {
	board := make([][]string, n)
	ans := make([][]string, 0)

	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}

	generateBoard(0, n, board, &ans)
	return ans
}

func generateBoard(col, n int, board [][]string, ans *[][]string) {
	if col >= n {
		sol := genAns(board, n)
		*ans = append(*ans, sol)
		return
	}

	for i := 0; i < n; i++ {
		if canPlaceQueen(i, col, n, board) {
			board[i][col] = "Q"
			generateBoard(col+1, n, board, ans)
			board[i][col] = "."
		}
	}
}

func canPlaceQueen(i, j, n int, board [][]string) bool {
	// check  row
	x, y := i, j
	for y >= 0 {
		if board[x][y] == "Q" {
			return false
		}
		y--
	}

	// check top-left diagonal
	x, y = i, j
	for x >= 0 && y >= 0 {
		if board[x][y] == "Q" {
			return false
		}
		x--
		y--
	}

	// check bottom-left diagonal
	x, y = i, j
	for x < n && y >= 0 {
		if board[x][y] == "Q" {
			return false
		}
		x++
		y--
	}

	return true
}

func genAns(board [][]string, n int) []string {
	arr := make([]string, 0)

	for i := 0; i < n; i++ {
		temp := ""
		for j := 0; j < n; j++ {
			temp += board[i][j]
		}

		arr = append(arr, temp)
	}

	return arr
}

type Car struct {
	Position int
	Speed    int
}

func carFleet(target int, position, speed []int) int {
	arr := make([]Car, 0)

	for i := 0; i < len(position); i++ {
		arr = append(arr, Car{Position: position[i], Speed: speed[i]})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Position > arr[j].Position
	})

	prevTimeToReachTarget := 0.0
	fleets := 0
	for _, car := range arr {
		time := float64(target-car.Position) / float64(car.Speed)

		if time > prevTimeToReachTarget {
			fleets++
			prevTimeToReachTarget = time
		}
	}
	return fleets
}

func removeDuplicateLetter(s string) string {
	freqMap := make(map[rune]int)
	stackContains := make(map[rune]bool)
	stack := make([]rune, 0)

	for _, r := range s {
		freqMap[r] += 1
	}

	for _, r := range s {
		_, exist := stackContains[r]

		for !exist && len(stack) > 0 && stack[len(stack)-1] > r && freqMap[stack[len(stack)-1]] != 0 {
			char := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			delete(stackContains, char)
		}

		if !exist {
			stack = append(stack, r)
			stackContains[r] = true
		}

		freqMap[r] -= 1
	}

	return string(stack)
}

func repeatedNum(arr []int) int {
	fast := arr[0]
	slow := arr[0]

	for {
		fast = arr[arr[fast]]
		slow = arr[slow]

		if fast == slow {
			break
		}
	}

	fast = arr[0]
	for slow != fast {
		slow = arr[slow]
		fast = arr[fast]
	}

	return slow
}

func isSameTree(p, q *trees.Node) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// binary search method (find max int which on squaring is less than equal to n)
func sqrtOfInt(n int) int {
	left, right := 1, n

	for left < right {
		mid := (left + right) / 2

		if mid*mid <= n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

func minInRotatedSortedArray(arr []int) int {
	left, right := 0, len(arr)-1
	ans := math.MaxInt

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] >= arr[left] {
			ans = util.Min(ans, arr[left])
			left = mid + 1
		} else {
			ans = util.Min(ans, arr[mid])
			right = mid - 1
		}
	}

	return ans
}

func canJump(arr []int) bool {
	n := len(arr)
	max_reach := arr[0]
	lastIndex := n - 1

	for i := 0; i <= max_reach; i++ {
		if i+arr[i] >= lastIndex {
			return true
		} else {
			max_reach = util.Max(max_reach, i+arr[i])
		}
	}

	return false
}

func twoSum(nums []int, target int) []int {
	complementMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		complementMap[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		if index, exist := complementMap[complement]; exist && index != i {
			return []int{index, i}
		}
	}

	return []int{}
}

func candy(ratings []int) int {
	candyArr := make([]int, len(ratings))

	for i := 0; i < len(ratings); i++ {
		candyArr[i] = 1
	}

	// check left neighbour
	for i := 1; i < len(ratings)-1; i++ {
		if ratings[i] > ratings[i-1] {
			candyArr[i] = candyArr[i-1] + 1
		}
	}

	// check right neighbour
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candyArr[i] <= candyArr[i+1] {
			candyArr[i] = candyArr[i+1] + 1
		}
	}

	totalCandy := 0
	for i := 0; i < len(candyArr); i++ {
		totalCandy += candyArr[i]
	}

	return totalCandy
}

func buyChocolate(prices []int, money int) int {
	savings := -1
	ptr1 := 0
	ptr2 := len(prices) - 1

	for i := 0; i < len(prices); i++ {
		cost := prices[ptr1] + prices[ptr2]
		if money-cost > 0 {
			savings = util.Max(savings, money-cost)
		}

		if prices[ptr1] > prices[ptr2] {
			ptr1++
		} else {
			ptr2--
		}
	}

	if savings < 0 {
		return money
	}

	return savings
}

func firstAndLastIndex(arr []int, target int) []int {
	start, end := 0, len(arr)-1
	firstPos, lastPos := -1, -1

	// first find the first occurence of target
	for start <= end {
		mid := (start + end) / 2

		if arr[mid] == target {
			firstPos = mid
			end = mid - 1
		} else if arr[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	// reset the positions
	start, end = 0, len(arr)-1

	// now find the last occurence of the target
	for start <= end {
		mid := (start + end) / 2

		if arr[mid] == target {
			lastPos = mid
			start = mid + 1
		} else if arr[mid] > target {
			end = mid - 1
		} else {
			start = start + 1
		}
	}

	return []int{firstPos, lastPos}
}

func completeCircuit(cost, gas []int) int {
	currGain := 0
	totalGain := 0
	possiblePos := -1

	for i := 0; i < len(gas); i++ {
		totalGain += gas[i] - cost[i]
		currGain += gas[i] - cost[i]

		if currGain >= 0 && possiblePos == -1 {
			possiblePos = i
		} else if currGain < 0 {
			possiblePos = -1
			currGain = 0
		}
	}

	if totalGain >= 0 {
		return possiblePos
	}

	return -1
}

func addOne(digits []int) []int {
	carry := 1

	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		digits[i] = sum % 10
		carry = sum / 10
	}

	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}

	return digits
}

func maxWaterCapacityContainer(heights []int) int {
	start := 0
	end := len(heights) - 1
	maxArea := 0

	for start < end {
		short := util.Min(heights[start], heights[end])
		maxArea = util.Max(maxArea, (end-start)*short)

		if heights[start] < heights[end] {
			start++
		} else {
			end--
		}
	}

	return maxArea
}

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	prefixProduct := make([]int, len(nums))
	postfixProduct := make([]int, len(nums))

	for i := 0; i < len(prefixProduct); i++ {
		prefixProduct[i] = 1
		postfixProduct[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		prefixProduct[i] = prefixProduct[i-1] * nums[i-1]
	}

	for i := len(nums) - 2; i > 0; i-- {
		postfixProduct[i] = postfixProduct[i+1] * nums[i+1]
	}

	for i := 0; i < len(postfixProduct); i++ {
		ans[i] = prefixProduct[i] * postfixProduct[i]
	}

	return ans
}

func prductExceptSelf2(nums []int) []int {
	ans := make([]int, len(nums))
	temp := 1

	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = nums[i-1] * ans[i-1]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] = ans[i] * temp
		temp = nums[i] * temp
	}

	return ans
}
