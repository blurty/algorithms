package permutation

import (
	"reflect"
	"sort"
	"testing"
)

type tester struct {
	flag   int
	expect []string
}

func checkEqual(got, expect []string) bool {
	gotLength, expectLength := len(got), len(expect)
	if gotLength != expectLength {
		return false
	}
	sort.Strings(got)
	sort.Strings(expect)
	return reflect.DeepEqual(got, expect)
}

func TestGenerateParenthesis(t *testing.T) {
	var tests = []tester{
		{3, []string{"((()))", "(()())", "()()()", "(())()", "()(())"}},
		{1, []string{"()"}},
		{0, []string{}},
	}
	for i := 0; i < len(tests); i++ {
		got := GenerateParenthesis(tests[i].flag)
		result := checkEqual(got, tests[i].expect)
		if !result {
			t.Fatalf("got:%v, want:%v", got, tests[i].expect)
		}
	}
}

func TestLongestValidParentheses(t *testing.T)  {
	type tester struct {
		intput string
		output int
	}
	tests := []tester{
		tester{"()))", 2},
		tester{"()())", 4},
		tester{")((", 0},
		tester{"(())(()", 4},
		tester{"(", 0},
	}
	for _, v := range tests {
		got := LongestValidParentheses(v.intput)
		if got != v.output {
			t.Errorf("want:%v, got:%v", v.output, got)
		}
	}
}