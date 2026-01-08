package backtrack

/**
给定一个不含重复数字的数组 nums ，
返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
用数组 nums，构造一个长度为 len(nums) 的全排列
怎么做？
假设数组长度为 n 也就是说，一共有 n 个数。
先从中选 1个数，再拿剩下的 n-1 个数构造一个长度为 n-1 的排列
这不把问题变成了规模更小的子问题
这种就是用递归来解决。
边界条件就是全都选完了，可以收集结果了。

*/

func permute(nums []int) [][]int {
	var res [][]int
	var path []int
	n := len(nums)

	used := make([]bool, n)

	var backtrack func()
	backtrack = func() {
		// 边界条件
		if len(path) == n {
			temp := make([]int, n)
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := 0; i < n; i++ {
			// 之前选过了，本次不选
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			//向下递归
			backtrack()
			// 回溯
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()
	return res
}
