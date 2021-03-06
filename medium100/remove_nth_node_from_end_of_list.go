package medium100

import "github.com/ray1888/leetcode-interview/datastructure"

func removeNthFromEnd(head *datastructure.ListNode, n int) *datastructure.ListNode {
	if head == nil {
		return nil
	}
	dummy := new(datastructure.ListNode)
	dummy.Next = head
	fast := dummy
	slow := dummy
	for n > 0 && fast.Next != nil {
		fast = fast.Next
		n--
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	slow = slow.Next
	return dummy.Next
}
