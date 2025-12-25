package twoPointers

func removeElement(nums []int, val int) int {
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] != val {
			//因为这里是相同元素一个都不保留
			nums[i] = nums[j]
			i++
		}
		j++
	}
	return i
}

// RemoveElement 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素
func RemoveElement(nums []int, val int) int {
	return removeElement(nums, val)
}
