package backtrack

import "sort"

/*
*
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
*/
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	path := []int{}
	sum := 0

	// 同一个数字都可以无限制重复被选取
	var backtrack func(int)
	backtrack = func(start int) {
		if sum > target {
			return
		}

		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i < len(candidates); i++ {
			sum += candidates[i]
			path = append(path, candidates[i])
			backtrack(i)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}

	// 从小到达排序
	sort.Ints(candidates)
	backtrack(0)

	return res
}
