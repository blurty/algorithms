package sort

import (
	"reflect"
	"testing"
)

type sortTester struct {
	input, output []int
}

var sortTests = []sortTester{
	sortTester{[]int{2, 3, 2, 5, 1}, []int{1, 2, 2, 3, 5}},
}

func TestBubbleSort(t *testing.T) {
	for _, v := range sortTests {
		BubbleSort(v.input)
		if !reflect.DeepEqual(v.input, v.output) {
			t.Errorf("want:%v, got:%v", v.output, v.input)
		}
	}
}

func TestQuickSort(t *testing.T) {
	for _, v := range sortTests {
		QuickSort(v.input)
		if !reflect.DeepEqual(v.input, v.output) {
			t.Errorf("want:%v, got:%v", v.output, v.input)
		}
	}
}