package enumerate

/*
*
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
*/
func maxProfit(prices []int) int {
	//维护这个数左边的最小值
	minValue := prices[0]
	res := 0
	for _, p := range prices {
		//对于 p 来说
		res = max(res, p-minValue)
		// 对于 p 后面的元素来说
		minValue = min(minValue, p)
	}
	return res
}

func MaxProfit(prices []int) int {
	return maxProfit(prices)
}
