package permutation

import (
	"reflect"
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
