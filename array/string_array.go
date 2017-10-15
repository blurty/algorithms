package array

import (
	"github.com/blurtheart/mylib/strutils"
)

// Given an array of strings, group anagrams together.
// All inputs should be in lower-case.
func GroupAnagrams(strs []string) [][]string {
	cn := make(map[string][]string, 0)
	for i := 0; i < len(strs); i++ {
		charArray := strutils.SortString(strs[i])
		if _, ok := cn[charArray]; ok {
			cn[charArray] = append(cn[charArray], strs[i])
		} else {
			cn[charArray] = []string{strs[i]}
		}
	}
	ret := make([][]string, len(cn))
	index := 0
	for _, v := range cn {
		ret[index] = v
		index++
	}
	return ret
}
