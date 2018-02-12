package main

import (
	"container/heap"
	"fmt"

	h "github.com/irlndts/go-callstats/heap"
)

// Median ...
type Median struct {
	MaxHeap *h.MaxHeap
	MinHeap *h.MinHeap
	Size    int
}

func InitMedianQueue(size int) *Median {
	return &Median{
		Size:    size,
		MaxHeap: h.NewMaxHeap(),
		MinHeap: h.NewMinHeap(),
	}
}

// AddDelay adds a delay to the slice
func (m *Median) AddDelay(delay int) {
	heap.Push(m.MaxHeap, delay)
	//heap.Push(m.MinHeap, int(m.MaxHeap.Top()))
	//heap.Pop(m.MaxHeap)

	if m.MaxHeap.Len() > m.MinHeap.Len() {
		heap.Push(m.MinHeap, heap.Pop(m.MaxHeap))
	}
	fmt.Println("Min Heap", m.MinHeap)
	fmt.Println("Max Heap", m.MaxHeap)
}

// GetMedian returns a median value of the slice
func (m Median) GetMedian() float64 {
	if m.MinHeap.Len() > m.MaxHeap.Len() {
		return float64(m.MinHeap.Top())
	}

	return float64(m.MaxHeap.Top()+m.MinHeap.Top()) * 0.5
}
