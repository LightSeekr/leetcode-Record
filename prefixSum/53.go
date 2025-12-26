package prefixSum

/*
*

给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组是数组中的一个连续部分。
[-2,1,-3,4,-1,2,1,-5,4]
数组中有负数，不能用滑动窗口
*/
func maxSubArray(nums []int) int {
	//n := len(nums)
	ans := nums[0]
	minPreSum := 0
	nowPreSum := 0
	// 连续子数组+求和 可以转成 两个前缀和相减
	for _, num := range nums {
		nowPreSum += num                      // 当前前缀和
		ans = max(ans, nowPreSum-minPreSum)   // 当前前缀和减去最小前缀和。就是两个对应的子数组之间的和
		minPreSum = min(minPreSum, nowPreSum) // 更新最小前缀和
	}
	return ans
}
