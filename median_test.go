package main

import (
	"reflect"
	"testing"
)

func TestGetMedian(t *testing.T) {
	data := []int{100, 102, 101, 110, 120, 115}
	tests := map[string]struct {
		slider int
		input  []int
		want   []float64
	}{
		"Slider size 1": {1, data, []float64{-1, -1, -1, -1, -1, -1}},
		"Slider size 3": {3, data, []float64{-1, 101, 101, 102, 110, 115}},
		"Slider size 6": {6, data, []float64{-1, 101, 101, 101.5, 102, 106}},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			median := InitMedianQueue(tt.slider)
			got := make([]float64, 0, len(tt.input))
			for _, metric := range tt.input {
				median.AddDelay(metric)
				got = append(got, median.GetMedian())
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("medians got=%#v; want=%#v;", got, tt.want)
			}
		})
	}
}
