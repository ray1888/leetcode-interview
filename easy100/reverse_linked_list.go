package easy100

import (
	"github.com/ray1888/leetcode-interview/datastructure"
)

func reverseList(head *datastructure.ListNode) *datastructure.ListNode {
	var cur *datastructure.ListNode
	var pre *datastructure.ListNode
	var next *datastructure.ListNode

	cur = head

	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
