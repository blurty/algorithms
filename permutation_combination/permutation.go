package permutation

import (
	"sort"
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
