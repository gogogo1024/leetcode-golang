package linkedlist

import "container/heap"

// 给定一个链表数组，每个链表都已经按升序排列。
// 将所有链表合并到一个升序链表中，返回合并后的链表
// 输入: lists = [[1,4,5],[1,3,4],[2,6]]
// 输出: [1,1,2,3,4,4,5,6]
// 1.先合并两个，
// 2.合并后在和剩余的其中任意一个合并

// 合并两个有序链表
func mergeTwoLinkedList[T Ordered](a, b *ListNode[T]) *ListNode[T] {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	var head, tail *ListNode[T]
	if a.Val < b.Val {
		head = a
		tail = a
		a = a.Next
	} else {
		head = b
		tail = b
		b = b.Next
	}
	for a != nil && b != nil {
		if a.Val < b.Val {
			tail.Next = a
			a = a.Next
		} else {
			tail.Next = b
			b = b.Next
		}
	}
	if a != nil {
		tail.Next = a
	}
	if b != nil {
		tail.Next = b
	}
	return head
}

// 合并k个有序链表
func mergeKthLinkedList[T Ordered](list []*ListNode[T]) *ListNode[T] {
	var ans *ListNode[T]
	for _, l := range list {
		ans = mergeTwoLinkedList[T](ans, l)
	}
	return ans
}

type MinStack[T Ordered] []*ListNode[T]

func (s MinStack[T]) Len() int {
	return len(s)
}
func (s MinStack[T]) Less(i, j int) bool {
	return s[i].Val < s[j].Val
}
func (s MinStack[T]) Swap(i, j int) {
	s[i], s[j] = s[i], s[j]
}
func (ms *MinStack[T]) Push(x interface{}) {
	*ms = append(*ms, x.(*ListNode[T]))
}
func (ms *MinStack[T]) Pop() interface{} {
	old := *ms
	n := len(old)
	x := old[n-1]
	*ms = old[0 : n-1]
	return x
}

// 利用golang自带的最小堆来实现每次pop出来的都是最小的
func mergeKthLinkedListStack[T Ordered](list []*ListNode[T]) *ListNode[T] {
	ms := make(MinStack[T], 0)
	heap.Init(&ms)
	for _, node := range list {
		heap.Push(&ms, node)
	}
	dummyNode := &ListNode[T]{}
	tail := dummyNode
	for ms.Len() > 0 {
		node := heap.Pop(&ms).(*ListNode[T])
		tail.Next = node
		tail = node
		if node.Next != nil {
			heap.Push(&ms, node.Next)
		}
	}
	return dummyNode.Next
}
