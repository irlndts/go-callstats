package main

import (
	"fmt"
)

var data = []int{100, 102, 101, 110, 120, 115}

func main() {
	median := InitMedianQueue(3)

	for _, d := range data {
		median.AddDelay(d)
		fmt.Println("MEDIAN: ", median.GetMedian())
	}
}
