package permutation

import (
	"sort"
	"strconv"
)

// Implement next permutation, which rearranges numbers into
// the lexicographically next greater permutation of numbers.
// 1,2,3 → 1,3,2
// 3,2,1 → 1,2,3
// 1,1,5 → 1,5,1
func NextPermutation(nums []int) {
	length := len(nums)
	if length < 2 {
		return
	} else if length == 2 {
		nums[0], nums[1] = nums[1], nums[0]
		return
	}
	var flag bool
	var exchange bool
	for i := length - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			for j := i + 1; j < length; j++ {
				if nums[j] <= nums[i] {
					nums[i], nums[j-1] = nums[j-1], nums[i]
					exchange = true
					break
				}
			}
			if !exchange {
				nums[i], nums[length-1] = nums[length-1], nums[i]
			}

			sort.Ints(nums[i+1:])
			flag = true
			break
		}
	}
	if !flag {
		sort.Ints(nums)
	}
}

// return all possible permutations of nums
func Permute(nums []int) [][]int {
	ret := [][]int{}

	if len(nums) == 1 {
		ret = append(ret, nums)
	} else {
		for i := 0; i < len(nums); i++ {
			rem := append([]int{}, nums...)
			// pick 1 element (i)
			// and get permutations for the rest
			rem = append(rem[:i], rem[i+1:]...)
			arr := Permute(rem)
			// prepend element (i)
			for j := 0; j < len(arr); j++ {
				//arr[j] = append(nums[i:i+1], arr[j]...)
				k := append([]int{nums[i]}, arr[j]...)
				ret = append(ret, k)
			}
		}
	}
	return ret
}

// nums is a sorted int slice
// return all possible unique permutations of nums
// just similar to Permute
func PermuteUnique(nums []int) [][]int {
	ret := [][]int{}

	if len(nums) == 1 {
		ret = append(ret, nums)
	} else {
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			rem := append([]int{}, nums...)
			// pick 1 element (i)
			// and get permutations for the rest
			rem = append(rem[:i], rem[i+1:]...)
			arr := PermuteUnique(rem)
			// prepend element (i)
			for j := 0; j < len(arr); j++ {
				//arr[j] = append(nums[i:i+1], arr[j]...)
				k := append([]int{nums[i]}, arr[j]...)
				ret = append(ret, k)
			}
		}
	}
	return ret
}

// calculate the factorial of n
func factorial(n int) int {
	if n < 1 {
		return 0
	} else if n == 1 {
		return n
	}
	return n * factorial(n-1)
}

// find the kth permutation sequence produced by nums recursively
// n is the number of nums
func restString(nums []string, n, k int) string {
	var ret string
	if n < 2 {
		return nums[0]
	}
	divider := factorial(n - 1)
	quotient, remainder := k/divider, k%divider
	index := quotient
	if remainder == 0 {
		index--
	}
	ret += nums[index]
	rem := append([]string{}, nums...)
	rem = append(rem[:index], rem[index+1:]...)
	ret += restString(rem, n-1, k-index*divider)

	return ret
}

// n represent the base of producing permutations
// k represent the index of permutations
// return the kth permutation sequence.
func GetPermutation(n int, k int) string {
	/*
		if n == 0 {
			return ""
		}
		data := make([]int, n)
		data[0] = 1
		iToS := []byte("123456789")
		for i := 1; i < n; i++ {
			data[i] = data[i-1] * i
		}
		ret := ""
		k--
		for i := n - 1; i >= 0; i-- {
			ret += string(iToS[k/data[i]])
			iToS = append(iToS[:k/data[i]], iToS[k/data[i]+1:]...)
			k = k % data[i]
		}
		return ret
	*/
	var ret string
	if k > factorial(n) || k < 1 || n > 9 {
		return ret
	}

	nums := make([]string, n)
	for i := 0; i < n; i++ {
		nums[i] = strconv.Itoa(i + 1)
	}
	return restString(nums, n, k)
}
