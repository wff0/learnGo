package offer

import (
	. "learnGo/leetcode/dataStruct"
)

// https://leetcode-cn.com/problems/LGjMqU/
// 主要用到了快慢指针，反转链表，以及合并链表
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleOfLis(head)
	l1 := head
	l2 := reverse(mid.Next)
	mid.Next = nil
	mergeList(l1, l2)
	return
}

func middleOfLis(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

//func reverse(head *ListNode) *ListNode {
//	var pre *ListNode
//	for head != nil {
//		next := head.Next
//		head.Next = pre
//		pre = head
//		head = next
//	}
//	return pre
//}

func mergeList(l1, l2 *ListNode) {
	for l1 != nil && l2 != nil {
		l1Next := l1.Next
		l2Next := l2.Next

		l1.Next = l2
		l1 = l1Next

		l2.Next = l1
		l2 = l2Next
	}
}
