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
