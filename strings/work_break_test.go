package strings

import (
	"testing"
	"reflect"
)

func TestFindAllSubStringIndex(t *testing.T) {
	testCases := []struct {
		s string
		sub string
		indexes []int
	} {
		{
			s:"applepenapple",
			sub:"apple",
			indexes:[]int{0, 8},
		},
		{
			s:"apple",
			sub:"cat",
			indexes:[]int{},
		},
	}
	for _, testCase := range testCases {
		gotResult := findAllSubStringIndex(testCase.s, testCase.sub)
		if reflect.DeepEqual(testCase.indexes, gotResult) {
			t.Errorf("findAllSubStringIndex(%v,%v) want %v, got %v", testCase.s, testCase.sub, testCase.indexes, gotResult)
		}
	}
}

func TestWordBreak(t *testing.T) {
	testCases := []struct {
		s string
		d []string
		wantResult bool
	} {
		{
			s:"applepenapple",
			d:[]string{"apple", "pen"},
			wantResult:true,
		},
		{
			s:"leetcode",
			d:[]string{"leet", "code"},
			wantResult:true,
		},
		{
			s:"catsandog",
			d:[]string{"cats", "dog", "sand", "and", "cat"},
			wantResult:false,
		},
		{
			s:"catnamecat",
			d:[]string{"cat", "catname"},
			wantResult:true,
		},
	}
	for _, testCase := range testCases {
		gotResult := WordBreak(testCase.s, testCase.d)
		if gotResult != testCase.wantResult {
			t.Errorf("WordBreak(%v,%v) want %v, got %v", testCase.s, testCase.d, testCase.wantResult, gotResult)
		}
	}
}
