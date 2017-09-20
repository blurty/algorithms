package strings

// returns the index of the first occurrence of needle in haystack,
// or -1 if needle is not part of haystack.
func StrStr(haystack string, needle string) int {
	hayLen, neeLen := len(haystack), len(needle)
outloop:
	for i := 0; i <= hayLen-neeLen; i++ {
		for j := 0; j < neeLen; j++ {
			if haystack[i+j] != needle[j] {
				continue outloop
			}
		}
		return i
	}

	return -1
}
