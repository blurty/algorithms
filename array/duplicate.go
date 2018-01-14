package array

// remove the duplicates in place such that each element
// appear only once and return the new length in sorted array
func RemoveDuplicates(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}
	deleted_num := 0
	for i := 1; i < length; i++ {
		if nums[i-1-deleted_num] == nums[i] {
			deleted_num++
		} else if deleted_num > 0 {
			nums[i-deleted_num] = nums[i]
		}
	}
	return length - deleted_num
}

// each element are allowed twice in sorted array
// return the new length
func RemoveDuplicates2(nums []int) int {
	length := len(nums)
	if length < 3 {
		return length
	}
	deleted_num := 0
	for i := 2; i < length; i++ {
		if nums[i] == nums[i-2-deleted_num] {
			deleted_num++
		} else if deleted_num > 0 {
			nums[i-deleted_num] = nums[i]
		}
	}
	return length - deleted_num
}
