package main

import (
	"container/heap"

	"github.com/irlndts/go-callstats/heaps"
)

// Median contains information about metrics queue
type Median struct {
	Size    int // slider size
	MaxHeap *heaps.MaxHeap
	MinHeap *heaps.MinHeap
	Queue   *heaps.MetricsQueue
}

// InitMedianQueue returns inialized Median
func InitMedianQueue(size int) *Median {
	return &Median{
		Size:    size,
		MaxHeap: heaps.NewMaxHeap(),
		MinHeap: heaps.NewMinHeap(),
		Queue:   heaps.NewMetricsQueue(),
	}
}

// AddDelay adds a delay to the slice
func (m *Median) AddDelay(index, delay int) {
	// check if there stack is not full
	if m.MaxHeap.Len()+m.MinHeap.Len() >= m.Size {
		// remove the first element from queue to keep it in slider area
		m.remove()
	}

	// add value to the max heap
	heap.Push(m.Queue, &heaps.Metric{Value: delay, Priority: index})

	// balance max and mean heap
	heap.Push(m.MaxHeap, delay)
	heap.Push(m.MinHeap, m.MaxHeap.Top())
	heap.Pop(m.MaxHeap)

	if m.MaxHeap.Len() < m.MinHeap.Len() {
		heap.Push(m.MaxHeap, m.MinHeap.Top())
		heap.Pop(m.MinHeap)
	}
}

// GetMedian returns a median value of the slider slice
func (m Median) GetMedian() float64 {
	if m.MinHeap.Len()+m.MaxHeap.Len() <= 1 {
		return float64(-1)
	}
	if m.MinHeap.Len() < m.MaxHeap.Len() {
		return float64(m.MaxHeap.Top())
	}
	return float64(m.MaxHeap.Top()+m.MinHeap.Top()) * 0.5
}

// remove removes an element from Queue and
// also from one of two heaps
func (m *Median) remove() {
	item := heap.Pop(m.Queue).(*heaps.Metric)
	if item.Value < m.MaxHeap.Top() {
		index := m.MinHeap.Index(item.Value)
		heap.Remove(m.MinHeap, index)
	} else {
		index := m.MaxHeap.Index(item.Value)
		heap.Remove(m.MaxHeap, index)
	}
}
