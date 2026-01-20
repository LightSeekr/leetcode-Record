package trick

func majorityElement(nums []int) int {
	candidate := 0
	count := 0

	for _, num := range nums {
		// 如果当前票数为 0，更换候选人
		if count == 0 {
			candidate = num
		}

		// 投票逻辑
		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}
