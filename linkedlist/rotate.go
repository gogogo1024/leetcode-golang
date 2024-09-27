package linkedlist

//旋转链表，将链表每个节点向右移动 k 个位置。
//输入: head = [1,2,3,4,5], k = 2
//输出: [4,5,1,2,3]

// 最后一个节点Next指向头节点
// 位置 length - k%length处的节点的next指向nil
//

// 链表每个节点向右移动 k 个位置
func rotate[T Ordered](head *ListNode[T], k int) *ListNode[T] {
	if head == nil || k == 0 || head.Next == nil {
		return head
	}
	length := 1
	node := head
	for node.Next != nil {
		node = node.Next
		length++
	}
	move := length - k%length
	if move == length {
		return head
	}
	// 最后一个节点指向头节点，此时形成链表环
	node.Next = head

	for move > 0 {
		node = node.Next
		move--
	}
	newHead := node.Next
	node.Next = nil
	return newHead
}
