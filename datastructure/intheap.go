package datastructure

type MinIntHeap []int

func (mh MinIntHeap) Len() int           { return len(mh) }
func (mh MinIntHeap) Less(i, j int) bool { return mh[i] < mh[j] }
func (mh *MinIntHeap) Swap(i, j int) {
	(*mh)[i], (*mh)[j] = (*mh)[j], (*mh)[i]
}

func (mh *MinIntHeap) Push(node interface{}) {
	*mh = append(*mh, node.(int))
}

func (mh *MinIntHeap) Pop() interface{} {
	tmp := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return tmp
}

type MaxIntHeap []int

func (mh MaxIntHeap) Len() int           { return len(mh) }
func (mh MaxIntHeap) Less(i, j int) bool { return mh[i] > mh[j] }
func (mh *MaxIntHeap) Swap(i, j int) {
	(*mh)[i], (*mh)[j] = (*mh)[j], (*mh)[i]
}

func (mh *MaxIntHeap) Push(node interface{}) {
	*mh = append(*mh, node.(int))
}

func (mh *MaxIntHeap) Pop() interface{} {
	tmp := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return tmp
}
