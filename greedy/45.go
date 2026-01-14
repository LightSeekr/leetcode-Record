package greedy

func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	steps := 0  // 跳跃次数
	end := 0    // 当前这一跳能到达的最远边界
	maxPos := 0 // 在当前范围内，能探测到的最远位置（为下一跳做准备）

	for i := 0; i < n-1; i++ {
		maxPos = max(maxPos, nums[i]+i)
		if i == end {
			end = maxPos
			steps++
			if end >= n-1 {
				break
			}
		}
	}
	return steps
}
