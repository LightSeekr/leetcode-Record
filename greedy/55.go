package greedy

/*
*
给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。
*/
func canJump(nums []int) bool {
	n := len(nums)
	// 维护一个能到达的最远距离
	maxReach := 0
	for i := 0; i < n; i++ {
		if i > maxReach {
			return false
		}
		maxReach = max(maxReach, i+nums[i])
		// 说明一定能到
		if maxReach >= n-1 {
			return true
		}
	}
	return true
}
