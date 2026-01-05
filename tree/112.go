package tree

/**
给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。
判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。
如果存在，返回 true ；否则，返回 false 。
*/

func hasPathSum(root *TreeNode, targetSum int) bool {
	// 1. 递归终止条件：空节点
	// 注意：即使 targetSum 为 0，空树也不包含路径，所以返回 false
	if root == nil {
		return false // 空节点，就算 targetSum==0 也不行
	}
	// 2. 判断是否是叶子节点
	// 如果是叶子节点，判断当前值是否等于剩下的目标值
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return root.Val == targetSum
	}
	// 3. 递归处理左右子树
	// 只要有一边能找到路径即可，所以用 || (或)
	// 传入新的 targetSum = targetSum - root.Val
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)

}
