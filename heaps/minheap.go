package heaps

import (
	"container/heap"
)

// MinHeap is a min heap of ints.
type MinHeap []int

// NewMinHeap initialized MinHeap
func NewMinHeap() *MinHeap {
	pq := &MinHeap{}
	heap.Init(pq)
	return pq
}

// Top returns the first value from the MaxHeap
func (h *MinHeap) Top() int {
	a := *h
	if len(a) != 0 {
		return int(a[0])
	}
	return 0
}

// // Index returns index of the passed element
func (h *MinHeap) Index(x interface{}) int {
	for i := 0; i < h.Len(); i++ {
		if x == (*h)[i] {
			return i
		}
	}
	return -1
}

// Methods bellow are inplementations of heap interface methods
// based on https://golang.org/pkg/container/heap/#Interface

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
