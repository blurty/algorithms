package permutation

import (
	"reflect"
	"sort"
	"testing"
)

type nextTester struct {
	input, output []int
}

func TestNextPermutation(t *testing.T) {
	tests := []nextTester{
		nextTester{[]int{1}, []int{1}},
		nextTester{[]int{1, 2}, []int{2, 1}},
		nextTester{[]int{1, 2, 3}, []int{1, 3, 2}},
		nextTester{[]int{1, 3, 2}, []int{2, 1, 3}},
		nextTester{[]int{3, 2, 1}, []int{1, 2, 3}},
		nextTester{[]int{1, 5, 1}, []int{5, 1, 1}},
		nextTester{[]int{2, 5, 3, 1}, []int{3, 1, 2, 5}},
		nextTester{[]int{2, 1, 7, 5, 3}, []int{2, 3, 1, 5, 7}},
		nextTester{[]int{2, 6, 7, 5, 3}, []int{2, 7, 3, 5, 6}},
	}

	for _, nums := range tests {
		NextPermutation(nums.input)
		if !reflect.DeepEqual(nums.input, nums.output) {
			t.Errorf("want:%v, got:%v", nums.input, nums.output)
		}
	}
}

type permuteTester struct {
	input  []int
	output [][]int
}

func compareSlices(a, b [][]int) bool {
	var f func(m [][]int) func(x, y int) bool
	f = func(m [][]int) func(x, y int) bool {
		return func(i, j int) bool {
			for k := 0; k < len(m[i]); k++ {
				if m[i][k] < m[j][k] {
					return true
				}
			}
			return false
		}
	}
	sort.SliceStable(a, f(a))
	sort.SliceStable(b, f(b))
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if b[i][j] != a[i][j] {
				return false
			}
		}
	}
	return true
}

func TestPermute(t *testing.T) {
	tests := []permuteTester{
		permuteTester{
			input: []int{1, 2, 3},
			output: [][]int{
				[]int{1, 2, 3},
				[]int{1, 3, 2},
				[]int{2, 1, 3},
				[]int{2, 3, 1},
				[]int{3, 1, 2},
				[]int{3, 2, 1},
			},
		},
	}

	for _, v := range tests {
		got := Permute(v.input)
		if !compareSlices(got, v.output) {
			t.Fatalf("want:%v, got:%v", v.output, got)
		}
	}
}

func TestPermuteUnique(t *testing.T) {
	tests := []permuteTester{
		permuteTester{
			input: []int{1, 1, 2},
			output: [][]int{
				[]int{1, 1, 2},
				[]int{1, 2, 1},
				[]int{2, 1, 1},
			},
		},
	}

	for _, v := range tests {
		got := PermuteUnique(v.input)
		if !compareSlices(got, v.output) {
			t.Fatalf("want:%v, got:%v", v.output, got)
		}
	}
}

func TestGetPermutation(t *testing.T) {
	type tester struct {
		input  int
		index  int
		output string
	}
	tests := []tester{
		tester{
			input:  3,
			index:  5,
			output: "312",
		},
		tester{
			input:  4,
			index:  13,
			output: "3124",
		},
		tester{
			input:  3,
			index:  2,
			output: "132",
		},
	}
	for _, v := range tests {
		got := GetPermutation(v.input, v.index)
		if got != v.output {
			t.Errorf("want:%s, got:%s", v.output, got)
		}
	}
}
