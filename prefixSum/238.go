package prefixSum

/*
*
给你一个整数数组 nums，返回 数组 answer ，
其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
请 不要使用除法
不用除法。。。那还玩个6啊
那就是用前缀积
左边的，右边的，分别求一次前缀积
可以进一步优化，然后 不用特判 0 和 n-1 的情况
前后缀分解
*/
//func productExceptSelf(nums []int) []int {
//	n := len(nums)
//	left := make([]int, n)
//	for i := 0; i < n; i++ {
//		if i == 0 {
//			left[i] = nums[i]
//		} else {
//			left[i] = left[i-1] * nums[i]
//		}
//	}
//	right := make([]int, n)
//	for i := n - 1; i >= 0; i-- {
//		if i == n-1 {
//			right[i] = nums[i]
//		} else {
//			right[i] = right[i+1] * nums[i]
//		}
//	}
//	res := make([]int, n)
//	for i := 0; i < n; i++ {
//		if i == 0 {
//			res[i] = right[i+1]
//		} else if i == n-1 {
//			res[i] = left[i-1]
//		} else {
//			res[i] = left[i-1] * right[i+1]
//		}
//	}
//	return res
//}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	// 构建左右的前缀积
	// left[i] 表示 nums[0] * nums[1] * ... * nums[i]
	left := make([]int, n)
	right := make([]int, n)
	left[0] = nums[0]
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i]
	}
	// right[i] 表示 nums[i] * nums[i+2] * ... * nums[n-1]
	right[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i]
	}
	//fmt.Println(right)

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = 1
		if i-1 >= 0 {
			ans[i] *= left[i-1]
		}
		if i+1 < n {
			ans[i] *= right[i+1]
		}
	}
	return ans
}
