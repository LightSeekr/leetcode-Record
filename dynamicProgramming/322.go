package dynamicProgramming

/*
*
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。
*/
func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 // 使用 amount+1 代替 math.MaxInt，避免溢出
	}

	dp[0] = 0
	// 凑这个数的话，只需要 1枚，因为有
	//for _, coin := range coins {
	//	dp[coin] = 1
	//}
	/**
	上面这样初始化的话，过不了
	coins:  []int{2},
	amount: 1,
	这个例子
	*/
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			// 这里借助 dp[0] 完成初始化，这里不可以是 <=
			if i-coin < 0 {
				continue
			}
			// 对于这个数 i
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	// 考虑凑不出来的情况
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
