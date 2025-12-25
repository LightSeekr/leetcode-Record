package twoPointers

import "sort"

/*
*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]]
满足 i != j、i != k 且 j != k ，
同时还满足 nums[i] + nums[j] + nums[k] == 0 。
请你返回所有和为 0 且不重复的三元组。
待定：这题要的是值，而不是下标，可以尝试下排序会更有帮助吗？

一开始用 固定左，枚举右 但是没法去重复，然后看到了 题目要的是值，而不是下标，想到了可以排序。

然后用 AI 一把梭哈了。

排序有帮助，自己也想出来了正确的解法
可惜提前问 AI 了
这题还有点小细节要处理，手写估计还挺能锻炼的
*/

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)
	if n < 3 {
		return res
	}

	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		// 去重：固定第一个数
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 剪枝：最小值都 > 0，后面不可能凑出 0
		if nums[i] > 0 {
			break
		}

		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			switch {
			case sum == 0:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 去重：移动 l / r 时跳过相同值
				leftVal, rightVal := nums[l], nums[r]
				// 这里就是小细节 为的是去重复
				for l < r && nums[l] == leftVal {
					l++
				}
				for l < r && nums[r] == rightVal {
					r--
				}
			case sum < 0:
				l++
			default: // sum>0
				r--
			}
		}
	}
	return res
}

func threeSum_b(nums []int) [][]int {
	// 暴力解法（带去重）不建议
	res := make([][]int, 0)
	n := len(nums)
	seen := make(map[[3]int]struct{}, 0)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					a, b, c := nums[i], nums[j], nums[k]
					if a > b {
						a, b = b, a
					}
					if b > c {
						b, c = c, b
					}
					if a > b {
						a, b = b, a
					}

					key := [3]int{a, b, c}
					if _, ok := seen[key]; ok {
						continue
					}
					seen[key] = struct{}{}
					res = append(res, []int{a, b, c})
				}
			}
		}
	}
	return res
}
