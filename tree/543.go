package tree

/**
1. 直径的定义：路径长度 = 路径上的边数。
2. 关键转化：对于任意一个节点 node，经过该节点的“最长路径”长度等于：左子树的最大深度 + 右子树的最大深度
3. 全局最优：整棵树的直径，就是所有节点中，上述计算结果（左+右）的最大值。

注意：最长路径不一定经过根节点 root，它可能完全包含在某个子树内部。
我们需要一个递归函数（DFS），它要做两件事：
1. 向上汇报：告诉父节点，“以我为起点的最长单侧路径（深度）”是多少。（为了给父节点算它的直径用）。
2. 更新全局最大值：计算“穿过当前节点的最长路径”，并尝试更新全局的 ans。
*/

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0

	// 定义递归函数
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		// 递归终止条件
		if node == nil {
			return 0
		}

		// 1. 获取左子树的最大深度（单链长度）
		leftDepth := dfs(node.Left)

		// 2. 获取右子树的最大深度（单链长度）
		rightDepth := dfs(node.Right)

		// 3. 计算穿过当前节点的最长路径长度
		// 路径长度 = 左边深度 + 右边深度
		// 尝试更新全局最大值

		maxDiameter = max(maxDiameter, leftDepth+rightDepth)

		// 4. 返回值：返回给父节点的是以当前节点为端点的最长链长度
		// 必须选左边或右边更长的一条，并加上当前节点自己 (+1)
		return max(leftDepth, rightDepth) + 1
	}
	dfs(root)
	return maxDiameter
}
