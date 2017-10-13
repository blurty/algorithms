package array

// Given an array of strings, group anagrams together.
// All inputs should be in lower-case.
func GroupAnagrams(strs []string) [][]string {
	ret := make([][]string, 0)
	for i := 0; i < len(strs); i++ {
		flag := false
		for j := 0; j < len(ret); j++ {
			for k := 0; k < len(ret[j]); k++ {
				if same(strs[i], ret[j][k]) {
					ret[j] = append(ret[j], strs[i])
					flag = true
					break
				}
			}
		}
		if !flag {
			ret = append(ret, []string{strs[i]})
		}
	}
	return ret
}

func same(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	visit := make([]bool, len(b))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if !visit[j] && a[i] == b[j] {
				visit[j] = true
				break
			}
		}
	}
	for i := 0; i < len(visit); i++ {
		if !visit[i] {
			return false
		}
	}
	return true
}
