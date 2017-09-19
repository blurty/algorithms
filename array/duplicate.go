package array

func RemoveDuplicates(nums []int) int {
	length := len(nums)
    if length < 2 {
        return length
    }
    deleted_num := 0
    for i:=1; i<length; i++ {
        if nums[i - 1 - deleted_num] == nums[i] {
            deleted_num++
        } else if (deleted_num > 0) {
            nums[i - deleted_num] = nums[i]
        }
    }
    return length - deleted_num
}