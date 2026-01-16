package dynamicProgramming

/**
给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续 子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
测试用例的答案是一个 32-位 整数。
请注意，一个只包含一个元素的数组的乘积是这个元素的值。
*/

func maxProduct(nums []int) int {
	// 有负数，可能就要 负负得正
	// 需要同时维护最大值和最小值，因为负数×最小值可能得到最大值
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dpMax := make([]int, n) // 以 i 结尾的子数组的最大乘积
	dpMin := make([]int, n) // 以 i 结尾的子数组的最小乘积
	dpMax[0] = nums[0]
	dpMin[0] = nums[0]
	res := dpMax[0]

	for i := 1; i < n; i++ {
		// 当前数为正数时：最大值 = max(当前数, 当前数×前面最大值)
		// 当前数为负数时：最大值 = max(当前数, 当前数×前面最小值) - 负负得正
		dpMax[i] = max(nums[i], max(nums[i]*dpMax[i-1], nums[i]*dpMin[i-1]))
		dpMin[i] = min(nums[i], min(nums[i]*dpMax[i-1], nums[i]*dpMin[i-1]))
		res = max(res, dpMax[i])
	}
	return res
}
