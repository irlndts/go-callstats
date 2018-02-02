package main

import (
	"container/heap"
	"fmt"
	"sort"
)

var data = []int{100, 102, 101, 110, 120, 115}

func main() {
	// Create a priority queue, put the metrics in it, and
	// establish the priority queue (heap) invariants.
	mq := make(MetricsQueue, len(data))
	for i, value := range data {
		mq[i] = &Metric{
			value:    value,
			priority: i,
			index:    i,
		}
	}
	heap.Init(&mq)

	/*
		// Insert a new metric and then modify its priority.
		metric := &Metric{
			value:    125,
			priority: len(data),
		}
		heap.Push(&mq, metric)
		mq.update(metric, metric.value, 7)
	*/
	fmt.Println(data)
	for k, v := range mq {
		fmt.Println(k, v)
	}
	sort.Sort(MetricsQueue(mq))
	//heap.Remove(&mq, 2)
	fmt.Println("####")
	for k, v := range mq {
		fmt.Println(k, v)
	}

	// Take the metrics out; they arrive in decreasing priority order.
	for mq.Len() > 0 {
		metric := heap.Pop(&mq).(*Metric)
		fmt.Printf("%.2d:%d ", metric.priority, metric.value)
	}
}
