package slideWindow

// 滑动窗口的题目，是否与子数组相关呢？ 是的
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	sum := 0
	ans := n + 1
	l := 0
	for r := 0; r < n; r++ {
		sum += nums[r]
		//找出该数组中满足其总和大于等于 target 的长度最小的 子数组,并返回长度
		for sum >= target {
			//[left,right] 满足条件
			temp := r - l + 1
			ans = min(ans, temp) // min是新用法啊
			// 下一步，左移left 找最优解
			sum -= nums[l]
			l++
		}
	}

	// 子数组长度最多为n,不可能为 n+1
	if ans == n+1 {
		return 0
	}
	return ans
}

func MinSubArrayLen(target int, nums []int) int {
	return minSubArrayLen(target, nums)
}

//给定一个含有 n 个正整数的数组和一个正整数 target 。
//
//找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
