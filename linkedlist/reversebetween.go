package linkedlist

// 给你单链表的头节点 head,left,right (left<=right)，反转从位置left到right的链表，并返回反转后的链表
// 输入：head=[1,2,3,4,5], left=2 ,right =4
// 输出：[1,4,3,2,5]
// 1.找到位置left节点以及它的前一个节点 (位置从1开始)
// 2.找到位置right节点以及它的前后一个节点
// 3.从链表中断开从位置left到right的作为新链表
// 4.缝合
func reversebetwwen[T any](head *ListNode[T], left, right int) *ListNode[T] {
	var zero T
	dummy := &ListNode[T]{Val: zero, Next: head}

	prev := dummy // 前一个节点
	// 找到位置left节点的前一个节点
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}
	rightNode := prev

	//  找到位置right节点
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	// 位置left节点
	leftNode := prev.Next
	// 位置right节点的下一个节点
	rightNodeNext := rightNode.Next

	// 从位置left到right把链表中断开
	prev.Next = nil
	rightNode.Next = nil

	reverseList[T](leftNode)
	prev.Next = rightNode
	leftNode.Next = rightNodeNext

	return dummy.Next
}
