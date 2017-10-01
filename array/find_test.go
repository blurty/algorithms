package array

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	type tester struct {
		input []int
		index int
	}
	tests := []tester{
		tester{[]int{1, 2, 0}, 3},
		tester{[]int{3, 4, -1, 1}, 2},
	}
	for _, v := range tests {
		got := FirstMissingPositive(v.input)
		if got != v.index {
			t.Errorf("want:%d, got:%d", v.index, got)
		}
	}
}
