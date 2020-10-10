package easy100

import (
	"github.com/ray1888/leetcode-interview/datastructure"
)

func removeElements(head *datastructure.ListNode, val int) *datastructure.ListNode {
	if head == nil {
		return head
	}
	dummy := new(datastructure.ListNode)
	dummy.Next = head
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
