package array

import (
	"strings"

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

// format the words such that each line has exactly maxWidth characters and is fully (left and right) justified.
func FullJustify(words []string, maxWidth int) []string {
	ret := make([]string, 0)
	var wordNum int    // number of words in each line
	var wordLength int // total length of all words in each line
	for i := 0; i < len(words); i += wordNum {
		for wordLength, wordNum = 0, 0; i+wordNum < len(words) && wordLength+len(words[i+wordNum]) <= maxWidth-wordNum; wordNum++ {
			wordLength += len(words[i+wordNum])
		}
		line := words[i]      // each line
		var averageSpaces int // average number of spaces between every two words in each line
		var extraSpaces int   // extra number of spaces between every two words in each line
		if wordNum > 1 {
			averageSpaces = (maxWidth - wordLength) / (wordNum - 1)
			extraSpaces = (maxWidth - wordLength) % (wordNum - 1)
		}
		for j := 0; j < wordNum-1; j++ {
			if i+wordNum >= len(words) {
				line += " "
			} else {
				if j < extraSpaces {
					line += strings.Repeat(" ", averageSpaces+1)
				} else {
					line += strings.Repeat(" ", averageSpaces)
				}
			}
			line += words[i+j+1]
		}
		line += strings.Repeat(" ", maxWidth-len(line))
		ret = append(ret, line)
	}
	return ret
}
