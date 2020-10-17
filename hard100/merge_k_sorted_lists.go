package hard100

import (
	"container/heap"

	"github.com/ray1888/leetcode-interview/datastructure"
)

type listNodeHeap []*datastructure.ListNode

func (LNH listNodeHeap) Len() int           { return len(LNH) }
func (LNH listNodeHeap) Less(i, j int) bool { return LNH[i].Val < LNH[j].Val }
func (LNH *listNodeHeap) Swap(i, j int) {
	(*LNH)[i], (*LNH)[j] = (*LNH)[j], (*LNH)[i]
}

func (LNH *listNodeHeap) Push(node interface{}) {
	*LNH = append(*LNH, node.(*datastructure.ListNode))
}

func (LNH *listNodeHeap) Pop() interface{} {
	tmp := (*LNH)[len(*LNH)-1]
	*LNH = (*LNH)[:len(*LNH)-1]
	return tmp
}

// 使用堆来进行排序，堆的大小为len(lists)的数量
func mergeKLists(lists []*datastructure.ListNode) *datastructure.ListNode {
	if len(lists) == 0 {
		return nil
	}
	dummy := new(datastructure.ListNode)

	nodeHeap := listNodeHeap{}
	heap.Init(&nodeHeap)
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(&nodeHeap, lists[i])
		}
	}

	work := dummy
	for nodeHeap.Len() > 0 {
		node := heap.Pop(&nodeHeap).(*datastructure.ListNode)
		work.Next = node
		work = work.Next
		if node.Next != nil {
			nodeHeap.Push(node.Next)
			heap.Fix(&nodeHeap, len(nodeHeap)-1)
		}
	}

	return dummy.Next
}
