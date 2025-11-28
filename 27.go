package leetcode

func RemoveElement(nums []int, val int) int {
	tail := len(nums) - 1
	// 边界 case
	if tail == -1 {
		return 0
	}
	stackSize := 0
	for _, num := range nums {
		if num != val {
			nums[stackSize] = num
			stackSize++
		}
	}
	return stackSize
}
