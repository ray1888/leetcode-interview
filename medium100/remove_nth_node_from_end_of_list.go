package medium100

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := new(ListNode)
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
