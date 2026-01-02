package tree

/*
*
一个二叉树的最大深度 = max(左子树深度, 右子树深度) + 1。
这是一个典型的“后序遍历”逻辑：先算出左右孩子的深度，再汇总给自己。
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 1. 计算左子树的深度
	leftDepth := maxDepth(root.Left)

	// 2. 计算右子树的深度
	rightDepth := maxDepth(root.Right)

	// 3. 取最大值并加 1 (加 1 是因为要算上当前根节点这一层)
	return max(leftDepth, rightDepth) + 1

}
