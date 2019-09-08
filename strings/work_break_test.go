package strings

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindAllSubStringIndex(t *testing.T) {
	testCases := []struct {
		s       string
		sub     string
		indexes []int
	}{
		{
			s:       "applepenapple",
			sub:     "apple",
			indexes: []int{0, 8},
		},
		{
			s:       "apple",
			sub:     "cat",
			indexes: []int{},
		},
	}
	for _, testCase := range testCases {
		gotResult := findAllSubStringIndex(testCase.s, testCase.sub)
		if !reflect.DeepEqual(testCase.indexes, gotResult) {
			t.Errorf("findAllSubStringIndex(%v,%v) want %v, got %v", testCase.s, testCase.sub, testCase.indexes, gotResult)
		}
	}
}

func TestWordBreak(t *testing.T) {
	testCases := []struct {
		s          string
		d          []string
		wantResult bool
	}{
		{
			s:          "applepenapple",
			d:          []string{"apple", "pen"},
			wantResult: true,
		},
		{
			s:          "leetcode",
			d:          []string{"leet", "code"},
			wantResult: true,
		},
		{
			s:          "catsandog",
			d:          []string{"cats", "dog", "sand", "and", "cat"},
			wantResult: false,
		},
		{
			s:          "catnamecat",
			d:          []string{"cat", "catname"},
			wantResult: true,
		},
	}
	for _, testCase := range testCases {
		gotResult := WordBreak(testCase.s, testCase.d)
		if gotResult != testCase.wantResult {
			t.Errorf("WordBreak(%v,%v) want %v, got %v", testCase.s, testCase.d, testCase.wantResult, gotResult)
		}
	}
}

func TestWordBreak2(t *testing.T) {
	testCases := []struct {
		s          string
		d          []string
		wantResult []string
	}{
		{
			s:          "applepenapple",
			d:          []string{"apple", "pen"},
			wantResult: []string{"apple pen apple"},
		},
		{
			s:          "leetcode",
			d:          []string{"leet", "code"},
			wantResult: []string{"leet code"},
		},
		{
			s:          "catsandog",
			d:          []string{"cats", "dog", "sand", "and", "cat"},
			wantResult: []string{},
		},
		{
			s:          "catnamecat",
			d:          []string{"cat", "catname", "name"},
			wantResult: []string{"catname cat", "cat name cat"},
		},
		{
			s:          "catsanddog",
			d:          []string{"cat", "cats", "and", "sand", "dog"},
			wantResult: []string{"cat sand dog", "cats and dog"},
		},
	}
	for _, testCase := range testCases {
		gotResult := WordBreak2(testCase.s, testCase.d)
		sort.Strings(gotResult)
		sort.Strings(testCase.wantResult)
		if !reflect.DeepEqual(gotResult, testCase.wantResult) {
			t.Errorf("WordBreak2(%v,%v) want %v, got %v", testCase.s, testCase.d, testCase.wantResult, gotResult)
		}
	}
}
