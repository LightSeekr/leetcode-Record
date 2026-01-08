package backtrack

/*
*
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
这里的不能重复很恶心啊
*/
func subsets(nums []int) [][]int {
	var res [][]int
	var path []int
	var backtrack func(int)

	backtrack = func(start int) {
		// 加入结果
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)

		//if start == len(nums) {
		//	return
		//}
		// 这里的边界 条件就是 start==n
		//但是 start ==n 的时候，也进不去下面那个循环

		// 只允许从这个数往后选，避免重复
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return res
}
