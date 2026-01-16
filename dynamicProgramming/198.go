package dynamicProgramming

/**


状态转移方程：
$$ dp[i] = \max(dp[i-1], \quad nums[i] + dp[i-2]) $$
(不偷第 i 间 vs 偷第 i 间)

初始条件：

dp[0] = nums[0] （只有一间房，必偷）
dp[1] = max(nums[0], nums[1]) （有两间房，选钱多的一间偷，因为相邻不能一起偷）
*/

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	res := max(dp[0], dp[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		res = max(res, dp[i])
	}
	return res
}
