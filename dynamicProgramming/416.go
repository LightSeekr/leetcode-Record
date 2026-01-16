package dynamicProgramming

/*
*
给你一个 只包含正整数 的 非空 数组 nums 。
请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
*/
func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	//奇数
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		for j := target; j >= nums[i]; j-- {
			if dp[j-nums[i]] {
				dp[j] = true
			}
		}
	}
	return dp[target]
}
