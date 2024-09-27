package linkedlist

//输入: head = [1,4,3,2,5,2], x = 3
//输出: [1,2,2,4,3,5]
//注意我们只要求比 x 小的值挪到左边，并没有要求大于等于 x 的值挪到右边。

// 用两个链表 small链表按顺序存储所有小于x的节点
// large链表按顺序存储所有大于等于x的节点

func partition[T Ordered](head *ListNode[T], x T) *ListNode[T] {
	smallHead := &ListNode[T]{}
	largeHead := &ListNode[T]{}
	small := smallHead
	large := largeHead
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	large.Next = nil
	small.Next = largeHead.Next
	return smallHead.Next
}
