package prefixSum

import "fmt"

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
func productExceptSelf(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			left[i] = nums[i]
		} else {
			left[i] = left[i-1] * nums[i]
		}
	}
	right := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			right[i] = nums[i]
		} else {
			right[i] = right[i+1] * nums[i]
		}
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			res[i] = right[i+1]
		} else if i == n-1 {
			res[i] = left[i-1]
		} else {
			res[i] = left[i-1] * right[i+1]
		}
	}
	return res
}

func answer(nums []int) []int {
	res := make([]int, len(nums))
	n := 1
	for _, num := range nums {
		n *= num
	}
	for i := 0; i < len(nums); i++ {
		res[i] = n / nums[i]
	}
	fmt.Println("answer: ", res)
	return res
}
