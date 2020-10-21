package easy100

import "github.com/ray1888/leetcode-interview/datastructure"

func getIntersectionNode(headA, headB *datastructure.ListNode) *datastructure.ListNode {
	if headB == nil {
		return headA
	} else if headA == nil {
		return headB
	}
	wA := headA
	wB := headB
	for wA != wB {
		if wA == nil {
			wA = headB
		} else {
			wA = wA.Next
		}
		if wB == nil {
			wB = headA
		} else {
			wB = wB.Next
		}
	}
	return wA
}
