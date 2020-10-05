package easy100

import (
	"github.com/ray1888/leetcode-interview/datastructure"
)

func mergeTwoLists(l1 *datastructure.ListNode, l2 *datastructure.ListNode) *datastructure.ListNode {
	var head = new(datastructure.ListNode)
	cur := head
	for l1 != nil && l2 != nil {
		newNode := new(datastructure.ListNode)
		if l1.Val > l2.Val {
			newNode.Val = l2.Val
			l2 = l2.Next
		} else {
			newNode.Val = l1.Val
			l1 = l1.Next
		}
		cur.Next = newNode
		cur = cur.Next
	}
	if l1 != nil {
		for l1 != nil {
			cur.Next = l1
			l1 = l1.Next
			cur = cur.Next
		}
	}
	if l2 != nil {
		for l2 != nil {
			cur.Next = l2
			l2 = l2.Next
			cur = cur.Next
		}
	}
	return head
}
