package array

import (
	"reflect"
	"sort"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	type tester struct {
		input  []int
		output int
	}
	tests := []tester{
		tester{
			input:  []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			output: 6,
		},
	}
	for _, v := range tests {
		got := MaxSubArray(v.input)
		if got != v.output {
			t.Errorf("want:%d, got:%d", v.output, got)
		}
	}
}

func TestMergeInterval(t *testing.T) {
	type tester struct {
		input, output []Interval
	}
	tests := []tester{
		tester{
			input: []Interval{
				Interval{1, 3},
				Interval{2, 6},
				Interval{8, 10},
				Interval{15, 18},
			},
			output: []Interval{
				Interval{1, 6},
				Interval{8, 10},
				Interval{15, 18},
			},
		},
		tester{
			input: []Interval{
				Interval{2, 3},
				Interval{4, 5},
				Interval{6, 7},
				Interval{8, 9},
				Interval{1, 10},
			},
			output: []Interval{
				Interval{1, 10},
			},
		},
	}
	for _, v := range tests {
		got := MergeInterval(v.input)
		sort.SliceStable(got, func(i, j int) bool {
			return got[i].Start < got[j].Start
		})
		sort.SliceStable(v.output, func(i, j int) bool {
			return v.output[i].Start < v.output[j].Start
		})
		if !reflect.DeepEqual(got, v.output) {
			t.Errorf("want:%v, got:%v", v.output, got)
		}
	}
}

func TestInsertInterval(t *testing.T) {
	type tester struct {
		input  []Interval
		new    Interval
		output []Interval
	}
	tests := []tester{
		tester{
			input: []Interval{
				Interval{1, 3},
				Interval{6, 9},
			},
			new: Interval{2, 5},
			output: []Interval{
				Interval{1, 5},
				Interval{6, 9},
			},
		},
		tester{
			input: []Interval{
				Interval{1, 2},
				Interval{3, 5},
				Interval{6, 7},
				Interval{8, 10},
				Interval{12, 16},
			},
			new: Interval{4, 9},
			output: []Interval{
				Interval{1, 2},
				Interval{3, 10},
				Interval{12, 16},
			},
		},
	}
	for _, v := range tests {
		got := InsertInterval(v.input, v.new)
		sort.SliceStable(got, func(i, j int) bool {
			return got[i].Start < got[j].Start
		})
		sort.SliceStable(v.output, func(i, j int) bool {
			return v.output[i].Start < v.output[j].Start
		})
		if !reflect.DeepEqual(got, v.output) {
			t.Errorf("want:%v, got:%v", v.output, got)
		}
	}
}
