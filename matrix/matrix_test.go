package matrix

import (
	"reflect"
	"testing"
)

type rotateTester struct {
	input, output [][]int
}

func TestRotate(t *testing.T) {
	tests := []rotateTester{
		rotateTester{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			[][]int{
				[]int{7, 4, 1},
				[]int{8, 5, 2},
				[]int{9, 6, 3},
			},
		},
	}
	for _, v := range tests {
		Rotate(v.input)
		for i := 0; i < len(v.input); i++ {
			if !reflect.DeepEqual(v.input[i], v.output[i]) {
				t.Errorf("want:%v, got:%v", v.output, v.input)
			}
		}
	}
}

func TestAntiRotate(t *testing.T) {
	tests := []rotateTester{
		rotateTester{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			[][]int{
				[]int{3, 6, 9},
				[]int{2, 5, 8},
				[]int{1, 4, 7},
			},
		},
	}
	for _, v := range tests {
		AntiRotate(v.input)
		for i := 0; i < len(v.input); i++ {
			if !reflect.DeepEqual(v.input[i], v.output[i]) {
				t.Errorf("want:%v, got:%v", v.output, v.input)
			}
		}
	}
}
