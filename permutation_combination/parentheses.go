package permutation

import (
	"container/list"
)

/*
[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/
func GenerateParenthesis(n int) []string {
	if n < 1 {
		return []string{}
	}
	ret := make([]string, 0)
	recursive(&ret, "", n, 0, 0)
	return ret
}

func recursive(list *[]string, str string, sum, left, right int) {
	if len(str) == sum*2 {
		*list = append(*list, str)
		return
	}
	if right < left {
		recursive(list, str+")", sum, left, right+1)
	}
	if left < sum {
		recursive(list, str+"(", sum, left+1, right)
	}
}

// check if parentheses is valid
func IsValid(s string) bool {
	stack := list.New()
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack.PushBack(s[i])
		} else {
			item := stack.Back()
			if item == nil {
				return false
			}
			want := item.Value.(byte)
			if s[i] == ')' {
				if want == '(' {
					stack.Remove(item)
				} else {
					return false
				}
			} else if s[i] == ']' {
				if want == '[' {
					stack.Remove(item)
				} else {
					return false
				}
			} else if s[i] == '}' {
				if want == '{' {
					stack.Remove(item)
				} else {
					return false
				}
			} else {
				panic("unknown parentheses")
			}
		}
	}
	if stack.Len() == 0 {
		return true
	}
	return false
}

// return the length of the longest valid parentheses substring
func LongestValidParentheses(s string) int {
	length := len(s)
	for i:=length; i>0; i-- {
		for j:=0; j<=length-i; j++ {
			if s[j] == '(' && s[j+i-1]==')' && IsBracketClosed(s[j:j+i]) {
				return i
			}
		}
	}
	return 0
}

// check if bracket is closed
func IsBracketClosed(s string) bool {
	left, right := 0, 0
	for i:=0; i<len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if right > left {
			return false
		}
	}
	if right != left {
		return false
	}
	return true
}