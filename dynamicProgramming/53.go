package dynamicProgramming

/**


给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组是数组中的一个连续部分。
[-2,1,-3,4,-1,2,1,-5,4]
最大连续和的子数组
*/

/*
*
数组中有负数，所以不能用滑动窗口
连续子数组，注意这个连续

dp[i]  表示以 nums[i] 结尾的连续子数组的 最大和
dp[0]=-2
if dp[i-1] >0{
dp[i]=dp[i-1]+nums[i]// 如果和之前的子数组可以连接上，于是可以
}else{
dp[i]=nums[i]
}
*/
func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	ans := nums[0]
	for i := 1; i < n; i++ {
		//dp[i] = max(nums[i], dp[i-1]+nums[i])
		dp[i] = max(dp[i-1], 0) + nums[i]
		ans = max(ans, dp[i])
	}
	return ans
}
