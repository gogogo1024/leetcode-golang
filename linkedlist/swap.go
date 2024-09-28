package linkedlist

//给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。必须在不修
//改节点内部的值的情况下完成本题（即，只能进行节点交换)
//输入: head = [1,2,3,4]
//输出: [2,1,4,3]

// 递归两两交换节点
func swapRecursion[T Ordered](head *ListNode[T]) *ListNode[T] {
	if head == nil || head.Next == nil {
		return head
	}
	//保存当前头节点的Next为新的头节点
	newHead := head.Next
	// 新的头节点交换后是当前头节点的Next
	head.Next = swapRecursion(newHead.Next)
	// 新的头节点的Next是当前头节点
	newHead.Next = head
	return newHead
}

// 迭代两两交换节点
func swapIteration[T Ordered](head *ListNode[T]) *ListNode[T] {
	if head == nil || head.Next == nil {
		return head
	}

	dummyNode := &ListNode[T]{Next: head}
	temp := dummyNode
	for temp.Next != nil && temp.Next.Next != nil {
		first := temp.Next
		second := temp.Next.Next

		// temp -> first -> second
		// temp -> second ->first
		temp.Next = second
		second.Next = first.Next
		first.Next = second

		// 移动到 first来处理后续的两个节点
		temp = first
	}

	return dummyNode.Next
}
