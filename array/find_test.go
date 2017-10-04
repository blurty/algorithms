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

func TestJump(t *testing.T) {
	type tester struct {
		input []int
		index int
	}
	tests := []tester{
		tester{[]int{0}, 0},
		tester{[]int{1, 2, 3}, 2},
		tester{[]int{2, 2, 3}, 1},
		tester{[]int{2, 3, 1}, 1},
	}
	for _, v := range tests {
		got := Jump(v.input)
		if got != v.index {
			t.Errorf("want:%d, got:%d", v.index, got)
		}
	}
}

func TestCanJump(t *testing.T) {
	type tester struct {
		input  []int
		result bool
	}
	tests := []tester{
		tester{[]int{0}, true},
		tester{[]int{2, 3, 1, 1, 4}, true},
		tester{[]int{3, 2, 1, 0, 4}, false},
		tester{[]int{0, 0}, false},
	}
	for _, v := range tests {
		got := CanJump(v.input)
		if got != v.result {
			t.Errorf("want:%d, got:%d", v.result, got)
		}
	}
}
