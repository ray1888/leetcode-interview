package medium100

import (
	"container/heap"

	"github.com/ray1888/leetcode-interview/datastructure"
)

type MedianFinder struct {
	minHeap datastructure.MinIntHeap
	maxHeap datastructure.MaxIntHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	mf := MedianFinder{}
	mf.minHeap = make(datastructure.MinIntHeap, 0)
	heap.Init(&mf.minHeap)
	mf.maxHeap = make(datastructure.MaxIntHeap, 0)
	heap.Init(&mf.maxHeap)
	return mf
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.minHeap, num)
	heap.Push(&this.maxHeap, heap.Pop(&this.minHeap))
	if this.maxHeap.Len() > this.minHeap.Len()+1 {
		heap.Push(&this.minHeap, heap.Pop(&this.maxHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() > this.minHeap.Len() {
		return (float64(this.maxHeap[0]) + float64(this.minHeap[0])) / 2
	} else {
		return float64(this.minHeap[0])
	}
}
