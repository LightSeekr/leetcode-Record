package tree

/*
*
二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。
*
*/

func maxPathSum(root *TreeNode) int {
	maxSum := root.Val // 初始化为根节点的值，因为至少包含一个节点

	// 辅助函数：计算以当前节点为起点的最大路径和
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归计算左右子树的最大贡献值
		// 如果贡献值为负数，则不选择该子树（取0）
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		// 当前节点作为路径的最高点时的路径和
		// 这条路径包括：左子树的最大路径 + 当前节点 + 右子树的最大路径
		priceNewPath := node.Val + leftGain + rightGain

		// 更新全局最大路径和
		maxSum = max(maxSum, priceNewPath)

		// 返回当前节点的最大贡献值（只能选择左或右其中一条路径）
		// 因为路径不能分叉，所以只能选择左或右中较大的一个
		return node.Val + max(leftGain, rightGain)
	}

	maxGain(root)
	return maxSum
}
