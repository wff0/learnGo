package offer

import (
	. "learnGo/leetcode/dataStruct"
)

// https://leetcode-cn.com/problems/4ueAj6/submissions/
func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		cur := &Node{Val: x}
		cur.Next = cur
		return cur
	}
	cur := aNode
	for cur.Next != aNode {
		if cur.Next.Val < cur.Val {
			if cur.Next.Val >= x {
				break
			} else if cur.Val <= x {
				break
			}
		}
		if cur.Val <= x && cur.Next.Val >= x {
			break
		}
		cur = cur.Next
	}
	cur.Next = &Node{Val: x, Prev: cur.Next}
	return aNode
}
