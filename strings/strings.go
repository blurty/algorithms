package strings

import (
	"sort"
)

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

// concatenate some strings
// such as: ["hello", "world"] concat to ["helloworld", "worldhello"]
func ConcatStrings(words []string) []string {
	if len(words) < 2 {
		return words
	}
	// for avoid duplicate
	sort.Strings(words)

	ret := make([]string, 0)
	for i := 0; i < len(words); i++ {
		// strip duplicate string
		if i > 0 && words[i] == words[i-1] {
			continue
		}
		oldWord := words[i]
		newWords := make([]string, 0)
		newWords = append(newWords, words[:i]...)
		newWords = append(newWords, words[i+1:]...)
		newRet := ConcatStrings(newWords)
		for _, v := range newRet {
			ret = append(ret, oldWord+v)
		}
	}
	return ret
}

// nums1 and num2 both are representation of non-negative integers
// both's length are < 110
// both contain only digits 0-9
// both don't contain any leading zero
func Multiply(num1 string, num2 string) string {
	sum := make([]byte, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		var carry, bit byte
		for j := len(num2) - 1; j >= 0; j-- {
			bit = byte(num1[i]-'0')*byte(num2[j]-'0') + carry
			if sum[i+j+1] != 0 {
				bit += byte(sum[i+j+1] - '0')
			}
			sum[i+j+1] = bit%10 + '0'
			carry = bit / 10
		}
		if carry != 0 {
			sum[i] = carry + '0'
		}
	}

	for i := 0; i < len(sum); i++ {
		if sum[i] != 0 && sum[i]-'0' > 0 {
			return string(sum[i:])
		}
	}
	return "0"
}

// return sum of two binary digits represented as string
// such as "11" + "1" = "100"
func AddBinary(a string, b string) string {
	lenA, lenB := len(a), len(b)
	if lenA > lenB {
		return AddBinary(b, a)
	}
	interval := lenB - lenA
	var ret string
	var carry bool
	for i := lenA - 1; i >= 0; i-- {
		if !carry {
			if a[i] == '1' && b[i+interval] == '1' {
				ret = "0" + ret
				carry = true
			} else if a[i] == '1' && b[i+interval] == '0' || (a[i] == '0' && b[i+interval] == '1') {
				ret = "1" + ret
				carry = false
			} else if a[i] == '0' && b[i+interval] == '0' {
				ret = "0" + ret
				carry = false
			} else {
				panic("wrong input")
			}
		} else {
			if a[i] == '1' && b[i+interval] == '1' {
				ret = "1" + ret
				carry = true
			} else if a[i] == '1' && b[i+interval] == '0' || (a[i] == '0' && b[i+interval] == '1') {
				ret = "0" + ret
				carry = true
			} else if a[i] == '0' && b[i+interval] == '0' {
				ret = "1" + ret
				carry = false
			} else {
				panic("wrong input")
			}
		}
	}
	for i := interval - 1; i >= 0; i-- {
		if !carry {
			ret = string([]rune(b)[0:i+1]) + ret
			carry = false
			break
		} else {
			if b[i] == '1' {
				ret = "0" + ret
				carry = true
			} else if b[i] == '0' {
				ret = "1" + ret
				carry = false
			} else {
				panic("wrong input")
			}
		}
	}
	if carry {
		ret = "1" + ret
	}
	return ret
}
