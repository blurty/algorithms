package maze

import (
	"testing"
)

func TestUniquePaths(t *testing.T) {
	type tester struct {
		x, y, z int
	}
	tests := []tester{
		tester{23, 12, 193536720},
	}
	for _, v := range tests {
		got := UniquePaths(v.x, v.y)
		if got != v.z {
			t.Errorf("want:%d, got:%d", v.z, got)
		}
	}
}

func TestUniquePathsWithObstacles(t *testing.T) {
	type tester struct {
		input  [][]int
		output int
	}
	tests := []tester{
		tester{
			input: [][]int{
				[]int{1},
			},
			output: 0,
		},
		tester{
			input:  [][]int{},
			output: 0,
		},
		tester{
			input: [][]int{
				[]int{0, 0, 0},
				[]int{0, 1, 0},
				[]int{0, 0, 0},
			},
			output: 2,
		},
	}
	for _, v := range tests {
		got := UniquePathsWithObstacles(v.input)
		if got != v.output {
			t.Errorf("want:%d, got:%d", v.output, got)
		}
	}
}

func TestMinPathSum(t *testing.T) {
	type tester struct {
		input  [][]int
		output int
	}
	tests := []tester{
		tester{
			input: [][]int{
				[]int{1},
			},
			output: 1,
		},
		tester{
			input:  [][]int{},
			output: 0,
		},
		tester{
			input: [][]int{
				[]int{1, 3, 1},
				[]int{1, 5, 1},
				[]int{4, 2, 1},
			},
			output: 7,
		},
	}
	for _, v := range tests {
		got := MinPathSum(v.input)
		if got != v.output {
			t.Errorf("want:%d, got:%d", v.output, got)
		}
	}
}
