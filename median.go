package main

import (
	"container/heap"

	h "github.com/irlndts/go-callstats/heap"
)

// Median ...
type Median struct {
	Size    int
	MaxHeap *h.MaxHeap
	MinHeap *h.MinHeap
	Queue   *MetricsQueue
}

func InitMedianQueue(size int) *Median {
	return &Median{
		Size:    size,
		MaxHeap: h.NewMaxHeap(),
		MinHeap: h.NewMinHeap(),
		Queue:   NewMetricsQueue(),
	}
}

// AddDelay adds a delay to the slice
func (m *Median) AddDelay(index, delay int) {
	if m.MaxHeap.Len()+m.MinHeap.Len() >= m.Size {
		// remove the first element from queue to keep it in slider area
		m.remove()
	}

	heap.Push(m.Queue, &Metric{Value: delay, Priority: index})

	heap.Push(m.MaxHeap, delay)
	heap.Push(m.MinHeap, m.MaxHeap.Top())
	heap.Pop(m.MaxHeap)

	if m.MaxHeap.Len() < m.MinHeap.Len() {
		heap.Push(m.MaxHeap, m.MinHeap.Top())
		heap.Pop(m.MinHeap)
	}
}

func (m *Median) remove() {
	item := heap.Pop(m.Queue).(*Metric)
	if item.Value < m.MaxHeap.Top() {
		index := m.MinHeap.Index(item.Value)
		heap.Remove(m.MinHeap, index)
	} else {
		index := m.MaxHeap.Index(item.Value)
		heap.Remove(m.MaxHeap, index)
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
