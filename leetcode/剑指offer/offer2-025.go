package offer

import (
	. "learnGo/leetcode/dataStruct"
)

// https://leetcode-cn.com/problems/lMSNwu/
// 反转链表做法，不反转的话可以用栈来做
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1, l2 = reverse(l1), reverse(l2)
	dummyNode := &ListNode{Val: -1}
	cur := dummyNode
	carry := 0
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		cur = cur.Next
	}
	if carry != 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return reverse(dummyNode.Next)
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
