package hard100

type MedianFinderOfZeroToHundred struct {
	Arr    []int
	Length int
	n      int
}

func NewMedianFinderOfZeroToHunder() *MedianFinderOfZeroToHundred {
	m := new(MedianFinderOfZeroToHundred)
	m.Arr = make([]int, 101)
	m.Length = 101
	return m
}

func (m *MedianFinderOfZeroToHundred) FindKth(k int) int {
	cnt := 0
	for i := 0; i < k; i++ {
		cnt += m.Arr[i]
		if cnt > k {
			return i
		}
	}
	return -1
}

func (m *MedianFinderOfZeroToHundred) addNum(num int) {
	m.Arr[num]++
	m.n++
}

func (m *MedianFinderOfZeroToHundred) getNum() int {
	return m.n
}

func (m *MedianFinderOfZeroToHundred) FindMedian() int {
	if m.n%2 == 1 {
		return m.Arr[m.n/2+1]
	}
	return (m.Arr[m.n/2] + m.Arr[m.n/2+1]) / 2
}

type MedianArray struct {
	Arr []int
	n   int
}

func NewMedianArray() *MedianArray {
	ma := new(MedianArray)
	ma.Arr = make([]int, 0)
	ma.n = 0
	return ma
}

func (m *MedianArray) FindKth(k int) int {
	if k > len(m.Arr) {
		return -1
	}
	return m.Arr[k-1]
}

func (m *MedianArray) addNum(num int) {
	i := len(m.Arr) - 1
	for i >= 0 && m.Arr[i] > num {
		i--
	}
	m.Arr[i]++
}

func (m *MedianArray) getNum() int {
	return len(m.Arr)
}

type Median99Percent struct {
	Left   *MedianArray
	Right  *MedianArray
	Middle *MedianFinderOfZeroToHundred
}

func NewMedian99Percent() *Median99Percent {
	m := new(Median99Percent)
	m.Left = NewMedianArray()
	m.Right = NewMedianArray()
	m.Middle = NewMedianFinderOfZeroToHunder()
	return m
}

func (m99 *Median99Percent) addNum(num int) {
	if num < 0 {
		m99.Left.addNum(num)
	} else if num >= 0 && num <= 100 {
		m99.Middle.addNum(num)
	} else {
		m99.Right.addNum(num)
	}
}

func (m99 *Median99Percent) findMedian() int {
	middleSize := m99.Middle.getNum()
	leftSize := m99.Left.getNum()
	rightSize := m99.Right.getNum()

	n := middleSize + leftSize + rightSize

	if (n & 1) != 1 {
		k := n/2 + 1
		if k > leftSize && k < leftSize+middleSize {
			return m99.Middle.FindKth(k - leftSize)
		} else if k <= leftSize {
			return m99.Left.FindKth(k)
		} else {
			return m99.Right.FindKth(k - leftSize - middleSize)
		}
	} else {
		k := n / 2
		if k > leftSize && k+1 < leftSize+middleSize {
			k = k - leftSize
			return (m99.Middle.FindKth(k) + m99.Middle.FindKth(k+1)) / 2.0
		} else if k+1 <= leftSize {
			return (m99.Left.FindKth(k) + m99.Left.FindKth(k+1)) / 2.0
		} else if k == leftSize && middleSize > 0 {
			return (m99.Left.FindKth(k) + m99.Middle.FindKth(1)) / 2.0
		} else if k == leftSize && middleSize == 0 {
			return (m99.Left.FindKth(k) + m99.Right.FindKth(1)) / 2.0
		} else if k == leftSize+middleSize {
			return (m99.Middle.FindKth(k) + m99.Right.FindKth(1)) / 2
		} else {
			k = k - leftSize - rightSize
			return (m99.Right.FindKth(k) + m99.Right.FindKth(k+1)) / 2
		}
	}
	return -1
}
