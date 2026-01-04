package tree

/*
*
给定一个二叉树，判断它是否是 平衡二叉树
*
*/

func isBalanced(root *TreeNode) bool {
	_, balanced := checkBalanced(root)
	return balanced
}

func checkBalanced(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	leftH, leftOk := checkBalanced(node.Left)
	if !leftOk {
		return 0, false
	} // 提前退出

	rightH, rightOk := checkBalanced(node.Right)
	if !rightOk {
		return 0, false
	} // 提前退出

	// 检查差值
	diff := leftH - rightH
	if diff < -1 || diff > 1 {
		return 0, false
	}

	// 返回 max(left, right) + 1 和 true
	return max(leftH, rightH) + 1, true
}
