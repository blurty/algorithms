package permutation

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
