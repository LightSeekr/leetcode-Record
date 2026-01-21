package greedy

/*
*
初始化 totalSurplus (总剩余油量) 和 currentSurplus (当前段剩余油量) 为 0。
初始化 start (起点) 为 0。
遍历每一个加油站 i：
计算当前的净收益 diff = gas[i] - cost[i]。
累加到 totalSurplus 和 currentSurplus。
关键贪心逻辑：如果 currentSurplus < 0，说明从当前的 start 出发，走到 i 就没油了。
那么 start 到 i 之间的所有点都不可能是起点。
我们将起点重置为 i+1。
重置 currentSurplus = 0 (因为我们要从新的起点重新开始算积累)。
遍历结束后，如果 totalSurplus < 0，说明总油量不够，返回 -1。
否则，返回 start。
*/
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	totalSurplus, currentSurplus := 0, 0
	start, diff := 0, 0
	for i := 0; i < n; i++ {
		diff = gas[i] - cost[i]
		totalSurplus += diff
		currentSurplus += diff
		if currentSurplus < 0 {
			start = i + 1
			currentSurplus = 0
		}
	}
	if totalSurplus < 0 {
		return -1
	}
	return start
}
