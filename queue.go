package main

import "container/heap"

// Metric is a single element of metrics queue
type Metric struct {
	value    int // The value of the metric.
	priority int // The priority of the metric in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the metric in the heap.
}

// MetricsQueue implements heap.Interface and holds Metrics.
type MetricsQueue []*Metric

func (mq MetricsQueue) Len() int { return len(mq) }
func (mq MetricsQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return mq[i].value < mq[j].value
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

// update modifies the priority and value of an Metric in the queue.
func (pq *MetricsQueue) update(metric *Metric, value, priority int) {
	metric.value = value
	metric.priority = priority
	heap.Fix(pq, metric.index)
}
