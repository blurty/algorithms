package permutation

import "fmt"

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

// return all possible combinations of k numbers out of 1 ... n
func CombineK(n int, k int) [][]int {
	if k < 1 || n < k {
		return nil
	}
	contain := make([]int, n)
	for i := 0; i < n; i++ {
		contain[i] = i + 1
	}
	if n == k {
		return [][]int{contain}
	}

	var ret [][]int
	for i := 0; i < n+1-k; i++ {
		if k == 1 {
			ret = append(ret, []int{contain[i]})
			continue
		}
		res := combineKRecursive(contain[i+1:], k-1)
		for _, v := range res {
			tmp := append([]int{contain[i]}, v...)
			ret = append(ret, tmp)
		}
	}
	return ret
}

func combineKRecursive(list []int, n int) [][]int {
	var ret [][]int
	if n < 1 {
		return nil
	}

	total := len(list)

	for i := 0; i < total+1-n; i++ {
		if n == 1 {
			ret = append(ret, []int{list[i]})
			continue
		}
		res := combineKRecursive(list[i+1:], n-1)
		fmt.Println("i:", i, "res:", res)
		for _, v := range res {
			tmp := append([]int{list[i]}, v...)
			ret = append(ret, tmp)
		}
	}
	return ret
}
