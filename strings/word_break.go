package strings

import (
	ostrings "strings" // official strings library
)

/*
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words,
determine if s can be segmented into a space-separated sequence of one or more dictionary words.
Note:
- The same word in the dictionary may be reused multiple times in the segmentation.
- You may assume the dictionary does not contain duplicate words.
*/
func WordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i, _ := range s {
		for j := i; j < len(s); j++ {
			if dp[i] && inSliceString(s[i:j+1], wordDict) {
				dp[j+1] = true
			}
		}
	}
	return dp[len(s)]
}

func inSliceString(s string, ss []string) bool {
	for _, item := range ss {
		if s == item {
			return true
		}
	}
	return false
}

/*
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words,
add spaces in s to construct a sentence where each word is a valid dictionary word. Return
all such possible sentences.
Note:
- The same word in the dictionary may be reused multiple times in the segmentation.
- You may assume the dictionary does not contain duplicate words.
*/
func WordBreak2(s string, wordDict []string) []string {
	results := make([]string, 0)
	for i := 0; i <= len(s); i++ {
		if inSliceString(s[:i], wordDict) {
			if i == len(s) {
				results = append(results, s)
				continue
			}
			subResults := WordBreak2(s[i:], wordDict)
			for _, subResult := range subResults {
				results = append(results, s[:i]+" "+subResult)
			}
		}
	}
	return results
}

func findAllSubStringIndex(s string, sub string) []int {
	i := 0
	cs := s
	result := make([]int, 0)
	for i < len(s) {
		cs = cs[i:]
		idx := ostrings.Index(cs, sub)
		if idx != -1 {
			num := idx
			if len(result) > 0 {
				num += result[len(result)-1] + len(sub)
			}
			result = append(result, num)
			i = idx + len(sub)
		} else {
			break
		}
	}
	return result
}
