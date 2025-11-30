package TwoPointers

func removeDuplicates(nums []int) int {
	i, j := 0, 1
	for j < len(nums) {
		for j < len(nums) && nums[i] == nums[j] {
			j++
		}
		if j < len(nums) {
			i++
			nums[i] = nums[j]
			j++
		}
	}
	return i + 1
}

func RemoveDuplicates(nums []int) int {
	return removeDuplicates(nums)
}
