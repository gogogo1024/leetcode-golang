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

// 当我们使用快慢指针法检测到链表中存在环时，快指针和慢指针相遇的位置是在环内部的某个节点。
// 从链表头到环的入口的距离： a。
// 从环的入口到快慢指针相遇的节点的距离： b。
// 环的长度： L

// 相遇的数学关系
// 如前所述，当快慢指针相遇时，快指针比慢指针多走的距离可以表示为：
// 快指针走的距离=a+b+nL
// 慢指针走的距离=a+b

// 由于快指针的速度是慢指针的两倍，我们可以得到以下关系：a+b+nL=2(a+b)
// 通过简化，我们得到：a+b=nL
// 这说明从相遇点继续走 a 步（慢指针从起点也走 a 步）就会到达环的入口，两者会在环的入口节点相遇

// 查找链表环的入口点
func findCycleEntry[T Ordered](head *ListNode[T]) *ListNode[T] {
	// 快慢指针初始化
	slow, fast := head, head

	// 第一步：快慢指针相遇，检测是否有环
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // 慢指针移动一步
		fast = fast.Next.Next // 快指针移动两步
		if slow == fast {     // 相遇时，说明有环
			break
		}
	}

	// 如果没有环，直接返回 nil
	if fast == nil || fast.Next == nil {
		return nil
	}

	// 第二步：将一个指针移回到起点，两个指针都每次移动一步
	slow = head
	for slow != fast {
		slow = slow.Next // 慢指针从头开始移动
		fast = fast.Next // 快指针从相遇点开始移动
	}

	// 相遇点就是环的入口点
	return slow
}

// 通过指针的切换，使两个指针走过相同的路径长度（两个链表首尾相连）
// 从而能够在相交节点相遇
// 找到2个链表中的相交起始点
func findIntersection[T Ordered](head1 *ListNode[T], head2 *ListNode[T]) *ListNode[T] {
	if head1 == nil || head2 == nil {
		return nil
	}
	headA, headB := head1, head2
	for headA != headB {
		if headA == nil {
			headA = head2
		} else {
			headA = headA.Next
		}
		if headB == nil {
			headB = head1
		} else {
			headB = headB.Next
		}
	}
	return headA
}
