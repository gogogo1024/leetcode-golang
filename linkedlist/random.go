package linkedlist

import "math/rand"

//给你一个单链表，随机选择链表的一个节点，并返回相应的节点值。每个节点被选中的 概率一样。
//实现 Solution 类：
//Solution(ListNode head) 使用整数数组初始化对象。
//int getRandom() 从链表中随机选择一个节点并返回该节点的值。链表中所有节点被选中的概率相等。

// Solution 蓄水池抽样算法
type Solution[T Ordered] struct {
	head *ListNode[T]
}

func (s *Solution[T]) getRandom() (ans T) {
	if s.head == nil {
		var zero T // 返回 T 类型的零值
		return zero
	}
	for node, i := s.head, 1; node != nil; node = node.Next {
		if rand.Intn(i) == 0 {
			ans = node.Val
		}
		i++
	}
	return ans
}
