package array

import "testing"

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
