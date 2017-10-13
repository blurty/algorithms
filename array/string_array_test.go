package array

import (
	"testing"
)

func compareStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	visit := make([]bool, len(b))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if !visit[j] && b[j] == a[i] {
				visit[j] = true
				break
			}
		}
	}
	for i := 0; i < len(visit); i++ {
		if !visit[i] {
			return false
		}
	}
	return true
}

func compareStringSlices2(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	visit := make([]bool, len(b))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if !visit[j] && compareStringSlice(a[i], b[j]) {
				visit[j] = true
				break
			}
		}
	}
	for i := 0; i < len(visit); i++ {
		if !visit[i] {
			return false
		}
	}
	return true
}

func TestGroupAnagrams(t *testing.T) {
	type groupTester struct {
		input  []string
		output [][]string
	}
	tests := []groupTester{
		groupTester{
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			output: [][]string{
				[]string{"ate", "eat", "tea"},
				[]string{"nat", "tan"},
				[]string{"bat"},
			},
		},
	}
	for _, v := range tests {
		got := GroupAnagrams(v.input)
		if !compareStringSlices2(got, v.output) {
			t.Errorf("want:%v, got:%v", v.output, got)
		}
	}
}
