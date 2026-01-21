package greedy

func maxProfit2(prices []int) int {
	profits := 0
	n := len(prices)
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			profits += prices[i] - prices[i-1]
		}
	}
	return profits
}
