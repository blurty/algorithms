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

func findAllSubStringIndex(s string, sub string) (result []int) {
	i := 0
	cs := s
	for i < len(s) {
		cs = cs[i:]
		idx := ostrings.Index(cs, sub)
		if idx != -1 {
			result = append(result, idx)
			i = idx + len(sub)
		} else {
			break
		}
	}
	return
}
