package enumerate

import "fmt"

/*
*
给你一个下标从 0 开始的整数数组 nums ，该数组的大小为 n ，请你计算 nums[j] - nums[i] 能求得的 最大差值 ，
其中 0 <= i < j < n 且 nums[i] < nums[j] 。
返回 最大差值 。如果不存在满足要求的 i 和 j ，返回 -1 。
*/
func maximumDifference(prices []int) int {
	//维护这个数左边的最小值
	minValue := prices[0]
	res := -1
	for _, p := range prices {
		//对于 p 来说
		res = max(res, p-minValue)
		fmt.Printf("p:%d minValue:%d res:%d\n", p, minValue, res)
		// 对于 p 后面的元素来说
		minValue = min(minValue, p)
	}
	// 这里就是需要特殊判断的
	if res == 0 {
		return -1
	}
	return res
}

func MaximumDifference(nums []int) int {
	return maximumDifference(nums)
}
