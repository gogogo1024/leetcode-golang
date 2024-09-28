package linkedlist

//给你一个链表的头节点 head ，判断链表中是否有环。

//输入: head = [3,2,0,-4]
//输出: true

// 快慢指针移动，当相等时链表必定有环，否则无环
func cycle[T Ordered](head *ListNode[T]) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // 慢指针移动一步
		fast = fast.Next.Next // 快指针移动两步
		if slow == fast {     // 如果相遇，则说明有环
			return true
		}
	}
	return false
}
