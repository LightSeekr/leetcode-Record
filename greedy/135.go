package greedy

/**
初始化一个数组 candies，长度为 n，每个位置先填 1（满足每个孩子至少一个）。
从左向右遍历：
如果 ratings[i] > ratings[i-1]，则 candies[i] = candies[i-1] + 1。
从右向左遍历：
如果 ratings[i] > ratings[i+1]，说明第 i 个应该比右边多。
此时需要检查：“我现有的糖果是不是已经比右边多了？”
如果现有的不够（candies[i] <= candies[i+1]），那就更新为 candies[i+1] + 1。
如果现有的已经够了（因为左规则给了很多），就保持不变。
即：candies[i] = max(candies[i], candies[i+1] + 1)。
计算总和。
*/

func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i+1]+1, candies[i])
		}
	}
	res := 0
	for _, num := range candies {
		res += num
	}
	return res
}
