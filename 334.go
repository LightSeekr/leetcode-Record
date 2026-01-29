package leetcode

/*
*
给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。

如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k] ，返回 true ；否则，返回 false 。
*/
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	// 维护两个变量：最小值和第二小值
	// first: 目前遇到的最小值
	// second: 在找到first之后遇到的最小值（即第二小的值）
	first := nums[0]
	second := int(^uint(0) >> 1) // 初始化为int最大值

	for i := 1; i < len(nums); i++ {
		if nums[i] > second {
			// 找到了比second大的数，说明存在递增三元组
			return true
		} else if nums[i] > first {
			// 更新second为当前值
			second = nums[i]
		} else {
			// 更新first为更小的值
			first = nums[i]
		}
	}

	return false
}
