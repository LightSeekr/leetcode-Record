package dynamicProgramming

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0          // 没股票
	dp[0][1] = -prices[0] // 有股票 (买入花了钱)

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return max(dp[n-1][0], dp[n-1][1])
}
