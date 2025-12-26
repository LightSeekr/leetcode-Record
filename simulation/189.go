package simulation

import "slices"

/**
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
1,2,3,4,5,6,7   k = 3
5 6 7 1 2 3 4
*/

// 原地替换 使用 O(1) 的空间
func rotate(nums []int, k int) {
	k %= len(nums) // 轮转 k 次等于轮转 k % n 次
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}

// 下面是使用 切片的性质做的
func rotate_b(nums []int, k int) {
	n := len(nums)
	k = k % n // k 要取余数

	t2 := n - k
	res := make([]int, 0)
	res = append(res, nums[t2:]...)
	res = append(res, nums[0:t2]...)
	for i := 0; i < n; i++ {
		nums[i] = res[i]
	}
	return
}
