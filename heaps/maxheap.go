package heaps

import (
	"container/heap"
)

// MaxHeap is a max heap of ints.
type MaxHeap []int

// NewMaxHeap initialized MaxHeap
func NewMaxHeap() *MaxHeap {
	pq := &MaxHeap{}
	heap.Init(pq)
	return pq
}

// Index returns index of the passed element
func (h *MaxHeap) Index(x interface{}) int {
	for i := 0; i < h.Len(); i++ {
		if x == (*h)[i] {
			return i
		}
	}
	return -1
}

// Top returns the first value from the MaxHeap
func (h *MaxHeap) Top() int {
	a := *h
	if len(a) != 0 {
		return int(a[0])
	}
	return 0
}

// Methods bellow are inplementations of heap interface methods
// based on https://golang.org/pkg/container/heap/#Interface

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
