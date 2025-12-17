package SlideWindow

// 请你返回子数组内所有元素的乘积严格小于 k 的连续子数组的数目
func numSubarrayProductLessThanK(nums []int, k int) int {
	//n := len(nums)
	// l 是左端点
	//ans 是答案
	//product 是乘积
	l, ans, product := 0, 0, 1
	for r, num := range nums {
		product *= num
		for l <= r && product >= k {
			product /= nums[l]
			l++
		}
		if product < k {
			//product <k
			//[l,r] 中间的都符合条件
			ans += r - l + 1
		}

	}
	return ans
}

func NumSubarrayProductLessThanK(nums []int, k int) int {
	return numSubarrayProductLessThanK(nums, k)
}
