package priorityqueue

import "container/heap"

type PriorityQueue struct {
	bh binaryHeap
}

func New(compare func(interface{}, interface{}) bool) *PriorityQueue {
	return NewWithCapacity(compare, 32)
}

func NewWithCapacity(compare func(interface{}, interface{}) bool, capacity int) *PriorityQueue {
	pq := &PriorityQueue{
		binaryHeap{
			data: make([]interface{}, 0, capacity),
			cmp:  compare,
		},
	}
	return pq
}

func (pq PriorityQueue) Push(x interface{}) {
	heap.Push(&pq.bh, x)
}

func (pq PriorityQueue) Pop() interface{} {
	return heap.Pop(&pq.bh)
}

func (pq PriorityQueue) Size() int {
	return len(pq.bh.data)
}

func (pq PriorityQueue) Empty() bool {
	return len(pq.bh.data) == 0
}

func (pq PriorityQueue) Clear() {
	pq.bh = binaryHeap{
		data: make([]interface{}, 0, cap(pq.bh.data)),
		cmp:  pq.bh.cmp,
	}
}
