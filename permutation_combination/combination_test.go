package permutation

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombineString(t *testing.T) {
	got := CombineString("abc", "def")
	expect := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	sort.Strings(got)
	sort.Strings(expect)
	if !reflect.DeepEqual(got, expect) {
		t.Fatalf("got:%v, want:%v", got, expect)
	}
}

func TestCombineK(t *testing.T) {
	type tester struct {
		total int
		num   int
		want  [][]int
	}
	tests := []tester{
		tester{
			total: 4,
			num:   2,
			want: [][]int{
				[]int{1, 2},
				[]int{1, 3},
				[]int{1, 4},
				[]int{2, 3},
				[]int{2, 4},
				[]int{3, 4},
			},
		},
		tester{
			total: 1,
			num:   1,
			want:  [][]int{[]int{1}},
		},
	}
	for _, v := range tests {
		got := CombineK(v.total, v.num)
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("want:%v, got:%v", v.want, got)
		}
	}
}
