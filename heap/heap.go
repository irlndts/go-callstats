package medianheap

import (
	"container/heap"
)

// MaxHeap is a max heap of ints.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap) Top() int {
	a := *h
	if len(a) != 0 {
		return int(a[0])
	}
	return 0
}

func NewMaxHeap() *MaxHeap {
	pq := &MaxHeap{}
	heap.Init(pq)
	return pq
}

// MinHeap is a max heap of ints.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap) Top() int {
	a := *h
	if len(a) != 0 {
		return int(a[0])
	}
	return 0
}

func NewMinHeap() *MinHeap {
	pq := &MinHeap{}
	heap.Init(pq)
	return pq
}