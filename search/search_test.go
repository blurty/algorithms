package search

import (
	"testing"
)

type searchTester struct {
	nums []int
	target, index int
}

func TestBinarySearch(t *testing.T) {
	tests := []searchTester{
		searchTester{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2, 1},
	}

	for _, v := range tests {
		got := BinarySearch(v.nums, v.target)
		if got != v.index {
			t.Errorf("want:%d, got:%d", v.index, got)
		}
	}
}