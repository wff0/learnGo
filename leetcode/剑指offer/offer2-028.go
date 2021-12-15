package offer

import (
	. "learnGo/leetcode/dataStruct"
)

// https://leetcode-cn.com/problems/Qv1Da2/
func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	dfs(root)
	return root
}

func dfs(cur *Node) *Node {
	pre := cur.Prev
	for cur != nil {
		if cur.Child != nil {
			next := cur.Next
			cur.Next = cur.Child
			cur.Child.Prev = cur
			cur.Child = nil
			ret := dfs(cur.Next)
			ret.Next = next
			if next != nil {
				next.Prev = ret
			}
			pre = ret
			cur = next
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return pre
}
