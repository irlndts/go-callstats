package main

import (
	"container/heap"

	"github.com/irlndts/go-callstats/heaps"
)

// Median contains information about metrics queue
type Median struct {
	Size    int // slider size
	maxHeap *heaps.MaxHeap
	minHeap *heaps.MinHeap
	queue   *heaps.MetricsQueue
}

// InitMedianQueue returns inialized Median
func InitMedianQueue(size int) *Median {
	return &Median{
		Size:    size,
		maxHeap: heaps.NewMaxHeap(),
		minHeap: heaps.NewMinHeap(),
		queue:   heaps.NewMetricsQueue(),
	}
}

// AddDelay adds a delay to the slice
func (m *Median) AddDelay(delay int) {
	// check if there stack is not full
	if m.maxHeap.Len()+m.minHeap.Len() >= m.Size {
		// remove the first element from queue to keep it in slider area
		m.remove()
	}

	// add value to the metrics queue
	heap.Push(m.queue, &heaps.Metric{Value: delay, Priority: m.queue.Len()})

	// add value to the max heap
	heap.Push(m.maxHeap, delay)
	heap.Push(m.minHeap, m.maxHeap.Top())
	heap.Pop(m.maxHeap)

	// balance max and min heap
	if m.maxHeap.Len() < m.minHeap.Len() {
		heap.Push(m.maxHeap, m.minHeap.Top())
		heap.Pop(m.minHeap)
	}
}

// GetMedian returns a median value of the slider slice
func (m Median) GetMedian() float64 {
	if m.minHeap.Len()+m.maxHeap.Len() <= 1 {
		return float64(-1)
	}
	if m.minHeap.Len() < m.maxHeap.Len() {
		return float64(m.maxHeap.Top())
	}
	return float64(m.maxHeap.Top()+m.minHeap.Top()) * 0.5
}

// remove removes an element from queue and
// also from one of two heaps
func (m *Median) remove() {
	item := heap.Pop(m.queue).(*heaps.Metric)
	if item.Value < m.maxHeap.Top() {
		index := m.minHeap.Index(item.Value)
		heap.Remove(m.minHeap, index)
	} else {
		index := m.maxHeap.Index(item.Value)
		heap.Remove(m.maxHeap, index)
	}
}
