package tree

/*
*
给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。

路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

*
*/
func pathSum(root *TreeNode, targetSum int) int {
	//用一个 哈希表 (Map) 来记录历史上出现过的所有前缀和及其出现次数。
	// key: 前缀和, value: 该前缀和出现的次数
	// 这里的 map 充当了“历史记录”的角色
	preSumMap := make(map[int]int)

	// 初始化边界条件：
	// 如果一条路径直接从根节点开始，且和刚好为 targetSum
	// 那么 currSum - targetSum = 0。我们需要 Map[0] = 1 来匹配这种情况。
	preSumMap[0] = 1

	return dfs(root, 0, targetSum, preSumMap)
}

func dfs(root *TreeNode, currSum int, target int, preSumMap map[int]int) int {
	if root == nil {
		return 0
	}

	// 1. 更新当前路径的前缀和
	currSum += root.Val

	// 2. 核心逻辑：查找有多少条满足条件的历史路径
	// 我们想找 oldSum，使得 currSum - oldSum = target
	// 所以 oldSum = currSum - target
	count := preSumMap[currSum-target]
	// 3. 将当前的前缀和记录到 Map 中，供子孙节点查询
	preSumMap[currSum]++
	// 4. 递归处理左右子树，并将结果累加
	count += dfs(root.Left, currSum, target, preSumMap)
	count += dfs(root.Right, currSum, target, preSumMap)
	// 5. 回溯 (Backtracking) —— 非常重要！
	// 离开当前节点前，必须把它从前缀和记录中移除。
	// 因为这棵子树遍历完了，它的前缀和不能影响它的兄弟节点（右子树）。
	preSumMap[currSum]--

	return count
}
