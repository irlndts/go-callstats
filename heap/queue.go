package main

import (
	"container/heap"
)

type Metric struct {
	Value    int // The value of the metric.
	Priority int
	index    int // The index of the metric in the heap.
}

// MetricsQueue implements heap.Interface and holds Metrics.
type MetricsQueue []*Metric

func (mq MetricsQueue) Len() int { return len(mq) }
func (mq MetricsQueue) Less(i, j int) bool {
	return mq[i].Priority < mq[j].Priority
}
func (mq MetricsQueue) Swap(i, j int) {
	mq[i], mq[j] = mq[j], mq[i]
	mq[i].index = i
	mq[j].index = j
}

func (mq *MetricsQueue) Push(x interface{}) {
	n := len(*mq)
	metric := x.(*Metric)
	metric.index = n
	*mq = append(*mq, metric)
}

func (mq *MetricsQueue) Pop() interface{} {
	old := *mq
	n := len(old)
	metric := old[n-1]
	metric.index = -1 // for safety
	*mq = old[0 : n-1]
	return metric
}

func (mq *MetricsQueue) Remove(x interface{}, i int) interface{} {
	old := *mq
	n := len(old)
	metric := old[n-1]
	metric.index = -1 // for safety
	*mq = old[0 : n-1]
	return metric
}

func NewMetricsQueue() *MetricsQueue {
	mq := &MetricsQueue{}
	heap.Init(mq)
	return mq
}
