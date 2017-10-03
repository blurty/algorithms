package array

// find the first missing positive integer
func FirstMissingPositive(nums []int) int {
	length := len(nums)
	for i := 0; i < length; i++ {
		digit := nums[i]
		for digit > 0 && digit <= length && digit != nums[digit-1] {
			nums[i], nums[digit-1] = nums[digit-1], nums[i]
			digit = nums[i]
		}
	}
	for i := 0; i < length; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return length + 1
}

// the steps jump from the first item to the final item
// jump rule:
// 		Each element in the array represents maximum jump length at that position.
func Jump(nums []int) int {
	ret := 0
	curMax := 0
	curR := 0
	for i := 0; i < len(nums); i++ {
		if curR < i {
			ret++
			curR = curMax
		}
		if curMax < nums[i]+i {
			curMax = nums[i] + i
		}
	}
	return ret
}
