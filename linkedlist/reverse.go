package linkedlist

//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表
//输入： 1->2->3->4->5
//输出： 5->4->3->2->1

type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
}

func reverseList[T any](head *ListNode[T]) *ListNode[T] {
	// 在head节点前面添加一个空节点，保存当前节点的前一个节点
	// 理解为隐藏的Pre，方便移动
	var prev *ListNode[T]
	curr := head
	for curr != nil {
		next := curr.Next // 保存下一个节点
		curr.Next = prev  // 当前节点指向当前节点的前一个节点
		prev = curr       // 更新前一个节点为当前节点
		curr = next       // 当前节点更新为下一个节点
	}
	return prev
}
