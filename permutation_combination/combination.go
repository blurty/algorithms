package permutation

//
func doCombine(strs []string, str string) []string {
	if len(strs) == 0 {
		ret := make([]string, len(str))
		for i := 0; i < len(str); i++ {
			ret[i] = string(str[i])
		}
		return ret
	}

	ret := make([]string, len(strs)*len(str))
	for i := 0; i < len(strs); i++ {
		for j := 0; j < len(str); j++ {
			ret[i*len(str)+j] = strs[i] + string(str[j])
		}
	}
	return ret
}

// 对多个数组进行排列组合
// input: a:"abc"	b "def"
// output: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
func CombineString(strs ...string) []string {
	var ret []string
	for _, str := range strs {
		ret = doCombine(ret, str)
	}
	return ret
}
