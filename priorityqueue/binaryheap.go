package priorityqueue

type binaryHeap struct {
	data []interface{}
	cmp  func(interface{}, interface{}) bool
}

func (bh binaryHeap) Len() int { return len(bh.data) }

func (bh binaryHeap) Less(i, j int) bool {
	return bh.cmp(bh.data[i], bh.data[j])
}

func (bh binaryHeap) Swap(i, j int) {
	bh.data[i], bh.data[j] = bh.data[j], bh.data[i]
}

func (bh binaryHeap) Push(x interface{}) {
	bh.data = append(bh.data, x)
}
func (bh binaryHeap) Pop() interface{} {
	n := len(bh.data)
	result := bh.data[n-1]
	bh.data = bh.data[:n-1]
	return result
}
